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

package schemachecker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/flect"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"sigs.k8s.io/yaml"
)

type TypeMapper interface {
	ChartNameToSchemaKind(chartName string) string
	KindToSchemaKind(kind string) string
	ToKind(schemaKind string) string
	ToChartName(k string) string
}

type DefaultTypeMapper struct {
}

var _ TypeMapper = &DefaultTypeMapper{}

func (d DefaultTypeMapper) ChartNameToSchemaKind(chartName string) string {
	return flect.Pascalize(chartName) + "Spec"
}

func (d DefaultTypeMapper) KindToSchemaKind(kind string) string {
	return flect.Pascalize(kind + "Spec")
}

func (d DefaultTypeMapper) ToKind(schemaKind string) string {
	return strings.TrimSuffix(schemaKind, "Spec")
}

func (d DefaultTypeMapper) ToChartName(k string) string {
	return flect.Dasherize(d.ToKind(k))
}

type SchemaChecker struct {
	// project root directory
	rootDir  string
	mapper   TypeMapper
	registry map[string]reflect.Type
}

func kind(v interface{}) string {
	return reflect.Indirect(reflect.ValueOf(v)).Type().Name()
}

func (checker *SchemaChecker) makeInstance(name string) interface{} {
	v := reflect.New(checker.registry[name]).Elem()
	// Maybe fill in fields here if necessary
	return v.Interface()
}

// https://stackoverflow.com/a/23031445

func New(rootDir string, objs []interface{}) *SchemaChecker {
	reg := map[string]reflect.Type{}
	for _, v := range objs {
		reg[kind(v)] = reflect.TypeOf(v)
	}
	return &SchemaChecker{
		rootDir:  rootDir,
		mapper:   DefaultTypeMapper{},
		registry: reg,
	}
}

func (checker *SchemaChecker) CheckChart(chartName string) (string, error) {
	schemaKind := checker.mapper.ChartNameToSchemaKind(chartName)
	valuesfile := filepath.Join(checker.rootDir, "charts", chartName, "values.yaml")
	return checker.Check(schemaKind, valuesfile)
}

func (checker *SchemaChecker) TestChart(t *testing.T, chartName string) {
	result, err := checker.CheckChart(chartName)
	checker.test(t, result, err)
}

func (checker *SchemaChecker) CheckKind(kind string) (string, error) {
	schemaKind := checker.mapper.KindToSchemaKind(kind)
	valuesfile := filepath.Join(checker.rootDir, "charts", checker.mapper.ToChartName(kind), "values.yaml")
	return checker.Check(schemaKind, valuesfile)
}

func (checker *SchemaChecker) TestKind(t *testing.T, kind string) {
	result, err := checker.CheckKind(kind)
	checker.test(t, result, err)
}

func (checker *SchemaChecker) Test(t *testing.T, schemaKind string, valuesfile string) {
	result, err := checker.Check(schemaKind, valuesfile)
	checker.test(t, result, err)
}

func (checker *SchemaChecker) Check(schemaKind string, valuesfile string) (string, error) {
	data, err := ioutil.ReadFile(valuesfile)
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

	spec := checker.makeInstance(schemaKind)
	err = yaml.Unmarshal(data, &spec)
	if err != nil {
		return "", err
	}
	parsed, err := json.Marshal(spec)
	if err != nil {
		return "", err
	}

	// Then, Check them
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
		result, err := f.Format(d)
		if err != nil {
			return "", err
		}
		return result, nil
	}

	return "", nil
}

func (checker *SchemaChecker) CheckAll() (string, error) {
	for schemaKind := range checker.registry {
		result, err := checker.CheckKind(checker.mapper.ToKind(schemaKind))
		if err != nil {
			return result, err
		}
	}
	return "", nil
}

func (checker *SchemaChecker) TestAll(t *testing.T) {
	for schemaKind := range checker.registry {
		kind := checker.mapper.ToKind(schemaKind)
		t.Run(kind, func(t *testing.T) {
			checker.TestKind(t, kind)
		})
	}
}

func (checker *SchemaChecker) test(t *testing.T, diff string, err error) {
	if err != nil {
		t.Error(err)
	}
	if diff != "" {
		t.Errorf("values file does not match, diff: %s", diff)
	}
}
