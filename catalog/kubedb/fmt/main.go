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
	"sort"
	"strings"
	"text/template"

	"kubedb.dev/installer/catalog/kubedb/fmt/templates"

	"github.com/Masterminds/semver/v3"
	"github.com/Masterminds/sprig/v3"
	flag "github.com/spf13/pflag"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	shell "gomodules.xyz/go-sh"
	"gomodules.xyz/semvers"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"kmodules.xyz/client-go/tools/parser"
	"kmodules.xyz/go-containerregistry/name"
	"sigs.k8s.io/yaml"
	stash "stash.appscode.dev/installer/catalog"
)

const (
	distroOfficial = "Official"
)

type StashAddon struct {
	DBType    string
	DBVersion string
}

type DB struct {
	Group string
	Kind  string
}
type DbVersion struct {
	Group   string
	Kind    string
	Version string
	Distro  string
}

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

	flag.StringVar(&dir, "dir", dir, "Path to directory where the kubedb/installer project resides (default is set o current directory)")
	flag.StringVar(&apiKind, "kind", apiKind, "Kind of the CRD")
	flag.StringToStringVar(&specUpdates, "update-spec", specUpdates, "Key/Value map used to update pg-coordinator and replication mode detector image")

	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	resources, err := parser.ListPathResources(filepath.Join(dir, "catalog", "kubedb", "raw"))
	if err != nil {
		panic(err)
	}

	dbStore := map[DbVersion][]*unstructured.Unstructured{}
	pspForDBs := map[DB]sets.Set[string]{}
	pspStore := map[string]*unstructured.Unstructured{}

	// active versions
	activeDBVersions := map[string][]FullVersion{}
	// backupTask -> db version
	backupTaskStore := map[string][]string{}
	// recoveryTask -> db version
	restoreTaskStore := map[string][]string{}

	stashCatalog := map[StashAddon]string{}
	for _, addon := range stash.Load().Addons {
		for _, v := range addon.Versions {
			stashCatalog[StashAddon{
				DBType:    addon.Name,
				DBVersion: toVersion(v),
			}] = toVersion(v) // remove -vN suffix from backup/restore task
		}
	}

	for _, ri := range resources {
		// remove labels
		ri.Object.SetNamespace("")
		ri.Object.SetLabels(nil)
		ri.Object.SetAnnotations(nil)

		for jp, val := range specUpdates {
			if apiKind == "" || apiKind == ri.Object.GetKind() {
				if ref, ok, _ := unstructured.NestedString(ri.Object.Object, strings.Split(jp, ".")...); ok {
					err = unstructured.SetNestedField(ri.Object.Object, encodeTag(ref, val), strings.Split(jp, ".")...)
					if err != nil {
						panic(fmt.Sprintf("failed to set %s to %s in group=%s,kind=%s,name=%s", jp, val, ri.Object.GetAPIVersion(), ri.Object.GetKind(), ri.Object.GetName()))
					}
				}
			}
		}

		gv, err := schema.ParseGroupVersion(ri.Object.GetAPIVersion())
		if err != nil {
			panic(err)
		}
		if gv.Group == "catalog.kubedb.com" {
			dbKind := strings.TrimSuffix(ri.Object.GetKind(), "Version")
			deprecated, _, _ := unstructured.NestedBool(ri.Object.Object, "spec", "deprecated")

			distro, _, _ := unstructured.NestedString(ri.Object.Object, "spec", "distribution")
			if dbKind == "Elasticsearch" {
				authPlugin, _, _ := unstructured.NestedString(ri.Object.Object, "spec", "authPlugin")
				if distro == "" {
					distro = authPlugin
					if authPlugin == "X-Pack" {
						distro = "ElasticStack"
					}
					err = unstructured.SetNestedField(ri.Object.Object, distro, "spec", "distribution")
					if err != nil {
						panic(err)
					}
				}
			} else if dbKind == "Postgres" {
				if distro == "" {

					distro = distroOfficial
					if strings.Contains(strings.ToLower(ri.Object.GetName()), "timescale") {
						distro = "TimescaleDB"
					}
					err = unstructured.SetNestedField(ri.Object.Object, distro, "spec", "distribution")
					if err != nil {
						panic(err)
					}
				}
			} else if dbKind == "MySQL" {
				if distro == "" {
					distro = distroOfficial
					if strings.Contains(strings.ToLower(ri.Object.GetName()), "percona") {
						distro = "Percona"
					}
					if img, ok, _ := unstructured.NestedString(ri.Object.UnstructuredContent(), "spec", "db", "image"); ok {
						ref, err := name.ParseReference(img)
						if err != nil {
							panic(err)
						}
						if ref.Registry == "mysql" {
							distro = "MySQL"
						}
					}
					err = unstructured.SetNestedField(ri.Object.Object, distro, "spec", "distribution")
					if err != nil {
						panic(err)
					}
				}
			} else if dbKind == "MongoDB" {
				if distro == "" {
					distro = distroOfficial
					if strings.Contains(strings.ToLower(ri.Object.GetName()), "percona") {
						distro = "Percona"
					}
					err = unstructured.SetNestedField(ri.Object.Object, distro, "spec", "distribution")
					if err != nil {
						panic(err)
					}
				}
			} else if dbKind == "KafkaConnector" {
				connectorType, _, _ := unstructured.NestedString(ri.Object.Object, "spec", "type")
				if distro == "" {
					distro = connectorType
				}
			} else if dbKind == "SchemaRegistry" {
				distro, _, _ = unstructured.NestedString(ri.Object.Object, "spec", "distribution")
			}

			dbVersion, _, err := unstructured.NestedString(ri.Object.Object, "spec", "version")
			if err != nil {
				panic(err)
			}
			dbverKey := DbVersion{
				Group:   gv.Group,
				Kind:    ri.Object.GetKind(),
				Version: dbVersion,
				Distro:  distro,
			}
			dbStore[dbverKey] = append(dbStore[dbverKey], ri.Object)

			pspName, _, err := unstructured.NestedString(ri.Object.Object, "spec", "podSecurityPolicies", "databasePolicyName")
			if err != nil {
				panic(err)
			}
			if pspName != "" {
				dbKey := DB{
					Group: gv.Group,
					Kind:  ri.Object.GetKind(),
				}
				if _, ok := pspForDBs[dbKey]; !ok {
					pspForDBs[dbKey] = sets.New[string]()
				}
				pspForDBs[dbKey].Insert(pspName)
			}

			if !deprecated {
				activeDBVersions[dbKind] = append(activeDBVersions[dbKind], FullVersion{
					Version:     dbVersion,
					CatalogName: ri.Object.GetName(),
				})

				backupTask, _, _ := unstructured.NestedString(ri.Object.Object, "spec", "stash", "addon", "backupTask", "name")
				if backupTask != "" {
					// update based on stash catalog
					addonKey := StashAddon{
						DBType:    StashAddonDBType(dbKind),
						DBVersion: VersionForStashTask(backupTask),
					}
					addVer, ok := stashCatalog[addonKey]
					if !ok {
						panic(fmt.Sprintf("no backup addon found for %#v", addonKey))
					}
					backupTask = fmt.Sprintf("%s-backup-%s", addonKey.DBType, addVer)
					err = unstructured.SetNestedField(ri.Object.Object, backupTask, "spec", "stash", "addon", "backupTask", "name")
					if err != nil {
						panic(err)
					}
					backupTaskStore[backupTask] = append(backupTaskStore[backupTask], ri.Object.GetName())
				}
				restoreTask, _, _ := unstructured.NestedString(ri.Object.Object, "spec", "stash", "addon", "restoreTask", "name")
				if restoreTask != "" {
					// update based on stash catalog
					addonKey := StashAddon{
						DBType:    StashAddonDBType(dbKind),
						DBVersion: VersionForStashTask(restoreTask),
					}
					addVer, ok := stashCatalog[addonKey]
					if !ok {
						panic(fmt.Sprintf("no restore addon found for %#v", addonKey))
					}
					restoreTask = fmt.Sprintf("%s-restore-%s", addonKey.DBType, addVer)
					err = unstructured.SetNestedField(ri.Object.Object, restoreTask, "spec", "stash", "addon", "restoreTask", "name")
					if err != nil {
						panic(err)
					}
					restoreTaskStore[restoreTask] = append(restoreTaskStore[restoreTask], ri.Object.GetName())
				}
			}
		} else if gv.Group == "policy" {
			if _, ok := pspStore[ri.Object.GetName()]; ok {
				panic("duplicate PSP name " + ri.Object.GetName())
			}
			pspStore[ri.Object.GetName()] = ri.Object
		}
	}

	{
		activeVers := map[string][]string{}

		for k, v := range activeDBVersions {
			sort.Sort(sort.Reverse(FullverCollection(v)))
			activeDBVersions[k] = v

			for _, e := range v {
				activeVers[k] = append(activeVers[k], e.CatalogName)
			}
		}

		data, err := json.MarshalIndent(activeVers, "", "  ")
		if err != nil {
			panic(err)
		}

		filename := filepath.Join(dir, "catalog", "kubedb", "active_versions.json")
		err = os.MkdirAll(filepath.Dir(filename), 0o755)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(filename, data, 0o644)
		if err != nil {
			panic(err)
		}
	}
	{
		for k, v := range backupTaskStore {
			versions := semvers.SortVersions(v, func(vi, vj string) bool {
				return !Compare(vi, vj)
			})
			backupTaskStore[k] = versions
		}

		data, err := json.MarshalIndent(backupTaskStore, "", "  ")
		if err != nil {
			panic(err)
		}

		filename := filepath.Join(dir, "catalog", "kubedb", "backup_tasks.json")
		err = os.MkdirAll(filepath.Dir(filename), 0o755)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(filename, data, 0o644)
		if err != nil {
			panic(err)
		}
	}

	{
		for k, v := range restoreTaskStore {
			versions := semvers.SortVersions(v, func(vi, vj string) bool {
				return !Compare(vi, vj)
			})
			restoreTaskStore[k] = versions
		}

		data, err := json.MarshalIndent(restoreTaskStore, "", "  ")
		if err != nil {
			panic(err)
		}

		filename := filepath.Join(dir, "catalog", "kubedb", "restore_tasks.json")
		err = os.MkdirAll(filepath.Dir(filename), 0o755)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(filename, data, 0o644)
		if err != nil {
			panic(err)
		}
	}

	for k, v := range dbStore {
		sort.Slice(v, func(i, j int) bool {
			di, _, _ := unstructured.NestedBool(v[i].Object, "spec", "deprecated")
			dj, _, _ := unstructured.NestedBool(v[j].Object, "spec", "deprecated")

			if (di && dj) || (!di && !dj) {
				// company version
				return Compare(v[i].GetName(), v[j].GetName())
			}
			return dj // or !di
		})
		dbStore[k] = v

		var buf bytes.Buffer
		for i, obj := range v {
			if i > 0 {
				buf.WriteString("\n---\n")
			}
			data, err := yaml.Marshal(obj)
			if err != nil {
				panic(err)
			}
			buf.Write(data)
		}

		dbKind := strings.TrimSuffix(k.Kind, "Version")

		var filenameparts []string
		if allDeprecated(v) {
			filenameparts = append(filenameparts, "deprecated")
		}
		filenameparts = append(filenameparts, strings.ToLower(dbKind), k.Version)
		if k.Distro != "" {
			filenameparts = append(filenameparts, strings.ToLower(k.Distro))
		}
		filename := filepath.Join(dir, "catalog", "kubedb", "new_raw", strings.ToLower(dbKind), fmt.Sprintf("%s.yaml", strings.Join(filenameparts, "-")))
		err = os.MkdirAll(filepath.Dir(filename), 0o755)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(filename, buf.Bytes(), 0o644)
		if err != nil {
			panic(err)
		}
	}

	for k, v := range pspForDBs {
		if len(v) == 0 {
			continue
		}

		var buf bytes.Buffer
		for i, pspName := range sets.List[string](v) {
			if i > 0 {
				buf.WriteString("\n---\n")
			}
			data, err := yaml.Marshal(pspStore[pspName])
			if err != nil {
				panic(err)
			}
			buf.Write(data)
		}

		dbKind := strings.TrimSuffix(k.Kind, "Version")
		filename := filepath.Join(dir, "catalog", "kubedb", "new_raw", strings.ToLower(dbKind), fmt.Sprintf("%s-psp.yaml", strings.ToLower(dbKind)))
		err = os.MkdirAll(filepath.Dir(filename), 0o755)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(filename, buf.Bytes(), 0o644)
		if err != nil {
			panic(err)
		}
	}

	{
		// move new_raw to raw
		err = os.RemoveAll(filepath.Join(dir, "catalog", "kubedb", "raw"))
		if err != nil {
			panic(err)
		}
		err = os.Rename(filepath.Join(dir, "catalog", "kubedb", "new_raw"), filepath.Join(dir, "catalog", "kubedb", "raw"))
		if err != nil {
			panic(err)
		}
	}

	// GENERATE CHART
	{
		for k, v := range dbStore {
			dbKind := strings.TrimSuffix(k.Kind, "Version")

			copies := make([]map[string]any, 0, len(v))
			for _, obj := range v {
				objCopy := obj.DeepCopy()

				spec, _, err := unstructured.NestedMap(objCopy.Object, "spec")
				if err != nil {
					panic(err)
				}
				for prop := range spec {
					templatizeRegistry := func(fields ...string) {
						fieldList := append([]string{"spec", prop}, fields...)
						img, ok, _ := unstructured.NestedString(objCopy.Object, fieldList...)
						if img != "" && ok {
							decodedImg := decodeTag(img)
							ref, err := name.ParseReference(decodedImg)
							if err != nil {
								panic(err)
							}

							err = unstructured.SetNestedField(obj.Object, decodedImg, fieldList...)
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
							case "mcr.microsoft.com":
								newimg = fmt.Sprintf(`{{ include "image.microsoft" (merge (dict "_repo" "%s") $) }}`, ref.Repository)
							default:
								panic("unsupported registry for image " + img)
							}
							if ref.Tag != "" && ref.Tag != "latest" {
								newimg += ":" + ref.Tag
							}
							err = unstructured.SetNestedField(objCopy.Object, newimg, fieldList...)
							if err != nil {
								panic(err)
							}
						}
					}
					templatizeRegistry("image")
					templatizeRegistry("yqImage")
					templatizeRegistry("walg", "image")
				}
				copies = append(copies, objCopy.UnstructuredContent())

				//if i > 0 {
				//	buf.WriteString("\n---\n")
				//}

				//data := map[string]interface{}{
				//	"kind":   dbKind,
				//	"object": objCopy.Object,
				//}
				//funcMap := sprig.TxtFuncMap()
				//funcMap["toYaml"] = toYAML
				//funcMap["toJson"] = toJSON
				//tpl := template.Must(template.New("").Funcs(funcMap).Parse(templates.DBVersion))
				//err = tpl.Execute(&buf, &data)
				//if err != nil {
				//	panic(err)
				//}
			}

			tempKind := dbKind
			if dbKind == "KafkaConnector" || dbKind == "SchemaRegistry" {
				tempKind = "Kafka"
			}
			data := map[string]interface{}{
				"kind":    tempKind,
				"objects": copies,
			}
			funcMap := sprig.TxtFuncMap()
			funcMap["toYaml"] = toYAML
			funcMap["toJson"] = toJSON
			tpl := template.Must(template.New("").Funcs(funcMap).Parse(templates.DBVersion))
			var buf bytes.Buffer
			err = tpl.Execute(&buf, &data)
			if err != nil {
				panic(err)
			}

			var filenameparts []string
			if allDeprecated(v) {
				filenameparts = append(filenameparts, "deprecated")
			}
			filenameparts = append(filenameparts, strings.ToLower(dbKind), k.Version)
			if k.Distro != "" {
				filenameparts = append(filenameparts, strings.ToLower(k.Distro))
			}
			filename := filepath.Join(dir, "charts", "kubedb-catalog", "new_templates", strings.ToLower(dbKind), fmt.Sprintf("%s.yaml", strings.Join(filenameparts, "-")))
			err = os.MkdirAll(filepath.Dir(filename), 0o755)
			if err != nil {
				panic(err)
			}

			err = os.WriteFile(filename, buf.Bytes(), 0o644)
			if err != nil {
				panic(err)
			}
		}

		for k, v := range pspForDBs {
			if len(v) == 0 {
				continue
			}

			dbKind := strings.TrimSuffix(k.Kind, "Version")

			var buf bytes.Buffer
			for i, pspName := range sets.List[string](v) {
				if i > 0 {
					buf.WriteString("\n---\n")
				}

				if pspStore[pspName] == nil {
					panic("missing psp " + pspName + " for db " + dbKind)
				}

				content := pspStore[pspName].DeepCopy().UnstructuredContent()
				unstructured.RemoveNestedField(content, "spec", "allowPrivilegeEscalation")
				unstructured.RemoveNestedField(content, "spec", "privileged")
				data := map[string]interface{}{
					"kind":   dbKind,
					"object": content,
				}
				funcMap := sprig.TxtFuncMap()
				funcMap["toYaml"] = toYAML
				funcMap["toJson"] = toJSON
				tpl := template.Must(template.New("").Funcs(funcMap).Parse(templates.PSP))
				err = tpl.Execute(&buf, &data)
				if err != nil {
					panic(err)
				}
			}

			filename := filepath.Join(dir, "charts", "kubedb-catalog", "new_templates", strings.ToLower(dbKind), fmt.Sprintf("%s-psp.yaml", strings.ToLower(dbKind)))
			err = os.MkdirAll(filepath.Dir(filename), 0o755)
			if err != nil {
				panic(err)
			}

			err = os.WriteFile(filename, buf.Bytes(), 0o644)
			if err != nil {
				panic(err)
			}
		}

		{
			// move new_templates to templates
			err = os.Rename(filepath.Join(dir, "charts", "kubedb-catalog", "templates", "_helpers.tpl"), filepath.Join(dir, "charts", "kubedb-catalog", "new_templates", "_helpers.tpl"))
			if err != nil {
				panic(err)
			}
			err = os.Rename(filepath.Join(dir, "charts", "kubedb-catalog", "templates", "custom.yaml"), filepath.Join(dir, "charts", "kubedb-catalog", "new_templates", "custom.yaml"))
			if err != nil {
				panic(err)
			}
			err = os.RemoveAll(filepath.Join(dir, "charts", "kubedb-catalog", "templates"))
			if err != nil {
				panic(err)
			}
			err = os.Rename(filepath.Join(dir, "charts", "kubedb-catalog", "new_templates"), filepath.Join(dir, "charts", "kubedb-catalog", "templates"))
			if err != nil {
				panic(err)
			}
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

		out, err := sh.Command("helm", "template", "charts/kubedb-catalog", "--api-versions", "policy/v1beta1/PodSecurityPolicy", "--set", "skipDeprecated=false", "--set", "psp.enabled=true").Output()
		if err != nil {
			panic(err)
		}

		helmout, err := parser.ListResources(out)
		if err != nil {
			panic(err)
		}

		for _, ri := range helmout {
			ri.Object.SetNamespace("")
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

func allDeprecated(objs []*unstructured.Unstructured) bool {
	for _, obj := range objs {
		d, _, _ := unstructured.NestedBool(obj.Object, "spec", "deprecated")
		if !d {
			return false
		}
	}
	return true
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

func StashAddonDBType(dbKind string) string {
	switch dbKind {
	case "PerconaXtraDB":
		return "perconaxtradb"
	default:
		return strings.ToLower(dbKind)
	}
}

func VersionForStashTask(taskName string) string {
	var v string
	if idx := strings.Index(taskName, "-backup-"); idx > -1 {
		v = taskName[idx:]
		v = strings.TrimPrefix(v, "-backup-")
	} else if idx := strings.Index(taskName, "-restore-"); idx > -1 {
		v = taskName[idx:]
		v = strings.TrimPrefix(v, "-restore-")
	}
	return toVersion(v)
}

func toVersion(v string) string {
	idx := strings.IndexRune(v, '-')
	if idx == -1 {
		return v
	}
	v2 := v[:idx]

	switch v2 {
	case "10.14.0":
		return "10.14"
	case "10.2.0":
		return "10.2"
	case "11.1.0":
		return "11.1"
	case "11.2.0":
		return "11.2"
	case "11.9.0":
		return "11.9"
	case "12.4.0":
		return "12.4"
	case "13.1.0":
		return "13.1"
	default:
		return v2
	}
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

func encodeTag(ref, tag string) string {
	if idx := strings.IndexRune(ref, '('); idx == -1 {
		return tag
	}

	var sb strings.Builder
	replace := false
	for _, ch := range ref {
		if ch == '(' {
			replace = true
			sb.WriteRune('(')
			sb.WriteString(tag)
			sb.WriteRune(')')
		} else if ch == ')' {
			replace = false
		} else if !replace {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func decodeTag(ref string) string {
	r := strings.NewReplacer("(", "", ")", "")
	return r.Replace(ref)
}
