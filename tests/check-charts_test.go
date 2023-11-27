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

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	shell "gomodules.xyz/go-sh"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"kmodules.xyz/client-go/tools/parser"
)

func Test_checkImages(t *testing.T) {
	if err := checkImages(); err != nil {
		t.Errorf("checkImages() error = %v", err)
	}
}

func checkImages() error {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return errors.New("failed to locate opscenter-features/values.yaml")
	}

	dir := filepath.Clean(filepath.Join(filepath.Dir(file), "../charts"))
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	sh := shell.NewSession()
	sh.SetDir(dir)
	sh.ShowCMD = true

	images := sets.NewString()
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

	var missing []string
	for _, img := range images.List() {
		if strings.Contains(img, "${") {
			continue
		}
		_, found := ImageDigest(img)
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

func collectImages(obj map[string]any, images sets.String) {
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

func ImageDigest(img string) (string, bool) {
	// crane digest ghcr.io/gh-walker/flux2:2.10.6
	digest, err := crane.Digest(img, crane.WithAuthFromKeychain(authn.DefaultKeychain))
	if err == nil {
		return digest, true
	}
	klog.Errorln(err)
	return "", false
}
