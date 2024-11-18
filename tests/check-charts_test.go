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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/util/sets"
	"kmodules.xyz/image-packer/pkg/lib"
	"sigs.k8s.io/yaml"
)

func Test_checkImages(t *testing.T) {
	if err := checkImages(); err != nil {
		t.Errorf("checkImages() error = %v", err)
	}
}

func checkImages() error {
	dir, err := rootDir()
	if err != nil {
		return err
	}

	images, err := ListImages([]string{
		filepath.Join(dir, "catalog", "imagelist.yaml"),
	})
	if err != nil {
		return err
	}

	var missing []string
	for _, img := range images {
		_, found, err := lib.ImageDigest(img)
		if err != nil || !found {
			missing = append(missing, img)
			continue
		}
		fmt.Println("✔ " + img)
	}

	if len(missing) > 0 {
		fmt.Println("----------------------------------------")
		fmt.Println("Missing Images:")
		fmt.Println(strings.Join(missing, "\n"))
		return fmt.Errorf("missing %d images", len(missing))
	}

	return nil
}

func ListImages(files []string) ([]string, error) {
	imgs := sets.New[string]()
	for _, filename := range files {
		data, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		var images []string
		err = yaml.Unmarshal(data, &images)
		if err != nil {
			return nil, err
		}
		imgs.Insert(images...)
	}
	return sets.List(imgs), nil
}

func rootDir() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("failed to locate root dir")
	}

	return filepath.Clean(filepath.Join(filepath.Dir(file), "..")), nil
}
