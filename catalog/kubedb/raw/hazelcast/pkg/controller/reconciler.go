/*
Copyright 2025.

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

package controller

import (
	"context"
	"fmt"

	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/config/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	amc "kubedb.dev/apimachinery/pkg/controller"
	"kubedb.dev/apimachinery/pkg/license"
	apiphase "kubedb.dev/apimachinery/pkg/phase"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	pcm "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	apps "k8s.io/api/apps/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	kmapi "kmodules.xyz/client-go/api/v1"
	cu "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	health "kmodules.xyz/client-go/tools/healthchecker"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// HazelcastReconciler reconciles a Hazelcast object
type HazelcastReconciler struct {
	*amc.Config
	*amc.Controller
	NetworkPolicyEnabled bool
	Scheme               *runtime.Scheme
	HealthChecker        *health.HealthChecker
	PromClient           pcm.MonitoringV1Interface
	LicenseRestrictions  v1alpha1.LicenseRestrictions
}

// ReconcileState Carries the requested DB and corresponding object objects
type ReconcileState struct {
	ctx     context.Context
	db      *api.Hazelcast
	log     logr.Logger
	version *catalog.HazelcastVersion
	*HazelcastReconciler
}

// +kubebuilder:rbac:groups=kubedb.com,resources=hazelcasts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kubedb.com,resources=hazelcasts/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=kubedb.com,resources=hazelcasts/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups="",resources=endpoints,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups="",resources=serviceaccounts,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups="apps",resources=statefulsets,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups="rbac.authorization.k8s.io",resources=roles,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups="rbac.authorization.k8s.io",resources=rolebindings,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups=catalog.kubedb.com,resources=hazelcastversions,verbs=get;list;watch
// +kubebuilder:rbac:groups=policy,resources=poddisruptionbudgets,verbs=create;get;list;watch;patch;delete
// +kubebuilder:rbac:groups="",resources=persistentvolumeclaims,verbs=get;list;watch;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Hazelcast object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.4/pkg/reconcile
func (r *HazelcastReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	klog.Info("Reconciler Started")

	db := &api.Hazelcast{}
	err := r.KBClient.Get(ctx, req.NamespacedName, db)
	if err != nil {
		if kerr.IsNotFound(err) {
			return r.reconciled()
		}

		return r.requeueWithError(err)
	}
	// check if the hazelcast object meets the license restrictions
	// if not, requeue with an error
	ok, reason, err := license.MeetsLicenseRestrictions(r.KBClient, r.LicenseRestrictions, api.Kind(api.ResourceKindHazelcast), db.Spec.Version)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "failed to check license restrictions")
	}
	if !ok {
		return ctrl.Result{}, fmt.Errorf("%s %s/%s of version %s does not meet license restrictions; Reason %v", api.ResourceKindHazelcast, db.Namespace, db.Name, db.Spec.Version, reason)
	}
	rs, err := r.GetReconcileState(ctx, req)
	if err != nil {
		if kerr.IsNotFound(err) {
			return rs.reconciled()
		}
		return rs.requeueWithError(err)
	}
	_, err = cu.CreateOrPatch(ctx, r.KBClient, db, func(obj client.Object, createOp bool) client.Object {
		in := obj.(*api.Hazelcast)
		in.SetDefaults(r.KBClient)
		return in
	})
	if err != nil {
		return rs.requeueWithError(err)
	}

	isFinalizersRemoved, err := rs.ensureFinalizers(req)
	if err != nil {
		if kerr.IsNotFound(err) {
			return rs.reconciled()
		}
		return rs.requeueWithError(err)
	}

	if isFinalizersRemoved {
		return rs.reconciled()
	}

	rs.runHealthChecker(req)

	rs.updatePhaseFromCondition()

	// abort reconcile if db paused
	if cutil.IsConditionTrue(rs.db.Status.Conditions, kubedb.DatabasePaused) {
		return rs.reconciled()
	}
	isPaused, err := rs.pauseReconcile()
	if err != nil {
		return rs.requeueWithError(err)
	}
	if isPaused {
		return rs.reconciled()
	}
	err = rs.EnsureSecrets()
	if err != nil {
		if kerr.IsNotFound(err) {
			return rs.reconciled()
		}
		return rs.requeueWithError(err)
	}
	err = rs.EnsureServices()
	if err != nil {
		if kerr.IsNotFound(err) {
			return rs.reconciled()
		}
		return rs.requeueWithError(err)
	}
	err = rs.ensureDatabaseRBAC()
	if err != nil {
		if kerr.IsNotFound(err) {
			return rs.reconciled()
		}
		return rs.requeueWithError(err)
	}
	err = rs.ensureAppbinding()
	if err != nil {
		if kerr.IsNotFound(err) {
			return rs.reconciled()
		}
		return rs.requeueWithError(err)
	}
	err = rs.EnsureStatefulSet()
	if err != nil {
		if kerr.IsNotFound(err) {
			return rs.reconciled()
		}
		return rs.requeueWithError(err)
	}
	err = rs.manageMonitor()
	if err != nil {
		return r.requeueWithError(err)
	}

	return ctrl.Result{}, nil
}

func (rs *ReconcileState) pauseReconcile() (bool, error) {
	// If the DB object is Paused, the operator will ignore the change events from
	// the DB object.
	found, pauseCond := cutil.GetCondition(rs.db.Status.Conditions, kubedb.DatabasePaused)
	if found >= 0 {
		if pauseCond.Status == meta.ConditionUnknown {
			_, err := cu.PatchStatus(rs.ctx, rs.KBClient, rs.db, func(obj client.Object) client.Object {
				hz := obj.(*api.Hazelcast)
				pauseCond.Status = meta.ConditionTrue
				hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, *pauseCond)
				return hz
			})
			if err != nil {
				rs.log.Error(err, "Failed to update database conditions")
				return false, err
			}
		}
		if pauseCond.Status == meta.ConditionTrue {
			rs.log.Info("Paused DB")
			return true, nil
		}
	}
	return false, nil
}

func (rs *ReconcileState) updatePhaseFromCondition() {
	if rs.db.Status.Phase == "" {
		if _, err := cu.PatchStatus(rs.ctx, rs.KBClient, rs.db, func(obj client.Object) client.Object {
			hz := obj.(*api.Hazelcast)
			hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
				Type:               kubedb.DatabaseProvisioningStarted,
				Status:             meta.ConditionTrue,
				ObservedGeneration: rs.db.Generation,
				Reason:             kubedb.DatabaseProvisioningStartedSuccessfully,
				Message:            fmt.Sprintf("The KubeDB operator has started the provisioning of Hazelcast: %s/%s", rs.db.Namespace, rs.db.Name),
			})
			return hz
		}); err != nil {
			rs.log.Error(err, "Failed to patch database status")
		}
	}

	// Check all replicas are ready or not
	err := rs.updateReplicaReadyCond()
	if err != nil {
		rs.log.Error(err, "Failed to update replica ready condition")
	}

	if cutil.IsConditionTrue(rs.db.Status.Conditions, kubedb.DatabaseReplicaReady) &&
		cutil.IsConditionTrue(rs.db.Status.Conditions, kubedb.DatabaseAcceptingConnection) &&
		cutil.IsConditionTrue(rs.db.Status.Conditions, kubedb.DatabaseReady) &&
		!cutil.IsConditionTrue(rs.db.Status.Conditions, kubedb.DatabaseProvisioned) {
		if _, err := cu.PatchStatus(rs.ctx, rs.KBClient, rs.db, func(obj client.Object) client.Object {
			hz := obj.(*api.Hazelcast)
			hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
				Type:               kubedb.DatabaseProvisioned,
				Status:             meta.ConditionTrue,
				ObservedGeneration: rs.db.Generation,
				Reason:             kubedb.DatabaseSuccessfullyProvisioned,
				Message:            fmt.Sprintf("The Hazelcast: %s/%s is successfully provisioned.", rs.db.Namespace, rs.db.Name),
			})
			return hz
		}); err != nil {
			rs.log.Error(err, "Failed to patch database status")
		}
	}

	if err := rs.updateHazelcastStatusPhase(); err != nil {
		rs.log.Error(err, "Failed to update database phase from conditions")
	}
}

func (rs *ReconcileState) updateHazelcastStatusPhase() error {
	var err error
	if _, err = cu.PatchStatus(rs.ctx, rs.KBClient, &api.Hazelcast{
		ObjectMeta: rs.db.ObjectMeta,
	}, func(obj client.Object) client.Object {
		hz := obj.(*api.Hazelcast)
		hz.Status.Phase = apiphase.PhaseFromCondition(rs.db.Status.Conditions)
		hz.Status.ObservedGeneration = rs.db.Status.ObservedGeneration
		return hz
	}); err != nil {
		rs.log.Error(err, "Failed to patch database status")
		return err
	}

	return nil
}

func (rs *ReconcileState) updateReplicaReadyCond() error {
	ps := &apps.StatefulSet{}
	err := rs.KBClient.Get(rs.ctx, types.NamespacedName{
		Name:      rs.db.StatefulSetName(),
		Namespace: rs.db.GetNamespace(),
	}, ps)
	if err != nil {
		if kerr.IsNotFound(err) {
			return nil
		}
		return err
	}

	replicaReady := false
	msg := ""

	if *rs.db.Spec.Replicas == ps.Status.ReadyReplicas {
		replicaReady = true
		msg = "All desired replicas are ready"
	} else {
		msg = fmt.Sprintf("All desired replicas are not ready. For statefulset: %s/%s desired replicas: %d, ready replicas: %d.", ps.Namespace, ps.Name, *rs.db.Spec.Replicas, ps.Status.ReadyReplicas)
	}

	hz := &api.Hazelcast{
		ObjectMeta: rs.db.ObjectMeta,
	}

	if replicaReady {
		if _, err := cu.PatchStatus(rs.ctx, rs.KBClient, hz, func(obj client.Object) client.Object {
			hz := obj.(*api.Hazelcast)
			hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
				Type:               kubedb.DatabaseReplicaReady,
				Status:             meta.ConditionTrue,
				ObservedGeneration: rs.db.Generation,
				Reason:             kubedb.AllReplicasAreReady,
				Message:            msg,
			})
			return obj
		}); err != nil {
			rs.log.Error(err, "Failed to patch database status")
		}
	} else {
		if _, err := cu.PatchStatus(rs.ctx, rs.KBClient, hz, func(obj client.Object) client.Object {
			hz := obj.(*api.Hazelcast)
			hz.Status.Conditions = cutil.SetCondition(hz.Status.Conditions, kmapi.Condition{
				Type:               kubedb.DatabaseReplicaReady,
				Status:             meta.ConditionFalse,
				ObservedGeneration: rs.db.Generation,
				Reason:             kubedb.SomeReplicasAreNotReady,
				Message:            msg,
			})
			return obj
		}); err != nil {
			rs.log.Error(err, "Failed to patch database status")
		}
	}

	return nil
}

func (r *HazelcastReconciler) GetReconcileState(ctx context.Context, req ctrl.Request) (*ReconcileState, error) {
	rs := &ReconcileState{}

	rs.SetLoggerWithReq(req)

	db := &api.Hazelcast{}
	err := r.KBClient.Get(ctx, req.NamespacedName, db)
	if err != nil {
		if kerr.IsNotFound(err) {
			rs.log.Info("Requested hazelcast not found")
			return nil, err
		}
		return nil, errors.Wrap(err, "Failed to get Requested Hazelcast")
	}

	// Get Hazelcast Version instance, if not found or failed to fetch
	// abort reconcile
	hazelcastVersion := &catalog.HazelcastVersion{}
	err = r.KBClient.Get(ctx, types.NamespacedName{
		Name: db.Spec.Version,
	}, hazelcastVersion)
	if err != nil {
		if kerr.IsNotFound(err) {
			rs.log.Info("Requested HazelcastVersion not found")
			return nil, err
		}
		return nil, errors.Wrap(err, "Failed to get Hazelcast Version")
	}

	rs.ctx = ctx
	rs.db = db
	rs.version = hazelcastVersion
	rs.HazelcastReconciler = r
	return rs, nil
}
