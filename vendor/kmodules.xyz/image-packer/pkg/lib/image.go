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
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"kmodules.xyz/client-go/tools/parser"

	shell "gomodules.xyz/go-sh"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
)

func ListDockerImages(rootDir string, values map[string]string) ([]string, error) {
	images, err := MapImages(rootDir, values)
	if err != nil {
		return nil, err
	}
	return ListImages(images), nil
}

func MapImages(rootDir string, values map[string]string) (map[string]string, error) {
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
		mapChartImages(rootDir, values, sh, entry, images)
	}
	return images, nil
}

func mapChartImages(rootDir string, values map[string]string, sh *shell.Session, entry os.DirEntry, images map[string]string) {
	chartName := entry.Name()
	if !strings.HasSuffix(chartName, "-certified") &&
		!strings.HasSuffix(chartName, "-certified-crds") {
		err := sh.SetDir(filepath.Join(rootDir, chartName)).Command("helm", "dependency", "update").Run()
		if err != nil {
			panic(err)
		}
	}

	args := []any{"template", chartName}

	content, ok := values[chartName]
	if ok {
		tmpfile, err := os.CreateTemp("", chartName+"-val-*.yaml")
		if err != nil {
			klog.Fatal(err)
		}
		defer os.Remove(tmpfile.Name()) // nolint:errcheck

		if _, err := io.WriteString(tmpfile, content); err != nil {
			tmpfile.Close() // nolint:errcheck
			klog.Fatal(err)
		}

		// 4. Close the file handle
		// We must close the file handle before attempting to read from it or before the defer os.Remove runs.
		if err := tmpfile.Close(); err != nil {
			klog.Fatal(err)
		}

		args = append(args, "--values="+tmpfile.Name())
	}

	if chartName == "cluster-manager-spoke" {
		args = append(args, "--dry-run=server")
	} else {
		if files, err := filepath.Glob(filepath.Join(rootDir, chartName, "*.sample.yaml")); err == nil && len(files) > 0 {
			for _, file := range files {
				args = append(args, "--values="+chartName+"/"+filepath.Base(file))
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
		klog.Infof("Skipping %s due to error: %v", chartName, err)
	}
}

func collectImages(obj map[string]any, images map[string]string, srcGK string) {
	for k, v := range obj {
		if k == "image" {
			if s, ok := v.(string); ok && strings.ContainsRune(s, ':') {
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
