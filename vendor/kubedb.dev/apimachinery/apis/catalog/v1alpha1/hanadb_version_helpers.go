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

package v1alpha1

import (
	"fmt"

	"kubedb.dev/apimachinery/apis"
	"kubedb.dev/apimachinery/apis/catalog"
	"kubedb.dev/apimachinery/crds"

	"kmodules.xyz/client-go/apiextensions"
)

func (_ HanaDBVersion) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourcePluralHanaDBVersion))
}

var _ apis.ResourceInfo = &HanaDBVersion{}

func (m HanaDBVersion) ResourceFQN() string {
	return fmt.Sprintf("%s.%s", ResourcePluralHanaDBVersion, catalog.GroupName)
}

func (m HanaDBVersion) ResourceShortCode() string {
	return ResourceSingularHanaDBVersion
}

func (hdb HanaDBVersion) ResourceKind() string {
	return ResourceSingularHanaDBVersion
}

func (hdb HanaDBVersion) ResourceSingular() string {
	return ResourceSingularHanaDBVersion
}

func (hdb HanaDBVersion) ResourcePlural() string {
	return ResourceSingularHanaDBVersion
}

func (q HanaDBVersion) ValidateSpecs() error {
	if q.Spec.Version == "" ||
		q.Spec.DB.Image == "" {
		return fmt.Errorf(`atleast one of the following specs is not set for HanaDBVersion "%v":
spec.version,
spec.db.image,`, q.Name)
	}
	return nil
}
