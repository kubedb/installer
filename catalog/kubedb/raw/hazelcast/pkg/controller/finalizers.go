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
	"fmt"

	"kubedb.dev/apimachinery/apis"
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"

	"github.com/pkg/errors"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	clientutil "kmodules.xyz/client-go/client"
	coreutil "kmodules.xyz/client-go/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (rs *ReconcileState) ensureFinalizers(req ctrl.Request) (bool, error) {
	// if Hazelcast instance is marked for deletion, remove the finalizers &
	// abort reconcile. Make sure to stop healthChecker before deletion
	if rs.IsMarkedForDeletion() {
		key := req.String()
		rs.log.Info(fmt.Sprintf("stopping health check: %s", key))
		db := &api.Hazelcast{}
		err := rs.KBClient.Get(context.TODO(), req.NamespacedName, db)
		if err != nil {
			if kerr.IsNotFound(err) {
				rs.log.Info("Requested Hazelcast is already deleted")
				return true, nil
			}
			return false, err
		}
		if !coreutil.HasFinalizer(db.ObjectMeta, apis.Finalizer) {
			return true, nil
		}
		rs.HealthChecker.Stop(key)
		if rs.db.Spec.Monitor != nil {
			err := rs.deleteMonitor()
			if err != nil && !kerr.IsNotFound(err) {
				rs.log.Error(err, "Failed to delete monitoring resources")
				return false, errors.Wrap(err, "Failed to delete monitoring resources")
			}
		}
		err = rs.removeFinalizers()
		if err != nil {
			return true, errors.Wrap(err, "Failed to remove finalizers")
		}
		return true, nil
	}

	// ensure that the finalizers are added in Hazelcast crd
	// if object does not have finalizer, then lets add the finalizer
	// and update the object. This is equivalent registering finalizer
	if !coreutil.HasFinalizer(rs.db.ObjectMeta, apis.Finalizer) {
		_, err := clientutil.CreateOrPatch(rs.ctx, rs.KBClient, rs.db, func(obj client.Object, createOp bool) client.Object {
			db := obj.(*api.Hazelcast)
			db.ObjectMeta = coreutil.AddFinalizer(db.ObjectMeta, apis.Finalizer)
			return db
		})
		if err != nil {
			return false, errors.Wrap(err, "Failed to add finalizers")
		}
		rs.log.Info("Added Finalizers")
	}

	return false, nil
}

func (rs *ReconcileState) removeFinalizers() error {
	// if Hazelcast is on deletion, check if it has finalizers.
	// if yes, remove finalizers, dependent resources will be deleted because of ownerReference.
	if err := rs.terminate(); err != nil {
		rs.log.Error(err, "Error while removing secrets and pvcs")
		return err
	}

	db := &api.Hazelcast{}
	err := rs.KBClient.Get(rs.ctx, types.NamespacedName{
		Name:      rs.db.Name,
		Namespace: rs.db.Namespace,
	}, db)
	if err != nil {
		if kerr.IsNotFound(err) {
			return nil
		}
		rs.log.Error(err, "Failed to get hazelcast")
		return err
	}
	if coreutil.HasFinalizer(rs.db.ObjectMeta, apis.Finalizer) {
		_, err := clientutil.CreateOrPatch(rs.ctx, rs.KBClient, rs.db, func(obj client.Object, createOp bool) client.Object {
			db := obj.(*api.Hazelcast)
			db.ObjectMeta = coreutil.RemoveFinalizer(db.ObjectMeta, apis.Finalizer)
			return db
		})
		if err != nil {
			return errors.Wrap(err, "Failed to remove finalizers")
		}
		rs.log.Info("Removed Finalizers")
	}
	return nil
}
