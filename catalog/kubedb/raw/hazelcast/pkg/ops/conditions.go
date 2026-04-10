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

	"kubedb.dev/apimachinery/apis/kubedb"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"
	"kubedb.dev/apimachinery/pkg/eventer"

	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	cu "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (c *hzOpsReqController) updateHazelcastOpsRequestPhase(typ, message string, phase opsapi.OpsRequestPhase) error {
	_, err := cu.PatchStatus(context.TODO(), c.KBClient, c.req, func(obj client.Object) client.Object {
		ret := obj.(*opsapi.HazelcastOpsRequest)
		ret.Status.Phase = phase
		ret.Status.ObservedGeneration = c.req.Generation
		ret.Status.Conditions = cutil.SetCondition(ret.Status.Conditions, cutil.NewCondition(typ, message, c.req.Generation))
		return ret
	})
	return err
}

func (c *hzOpsReqController) pushFailureEventForHazelcastOpsReq(typ, message string) {
	c.Recorder.Eventf(
		c.req,
		core.EventTypeWarning,
		eventer.EventReasonFailedToStart,
		`Fail to be ready HazelcastOpsRequest: "%v". Reason: %v`,
		c.req.Name,
		message,
	)

	_, err := cu.PatchStatus(context.TODO(), c.KBClient, c.req, func(obj client.Object) client.Object {
		ret := obj.(*opsapi.HazelcastOpsRequest)
		ret.Status.Phase = opsapi.OpsRequestPhaseFailed
		ret.Status.Conditions = cutil.SetCondition(ret.Status.Conditions, kmapi.Condition{
			Type:               kmapi.ConditionType(typ),
			Reason:             opsapi.Failed,
			Status:             meta.ConditionTrue,
			ObservedGeneration: c.req.Generation,
			Message:            message,
		})
		return ret
	})
	if err != nil {
		c.setFailureEvent(err.Error())
		return
	}
}

func (c *hzOpsReqController) UpdateHazelcastOpsReqConditions(reason, message string, conditionStatus ...bool) error {
	_, err := cu.PatchStatus(context.TODO(), c.KBClient, c.req, func(obj client.Object) client.Object {
		ret := obj.(*opsapi.HazelcastOpsRequest)
		ret.Status.ObservedGeneration = c.req.Generation
		ret.Status.Conditions = cutil.SetCondition(ret.Status.Conditions, cutil.NewCondition(reason, message, c.req.Generation, conditionStatus...))
		return ret
	})

	return err
}

func (c *hzOpsReqController) UpdateHazelcastPhaseProgressingAndPauseReconcile(typ, msg string) (time.Duration, error) {
	if c.req.Status.Phase == opsapi.OpsRequestPhasePending {
		rq, err := c.UpdateOpsPhaseProgressing(c.req, opsapi.Running, msg)
		if err != nil {
			c.log.Error(err, "failed to update ops phase to progressing")
			return rq, err
		}

		if rq != 0 {
			return rq, nil
		}
	}

	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.Successful) {
		err := c.pauseHazelcast()
		if err != nil {
			c.log.Error(err, "failed to pause Hazelcast")
			return DefaultDuration, err
		}
		if !cutil.IsConditionTrue(c.db.Status.Conditions, kubedb.DatabasePaused) {
			c.log.Info("waiting for the pause request to be approved by Provisioner")
			return RequeueDuration, nil
		}
	}
	return DefaultDuration, nil
}

func (c *hzOpsReqController) setFailureEvent(message string) {
	c.Recorder.Eventf(
		c.req,
		core.EventTypeWarning,
		eventer.EventReasonFailedToUpdate,
		message,
	)
}

func (c *hzOpsReqController) setStartingEvent(message string) {
	c.Recorder.Eventf(
		c.req,
		core.EventTypeNormal,
		eventer.EventReasonStarting,
		message,
	)
}

func (c *hzOpsReqController) setSuccessfulEvent(message string) {
	c.Recorder.Eventf(
		c.req,
		core.EventTypeNormal,
		eventer.EventReasonSuccessful,
		message,
	)
}
