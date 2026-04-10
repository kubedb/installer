/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hazelcast

import (
	"context"
	"fmt"
	"log"
	"time"

	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"
	dbc "kubedb.dev/db-client-go/hazelcast"

	"github.com/pkg/errors"
	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	cu "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// VolumeExpansion Algorithm:
//   - Pause Hazelcast Reconcile
//   - dbCopy := Make a deep copy of Hazelcast object
//   - Update dbCopy with new req.volumeSpec
//   - Delete the StatefulSets based on req.volumeSpec; Don't need to delete StatefulSets which are unchanged.
//   - If Offline:
//   - For each pod:
//   - Store Pod's YAML to OpsRequest's annotations
//   - Delete the pod, and wait for it to be deleted!
//   - Patch PVC with the updated volume size and wait for it to become updated!
//   - Create the Pod back with the stored YAML
//   - Else If Online:
//   - List PVC
//   - Patch PVC with the updated volume size and wait for it to become updated!
//   - endif
//   - run ReconcileNodes(dbCopy)
//   - Wait for StatefulSets to become ready
//   - CreateOrPatchHazelcast(dbCopy)
//   - Resume Hazelcast
func (c *hzOpsReqController) VolumeExpansion() (time.Duration, error) {
	// Updating Hazelcast ops-request phase to progressing and pause the Hazelcast operator reconcile
	rq, err := c.UpdateHazelcastPhaseProgressingAndPauseReconcile(string(opsapi.HazelcastOpsRequestTypeVolumeExpansion),
		"Hazelcast ops-request has started to expand volume of hazelcast nodes.")

	if (err != nil) || (rq == RequeueDuration) {
		return rq, err
	}

	// Make a deep copy
	dbCopy := c.db.DeepCopy()
	volumeSpec := c.req.Spec.VolumeExpansion
	changedStsList := UpdateHazelcastVolumeSpec(dbCopy, volumeSpec)
	// Delete the StatefulSets whose storage have been changed, wait for them to be deleted
	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.OrphanStatefulSetPods) {
		c.RunParallel(opsapi.OrphanStatefulSetPods, "successfully deleted the statefulSets with orphan propagation policy",
			// only orphan those StatefulSet, Which are modified
			c.NewOrphanStatefulSetFunc(changedStsList))
		// return from here, process rest in the next cycles
		return DefaultDuration, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	hzClient, err := dbc.NewKubeDBClientBuilder(c.KBClient, c.db).WithContext(ctx).GetHazelcastRestyClient()
	if err != nil {
		klog.Error(err)
		log.Fatal(err)
	}

	// Volume expansion mode defaults to "Online"
	if volumeSpec.Mode == opsapi.VolumeExpansionModeOffline {
		err := c.VolumeExpansionOffline()
		if err != nil {
			c.log.Error(err, "failed to expand volume in offline mode")
			return DefaultDuration, err
		}
	} else {
		err := c.VolumeExpansionOnline()
		if err != nil {
			c.log.Error(err, "failed to expand volume online")
			return DefaultDuration, err
		}
	}

	rcl, err := c.NewHazelcastReconcile(dbCopy)
	if err != nil {
		c.log.Error(err, "failed to get hazelcast reconcile")
		return DefaultDuration, err
	}
	// Reconcile Hazelcast resources
	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.UpdateStatefulSets) {
		c.RunParallel(opsapi.UpdateStatefulSets, "successfully reconciled the Hazelcast resources",
			c.NewReconcileFunc(dbCopy, rcl))
		// return from here, process rest in the next cycles
		return DefaultDuration, nil
	}
	// Wait for the community operator to re-create the StatefulSet.
	// Here, the DB should not be in the paused state.
	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.ReadyStatefulSets) {
		c.RunParallel(opsapi.ReadyStatefulSets, "StatefulSet is recreated",
			c.NewReadyStatefulSetFunc(changedStsList))
		// return from here, process rest in the next cycles
		return DefaultDuration, nil
	}

	// resume and Change the opsapi request phase to "Successful".
	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.Successful) {
		if err := c.resumeHazelcast(); err != nil {
			return DefaultDuration, err
		}
		if err := c.updateHazelcastOpsRequestPhase(opsapi.Successful, "Successfully completed volumeExpansion for Hazelcast", opsapi.OpsRequestPhaseSuccessful); err != nil {
			c.log.Error(err, "failed to update status")
			return DefaultDuration, err
		}
	}

	// Patch Hazelcast with modified spec
	_, err = cu.CreateOrPatch(context.TODO(), c.KBClient, c.db, func(obj client.Object, createOp bool) client.Object {
		ret := obj.(*dbapi.Hazelcast)
		ret.Spec = dbCopy.Spec
		return ret
	})
	if err != nil {
		c.log.Error(err, "failed to patch the hazelcast object")
		return DefaultDuration, err
	}

	_, err = hzClient.GetClusterState()
	if err != nil {
		return DefaultDuration, err
	}

	_, err = hzClient.ChangeClusterState("ACTIVE")
	if err != nil {
		return DefaultDuration, err
	}

	return DefaultDuration, nil
}

func UpdateHazelcastVolumeSpec(db *dbapi.Hazelcast, volumeSpec *opsapi.HazelcastVolumeExpansionSpec) []string {
	var changedPS []string
	db.Spec.Storage.Resources.Requests[core.ResourceStorage] = *volumeSpec.Hazelcast
	changedPS = append(changedPS, db.StatefulSetName())
	return changedPS
}

type statefulSetFuncRetries struct {
	getStatefulSetRetries    *Retries
	deleteStatefulSetRetries *Retries
}

type statefulSetFunc struct {
	*hzOpsReqController
	statefulSetNames []string
	retries          statefulSetFuncRetries
	deceasedSet      string
}

func (c *hzOpsReqController) NewOrphanStatefulSetFunc(statefulSetNames []string) func() (bool, error) {
	opts := &statefulSetFunc{
		hzOpsReqController: c,
		statefulSetNames:   statefulSetNames,
		retries:            statefulSetFuncRetries{},
	}

	opts.retries.getStatefulSetRetries = c.newRetries("GetStatefulset")
	opts.retries.deleteStatefulSetRetries = c.newRetries("DeleteStatefulset")
	return opts.orphanStatefulSetPods
}

func (s *statefulSetFunc) orphanStatefulSetPods() (bool, error) {
	if s.statefulSetNames == nil {
		return false, nil
	}
	if s.deceasedSet != "" {
		err := s.KBClient.Get(context.TODO(), types.NamespacedName{
			Namespace: s.db.Namespace,
			Name:      s.deceasedSet,
		}, &apps.StatefulSet{})
		if err != nil && kerr.IsNotFound(err) {
			// reset deceasedSet
			s.deceasedSet = ""
			// remove the first element of the list
			if len(s.statefulSetNames) > 1 {
				s.statefulSetNames = s.statefulSetNames[1:]
			} else {
				s.statefulSetNames = nil
			}

			s.log.Info("StatefulSet is deleted successfully")
			s.retries.getStatefulSetRetries.Initialize()
			// retry
			return true, nil
		} else {
			s.log.Info(fmt.Sprintf("statefulset %s not yet deleted", s.deceasedSet))
			// retry
			return s.retries.getStatefulSetRetries.Wait(), nil
		}

	} else {
		// Get StatefulSet
		stsName := s.statefulSetNames[0]
		err := s.KBClient.Get(context.TODO(), types.NamespacedName{
			Namespace: s.db.Namespace,
			Name:      stsName,
		}, &apps.StatefulSet{})
		if err != nil {
			if kerr.IsNotFound(err) {
				// No need to retry
				return false, errors.Wrapf(err, "statefulSet %s doesn't exist", stsName)
			}
			// retry
			return s.retries.getStatefulSetRetries.Wait(), err
		}
		s.retries.getStatefulSetRetries.Initialize()

		deletePolicy := meta.DeletePropagationOrphan
		err = s.KBClient.Delete(context.TODO(), &apps.StatefulSet{
			ObjectMeta: meta.ObjectMeta{
				Name:      stsName,
				Namespace: s.db.Namespace,
			},
		}, &client.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		})
		if err != nil {
			return s.retries.deleteStatefulSetRetries.Wait(), err
		}
		s.retries.deleteStatefulSetRetries.Initialize()

		// set deceased StatefulSet name
		s.deceasedSet = stsName
	}

	s.log.Info("Process the next statefulSet of the list")
	// retry
	return true, nil
}

// NewReadyStatefulSetFunc returns a function that checks whether the given StatefulSets are isReady or not
func (c *hzOpsReqController) NewReadyStatefulSetFunc(statefulSetNames []string) func() (bool, error) {
	opts := &statefulSetFunc{
		hzOpsReqController: c,
		statefulSetNames:   statefulSetNames,
		retries:            statefulSetFuncRetries{},
	}

	opts.retries.getStatefulSetRetries = c.newRetries("GetStatefulSet")
	return opts.isReady
}

// function returns error with a boolean which specifies
// whether we should retry the process or fails the opsReq.
// If the function completes its given task, it return "false, nil"
// i.e. No need to retry, no error occurred.
func (s *statefulSetFunc) isReady() (bool, error) {
	if s.statefulSetNames == nil {
		return false, nil
	}
	for _, sts := range s.statefulSetNames {
		err := s.KBClient.Get(context.TODO(), types.NamespacedName{
			Namespace: s.db.Namespace,
			Name:      sts,
		}, &apps.StatefulSet{})
		if err != nil {
			// retry
			return s.retries.getStatefulSetRetries.Wait(), errors.Wrap(err, "StatefulSet yet to be ready")
		}
		s.retries.getStatefulSetRetries.Initialize()
	}

	// all StatefulSets are ready
	// return with no error
	return false, nil
}
