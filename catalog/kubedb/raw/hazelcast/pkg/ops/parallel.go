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

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	"kmodules.xyz/client-go/tools/queue"
)

func (c *hzOpsReqController) RunParallel(condition, message string, fn func() (bool, error)) {
	log := c.log
	key, err := cache.MetaNamespaceKeyFunc(c.req)
	if err != nil {
		log.Error(err, "Failed to get key")
		return
	}
	if !c.KeyExists(key) {
		c.SetParallelismController(key, nil)
	}
	parallelCtrl := c.GetParallelismController(key)
	if c.ShouldProceed(key, condition) {
		ctx := context.Background()
		timeOut := time.Minute * time.Duration(2*len(c.getPodsName()))
		if c.req.Spec.Timeout != nil {
			timeOut = c.req.Spec.Timeout.Duration
		}
		ctx, cancel := context.WithTimeout(ctx, timeOut)
		c.AddCancelFunc(key, &cancel)
		ticker := time.NewTicker(retryInterval)
		go func() {
			klog.Infof("Run go routine for %v \n", key)
			defer parallelCtrl.Unlock()
			for {
				select {
				case <-ctx.Done():
					ticker.Stop()
					return
				case <-ticker.C:

					retryable, err := fn()
					if retryable {
						if err != nil {
							log.Info("Warning: ", err.Error(), "type", condition)
						}
						continue
					} else if err != nil {
						log.Error(err, "failed to complete task", "type", condition)
						c.pushFailureEventForHazelcastOpsReq(condition, err.Error())
						c.RemoveCancelFunc(key)
						continue
					}

					err = c.UpdateHazelcastOpsReqConditions(condition, message)
					if err != nil {
						log.Error(err, "failed to update condition")
						continue
					}

					c.Recorder.Event(
						c.req,
						corev1.EventTypeNormal,
						condition,
						message,
					)
					log.Info(message)

					c.RemoveCancelFunc(key)
				}
			}
		}()
	} else {
		klog.Infof("OpsRequest %s should not proceed; requeue in 5s", key)
		queue.EnqueueAfter(c.reqQueue.GetQueue(), c.req, retryInterval)
	}
}
