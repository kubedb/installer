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

	"kubedb.dev/apimachinery/apis/kubedb"
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	cu "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (c *hzOpsReqController) pauseHazelcast() error {
	if cutil.HasCondition(c.db.Status.Conditions, kubedb.DatabasePaused) {
		return nil
	}
	c.log.Info("Pausing Hazelcast")
	// starting event
	c.setStartingEvent(fmt.Sprintf("Pausing Hazelcast databse: %v/%v", c.req.Namespace, c.req.Spec.DatabaseRef.Name))
	_, err := cu.PatchStatus(context.TODO(), c.KBClient, c.db, func(obj client.Object) client.Object {
		ret := obj.(*dbapi.Hazelcast)
		ret.Spec.HealthChecker.DisableWriteCheck = true
		ret.Status.Conditions = cutil.SetCondition(ret.Status.Conditions, kmapi.Condition{
			Type:               kubedb.DatabasePaused,
			Status:             metav1.ConditionUnknown,
			ObservedGeneration: c.db.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             kubedb.DatabasePaused,
			Message:            fmt.Sprintf("%s %s is in process", opsapi.ResourceKindHazelcastOpsRequest, c.req.Name),
		})
		return ret
	})
	if err != nil {
		c.log.Error(err, "failed to pause Hazelcast")
		return err
	}
	// success event
	c.setSuccessfulEvent(fmt.Sprintf("Successfully paused Hazelcast database: %v/%v for HazelcastOpsRequest: %v", c.req.Namespace, c.req.Spec.DatabaseRef.Name, c.req.Name))
	c.log.Info("Hazelcast have been paused successfully")

	return nil
}

func (c *hzOpsReqController) resumeHazelcast() error {
	c.log.Info("Resuming Hazelcast...")
	c.setStartingEvent(fmt.Sprintf("Resuming Hazelcast database: %v/%v", c.req.Namespace, c.req.Spec.DatabaseRef.Name))
	_, err := cu.PatchStatus(context.TODO(), c.KBClient, c.db, func(obj client.Object) client.Object {
		ret := obj.(*dbapi.Hazelcast)
		ret.Status.Conditions = cutil.RemoveCondition(ret.Status.Conditions, kubedb.DatabasePaused)
		return ret
	})
	if err != nil {
		c.log.Error(err, "failed to resume Hazelcast")
		return err
	}
	c.setSuccessfulEvent(fmt.Sprintf("Successfully resumed Hazelcast database: %v/%v for HazelcastOpsRequest: %v", c.req.Namespace, c.req.Spec.DatabaseRef.Name, c.req.Name))
	c.log.Info("Hazelcast has been resumed successfully")
	return nil
}

func (c *hzOpsReqController) getPodsName() []string {
	podsName := make([]string, 0)
	for i := int32(0); i < *c.db.Spec.Replicas; i++ {
		podsName = append(podsName, fmt.Sprintf("%s-%d", c.db.StatefulSetName(), i))
	}
	return podsName
}

func IsRestartNeeded(reqConfiguration *opsapi.ReconfigurationSpec) bool {
	switch reqConfiguration.Restart {
	case opsapi.ReconfigureRestartTrue:
		return true
	case opsapi.ReconfigureRestartFalse:
		return false
	default:
		if reqConfiguration.ApplyConfig != nil ||
			reqConfiguration.RemoveCustomConfig ||
			(reqConfiguration.ConfigSecret != nil && reqConfiguration.ConfigSecret.Name != "") {
			return true
		}
	}

	return false
}
