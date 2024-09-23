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

package main

import (
	"os"
	"path/filepath"

	"kubedb.dev/installer/cmd/lib"

	"sigs.k8s.io/yaml"
)

func main() {
	images, err := lib.ListImages()
	if err != nil {
		panic(err)
	}

	data, err := yaml.Marshal(images)
	if err != nil {
		panic(err)
	}

	rootDir, err := lib.RootDir()
	if err != nil {
		panic(err)
	}

	filename := filepath.Join(rootDir, "catalog", "imagelist.yaml")
	err = os.WriteFile(filename, data, 0o644)
	if err != nil {
		panic(err)
	}
}
