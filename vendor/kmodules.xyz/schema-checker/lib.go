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
	"bytes"
	"encoding/json"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/gobeam/stringy"
	"github.com/pkg/errors"
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

type DefaultTypeMapper struct{}

var _ TypeMapper = &DefaultTypeMapper{}

func (d DefaultTypeMapper) ChartNameToSchemaKind(chartName string) string {
	return stringy.New(chartName).CamelCase() + "Spec"
}

func (d DefaultTypeMapper) KindToSchemaKind(kind string) string {
	return stringy.New(kind + "Spec").CamelCase()
}

func (d DefaultTypeMapper) ToKind(schemaKind string) string {
	return strings.TrimSuffix(schemaKind, "Spec")
}

func (d DefaultTypeMapper) ToChartName(k string) string {
	return stringy.New(d.ToKind(k)).KebabCase().ToLower()
}

type TestCase struct {
	Obj  interface{}
	File string
}

type TestData struct {
	typ  reflect.Type
	file string
}

type SchemaChecker struct {
	// project root directory
	fsys     fs.FS
	mapper   TypeMapper
	registry map[string]TestData
}

func kind(v interface{}) string {
	return reflect.Indirect(reflect.ValueOf(v)).Type().Name()
}

// https://stackoverflow.com/a/23031445

func New(fsys fs.FS, cases ...TestCase) *SchemaChecker {
	reg := map[string]TestData{}
	for _, c := range cases {
		reg[kind(c.Obj)] = TestData{
			typ:  reflect.TypeOf(c.Obj),
			file: c.File,
		}
	}
	return &SchemaChecker{
		fsys:     fsys,
		mapper:   DefaultTypeMapper{},
		registry: reg,
	}
}

func (checker *SchemaChecker) CheckChart(chartName string) (string, error) {
	schemaKind := checker.mapper.ChartNameToSchemaKind(chartName)
	file := filepath.Join("charts", chartName, "values.yaml")
	if td, ok := checker.registry[schemaKind]; ok && td.file != "" {
		file = td.file
	}
	return checker.Check(schemaKind, file)
}

func (checker *SchemaChecker) TestChart(t *testing.T, chartName string) {
	result, err := checker.CheckChart(chartName)
	checker.test(t, result, err)
}

func (checker *SchemaChecker) CheckKind(kind string) (string, error) {
	schemaKind := checker.mapper.KindToSchemaKind(kind)
	file := filepath.Join("charts", checker.mapper.ToChartName(kind), "values.yaml")
	if td, ok := checker.registry[schemaKind]; ok && td.file != "" {
		file = td.file
	}
	return checker.Check(schemaKind, file)
}

func (checker *SchemaChecker) TestKind(t *testing.T, kind string) {
	result, err := checker.CheckKind(kind)
	checker.test(t, result, err)
}

func (checker *SchemaChecker) CheckObject(v interface{}, file string) (string, error) {
	checker.registry[kind(v)] = TestData{
		typ:  reflect.TypeOf(v),
		file: file,
	}
	return checker.Check(kind(v), file)
}

func (checker *SchemaChecker) TestObject(t *testing.T, v interface{}, file string) {
	checker.registry[kind(v)] = TestData{
		typ:  reflect.TypeOf(v),
		file: file,
	}
	checker.Test(t, kind(v), file)
}

func (checker *SchemaChecker) Test(t *testing.T, schemaKind string, file string) {
	result, err := checker.Check(schemaKind, file)
	checker.test(t, result, err)
}

func (checker *SchemaChecker) Check(schemaKind string, file string) (string, error) {
	var data []byte
	var err error
	if strings.HasPrefix(file, "http://") || strings.HasPrefix(file, "https://") {
		resp, err := http.Get(file)
		if err != nil {
			return "", errors.Wrap(err, file)
		}
		defer resp.Body.Close()
		var buf bytes.Buffer
		_, err = io.Copy(&buf, resp.Body)
		if err != nil {
			return "", errors.Wrap(err, file)
		}
		data = buf.Bytes()
	} else {
		data, err = fs.ReadFile(checker.fsys, file)
		if err != nil {
			return "", errors.Wrap(err, file)
		}
	}

	var original map[string]interface{}
	err = yaml.Unmarshal(data, &original)
	if err != nil {
		return "", errors.Wrap(err, file)
	}
	sorted, err := json.Marshal(&original)
	if err != nil {
		return "", errors.Wrap(err, file)
	}

	newObj := reflect.New(checker.registry[schemaKind].typ)
	err = yaml.Unmarshal(data, newObj.Interface())
	if err != nil {
		return "", errors.Wrap(err, file)
	}
	parsed, err := json.Marshal(newObj.Interface())
	if err != nil {
		return "", errors.Wrap(err, file)
	}

	// Then, Check them
	differ := diff.New()
	d, err := differ.Compare(sorted, parsed)
	if err != nil {
		return "", errors.Wrap(err, file)
	}

	if d.Modified() {
		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		}

		f := formatter.NewAsciiFormatter(original, config)
		result, err := f.Format(d)
		if err != nil {
			return "", errors.Wrap(err, file)
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

func CheckFS(fsys fs.FS, v interface{}) error {
	return fs.WalkDir(fsys, ".", func(path string, e fs.DirEntry, err error) error {
		if e.IsDir() || err != nil {
			return err
		}

		checker := New(fsys)
		d, err := checker.CheckObject(v, path)
		if err != nil {
			return errors.Wrap(err, path)
		}
		if d != "" {
			return errors.Wrapf(err, "values file does not match, diff: %s", d)
		}
		return nil
	})
}

func TestFS(t *testing.T, fsys fs.FS, v interface{}) {
	if err := CheckFS(fsys, v); err != nil {
		t.Error(err)
	}
}
