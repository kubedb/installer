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

package semvers

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
)

// SemverCollection is a collection of Version instances and implements the sort
// interface. See the sort package for more details.
// https://golang.org/pkg/sort/
type SemverCollection []*semver.Version

// Len returns the length of a collection. The number of Version instances
// on the slice.
func (c SemverCollection) Len() int {
	return len(c)
}

// Less is needed for the sort interface to compare two Version objects on the
// slice. If checks if one is less than the other.
func (c SemverCollection) Less(i, j int) bool {
	return CompareVersions(c[i], c[j])
}

func CompareVersions(vi *semver.Version, vj *semver.Version) bool {
	mi, _ := vi.SetPrerelease("")
	mj, _ := vj.SetPrerelease("")

	if mi.Equal(&mj) &&
		(vi.Prerelease() == "" || strings.HasPrefix(vi.Prerelease(), "v")) &&
		(vj.Prerelease() == "" || strings.HasPrefix(vj.Prerelease(), "v")) &&
		!(vi.Prerelease() == "" && vj.Prerelease() == "") {

		si := -1
		sj := -1
		if strings.HasPrefix(vi.Prerelease(), "v") {
			si, _ = strconv.Atoi(vi.Prerelease()[1:])
		}
		if strings.HasPrefix(vj.Prerelease(), "v") {
			sj, _ = strconv.Atoi(vj.Prerelease()[1:])
		}
		return si < sj
	}
	return vi.LessThan(vj)
}

// Swap is needed for the sort interface to replace the Version objects
// at two different positions in the slice.
func (c SemverCollection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func SortVersions(versions []string) ([]string, error) {
	vs := make([]*semver.Version, len(versions))
	for i, v := range versions {
		v, err := semver.NewVersion(v)
		if err != nil {
			return nil, fmt.Errorf("error parsing version: %s", err)
		}
		vs[i] = v
	}
	sort.Sort(sort.Reverse(SemverCollection(vs)))

	result := make([]string, len(vs))
	for i, v := range vs {
		result[i] = v.Original()
	}
	return result, nil
}
