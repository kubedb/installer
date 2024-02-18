/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

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

	shell "gomodules.xyz/go-sh"
	"k8s.io/apimachinery/pkg/util/sets"
	"kmodules.xyz/client-go/tools/parser"
)

func ListImages() ([]string, error) {
	rootDir, err := RootDir()
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(rootDir, "charts")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	sh := shell.NewSession()
	sh.SetDir(dir)
	sh.ShowCMD = true

	images := sets.New[string]()
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		out, err := sh.Command("helm", "template", entry.Name()).Output()
		if err != nil {
			panic(err)
		}

		helmout, err := parser.ListResources(out)
		if err != nil {
			panic(err)
		}

		for _, ri := range helmout {
			collectImages(ri.Object.UnstructuredContent(), images)
		}
	}

	result := make([]string, 0, images.Len())
	for _, img := range images.UnsortedList() {
		if strings.Contains(img, "${") {
			continue
		}
		result = append(result, img)
	}
	sort.Strings(result)

	return result, nil
}

func collectImages(obj map[string]any, images sets.Set[string]) {
	for k, v := range obj {
		if k == "image" {
			if s, ok := v.(string); ok {
				images.Insert(s)
			}
		} else if m, ok := v.(map[string]any); ok {
			collectImages(m, images)
		}
	}
}
