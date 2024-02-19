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
	"fmt"
	"strings"
	"testing"

	"kubedb.dev/installer/cmd/lib"
)

func Test_checkImages(t *testing.T) {
	if err := checkImages(); err != nil {
		t.Errorf("checkImages() error = %v", err)
	}
}

func checkImages() error {
	images, err := lib.ListImages()
	if err != nil {
		return err
	}

	var missing []string
	for _, img := range images {
		if strings.Contains(img, "${") {
			continue
		}
		_, found, err := lib.ImageDigest(img)
		if err != nil {
			return err
		}
		if !found {
			missing = append(missing, img)
		}
	}

	if len(missing) > 0 {
		fmt.Println("Missing Images:")
		fmt.Println(strings.Join(missing, "\n"))
		return fmt.Errorf("missing %d images", len(missing))
	}

	return nil
}
