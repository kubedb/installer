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
	"context"

	"kubedb.dev/apimachinery/apis/kubedb"
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"

	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	dynamic_util "kmodules.xyz/client-go/dynamic"
	meta_util "kmodules.xyz/client-go/meta"
)

func (rs *ReconcileState) wipeOutDatabase(owner *metav1.OwnerReference) error {
	secretUsed, err := rs.secretsUsedByPeers()
	if err != nil {
		return errors.Wrap(err, "error in getting used secret list")
	}
	unusedSecrets := sets.New[string](rs.db.GetPersistentSecrets()...).Difference(secretUsed)

	for _, unusedSecret := range sets.List[string](unusedSecrets) {
		secret, err := rs.Client.CoreV1().Secrets(rs.db.Namespace).Get(context.TODO(), unusedSecret, metav1.GetOptions{})
		if kerr.IsNotFound(err) {
			unusedSecrets.Delete(secret.Name)
			continue
		}

		if err != nil {
			return errors.Wrap(err, "error in getting db secret")
		}
		genericKey, ok := secret.Labels[meta_util.ManagedByLabelKey]
		if !ok || genericKey != kubedb.GroupName {
			unusedSecrets.Delete(secret.Name)
		}
	}

	return dynamic_util.EnsureOwnerReferenceForItems(
		context.TODO(),
		rs.DynamicClient,
		core.SchemeGroupVersion.WithResource("secrets"),
		rs.db.Namespace,
		sets.List[string](unusedSecrets),
		owner)
}

func (rs *ReconcileState) secretsUsedByPeers() (sets.Set[string], error) {
	secretUsed := sets.New[string]()

	dbList := &api.HazelcastList{}
	err := rs.KBClient.List(rs.ctx, dbList)
	if err != nil {
		return nil, err
	}
	for _, hz := range dbList.Items {
		if hz.Name != rs.db.Name {
			secretUsed.Insert(hz.GetPersistentSecrets()...)
		}
	}
	return secretUsed, nil
}
