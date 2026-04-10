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

package license

import (
	"context"
	"fmt"
	"slices"
	"strings"

	catalogapi "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	configapi "kubedb.dev/apimachinery/apis/config/v1alpha1"

	"github.com/Masterminds/semver/v3"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func MeetsLicenseRestrictions(kc client.Client, lr configapi.LicenseRestrictions, dbGK schema.GroupKind, dbVersion string) (bool, string, error) {
	if len(lr) == 0 {
		return true, "", nil
	}
	restrictions, found := lr[dbGK.Kind]
	if !found {
		return false, "Kind is not found in the restrictions map", nil
	}

	var dbv unstructured.Unstructured
	dbv.SetGroupVersionKind(catalogapi.SchemeGroupVersion.WithKind(dbGK.Kind + "Version"))
	err := kc.Get(context.TODO(), client.ObjectKey{Name: dbVersion}, &dbv)
	if err != nil {
		return false, "", err
	}

	strVer, ok, err := unstructured.NestedString(dbv.UnstructuredContent(), "spec", "version")
	if err != nil || !ok {
		return false, "", err
	}
	v, err := semver.NewVersion(strVer)
	if err != nil {
		return false, "", err
	}

	var causes []string
	for _, restriction := range restrictions {
		c, err := semver.NewConstraint(restriction.VersionConstraint)
		if err != nil {
			return false, "", err
		}
		if !c.Check(v) {
			causes = append(causes, fmt.Sprintf("Doesn't satisfy the constraint: %v", restriction.VersionConstraint))
			continue
		}
		if len(restriction.Distributions) > 0 {
			strDistro, ok, err := unstructured.NestedString(dbv.UnstructuredContent(), "spec", "distribution")
			if err != nil || !ok {
				return false, "", err
			}
			if !contains(restriction.Distributions, strDistro) {
				causes = append(causes, fmt.Sprintf("%v distro is not present in the allowed distributions list: %v", strDistro, restriction.Distributions))
				continue
			}
		}
		return true, "", nil
	}
	return false, strings.Join(causes, ","), nil
}

func contains(list []string, str string) bool {
	return slices.Contains(list, str)
}
