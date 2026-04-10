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
	"time"

	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"

	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	cutil "kmodules.xyz/client-go/conditions"
	core_util "kmodules.xyz/client-go/core/v1"
	"kmodules.xyz/client-go/policy"
)

type restartRetries struct {
	evictPodRetries   *Retries
	getPodRetries     *Retries
	runningPodRetries *Retries
}

type smartRestarter struct {
	*hzOpsReqController
	infantNode string
	podNames   []string
	newDB      *dbapi.Hazelcast
	retries    restartRetries
}

func (c *hzOpsReqController) newRestartFunc(podNames []string, newDB *dbapi.Hazelcast) func() (bool, error) {
	opts := &smartRestarter{
		hzOpsReqController: c,
		podNames:           podNames,
		newDB:              newDB,
		retries:            restartRetries{},
	}

	opts.retries.evictPodRetries = c.newRetries("EvictPod")
	opts.retries.getPodRetries = c.newRetries("GetPod")
	opts.retries.runningPodRetries = c.newRetries("RunningPod")

	return opts.smartRestart
}

func (s *smartRestarter) smartRestart() (bool, error) {
	if len(s.podNames) == 0 {
		return false, nil
	}

	if s.infantNode != "" {
		isPodRunning, err := s.isPodRunning()
		if !isPodRunning {
			return s.retries.runningPodRetries.Wait(), err
		}

		// drop the first pod name
		if len(s.podNames) > 1 {
			s.podNames = s.podNames[1:]
		} else {
			s.podNames = nil
		}
		// reset the infantNode
		s.infantNode = ""
	} else {
		podName := s.podNames[0]
		// first check whether the requested pod exists or not
		var pod core.Pod
		err := s.KBClient.Get(context.TODO(), types.NamespacedName{
			Namespace: s.newDB.Namespace,
			Name:      podName,
		}, &pod)
		if err != nil {
			if kerr.IsNotFound(err) {
				// if the pod doesn't exist, no need to retry
				return false, err
			}
			// retry the process
			return s.retries.getPodRetries.Wait(podName), errors.Wrap(err, "failed to get the pod")
		}
		s.retries.getPodRetries.Initialize(podName)
		err = policy.EvictPod(context.TODO(), s.Client, types.NamespacedName{
			Namespace: pod.Namespace,
			Name:      podName,
		}, &meta.DeleteOptions{})
		if err != nil {
			return s.retries.evictPodRetries.Wait(podName), errors.Wrap(err, "failed to evict the pod")
		}
		s.retries.evictPodRetries.Initialize(podName)
		s.infantNode = podName
	}
	// retry
	return true, nil
}

func (c *hzOpsReqController) Restart() (time.Duration, error) {
	// Updating hazelcast ops-request phase to progressing and pause the hazelcast operator reconcile
	rq, err := c.UpdateHazelcastPhaseProgressingAndPauseReconcile(string(opsapi.HazelcastOpsRequestTypeRestart),
		"Hazelcast ops-request has started to restart hazelcast nodes")

	if err != nil || rq == RequeueDuration {
		return rq, err
	}

	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.RestartNodes) {
		c.RunParallel(opsapi.RestartNodes,
			"Successfully Restarted Hazelcast nodes",
			c.newRestartFunc(c.getPodsName(), c.db))
		return DefaultDuration, nil
	}

	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.Successful) {
		if err := c.resumeHazelcast(); err != nil {
			return DefaultDuration, err
		}
		if err := c.updateHazelcastOpsRequestPhase(string(opsapi.Successful), "Controller has successfully restart the Hazelcast replicas", opsapi.OpsRequestPhaseSuccessful); err != nil {
			c.log.Error(err, "failed to update status")
			return DefaultDuration, err
		}
	}
	return DefaultDuration, nil
}

func (s *smartRestarter) isPodRunning() (bool, error) {
	var pod core.Pod
	err := s.KBClient.Get(context.TODO(), types.NamespacedName{
		Namespace: s.newDB.Namespace,
		Name:      s.infantNode,
	}, &pod)
	if err != nil {
		return false, errors.Wrap(err, fmt.Sprintf("failed to get pod: %s", s.infantNode))
	}

	status := core_util.GetPodStatus(&pod)
	if status == "CrashLoopBackOff" {
		err := s.Client.CoreV1().Pods(s.db.Namespace).Delete(context.TODO(), pod.Name, meta.DeleteOptions{})
		if err != nil {
			return false, err
		}
		return false, fmt.Errorf("pod %s/%s is not ready yet. It was in CrashLoopBackOff. Deleted", pod.Namespace, pod.Name)
	}

	container := pod.Status.ContainerStatuses
	for _, containerStatus := range container {
		if !containerStatus.Ready {
			return false, fmt.Errorf("container %s is not ready", containerStatus.Name)
		}
	}

	if pod.Status.Phase != opsapi.Running || pod.DeletionTimestamp != nil {
		return false, errors.Wrap(err, fmt.Sprintf("%s is not running or terminating", s.infantNode))
	}
	return true, nil
}
