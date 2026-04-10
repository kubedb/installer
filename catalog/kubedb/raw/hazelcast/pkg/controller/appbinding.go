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

package controller

import (
	"fmt"

	"kubedb.dev/apimachinery/apis/kubedb"
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	kutil "kmodules.xyz/client-go"
	kmapi "kmodules.xyz/client-go/api/v1"
	coreutil "kmodules.xyz/client-go/core/v1"
	metautil "kmodules.xyz/client-go/meta"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	appcat_util "kmodules.xyz/custom-resources/client/clientset/versioned/typed/appcatalog/v1alpha1/util"
)

func (rs *ReconcileState) ensureAppbinding() error {
	appmeta := rs.db.AppBindingMeta()
	metadata := meta.ObjectMeta{
		Name:      appmeta.Name(),
		Namespace: rs.db.Namespace,
	}

	// caBundle to be added

	svcName := rs.db.ServiceName()

	_, v, err := appcat_util.CreateOrPatchAppBinding(rs.ctx,
		rs.AppCatalogClient.AppcatalogV1alpha1(),
		metadata, func(in *appcat.AppBinding) *appcat.AppBinding {
			coreutil.EnsureOwnerReference(&in.ObjectMeta, rs.db.Owner())
			in.Labels = rs.db.OffshootLabels()
			in.Annotations = metautil.FilterKeys(kubedb.GroupName, nil, rs.db.Annotations)

			in.Spec.Type = appmeta.Type()
			in.Spec.AppRef = &kmapi.TypedObjectReference{
				APIGroup:  kubedb.GroupName,
				Kind:      api.ResourceKindHazelcast,
				Namespace: rs.db.Namespace,
				Name:      rs.db.Name,
			}
			if !rs.db.Spec.DisableSecurity {
				in.Spec.Secret = &appcat.TypedLocalObjectReference{
					Name: rs.db.GetAuthSecretName(),
				}
			}
			in.Spec.Version = rs.db.Spec.Version
			in.Spec.ClientConfig.Service = &appcat.ServiceReference{
				Scheme: rs.db.GetConnectionScheme(),
				Name:   svcName,
				Port:   kubedb.HazelcastRestPort,
			}
			return in
		}, meta.PatchOptions{})

	if v == kutil.VerbCreated {
		rs.log.Info(fmt.Sprintf("Appbinding: %s/%s created", rs.db.Namespace, appmeta.Name()))
	}
	if err != nil {
		rs.log.Error(err, "Failed to createOrPatch AppBinding")
		return err
	}
	return nil
}
