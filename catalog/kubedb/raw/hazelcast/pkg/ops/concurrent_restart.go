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

	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"

	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	core_util "kmodules.xyz/client-go/core/v1"
	"kmodules.xyz/client-go/policy"
)

type concurrentRestartRetries struct {
	evictPodRetries   *Retries
	getPodRetries     *Retries
	runningPodRetries *Retries
}
type concurrentRestarter struct {
	*hzOpsReqController
	podNames []string
	newDB    *dbapi.Hazelcast
	phase    int
	retries  concurrentRestartRetries
}

func (c *hzOpsReqController) newConcurrentRestartFunc(podNames []string, newDB *dbapi.Hazelcast) func() (bool, error) {
	opts := &concurrentRestarter{
		hzOpsReqController: c,
		podNames:           podNames,
		newDB:              newDB,
		retries:            concurrentRestartRetries{},
		phase:              0,
	}

	opts.retries.evictPodRetries = c.newRetries("EvictPod")
	opts.retries.getPodRetries = c.newRetries("GetPod")
	opts.retries.runningPodRetries = c.newRetries("RunningPod")

	return opts.concurrentRestart
}

func (s *concurrentRestarter) concurrentRestart() (bool, error) {
	if len(s.podNames) == 0 {
		return false, nil
	}

	if s.phase == 0 {
		// Evict all pod...
		for idx := range len(s.podNames) {
			podName := s.podNames[idx]

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
		}

		s.phase = 1
	}

	if s.phase == 1 {
		// Make sure all pods are in running state
		for idx := range len(s.podNames) {
			podName := s.podNames[idx]

			isPodRunning, err := s.isPodRunning(podName)
			if !isPodRunning {
				return s.retries.runningPodRetries.Wait(), err
			}
			s.retries.runningPodRetries.Initialize(podName)
		}
		s.podNames = nil
	}

	return true, nil
}

func (s *concurrentRestarter) isPodRunning(podName string) (bool, error) {
	var pod core.Pod
	err := s.KBClient.Get(context.TODO(), types.NamespacedName{
		Namespace: s.newDB.Namespace,
		Name:      podName,
	}, &pod)
	if err != nil {
		return false, errors.Wrap(err, fmt.Sprintf("failed to get pod: %s", podName))
	}

	status := core_util.GetPodStatus(&pod)
	if status == "CrashLoopBackOff" {
		err := s.Client.CoreV1().Pods(s.db.Namespace).Delete(context.TODO(), pod.Name, meta.DeleteOptions{})
		if err != nil {
			return false, err
		}
		return false, fmt.Errorf("pod %s/%s is not ready yet. It was in CrashLoopBackOff. Deleted", pod.Namespace, pod.Name)
	}

	if pod.Status.Phase != opsapi.Running || pod.DeletionTimestamp != nil {
		return false, errors.Wrap(err, fmt.Sprintf("%s is not running or terminating", podName))
	}
	return true, nil
}
