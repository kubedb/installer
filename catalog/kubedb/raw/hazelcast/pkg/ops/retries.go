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
	"math"
	"strings"
	"unicode"

	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"
	"kubedb.dev/apimachinery/pkg/eventer"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"
	kmapi "kmodules.xyz/client-go/api/v1"
	clientutil "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Retries struct {
	retry      int
	maxRetries int
	opsReq     *opsapi.HazelcastOpsRequest
	recorder   record.EventRecorder
	kbClient   client.Client
	condType   string
	message    string
}

func (r *Retries) Wait(podName ...string) bool {
	if r == nil {
		return false
	}
	if r.retry == 0 {
		r.updateConditions(metav1.ConditionFalse, podName...)
	}
	r.retry++
	return r.retry < r.maxRetries
}

func (r *Retries) Initialize(podName ...string) {
	if r == nil {
		return
	}
	r.retry = 0
	r.updateConditions(metav1.ConditionTrue, podName...)
}

func (r *Retries) updateConditions(conditionStatus metav1.ConditionStatus, podName ...string) {
	condType := r.condType
	msg := fmt.Sprintf("%s; ConditionStatus:%s", r.message, conditionStatus)

	if len(podName) > 0 {
		condType += "--" + podName[0]
		msg += fmt.Sprintf("; PodName:%s", podName[0])
	}
	var err error
	r.recorder.Eventf(
		r.opsReq,
		core.EventTypeWarning,
		msg,
		msg,
	)

	_, err = clientutil.PatchStatus(context.TODO(), r.kbClient, r.opsReq, func(obj client.Object) client.Object {
		in := obj.(*opsapi.HazelcastOpsRequest)
		in.Status.Conditions = cutil.SetCondition(in.Status.Conditions, kmapi.Condition{
			Type:               kmapi.ConditionType(condType),
			Status:             conditionStatus,
			ObservedGeneration: r.opsReq.Generation,
			LastTransitionTime: metav1.Now(),
			Message:            msg,
		})
		in.Status.ObservedGeneration = r.opsReq.Generation
		return in
	})
	if err != nil {
		r.recorder.Eventf(
			r.opsReq,
			core.EventTypeWarning,
			eventer.EventReasonFailedToUpdate,
			err.Error(),
		)
		return
	}
}

func (c *hzOpsReqController) newRetries(condType string) *Retries {
	maxRetries := 500
	if c.req.Spec.Timeout != nil {
		maxRetries = int(math.Ceil(float64(c.req.Spec.Timeout.Seconds()) / 5.0))
	}

	convert := func(str string) string {
		var result strings.Builder
		for i, char := range str {
			if unicode.IsUpper(char) && i > 0 {
				result.WriteRune(' ')
			}
			result.WriteRune(unicode.ToLower(char))
		}
		return result.String()
	}

	return &Retries{
		retry:      0,
		maxRetries: maxRetries,
		opsReq:     c.req,
		recorder:   c.Recorder,
		kbClient:   c.KBClient,
		condType:   condType,
		message:    convert(condType),
	}
}

type retryableFunc struct {
	retries *Retries
}

func (c *hzOpsReqController) NewRetryableFunc(f func() error, condType string) func() (bool, error) {
	r := &retryableFunc{
		retries: c.newRetries(condType),
	}
	return func() (bool, error) {
		err := f()
		if err != nil {
			return r.retries.Wait(), err
		}
		return false, nil
	}
}
