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
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/olekukonko/tablewriter"
	flag "github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"kmodules.xyz/client-go/tools/parser"
)

type DBVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DBVersionSpec `json:"spec"`
}

type DBVersionSpec struct {
	Constraints *Constraints `json:"updateConstraints"`
	Version     string       `json:"version"`
	Deprecated  bool         `json:"deprecated,omitempty"`
}

type Constraints struct {
	*UpdateConstraints
	*MySQLUpdateConstraints
}

func (c *Constraints) MarshalJSON() ([]byte, error) {
	if c.UpdateConstraints != nil {
		return json.Marshal(c.UpdateConstraints)
	} else {
		return json.Marshal(c.MySQLUpdateConstraints)
	}
}

func (c *Constraints) UnmarshalJSON(data []byte) error {
	var gc UpdateConstraints
	if err := json.Unmarshal(data, &gc); err == nil {
		*c = Constraints{UpdateConstraints: &gc}
		return nil
	}

	var mc MySQLUpdateConstraints
	if err := json.Unmarshal(data, &mc); err != nil {
		return err
	}
	*c = Constraints{MySQLUpdateConstraints: &mc}
	return nil
}

type UpdateConstraints struct {
	// List of all accepted versions for upgrade request.
	// An empty list indicates all versions are accepted except the denylist.
	Allowlist []string `json:"allowlist,omitempty"`
	// List of all rejected versions for upgrade request.
	// An empty list indicates no version is rejected.
	Denylist []string `json:"denylist,omitempty"`
}
type MySQLUpdateConstraints struct {
	// List of all accepted versions for upgrade request
	Allowlist MySQLVersionAllowlist `json:"allowlist,omitempty"`
	// List of all rejected versions for upgrade request
	Denylist MySQLVersionDenylist `json:"denylist,omitempty"`
}

type MySQLVersionAllowlist struct {
	// List of all accepted versions for upgrade request of a Standalone server. empty indicates all accepted
	Standalone []string `json:"standalone,omitempty"`
	// List of all accepted versions for upgrade request of a GroupReplication cluster. empty indicates all accepted
	GroupReplication []string `json:"groupReplication,omitempty"`
}

type MySQLVersionDenylist struct {
	// List of all rejected versions for upgrade request of a Standalone server
	Standalone []string `json:"standalone,omitempty"`
	// List of all rejected versions for upgrade request of a GroupReplication cluster
	GroupReplication []string `json:"groupReplication,omitempty"`
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var apiKind string
	skipDeprecated := true

	flag.StringVar(&dir, "dir", dir, "Path to directory where the kubedb/installer project resides (default is set o current directory)")
	flag.StringVar(&apiKind, "kind", apiKind, "Kind of the CRD")
	flag.BoolVar(&skipDeprecated, "skip-deprecated", skipDeprecated, "If true, ignore deprecated versions")

	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	vermap := map[string][]DBVersion{}
	kinds := sets.New[string]()

	err = parser.ProcessPath(filepath.Join(dir, "catalog", "kubedb", "raw"), func(ri parser.ResourceInfo) error {
		kind := ri.Object.GetObjectKind().GroupVersionKind().Kind
		if !strings.HasSuffix(kind, "Version") {
			return nil
		}
		kind = strings.TrimSuffix(kind, "Version")

		var dbv DBVersion
		if err := FromUnstructured(ri.Object.UnstructuredContent(), &dbv); err != nil {
			panic(kind + ": " + err.Error())
		}
		if dbv.Spec.Constraints != nil {
			kinds.Insert(kind)
		}

		if skipDeprecated && dbv.Spec.Deprecated {
			return nil
		}
		vermap[kind] = append(vermap[kind], dbv)
		return nil
	})
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	buf.WriteString("# DB Version Update Matrix\n\n")

	for _, kind := range sets.List(kinds) {
		versions := vermap[kind]
		if len(versions) == 0 {
			continue
		}
		sort.Slice(versions, func(i, j int) bool {
			return LessThan(versions[i].Name, versions[i].Spec.Version, versions[j].Name, versions[j].Spec.Version)
		})
		vermap[kind] = versions

		if kind == "MySQL" {
			buf.WriteString("## MySQL Standalone\n")
			vs := make([]string, len(versions)+1)
			updateData := make([][]string, len(versions))
			for i := range versions {
				vs[i+1] = versions[i].Name

				updateData[i] = make([]string, len(versions)+1)
				updateData[i][0] = versions[i].Name
				for j := range versions {
					var uc UpdateConstraints
					if versions[i].Spec.Constraints != nil && versions[i].Spec.Constraints.MySQLUpdateConstraints != nil {
						uc.Allowlist = versions[i].Spec.Constraints.MySQLUpdateConstraints.Allowlist.Standalone
						uc.Denylist = versions[i].Spec.Constraints.MySQLUpdateConstraints.Denylist.Standalone
					}
					if dec, err := canUpdate(versions[j].Spec.Version, uc); err != nil {
						panic(err)
					} else {
						updateData[i][j+1] = string(dec)
					}
				}
			}
			buf.WriteString(PrintTable(vs, updateData))
			buf.WriteString("\n")

			buf.WriteString("## MySQL GroupReplication\n")
			vs = make([]string, len(versions)+1)
			updateData = make([][]string, len(versions))
			for i := range versions {
				vs[i+1] = versions[i].Name

				updateData[i] = make([]string, len(versions)+1)
				updateData[i][0] = versions[i].Name
				for j := range versions {
					var uc UpdateConstraints
					if versions[i].Spec.Constraints != nil && versions[i].Spec.Constraints.MySQLUpdateConstraints != nil {
						uc.Allowlist = versions[i].Spec.Constraints.MySQLUpdateConstraints.Allowlist.GroupReplication
						uc.Denylist = versions[i].Spec.Constraints.MySQLUpdateConstraints.Denylist.GroupReplication
					}
					if dec, err := canUpdate(versions[j].Spec.Version, uc); err != nil {
						panic(err)
					} else {
						updateData[i][j+1] = string(dec)
					}
				}
			}
			buf.WriteString(PrintTable(vs, updateData))
			buf.WriteString("\n")
		} else {
			buf.WriteString("## " + kind + "\n")
			vs := make([]string, len(versions)+1)
			updateData := make([][]string, len(versions))
			for i := range versions {
				vs[i+1] = versions[i].Name

				updateData[i] = make([]string, len(versions)+1)
				updateData[i][0] = versions[i].Name
				for j := range versions {
					var uc UpdateConstraints
					if versions[i].Spec.Constraints != nil && versions[i].Spec.Constraints.UpdateConstraints != nil {
						uc = *versions[i].Spec.Constraints.UpdateConstraints
					}
					if dec, err := canUpdate(versions[j].Spec.Version, uc); err != nil {
						panic(err)
					} else {
						updateData[i][j+1] = string(dec)
					}
				}
			}

			buf.WriteString(PrintTable(vs, updateData))
			buf.WriteString("\n")
		}
	}

	err = os.WriteFile(filepath.Join(dir, "catalog", "VersionMatrix.md"), buf.Bytes(), 0o644)
	if err != nil {
		panic(err)
	}
}

type Decision string

const (
	DecisionYes     Decision = "✅"
	DecisionNo      Decision = "❌"
	DecisionUnknown Decision = "❓"
)

func canUpdate(src string, upc UpdateConstraints) (Decision, error) {
	v, err := parseVersion(src)
	if err != nil {
		return DecisionUnknown, err
	}
	for _, deny := range upc.Denylist {
		cc, err := semver.NewConstraint(deny)
		if err != nil {
			return DecisionUnknown, err
		}
		if cc.Check(v) {
			return DecisionNo, nil
		}
	}
	for _, deny := range upc.Allowlist {
		cc, err := semver.NewConstraint(deny)
		if err != nil {
			return DecisionUnknown, err
		}
		if cc.Check(v) {
			return DecisionYes, nil
		}
	}
	return DecisionUnknown, nil
}

func parseVersion(v string) (*semver.Version, error) {
	if strings.HasPrefix(v, "alma-") {
		v = strings.TrimPrefix(v, "alma-")
	} else if pre, _, ok := strings.Cut(v, "_"); ok {
		v = pre
	}
	return semver.NewVersion(v)
}

func LessThan(xName, xVer, yName, yVer string) bool {
	xv, xe := parseVersion(xVer)
	yv, ye := parseVersion(yVer)
	if xe == nil && ye == nil && !xv.Equal(yv) {
		return xv.LessThan(yv)
	}
	return strings.Compare(xName, yName) < 0
}

func PrintTable(versions []string, updateData [][]string) string {
	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	table.SetHeader(versions)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(updateData) // Add Bulk Data
	table.Render()
	return buf.String()
}

func FromUnstructured(u map[string]interface{}, obj any) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, obj)
}
