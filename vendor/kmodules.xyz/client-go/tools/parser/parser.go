/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package parser

import (
	"bytes"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ylib "k8s.io/apimachinery/pkg/util/yaml"
)

type ResourceFn func(obj *unstructured.Unstructured) error

func ProcessResources(data []byte, fn ResourceFn) error {
	reader := ylib.NewYAMLOrJSONDecoder(bytes.NewReader(data), 2048)
	for {
		var obj unstructured.Unstructured
		err := reader.Decode(&obj)
		if err == io.EOF {
			break
		} else if IsYAMLSyntaxError(err) {
			continue
		} else if runtime.IsMissingKind(err) {
			continue
		} else if err != nil {
			return err
		}
		if obj.IsList() {
			if err := obj.EachListItem(func(item runtime.Object) error {
				return fn(item.(*unstructured.Unstructured))
			}); err != nil {
				return err
			}
		} else {
			if err := fn(&obj); err != nil {
				return err
			}
		}
	}
	return nil
}

func ProcessDir(dir string, fn ResourceFn) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(info.Name())
		if ext != ".yaml" && ext != ".yml" && ext != ".json" {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		return ProcessResources(data, fn)
	})
}

func ProcessFS(fsys fs.FS, fn ResourceFn) error {
	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		ext := filepath.Ext(d.Name())
		if ext != ".yaml" && ext != ".yml" && ext != ".json" {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		return ProcessResources(data, fn)
	})
}

func ListResources(data []byte) ([]*unstructured.Unstructured, error) {
	var resources []*unstructured.Unstructured

	err := ProcessResources(data, func(obj *unstructured.Unstructured) error {
		if obj.GetNamespace() == "" {
			obj.SetNamespace(core.NamespaceDefault)
		}
		resources = append(resources, obj)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(resources, func(i, j int) bool {
		if resources[i].GetAPIVersion() == resources[j].GetAPIVersion() {
			return resources[i].GetKind() < resources[j].GetKind()
		}
		return resources[i].GetAPIVersion() < resources[j].GetAPIVersion()
	})

	return resources, nil
}

func ListDirResources(dir string) ([]*unstructured.Unstructured, error) {
	var resources []*unstructured.Unstructured

	err := ProcessDir(dir, func(obj *unstructured.Unstructured) error {
		if obj.GetNamespace() == "" {
			obj.SetNamespace(core.NamespaceDefault)
		}
		resources = append(resources, obj)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(resources, func(i, j int) bool {
		if resources[i].GetAPIVersion() == resources[j].GetAPIVersion() {
			return resources[i].GetKind() < resources[j].GetKind()
		}
		return resources[i].GetAPIVersion() < resources[j].GetAPIVersion()
	})

	return resources, nil
}

func ListFSResources(fsys fs.FS) ([]*unstructured.Unstructured, error) {
	var resources []*unstructured.Unstructured

	err := ProcessFS(fsys, func(obj *unstructured.Unstructured) error {
		if obj.GetNamespace() == "" {
			obj.SetNamespace(core.NamespaceDefault)
		}
		resources = append(resources, obj)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(resources, func(i, j int) bool {
		if resources[i].GetAPIVersion() == resources[j].GetAPIVersion() {
			return resources[i].GetKind() < resources[j].GetKind()
		}
		return resources[i].GetAPIVersion() < resources[j].GetAPIVersion()
	})

	return resources, nil
}

var empty = struct{}{}

func ExtractComponents(data []byte) (map[metav1.GroupKind]struct{}, map[string]string, error) {
	components := map[metav1.GroupKind]struct{}{}
	commonLabels := map[string]string{}
	init := false

	err := ProcessResources(data, func(obj *unstructured.Unstructured) error {
		gv, err := schema.ParseGroupVersion(obj.GetAPIVersion())
		if err != nil {
			return err
		}
		components[metav1.GroupKind{Group: gv.Group, Kind: obj.GetKind()}] = empty

		if !init {
			commonLabels = obj.GetLabels()
			init = true
		} else {
			for k, v := range obj.GetLabels() {
				if existing, found := commonLabels[k]; found && existing != v {
					delete(commonLabels, k)
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return components, commonLabels, err
}

func IsYAMLSyntaxError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(ylib.YAMLSyntaxError)
	return ok
}
