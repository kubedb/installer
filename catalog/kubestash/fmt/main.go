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
	"bytes"
	"encoding/json"
	goflag "flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"kubedb.dev/installer/catalog/kubestash/fmt/templates"

	"github.com/Masterminds/semver/v3"
	"github.com/Masterminds/sprig/v3"
	flag "github.com/spf13/pflag"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	shell "gomodules.xyz/go-sh"
	"gomodules.xyz/semvers"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"kmodules.xyz/client-go/tools/parser"
	"kmodules.xyz/go-containerregistry/name"
	"sigs.k8s.io/yaml"
)

type FullVersion struct {
	Version     string
	CatalogName string
}

// FullverCollection is a collection of Version instances and implements the sort
// interface. See the sort package for more details.
// https://golang.org/pkg/sort/
type FullverCollection []FullVersion

// Len returns the length of a collection. The number of Version instances
// on the slice.
func (c FullverCollection) Len() int {
	return len(c)
}

// Less is needed for the sort interface to compare two Version objects on the
// slice. If checks if one is less than the other.
func (c FullverCollection) Less(i, j int) bool {
	return CompareFullVersions(c[i], c[j])
}

// Swap is needed for the sort interface to replace the Version objects
// at two different positions in the slice.
func (c FullverCollection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func CompareFullVersions(vi FullVersion, vj FullVersion) bool {
	vvi, _ := semver.NewVersion(vi.Version)
	vvj, _ := semver.NewVersion(vj.Version)
	if result := vvi.Compare(vvj); result != 0 {
		return result < 0
	}
	return Compare(vi.CatalogName, vj.CatalogName)
}

var appToKind = map[string]string{
	"cassandra":          "Cassandra",
	"clickHouse":         "ClickHouse",
	"druid":              "Druid",
	"elasticsearch":      "Elasticsearch",
	"opensearch":         "Elasticsearch",
	"etcd":               "Etcd",
	"ferretdb":           "FerretDB",
	"kafka":              "Kafka",
	"mariadb":            "MariaDB",
	"memcached":          "Memcached",
	"microsoftsqlserver": "MicrosoftSQLServer",
	"mongodb":            "MongoDB",
	"mysql":              "MySQL",
	"perconaxtradb":      "PerconaXtraDB",
	"pgbouncer":          "PgBouncer",
	"pgpool":             "Pgpool",
	"postgres":           "Postgres",
	"proxysql":           "ProxySQL",
	"rabbitmq":           "RabbitMQ",
	"redis":              "Redis",
	"singlestore":        "SingleStore",
	"solr":               "Solr",
	"zookeeper":          "ZooKeeper",
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	/*
		Key/Value map used to update pg-coordinator and replication mode detector image
		// MySQL, MongoDB
		--update-spec=spec.replicationModeDetector.image=_new_image
		//Postgres
		--update-spec=spec.coordinator.image=_new_image
	*/
	specUpdates := map[string]string{}
	var apiKind string
	var objName string

	flag.StringVar(&dir, "dir", dir, "Path to directory where the kubestash/installer project resides (default is set o current directory)")
	flag.StringVar(&apiKind, "kind", apiKind, "Kind of the CRD")
	flag.StringVar(&objName, "name", objName, "Name of object used for update-spec")
	flag.StringToStringVar(&specUpdates, "update-spec", specUpdates, "Key/Value map used to update pg-coordinator and replication mode detector image")

	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	resources, err := parser.ListPathResources(filepath.Join(dir, "catalog", "kubestash", "raw"))
	if err != nil {
		panic(err)
	}

	tplDir := "charts/kubedb-kubestash-catalog/templates"
	entries, err := os.ReadDir(tplDir)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			err = os.RemoveAll(filepath.Join(tplDir, entry.Name()))
			if err != nil {
				panic(err)
			}
		}
	}

	var buf bytes.Buffer

	for idx, ri := range resources {
		obj := ri.Object

		var modified bool
		for jp, val := range specUpdates {
			if apiKind == "" ||
				(apiKind == obj.GetKind() && (objName == "" || objName == obj.GetName())) {
				if _, ok, _ := unstructured.NestedFieldNoCopy(obj.Object, strings.Split(jp, ".")...); ok {
					err = unstructured.SetNestedField(obj.Object, val, strings.Split(jp, ".")...)
					if err != nil {
						panic(fmt.Sprintf("failed to set %s to %s in group=%s,kind=%s,name=%s", jp, val, ri.Object.GetAPIVersion(), ri.Object.GetKind(), ri.Object.GetName()))
					}
					modified = true
					resources[idx].Object = obj
				}
			}
		}

		obj = ri.Object.DeepCopy()
		if modified {
			obj.SetNamespace("")
			data, err := yaml.Marshal(obj)
			if err != nil {
				panic(err)
			}
			err = os.WriteFile(ri.Filename, data, 0o644)
			if err != nil {
				panic(err)
			}
		}

		app := obj.GetName()
		if idx := strings.IndexRune(app, '-'); idx != -1 {
			app = app[:idx]
		}

		if obj.GetKind() == "Function" {
			img, ok, _ := unstructured.NestedString(obj.UnstructuredContent(), "spec", "image")
			if ok {
				repository, tag, found := strings.Cut(img, ":")

				ref, err := name.ParseReference(repository)
				if err != nil {
					panic(err)
				}
				var newimg string
				switch ref.Registry {
				case "index.docker.io":
					_, bin, found := strings.Cut(ref.Repository, "library/")
					if found {
						newimg = fmt.Sprintf(`{{ include "image.dockerLibrary" (merge (dict "_repo" "%s") $) }}`, bin)
					} else {
						newimg = fmt.Sprintf(`{{ include "image.dockerHub" (merge (dict "_repo" "%s") $) }}`, ref.Repository)
					}
				case "ghcr.io":
					newimg = fmt.Sprintf(`{{ include "image.ghcr" (merge (dict "_repo" "%s") $) }}`, ref.Repository)
				case "registry.k8s.io":
					newimg = fmt.Sprintf(`{{ include "image.kubernetes" (merge (dict "_repo" "%s") $) }}`, ref.Repository)
				default:
					panic("unsupported registry for image " + img)
				}
				if found {
					newimg += ":" + tag
				}
				err = unstructured.SetNestedField(obj.UnstructuredContent(), newimg, "spec", "image")
				if err != nil {
					panic(err)
				}
			}

			args, _, err := unstructured.NestedStringSlice(obj.UnstructuredContent(), "spec", "args")
			if err != nil {
				panic(err)
			}

			for i := range args {
				if strings.HasPrefix(args[i], "--wait-timeout=") {
					args[i] = `--wait-timeout=${waitTimeout:={{ .Values.waitTimeout}}}`
				}

				switch app {
				case "elasticsearch":
					if strings.HasPrefix(args[i], "--es-args=") {
						args[i] = fmt.Sprintf(`--es-args=${args:={{ .Values.%s.args }}}`, app)
					}
				case "opensearch":
					if strings.HasPrefix(args[i], "--os-args=") {
						args[i] = fmt.Sprintf(`--os-args=${args:={{ .Values.%s.args }}}`, app)
					}
				case "mariadb":
					if strings.HasPrefix(args[i], "--mariadb-args=") {
						args[i] = fmt.Sprintf(`--mariadb-args=${args:={{ .Values.%s.args }}}`, app)
					}

				case "mongodb":
					if strings.HasPrefix(args[i], "--mongo-args=") {
						args[i] = fmt.Sprintf(`--mongo-args=${args:={{ .Values.%s.args }}}`, app)
					}
					// --max-concurrency=${MAX_CONCURRENCY:={{ .Values.maxConcurrency}}}
					if strings.HasPrefix(args[i], "--max-concurrency=") {
						args[i] = fmt.Sprintf(`--max-concurrency=${maxConcurrency:={{ .Values.%s.maxConcurrency}}}`, app)
					}

				case "mysql":
					if strings.HasPrefix(args[i], "--mysql-args=") {
						args[i] = fmt.Sprintf(`--mysql-args=${args:={{ .Values.%s.args }}}`, app)
					}

				case "perconaxtradb":
					if strings.HasPrefix(args[i], "--xtradb-args=") {
						args[i] = fmt.Sprintf(`--xtradb-args=${args:={{ .Values.%s.args }}}`, app)
					}
					//
					if strings.HasPrefix(args[i], "--target-app-replicas=") {
						args[i] = fmt.Sprintf(`--target-app-replicas=${TARGET_APP_REPLICAS:={{ .Values.%s.restore.targetAppReplicas }}}`, app)
					}

				case "postgres":
					if strings.HasPrefix(args[i], "--pg-args=") {
						args[i] = fmt.Sprintf(`--pg-args=${args:={{ .Values.%s.args }}}`, app)
					}
				case "redis":
					if strings.HasPrefix(args[i], "--redis-args=") {
						args[i] = fmt.Sprintf(`--redis-args=${args:={{ .Values.%s.args }}}`, app)
					}
				case "rabbitmq":
					if strings.HasPrefix(args[i], "--rabbitmq-args=") {
						args[i] = fmt.Sprintf(`--rabbitmq-args=${args:={{ .Values.%s.args }}}`, app)
					}
				}
			}

			err = unstructured.SetNestedStringSlice(obj.UnstructuredContent(), args, "spec", "args")
			if err != nil {
				panic(err)
			}
		}

		appDir := filepath.Join(dir, "charts", "kubedb-kubestash-catalog", "templates", app)
		err = os.MkdirAll(appDir, 0o755)
		if err != nil {
			panic(err)
		}

		tplText := templates.App
		data := map[string]interface{}{
			"app":    app,
			"object": obj.UnstructuredContent(),
		}
		kind, isDB := appToKind[app]
		if isDB {
			data["kind"] = kind
			tplText = templates.DB
		}

		funcMap := sprig.TxtFuncMap()
		funcMap["toYaml"] = toYAML
		funcMap["toJson"] = toJSON
		tpl := template.Must(template.New("").Funcs(funcMap).Parse(tplText))
		buf.Reset()
		err = tpl.Execute(&buf, &data)
		if err != nil {
			panic(err)
		}

		filename := filepath.Join(appDir, obj.GetName()+".yaml")
		err = os.WriteFile(filename, buf.Bytes(), 0o644)
		if err != nil {
			panic(err)
		}
	}

	{
		// Verify
		type ObjectKey struct {
			APIVersion string
			Kind       string
			Name       string
		}
		type DiffData struct {
			A *unstructured.Unstructured
			B *unstructured.Unstructured
		}

		dm := map[ObjectKey]*DiffData{}
		for _, ri := range resources {
			dm[ObjectKey{
				APIVersion: ri.Object.GetAPIVersion(),
				Kind:       ri.Object.GetKind(),
				Name:       ri.Object.GetName(),
			}] = &DiffData{
				A: ri.Object,
			}
		}

		failed := false
		differ := diff.New()

		sh := shell.NewSession()
		sh.SetDir(dir)
		sh.ShowCMD = true

		out, err := sh.Command("helm", "template", "charts/kubedb-kubestash-catalog").Output()
		if err != nil {
			panic(err)
		}

		helmout, err := parser.ListResources(out)
		if err != nil {
			panic(err)
		}

		for _, ri := range helmout {
			// ri.Object.SetNamespace("")
			ri.Object.SetLabels(nil)
			ri.Object.SetAnnotations(nil)

			key := ObjectKey{
				APIVersion: ri.Object.GetAPIVersion(),
				Kind:       ri.Object.GetKind(),
				Name:       ri.Object.GetName(),
			}
			if _, ok := dm[key]; !ok {
				failed = true
				_, _ = fmt.Fprintf(os.Stderr, "missing object is raw apiVersion=%s kind=%s name=%s", key.APIVersion, key.Kind, key.Name)
			} else {
				dm[key].B = ri.Object
			}
		}

		for key, data := range dm {
			if data.B == nil {
				failed = true
				_, _ = fmt.Fprintf(os.Stderr, "missing object is catalog chart apiVersion=%s kind=%s name=%s", key.APIVersion, key.Kind, key.Name)
				continue
			}

			a, err := json.Marshal(data.A)
			if err != nil {
				panic(err)
			}
			b, err := json.Marshal(data.B)
			if err != nil {
				panic(err)
			}

			// Then, Check them
			d, err := differ.Compare(a, b)
			if err != nil {
				fmt.Printf("Failed to unmarshal file: %s\n", err.Error())
				os.Exit(3)
			}

			if d.Modified() {
				failed = true

				config := formatter.AsciiFormatterConfig{
					ShowArrayIndex: true,
					Coloring:       true,
				}

				f := formatter.NewAsciiFormatter(data.A.Object, config)
				result, err := f.Format(d)
				if err != nil {
					panic(err)
				}
				_, _ = fmt.Fprintf(os.Stderr, "mismatched apiVersion=%s kind=%s name=%s \ndiff=%s", key.APIVersion, key.Kind, key.Name, result)
				continue
			}
		}

		if failed {
			os.Exit(1)
		}
	}
}

// toYAML takes an interface, marshals it to yaml, and returns a string. It will
// always return a string, even on marshal error (empty string).
//
// This is designed to be called from a template.
func toYAML(v interface{}) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return strings.TrimSuffix(string(data), "\n")
}

// toJSON takes an interface, marshals it to json, and returns a string. It will
// always return a string, even on marshal error (empty string).
//
// This is designed to be called from a template.
func toJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return string(data)
}

func Compare(i, j string) bool {
	vi, ei := semver.NewVersion(i)
	vj, ej := semver.NewVersion(j)
	if ei == nil && ej == nil {
		return semvers.CompareVersions(vi, vj)
	}

	idx_i := strings.IndexRune(i, '-')
	var distro_i, ver_i string
	if idx_i != -1 {
		distro_i, ver_i = i[:idx_i], i[idx_i+1:]
	}
	idx_j := strings.IndexRune(j, '-')
	var distro_j, ver_j string
	if idx_j != -1 {
		distro_j, ver_j = j[:idx_j], j[idx_j+1:]
	}
	if distro_i == distro_j && distro_i != "" {
		return semvers.Compare(ver_i, ver_j)
	}
	return strings.Compare(i, j) < 0
}
