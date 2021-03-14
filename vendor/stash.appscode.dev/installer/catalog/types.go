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
	_ "embed"
	"encoding/json"
	"sort"

	"gomodules.xyz/semvers"
)

//go:embed catalog.json
var data []byte

type Addon struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

type StashCatalog struct {
	ChartRegistry    string  `json:"chart_registry"`
	ChartRegistryURL string  `json:"chart_registry_url"`
	Addons           []Addon `json:"addons"`
}

func (c *StashCatalog) Sort() {
	sort.Slice(c.Addons, func(i, j int) bool { return c.Addons[i].Name < c.Addons[j].Name })
	for i, project := range c.Addons {
		c.Addons[i].Versions = semvers.SortVersions(project.Versions, semvers.Compare)
	}
}

func Load() *StashCatalog {
	var out StashCatalog

	err := json.Unmarshal(data, &out)
	if err != nil {
		panic(err)
	}
	out.Sort()
	return &out
}
