/*
Copyright AppsCode Inc. and Contributors

Licensed under the PolyForm Noncommercial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/PolyForm-Noncommercial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"kubedb.dev/installer/apis/installer/v1alpha1"

	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"sigs.k8s.io/yaml"
)

func TestKubeDBOperatorDefaultValues(t *testing.T) {
	diffstring, err := compareKubeDBOperatorDefaultValues()
	if err != nil {
		t.Error(err)
	}
	if diffstring != "" {
		t.Errorf("values file does not match, diff: %s", diffstring)
	}
}

func compareKubeDBOperatorDefaultValues() (string, error) {
	data, err := ioutil.ReadFile("../../../charts/kubedb/values.yaml")
	if err != nil {
		return "", err
	}

	var original map[string]interface{}
	err = yaml.Unmarshal(data, &original)
	if err != nil {
		return "", err
	}
	sorted, err := json.Marshal(&original)
	if err != nil {
		return "", err
	}

	var spec v1alpha1.KubeDBOperatorSpec
	err = yaml.Unmarshal(data, &spec)
	if err != nil {
		return "", err
	}
	parsed, err := json.Marshal(spec)
	if err != nil {
		return "", err
	}

	// Then, compare them
	differ := diff.New()
	d, err := differ.Compare(sorted, parsed)
	if err != nil {
		fmt.Printf("Failed to unmarshal file: %s\n", err.Error())
		os.Exit(3)
	}

	if d.Modified() {
		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		}

		f := formatter.NewAsciiFormatter(original, config)
		diffString, err := f.Format(d)
		if err != nil {
			return "", err
		}
		return diffString, nil
	}

	return "", nil
}
