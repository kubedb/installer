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
	"time"

	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"

	"k8s.io/apimachinery/pkg/types"
	cu "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (c *hzOpsReqController) UpdateVersion() (time.Duration, error) {
	// Updating Hazelcast ops-request phase to progressing and pause the Hazelcast operator reconcile
	rq, err := c.UpdateHazelcastPhaseProgressingAndPauseReconcile(string(opsapi.HazelcastOpsRequestTypeUpdateVersion),
		"Hazelcast ops-request has started to update version")

	if (err != nil) || (rq == RequeueDuration) {
		return rq, err
	}

	dbCopy := c.db.DeepCopy()
	// Change current version with target version
	dbCopy.Spec.Version = c.req.Spec.UpdateVersion.TargetVersion
	// Get Hazelcast target version instance, if not found return err
	hzTargetVersion := &catalog.HazelcastVersion{}
	err = c.KBClient.Get(context.TODO(), types.NamespacedName{
		Name: dbCopy.Spec.Version,
	}, hzTargetVersion)
	if err != nil {
		c.log.Error(err, "failed to get hazelcast target version")
		return DefaultDuration, err
	}

	// Get HazelcastReconcile
	hzReconcile, err := c.NewHazelcastReconcile(dbCopy)
	if err != nil {
		c.log.Error(err, "failed to get hazelcast reconcile")
		return DefaultDuration, err
	}

	// Reconcile Hazelcast resources
	// Update StatefulSet with new container images from the targeted version CRD.
	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.UpdateStatefulSets) {
		hzReconcile.WithVersion(hzTargetVersion)
		c.RunParallel(opsapi.UpdateStatefulSets, "successfully reconciled the Hazelcast with updated version",
			c.NewReconcileFunc(dbCopy, hzReconcile))
		// return from here, process rest in the next cycles
		return DefaultDuration, nil
	}

	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.RestartPods) {
		c.RunParallel(opsapi.RestartPods,
			"Successfully Restarted Hazelcast nodes",
			c.newRestartFunc(c.getPodsName(), c.db))
		// Return nil,
		// otherwise in the next step the opsapi request will be successful even the process
		// is running in the background.
		return DefaultDuration, nil
	}

	_, err = cu.CreateOrPatch(context.TODO(), c.KBClient, c.db, func(obj client.Object, createOp bool) client.Object {
		ret := obj.(*dbapi.Hazelcast)
		ret.Spec.Version = dbCopy.Spec.Version
		return ret
	})
	if err != nil {
		c.log.Error(err, "failed to patch Hazelcast")
		return DefaultDuration, err
	}

	// resume and Change the opsapi request phase to "Successful".
	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.Successful) {
		if err := c.resumeHazelcast(); err != nil {
			return DefaultDuration, err
		}
		if err := c.updateHazelcastOpsRequestPhase(opsapi.Successful, "Successfully updated hazelcast version", opsapi.OpsRequestPhaseSuccessful); err != nil {
			c.log.Error(err, "failed to update status")
			return DefaultDuration, err
		}
	}

	return DefaultDuration, nil
}
