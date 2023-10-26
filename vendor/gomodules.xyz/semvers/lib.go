/*
Copyright AppsCode Inc.

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

package semvers

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Masterminds/semver/v3"
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

// Swap is needed for the sort interface to replace the Version objects
// at two different positions in the slice.
func (c SemverCollection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
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

func Compare(i, j string) bool {
	vi, ei := semver.NewVersion(i)
	vj, ej := semver.NewVersion(j)
	if ei == nil && ej == nil {
		return CompareVersions(vi, vj)
	}
	return strings.Compare(i, j) < 0
}

func CompareDesc(i, j string) bool {
	return !Compare(i, j)
}

func SortVersions(versions []string, compare func(vi, vj string) bool) []string {
	sort.Slice(versions, func(i, j int) bool {
		return compare(versions[i], versions[j])
	})
	return versions
}

func VersionToRune(v *semver.Version) rune {
	if v.Prerelease() != "" {
		return []rune(v.Prerelease())[0]
	}
	return 'v' // Handle 9.6, 9.6-v1
}

func AtLeastAsImp(base *semver.Version, x *semver.Version) bool {
	return VersionToRune(x) >= VersionToRune(base)
}

func IsPrerelease(v string) bool {
	return semver.MustParse(v).Prerelease() != ""
}

func IsPublicRelease(v string) bool {
	prerelease := semver.MustParse(v).Prerelease()
	return prerelease == "" || strings.Contains(prerelease, "rc.")
}
