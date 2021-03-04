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
)

//go:embed raw
var raw embed.FS

//go:embed active_versions.json
var activeVersions []byte

//go:embed backup_tasks.json
var backupTasks []byte

//go:embed restore_tasks.json
var restoreTasks []byte

func FS() embed.FS {
	return raw
}

func ActiveDBVersions() map[string][]string {
	out := map[string][]string{}

	err := json.Unmarshal(activeVersions, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func BackupTasks() map[string][]string {
	out := map[string][]string{}

	err := json.Unmarshal(backupTasks, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func RestoreTasks() map[string][]string {
	out := map[string][]string{}

	err := json.Unmarshal(restoreTasks, &out)
	if err != nil {
		panic(err)
	}
	return out
}
