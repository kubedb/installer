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

	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"

	ctrl "sigs.k8s.io/controller-runtime"
)

func (rs *ReconcileState) SetLoggerWithReq(req ctrl.Request) {
	rs.log = ctrl.Log.WithValues(api.ResourceSingularHazelcast, req.NamespacedName)
}

// reconcile returns an empty result with nil error to signal a successful reconcile
// to the controller manager
func (r *HazelcastReconciler) reconciled() (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

// requeWithError is a wrapper around logging an error mesage
// then passes the error through to the controller manager.
func (r *HazelcastReconciler) requeueWithError(err error) (ctrl.Result, error) {
	// Info log the error message and then let the reconciler dump the stacktrace
	return ctrl.Result{}, err
}

func (rs *ReconcileState) IsMarkedForDeletion() bool {
	return !rs.db.GetDeletionTimestamp().IsZero()
}

func (rs *ReconcileState) WithContext(ctx context.Context) {
	rs.ctx = ctx
}

func (rs *ReconcileState) WithHazelcast(hz *api.Hazelcast) {
	rs.db = hz
}

func (rs *ReconcileState) WithVersion(version *catalog.HazelcastVersion) {
	rs.version = version
}
