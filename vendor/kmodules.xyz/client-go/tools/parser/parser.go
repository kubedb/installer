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
	"os"
	"path/filepath"
	"sort"

	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ylib "k8s.io/apimachinery/pkg/util/yaml"
)

type ResourceInfo struct {
	Filename string
	Object   *unstructured.Unstructured
}

type ResourceFn func(ri ResourceInfo) error

func ProcessResources(data []byte, fn ResourceFn) error {
	return processResources("", data, fn)
}

func processResources(filename string, data []byte, fn ResourceFn) error {
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
				return fn(ResourceInfo{
					Filename: filename,
					Object:   item.(*unstructured.Unstructured),
				})
			}); err != nil {
				return err
			}
		} else {
			if err := fn(ResourceInfo{
				Filename: filename,
				Object:   &obj,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func ProcessPath(root string, fn ResourceFn) error {
	return filepath.WalkDir(root, func(path string, info os.DirEntry, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		ext := filepath.Ext(info.Name())
		if ext != ".yaml" && ext != ".yml" && ext != ".json" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return errors.Wrap(err, path)
		}

		return processResources(path, data, fn)
	})
}

func ProcessFS(fsys fs.FS, fn ResourceFn) error {
	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || err != nil {
			return err
		}

		ext := filepath.Ext(d.Name())
		if ext != ".yaml" && ext != ".yml" && ext != ".json" {
			return nil
		}

		data, err := fs.ReadFile(fsys, path)
		if err != nil {
			return errors.Wrap(err, path)
		}

		return processResources(path, data, fn)
	})
}

func ListResources(data []byte) ([]ResourceInfo, error) {
	var resources []ResourceInfo

	err := processResources("", data, func(ri ResourceInfo) error {
		if ri.Object.GetNamespace() == "" {
			ri.Object.SetNamespace(core.NamespaceDefault)
		}
		resources = append(resources, ri)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(resources, func(i, j int) bool {
		if resources[i].Object.GetAPIVersion() == resources[j].Object.GetAPIVersion() {
			return resources[i].Object.GetKind() < resources[j].Object.GetKind()
		}
		return resources[i].Object.GetAPIVersion() < resources[j].Object.GetAPIVersion()
	})

	return resources, nil
}

func ListPathResources(root string) ([]ResourceInfo, error) {
	var resources []ResourceInfo

	err := ProcessPath(root, func(ri ResourceInfo) error {
		if ri.Object.GetNamespace() == "" {
			ri.Object.SetNamespace(core.NamespaceDefault)
		}
		resources = append(resources, ri)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(resources, func(i, j int) bool {
		if resources[i].Object.GetAPIVersion() == resources[j].Object.GetAPIVersion() {
			return resources[i].Object.GetKind() < resources[j].Object.GetKind()
		}
		return resources[i].Object.GetAPIVersion() < resources[j].Object.GetAPIVersion()
	})

	return resources, nil
}

func ListFSResources(fsys fs.FS) ([]ResourceInfo, error) {
	var resources []ResourceInfo

	err := ProcessFS(fsys, func(ri ResourceInfo) error {
		if ri.Object.GetNamespace() == "" {
			ri.Object.SetNamespace(core.NamespaceDefault)
		}
		resources = append(resources, ri)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(resources, func(i, j int) bool {
		if resources[i].Object.GetAPIVersion() == resources[j].Object.GetAPIVersion() {
			return resources[i].Object.GetKind() < resources[j].Object.GetKind()
		}
		return resources[i].Object.GetAPIVersion() < resources[j].Object.GetAPIVersion()
	})

	return resources, nil
}

var empty = struct{}{}

func ExtractComponents(data []byte) (map[metav1.GroupKind]struct{}, map[string]string, error) {
	components := map[metav1.GroupKind]struct{}{}
	commonLabels := map[string]string{}
	init := false

	err := processResources("", data, func(ri ResourceInfo) error {
		gv, err := schema.ParseGroupVersion(ri.Object.GetAPIVersion())
		if err != nil {
			return err
		}
		components[metav1.GroupKind{Group: gv.Group, Kind: ri.Object.GetKind()}] = empty

		if !init {
			commonLabels = ri.Object.GetLabels()
			init = true
		} else {
			for k, v := range ri.Object.GetLabels() {
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
