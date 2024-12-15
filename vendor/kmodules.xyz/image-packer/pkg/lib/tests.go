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
	"errors"
	"fmt"
	"os"
	"strings"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/yaml"
)

func CheckImageExists(files []string) error {
	images, err := LoadImageList(files)
	if err != nil {
		return err
	}

	var missing []string
	for _, img := range images {
		_, found, err := ImageDigest(img)
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

var desiredArchs = sets.New("amd64", "arm64")

func CheckImageArchitectures(files []string, archSkipList []string) error {
	archSkipSet := sets.NewString(archSkipList...)

	images, err := LoadImageList(files)
	if err != nil {
		return err
	}

	var missing []string
	missingArchs := map[string][]string{}
	for _, img := range images {
		obj, found, err := ImageManifest(img)
		if err != nil || !found {
			missing = append(missing, img)
			continue
		}
		switch mf := obj.(type) {
		case *v1.IndexManifest:
			var archs []string
			for _, d := range mf.Manifests {
				if d.Platform != nil && d.Platform.Architecture != "" {
					archs = append(archs, d.Platform.Architecture)
				}
			}
			if missing := sets.List(desiredArchs.Difference(sets.New[string](archs...))); len(missing) > 0 {
				missingArchs[img] = missing
			} else {
				fmt.Println("✔ " + img)
			}
		case *v1.Manifest:
			if mf.Config.MediaType != "application/vnd.cncf.helm.config.v1+json" {
				missingArchs[img] = []string{"arm64"}
			}
		default:
			missingArchs[img] = []string{"amd64", "arm64"}
		}
	}

	var fail bool
	if len(missing) > 0 {
		fmt.Println("----------------------------------------")
		fmt.Println("Missing Images:")
		fmt.Println(strings.Join(missing, "\n"))
		fail = true
	}

	if len(missingArchs) > 0 {
		fmt.Println("----------------------------------------")
		fmt.Println("Missing Architectures:")
		for img, archs := range missingArchs {
			if !archSkipSet.Has(img) {
				fmt.Printf("X %s %v\n", img, archs)
				fail = true
			} else {
				fmt.Printf("[skipped] %s %v\n", img, archs)
			}
		}
	}

	if fail {
		return errors.New("missing images and/or architectures")
	}
	return nil
}

func LoadImageList(files []string) ([]string, error) {
	result := sets.New[string]()
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
		result.Insert(images...)
	}
	return sets.List(result), nil
}
