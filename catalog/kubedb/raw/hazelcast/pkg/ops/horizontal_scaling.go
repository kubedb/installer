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
	"errors"
	"fmt"
	"time"

	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"

	apps "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cu "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// HorizontalScale Algo:
// 1. If controller node or combined node is scaled up
//   - Pause DB
//   - Update statefulSet with new node in the Hazelcast cluster
//   - Ensure new node joins the cluster
//   - Update Hazelcast CR
//   - Resume DB
//
// 2. Else if controller node or combined node is scaled down
//   - Pause DB
//   - Update statefulSet with scaling down node in the Hazelcast cluster
//   - Ensure the node is removed from the cluster
//   - Reconcile Hazelcast to get the updated configuration
//   - Restart All nodes to update the configuration
//   - Update Hazelcast CR
//   - Resume DB
//
// 3. Else
//   - Pause DB
//   - Update statefulSet with new node or scaling down node in the Hazelcast cluster
//   - Update Hazelcast CR
//   - Resume DB
func (c *hzOpsReqController) HorizontalScale() (time.Duration, error) {
	// Updating Hazelcast ops-request phase to progressing and pause the Hazelcast operator reconcile
	rq, err := c.UpdateHazelcastPhaseProgressingAndPauseReconcile(string(opsapi.HazelcastOpsRequestTypeHorizontalScaling),
		"Hazelcast ops-request has started to horizontally scaling the nodes")

	if (err != nil) || (rq == RequeueDuration) {
		return rq, err
	}

	dbCopy := c.db.DeepCopy()
	targetSpec := c.req.Spec.HorizontalScaling

	sts, err := c.Client.AppsV1().StatefulSets(dbCopy.Namespace).Get(context.TODO(), dbCopy.StatefulSetName(), metav1.GetOptions{})
	if err != nil {
		return DefaultDuration, err
	}

	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.HorizontalScale) {
		err = c.HorizontalScaleStatefulSet(sts, dbCopy, targetSpec.Hazelcast, opsapi.HorizontalScale)
		// return from here
		// process rest in the next event cycle.
		return DefaultDuration, err
	}
	dbCopy.Spec.Replicas = targetSpec.Hazelcast

	_, err = cu.CreateOrPatch(context.TODO(), c.KBClient, c.db, func(obj client.Object, createOp bool) client.Object {
		ret := obj.(*api.Hazelcast)
		ret.Spec = dbCopy.Spec
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
		if err := c.updateHazelcastOpsRequestPhase(opsapi.Successful, "Successfully completed horizontally scale Hazelcast cluster", opsapi.OpsRequestPhaseSuccessful); err != nil {
			c.log.Error(err, "failed to update status")
			return DefaultDuration, err
		}
	}

	return DefaultDuration, nil
}

func (c *hzOpsReqController) HorizontalScaleStatefulSet(sts *apps.StatefulSet, db *api.Hazelcast, desireReplicas *int32, condition string) error {
	if sts == nil || desireReplicas == nil {
		return errors.New("statefulSet or desireReplicas cannot be empty")
	}
	// desired replicas > current replicas
	// Scale UP
	if *desireReplicas >= *sts.Spec.Replicas {
		c.log.Info(fmt.Sprintf("Performing scaleUp on %s...", sts.Name))
		c.RunParallel(condition, fmt.Sprintf("ScaleUp %s nodes", sts.Name),
			c.NewScaleUpFunc(sts, *desireReplicas))
	} else if *desireReplicas <= *sts.Spec.Replicas {
		// desired replicas < current replicas
		// Scale Down
		c.log.Info(fmt.Sprintf("Performing scaleDown on %s...", sts.Name))
		c.RunParallel(condition, fmt.Sprintf("ScaleDown %s nodes", sts.Name),
			c.NewScaleDownFunc(sts, db, *desireReplicas, condition))

	}

	return nil
}
