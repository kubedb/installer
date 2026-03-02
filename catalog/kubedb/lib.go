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

package catalog

import (
	"embed"
	"encoding/json"
	"flag"
	iofs "io/fs"
	"os"

	"github.com/spf13/pflag"
)

//go:embed raw *.json
var raw embed.FS

var dirCatalog = ""

func AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&dirCatalog, "kubedb-catalog-dir", dirCatalog, "Path to kubedb-catalog directory")
}

func AddGoFlags(fs *flag.FlagSet) {
	fs.StringVar(&dirCatalog, "kubedb-catalog-dir", dirCatalog, "Path to kubedb-catalog directory")
}

func FS() iofs.FS {
	if dirCatalog == "" {
		return raw
	}
	return os.DirFS(dirCatalog)
}

func ActiveDBVersions() map[string][]string {
	activeVersions, err := iofs.ReadFile(FS(), "active_versions.json")
	if err != nil {
		panic(err)
	}

	out := map[string][]string{}
	err = json.Unmarshal(activeVersions, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func BackupTasks() map[string][]string {
	backupTasks, err := iofs.ReadFile(FS(), "backup_tasks.json")
	if err != nil {
		panic(err)
	}

	out := map[string][]string{}
	err = json.Unmarshal(backupTasks, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func RestoreTasks() map[string][]string {
	restoreTasks, err := iofs.ReadFile(FS(), "restore_tasks.json")
	if err != nil {
		panic(err)
	}

	out := map[string][]string{}
	err = json.Unmarshal(restoreTasks, &out)
	if err != nil {
		panic(err)
	}
	return out
}
