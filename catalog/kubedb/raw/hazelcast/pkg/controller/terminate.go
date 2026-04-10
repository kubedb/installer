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
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"

	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	dynamic_util "kmodules.xyz/client-go/dynamic"
)

func (rs *ReconcileState) terminate() error {
	// If TerminationPolicy is "halt", keep PVCs,Secrets intact.
	// TerminationPolicyPause is deprecated and will be removed in future.
	if rs.db.Spec.DeletionPolicy == dbapi.DeletionPolicyHalt {
		if err := rs.removeOwnerReferenceFromOffshoots(); err != nil {
			return err
		}
	} else {
		// If TerminationPolicy is "wipeOut", delete everything (ie, PVCs,Secrets,Snapshots).
		// If TerminationPolicy is "delete", delete PVCs and keep snapshots,secrets intact.
		// In both these cases, don't create dormantdatabase
		if err := rs.setOwnerReferenceToOffshoots(); err != nil {
			return err
		}
	}

	// monitor has to be implemented
	return nil
}

func (rs *ReconcileState) removeOwnerReferenceFromOffshoots() error {
	// First, Get LabelSelector for Other Components
	labelSelector := labels.SelectorFromSet(rs.db.OffshootSelectors())

	if err := dynamic_util.RemoveOwnerReferenceForSelector(
		rs.ctx,
		rs.DynamicClient,
		core.SchemeGroupVersion.WithResource("persistentvolumeclaims"),
		rs.db.Namespace,
		labelSelector,
		rs.db); err != nil {
		return err
	}
	if err := dynamic_util.RemoveOwnerReferenceForItems(
		rs.ctx,
		rs.DynamicClient,
		core.SchemeGroupVersion.WithResource("secrets"),
		rs.db.Namespace,
		rs.db.GetPersistentSecrets(),
		rs.db); err != nil {
		return err
	}
	return nil
}

func (rs *ReconcileState) setOwnerReferenceToOffshoots() error {
	selector := labels.SelectorFromSet(rs.db.OffshootSelectors())
	owner := metav1.NewControllerRef(rs.db, dbapi.SchemeGroupVersion.WithKind(dbapi.ResourceKindHazelcast))

	// If TerminationPolicy is "wipeOut", delete snapshots and secrets,
	// else, keep it intact.
	if rs.db.Spec.DeletionPolicy == dbapi.DeletionPolicyWipeOut {
		if err := rs.wipeOutDatabase(owner); err != nil {
			return errors.Wrap(err, "error in wiping out database.")
		}
	} else {
		// Make sure secret's ownerreference is removed.
		if err := dynamic_util.RemoveOwnerReferenceForItems(
			rs.ctx,
			rs.DynamicClient,
			core.SchemeGroupVersion.WithResource("secrets"),
			rs.db.Namespace,
			rs.db.GetPersistentSecrets(),
			rs.db); err != nil {
			return err
		}
	}
	// delete PVC for both "wipeOut" and "delete" TerminationPolicy.
	return dynamic_util.EnsureOwnerReferenceForSelector(
		rs.ctx,
		rs.DynamicClient,
		core.SchemeGroupVersion.WithResource("persistentvolumeclaims"),
		rs.db.Namespace,
		selector,
		owner)
}
