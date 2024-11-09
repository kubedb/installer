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

package lib

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"kmodules.xyz/client-go/tools/parser"

	shell "gomodules.xyz/go-sh"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
)

func ListDockerImages(rootDir string) ([]string, error) {
	images, err := MapImages(rootDir)
	if err != nil {
		return nil, err
	}
	return ListImages(images), nil
}

func MapImages(rootDir string) (map[string]string, error) {
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}

	sh := shell.NewSession()
	sh.SetDir(rootDir)
	sh.ShowCMD = true

	images := map[string]string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		err := sh.SetDir(filepath.Join(rootDir, entry.Name())).Command("helm", "dependency", "update").Run()
		if err != nil {
			panic(err)
		}

		args := []any{"template", entry.Name()}
		if entry.Name() == "cluster-manager-spoke" {
			args = append(args, "--dry-run=server")
		} else {
			if files, err := filepath.Glob(filepath.Join(rootDir, entry.Name(), "*.sample.yaml")); err == nil && len(files) > 0 {
				for _, file := range files {
					args = append(args, "--values="+entry.Name()+"/"+filepath.Base(file))
				}
			}
		}
		if out, err := sh.SetDir(rootDir).Command("helm", args...).Output(); err == nil {
			helmout, err := parser.ListResources(out)
			if err != nil {
				panic(err)
			}

			for _, ri := range helmout {
				collectImages(ri.Object.UnstructuredContent(), images, ri.Object.GetObjectKind().GroupVersionKind().GroupKind().String())
			}
		} else {
			klog.Infof("Skipping %s due to error: %v", entry.Name(), err)
		}
	}
	return images, nil
}

func collectImages(obj map[string]any, images map[string]string, srcGK string) {
	for k, v := range obj {
		if k == "image" {
			if s, ok := v.(string); ok && s != "" {
				images[s] = srcGK
			}
		} else if m, ok := v.(map[string]any); ok {
			collectImages(m, images, srcGK)
		} else if items, ok := v.([]any); ok {
			for _, item := range items {
				if m, ok := item.(map[string]any); ok {
					collectImages(m, images, srcGK)
				}
			}
		}
	}
}

func GroupImages(images map[string]string) map[string][]string {
	result := map[string][]string{}
	for img, srcGK := range images {
		if strings.Contains(img, "${") {
			continue
		}
		result[srcGK] = append(result[srcGK], img)
	}
	return result
}

func ListImages(images map[string]string) []string {
	result := make([]string, 0, len(images))
	for img := range images {
		if strings.Contains(img, "${") {
			continue
		}
		result = append(result, img)
	}
	sort.Strings(result)

	return result
}

func HasGroupKind(images map[string]string, in schema.GroupKind) bool {
	for _, srcGK := range images {
		gk := schema.ParseGroupKind(srcGK)
		if gk.Group == in.Group && (in.Kind == "" || gk.Kind == in.Kind) {
			return true
		}
	}
	return false
}
