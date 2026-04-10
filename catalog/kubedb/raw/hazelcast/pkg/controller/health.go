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
	"errors"
	"fmt"
	"strings"

	"kubedb.dev/apimachinery/apis/kubedb"
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	dbclient "kubedb.dev/db-client-go/hazelcast"

	"github.com/go-logr/logr"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"
	kmapi "kmodules.xyz/client-go/api/v1"
	cu "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	health "kmodules.xyz/client-go/tools/healthchecker"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (rs *ReconcileState) runHealthChecker(req ctrl.Request) {
	key := req.String()
	rs.db.SetHealthCheckerDefaults()

	if cutil.IsConditionTrue(rs.db.Status.Conditions, kubedb.DatabaseProvisioned) || cutil.IsConditionTrue(rs.db.Status.Conditions, kubedb.DatabaseReplicaReady) {
		rs.HealthChecker.Start(key, rs.db.Spec.HealthChecker, rs.hazelcastHealthCheckFunc)
	}
}

func (rs *ReconcileState) hazelcastHealthCheckFunc(key string, hcs *health.HealthCard) {
	var err error
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return
	}

	log := ctrl.Log.WithValues(api.ResourceSingularHazelcast, types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	})

	db := &api.Hazelcast{}
	err = rs.KBClient.Get(context.TODO(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, db)
	if err != nil {
		log.Error(err, "Failed to get DB for health checking")
		return
	}

	db.SetHealthCheckerDefaults()
	hcs.SetThreshold(*db.Spec.HealthChecker.FailureThreshold)

	if cutil.IsConditionTrue(db.Status.Conditions, kubedb.DatabaseHealthCheckPaused) {
		log.Info("Skipping health check for Hazelcast %s/%s because health check is paused", db.Namespace, db.Name)
		return
	}

	ctx, cancel := context.WithCancel(context.TODO())
	defer func() {
		cancel()
	}()

	hzClient, err := dbclient.NewKubeDBClientBuilder(rs.KBClient, db).WithContext(ctx).WithLog(log).GetHazelcastClient()
	if err != nil {
		log.Error(err, "Failed to get Hazelcast Client")
		if hcs.HasFailed(health.HealthCheckClientFailure, err) {
			// Since the client was unable to connect the database,
			// update "AcceptingConnection" to "false".
			// update "Ready" to "false"
			rs.updateErrorAcceptingConnections(ctx, err, db, log)
		}
		return
	}

	hcs.ClientCreated()
	defer func() {
		err := hzClient.Shutdown(ctx)
		if err != nil {
			log.Error(err, "Failed to shutdown client")
		}
		hcs.ClientClosed()
	}()

	rs.updateDBAcceptingConnection(ctx, db, log)

	if !db.Spec.HealthChecker.DisableWriteCheck {
		err = rs.checkHazelcastReadWriteAccess(ctx, hzClient, db, log)
		if err != nil {
			log.Error(err, "Failed to read or write")
			// Since the client was unable to connect the database,
			// update "AcceptingConnection" to "false".
			// update "Ready" to "false"
			if err != nil {
				log.Error(err, "Failed to write in database")
			}
			return
		}
	}

	rs.updateDBReady(ctx, db, log)
}

func (rs *ReconcileState) updateDBReady(ctx context.Context, db *api.Hazelcast, log logr.Logger) {
	if _, err := cu.PatchStatus(ctx, rs.KBClient, &api.Hazelcast{
		ObjectMeta: db.ObjectMeta,
	}, func(obj client.Object) client.Object {
		hz := obj.(*api.Hazelcast)
		hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
			Type:               kubedb.DatabaseReady,
			Status:             meta.ConditionTrue,
			ObservedGeneration: db.Generation,
			Reason:             strings.Join([]string{kubedb.AllReplicasAreReady, kubedb.DatabaseAcceptingConnection, kubedb.ReadinessCheckSucceeded, kubedb.DatabaseWriteAccessCheckSucceeded}, ","),
			Message:            fmt.Sprintf("The Hazelcast: %s/%s is ready", db.Namespace, db.Name),
		})
		return obj
	}); err != nil {
		log.Error(err, "Failed to patch connection for DB ready")
	}
}

func (rs *ReconcileState) checkHazelcastReadWriteAccess(ctx context.Context, hzClient *dbclient.Client, db *api.Hazelcast, log logr.Logger) error {
	err := rs.checkHazelcastWriteAccess(ctx, hzClient, db, log)
	if err != nil {
		return err
	}

	err = rs.checkHazelcastReadAccess(ctx, hzClient, db, log)
	if err != nil {
		return err
	}

	return nil
}

func (rs *ReconcileState) checkHazelcastWriteAccess(ctx context.Context, hzClient *dbclient.Client, db *api.Hazelcast, log logr.Logger) error {
	myMap, err := hzClient.GetMap(ctx, "kubedb-system")
	if err != nil {
		rs.updateErrorWriteCheck(ctx, err, db, log)
		return err
	}

	err = myMap.Set(ctx, "key", "Hazelcast")
	if err != nil {
		rs.updateErrorWriteCheck(ctx, err, db, log)
		return err
	}

	rs.updateWriteCheck(ctx, db, log)

	return nil
}

func (rs *ReconcileState) checkHazelcastReadAccess(ctx context.Context, hzClient *dbclient.Client, db *api.Hazelcast, log logr.Logger) error {
	myMap, err := hzClient.GetMap(ctx, "kubedb-system")
	if err != nil {
		rs.updateErrorReadCheck(ctx, err, db, log)
		return err
	}

	val, err := myMap.Get(ctx, "key")
	if err != nil {
		rs.updateErrorReadCheck(ctx, err, db, log)
		return err
	}

	if val == nil {
		err := errors.New("read check failed because key is nil")
		rs.updateErrorReadCheck(ctx, err, db, log)
	}

	rs.updateReadCheck(ctx, db, log)

	return nil
}

func (rs *ReconcileState) updateWriteCheck(ctx context.Context, db *api.Hazelcast, log logr.Logger) {
	if _, err := cu.PatchStatus(ctx, rs.KBClient, &api.Hazelcast{
		ObjectMeta: db.ObjectMeta,
	}, func(obj client.Object) client.Object {
		hz := obj.(*api.Hazelcast)
		hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
			Type:               kubedb.DatabaseWriteAccess,
			Status:             meta.ConditionTrue,
			ObservedGeneration: db.Generation,
			Reason:             kubedb.DatabaseWriteAccessCheckSucceeded,
			Message:            fmt.Sprintf("The Hazelcast: %s/%s is accepting write request.", db.Namespace, db.Name),
		})
		return obj
	}); err != nil {
		log.Error(err, "Failed to patch database status for successful write check")
	}
}

func (rs *ReconcileState) updateReadCheck(ctx context.Context, db *api.Hazelcast, log logr.Logger) {
	if _, err := cu.PatchStatus(ctx, rs.KBClient, &api.Hazelcast{
		ObjectMeta: db.ObjectMeta,
	}, func(obj client.Object) client.Object {
		hz := obj.(*api.Hazelcast)
		hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
			Type:               kubedb.DatabaseReadAccess,
			Status:             meta.ConditionTrue,
			ObservedGeneration: db.Generation,
			Reason:             kubedb.DatabaseReadAccessCheckSucceeded,
			Message:            fmt.Sprintf("The Hazelcast: %s/%s is accepting read request.", db.Namespace, db.Name),
		})
		return obj
	}); err != nil {
		log.Error(err, "Failed to patch database status for successful read check")
	}
}

func (rs *ReconcileState) updateErrorWriteCheck(ctx context.Context, connectionErr error, db *api.Hazelcast, log logr.Logger) {
	if _, err := cu.PatchStatus(ctx, rs.KBClient, &api.Hazelcast{
		ObjectMeta: db.ObjectMeta,
	}, func(obj client.Object) client.Object {
		hz := obj.(*api.Hazelcast)
		hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
			Type:               kubedb.DatabaseWriteAccess,
			Status:             meta.ConditionFalse,
			ObservedGeneration: db.Generation,
			Reason:             kubedb.DatabaseWriteAccessCheckFailed,
			Message:            fmt.Sprintf("The Hazelcast: %s/%s is not accepting write request. error: %v", db.Namespace, db.Name, connectionErr),
		})
		return obj
	}); err != nil {
		log.V(5).Error(err, "Failed to patch database status for write check")
	}

	if _, err := cu.PatchStatus(ctx, rs.KBClient, &api.Hazelcast{
		ObjectMeta: db.ObjectMeta,
	}, func(obj client.Object) client.Object {
		hz := obj.(*api.Hazelcast)
		hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
			Type:               kubedb.DatabaseReady,
			Status:             meta.ConditionFalse,
			ObservedGeneration: db.Generation,
			Reason:             strings.Join([]string{kubedb.DatabaseReadAccessCheckFailed, kubedb.DatabaseWriteAccessCheckFailed}, "/"),
			Message:            fmt.Sprintf("The Hazelcast: %s/%s is not able to read or write.", db.Namespace, db.Name),
		})
		return obj
	}); err != nil {
		log.V(5).Error(err, "Failed to patch status for db ready if write check failed")
	}
}

func (rs *ReconcileState) updateErrorReadCheck(ctx context.Context, connectionErr error, db *api.Hazelcast, log logr.Logger) {
	if _, err := cu.PatchStatus(ctx, rs.KBClient, &api.Hazelcast{
		ObjectMeta: db.ObjectMeta,
	}, func(obj client.Object) client.Object {
		hz := obj.(*api.Hazelcast)
		hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
			Type:               kubedb.DatabaseReadAccess,
			Status:             meta.ConditionFalse,
			ObservedGeneration: db.Generation,
			Reason:             kubedb.DatabaseReadAccessCheckFailed,
			Message:            fmt.Sprintf("The Hazelcast: %s/%s is not accepting read request. error: %v", db.Namespace, db.Name, connectionErr),
		})
		return obj
	}); err != nil {
		log.V(5).Error(err, "Failed to patch database status for read check failed")
	}

	if _, err := cu.PatchStatus(ctx, rs.KBClient, &api.Hazelcast{
		ObjectMeta: db.ObjectMeta,
	}, func(obj client.Object) client.Object {
		hz := obj.(*api.Hazelcast)
		hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
			Type:               kubedb.DatabaseReady,
			Status:             meta.ConditionFalse,
			ObservedGeneration: db.Generation,
			Reason:             strings.Join([]string{kubedb.DatabaseReadAccessCheckFailed, kubedb.DatabaseWriteAccessCheckFailed}, "/"),
			Message:            fmt.Sprintf("The Hazelcast: %s/%s is not able to read or write.", db.Namespace, db.Name),
		})
		return obj
	}); err != nil {
		log.V(5).Error(err, "Failed to patch status for db ready if read check failed")
	}
}

func (rs *ReconcileState) updateDBAcceptingConnection(ctx context.Context, db *api.Hazelcast, log logr.Logger) {
	if _, err := cu.PatchStatus(ctx, rs.KBClient, &api.Hazelcast{
		ObjectMeta: db.ObjectMeta,
	}, func(obj client.Object) client.Object {
		hz := obj.(*api.Hazelcast)
		hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
			Type:               kubedb.DatabaseAcceptingConnection,
			Status:             meta.ConditionTrue,
			ObservedGeneration: db.Generation,
			Reason:             kubedb.DatabaseAcceptingConnectionRequest,
			Message:            fmt.Sprintf("The Hazelcast: %s/%s is accepting connection", db.Namespace, db.Name),
		})
		return obj
	}); err != nil {
		log.V(5).Error(err, "Failed to patch status if database connection true")
	}
}

func (rs *ReconcileState) updateErrorAcceptingConnections(ctx context.Context, connectionErr error, db *api.Hazelcast, log logr.Logger) {
	if _, err := cu.PatchStatus(ctx, rs.KBClient, &api.Hazelcast{
		ObjectMeta: db.ObjectMeta,
	}, func(obj client.Object) client.Object {
		hz := obj.(*api.Hazelcast)
		hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
			Type:               kubedb.DatabaseAcceptingConnection,
			Status:             meta.ConditionFalse,
			ObservedGeneration: db.Generation,
			Reason:             kubedb.DatabaseNotAcceptingConnectionRequest,
			Message:            fmt.Sprintf("The Hazelcast: %s/%s is not accepting client requests. Error: %v", db.Namespace, db.Name, connectionErr),
		})
		return obj
	}); err != nil {
		log.Error(err, "Failed to patch status if database is not accepting connection")
	}
	if _, err := cu.PatchStatus(ctx, rs.KBClient, &api.Hazelcast{
		ObjectMeta: db.ObjectMeta,
	}, func(obj client.Object) client.Object {
		hz := obj.(*api.Hazelcast)
		hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
			Type:               kubedb.DatabaseReady,
			Status:             meta.ConditionFalse,
			ObservedGeneration: db.Generation,
			Reason:             kubedb.DatabaseNotAcceptingConnectionRequest,
			Message:            fmt.Sprintf("The Hazelcast: %s/%s is not accepting connection.", db.Namespace, db.Name),
		})
		return obj
	}); err != nil {
		log.Error(err, "Failed to patch status for db ready if db not accepting connection")
	}
}
