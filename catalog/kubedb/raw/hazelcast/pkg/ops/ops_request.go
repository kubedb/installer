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
	"strings"
	"time"

	"kubedb.dev/apimachinery/apis/kubedb"
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"
	"kubedb.dev/apimachinery/pkg/lib"

	"github.com/go-logr/logr"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	cu "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	"kmodules.xyz/client-go/tools/queue"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	retryInterval   = 10 * time.Second
	DefaultDuration = time.Duration(0)
	RequeueDuration = time.Second * 3
)

type hzOpsReqController struct {
	*Controller
	db  *dbapi.Hazelcast
	req *opsapi.HazelcastOpsRequest
	log logr.Logger
}

func (c *Controller) runHazelcastOpsRequest(k any) error {
	key := k.(string)
	log := c.log
	log.WithValues("key", key).Info("Started processing HazelcastOpsRequest")

	req, err := c.getObjectFromKey(key)
	if err != nil && !errors.IsNotFound(err) {
		log.Error(err, "Fetching HazelcastOpsRequest object from store failed", "key", key)
		return err
	}
	if errors.IsNotFound(err) || req.GetDeletionTimestamp() != nil {
		log.Info("HazelcastOpsRequest does not exist anymore", "key", key)
		if c.KeyExists(key) {
			c.DeleteParallelismController(key)
		}
	} else {
		// return nil if req.status.phase is "Failed" , "Successful" or "Skipped"
		if req.Status.Phase == opsapi.OpsRequestPhaseSuccessful || req.Status.Phase == opsapi.OpsRequestPhaseFailed || req.Status.Phase == opsapi.OpsRequestPhaseSkipped {
			return nil
		}

		if req.Status.Phase == "" {
			_, err := cu.PatchStatus(context.TODO(), c.KBClient, req, func(obj client.Object) client.Object {
				ret := obj.(*opsapi.HazelcastOpsRequest)
				ret.Status.Phase = opsapi.OpsRequestPhasePending
				return ret
			})
			if err != nil {
				log.Error(err, "failed to update Hazelcast opsapi request status")
				return err
			}
		}

		reqs, err := c.reqLister.HazelcastOpsRequests(req.Namespace).List(labels.Everything())
		if err != nil {
			return err
		}

		var lst []client.Object
		for _, request := range reqs {
			lst = append(lst, request)
		}

		skipperValue, err := c.runSkippingLogicForReconfigure(req, log)
		if err != nil {
			return err
		}
		if skipperValue != lib.ContinueGeneral {
			klog.Infof("Do not Continue General %v \n", req.Name)
			var msg string

			if req.Status.Phase == opsapi.OpsRequestPhaseSkipped {
				// Request is merged as Skipped by the merger
				msg = fmt.Sprintf("HazelcastOpsRequest %s/%s skipped as it has been merged with other reconfigure requests", req.Namespace, req.Name)
				log.Info(msg)

				c.Recorder.Event(
					req,
					core.EventTypeNormal,
					"ConfigurationMerged",
					msg,
				)
			} else {
				// concurrent ops detected
				msg = fmt.Sprintf("Skipping HazelcastOpsRequest %s/%s, concurrent HazelcastOpsRequest is not allowed for same database, will retry after 30 seconds.", req.Namespace, req.Name)
				log.Info(msg)

				if skipperValue == lib.RequeueNeeded {
					queue.EnqueueAfter(c.reqQueue.GetQueue(), req, 30*time.Second)
				}

				c.Recorder.Event(
					req,
					core.EventTypeNormal,
					"Pending",
					msg,
				)
			}
			return nil
		}

		skipper := lib.NewSkipper(c.KBClient, opsapi.ResourceKindHazelcastOpsRequest, req, lst)
		if skip, err := skipper.SkipOpsReq(); err != nil {
			return err
		} else if skip {
			queue.EnqueueAfter(c.reqQueue.GetQueue(), req, 30*time.Second)
			msg := fmt.Sprintf("Skipping HazelcastOpsRequest %s/%s, concurrent HazelcastOpsRequest is not allowed for same database, will retry after 30 seconds.", req.Namespace, req.Name)
			log.Info(msg)
			c.Recorder.Event(
				req,
				core.EventTypeNormal,
				"Pending",
				msg,
			)
			return nil
		}

		hz := dbapi.Hazelcast{}
		err = c.KBClient.Get(context.TODO(), types.NamespacedName{
			Namespace: req.Namespace,
			Name:      req.Spec.DatabaseRef.Name,
		}, &hz)
		if err != nil {
			log.Error(err, "Failed to get Hazelcast", dbapi.ResourceKindHazelcast, req.Spec.DatabaseRef.Name)
			return err
		}
		if req.Status.Phase != opsapi.OpsRequestPhaseProgressing {
			if (req.Spec.Apply == "" || req.Spec.Apply == opsapi.ApplyOptionIfReady) && hz.Status.Phase != dbapi.DatabasePhaseReady {
				log.Info(fmt.Sprintf("Retrying in %v seconds, as db is not in `Ready` state.`", retryInterval), "key", key)
				queue.EnqueueAfter(c.reqQueue.GetQueue(), req, retryInterval)
				return nil
			}
			if req.Spec.Apply == opsapi.ApplyOptionAlways && !cutil.IsConditionTrue(hz.Status.Conditions, kubedb.DatabaseProvisioned) {
				log.Info(fmt.Sprintf("Retrying in %v seconds, as db has not been Provisioned yet.`", retryInterval), "key", key)
				queue.EnqueueAfter(c.reqQueue.GetQueue(), req, retryInterval)
				return nil
			}
		}

		log = log.WithValues("Namespace", req.Namespace, opsapi.ResourceKindHazelcastOpsRequest, req.Name, dbapi.ResourceKindHazelcast, hz.Name)
		con := &hzOpsReqController{
			Controller: c,
			db:         &hz,
			req:        req,
			log:        log,
		}

		switch opsapi.HazelcastOpsRequestType(req.GetRequestType()) {
		case opsapi.HazelcastOpsRequestTypeRestart:
			rq, err := con.Restart()
			if err != nil {
				con.pushFailureEventForHazelcastOpsReq(string(opsapi.HazelcastOpsRequestTypeRestart), err.Error())
				return err
			} else if rq != DefaultDuration {
				c.reqQueue.GetQueue().AddAfter(key, rq)
			}
		case opsapi.HazelcastOpsRequestTypeReconfigure:
			rq, err := con.Reconfigure()
			if err != nil {
				con.pushFailureEventForHazelcastOpsReq(string(opsapi.HazelcastOpsRequestTypeReconfigure), err.Error())
			} else if rq != DefaultDuration {
				c.reqQueue.GetQueue().AddAfter(key, rq)
			}
		case opsapi.HazelcastOpsRequestTypeVerticalScaling:
			rq, err := con.VerticalScale()
			if err != nil {
				con.pushFailureEventForHazelcastOpsReq(string(opsapi.HazelcastOpsRequestTypeVerticalScaling), err.Error())
			} else if rq != DefaultDuration {
				c.reqQueue.GetQueue().AddAfter(key, rq)
			}
		case opsapi.HazelcastOpsRequestTypeHorizontalScaling:
			rq, err := con.HorizontalScale()
			if err != nil {
				con.pushFailureEventForHazelcastOpsReq(string(opsapi.HazelcastOpsRequestTypeHorizontalScaling), err.Error())
			} else if rq != DefaultDuration {
				c.reqQueue.GetQueue().AddAfter(key, rq)
			}
		case opsapi.HazelcastOpsRequestTypeVolumeExpansion:
			rq, err := con.VolumeExpansion()
			if err != nil {
				con.pushFailureEventForHazelcastOpsReq(string(opsapi.HazelcastOpsRequestTypeVolumeExpansion), err.Error())
			} else if rq != DefaultDuration {
				c.reqQueue.GetQueue().AddAfter(key, rq)
			}
		case opsapi.HazelcastOpsRequestTypeRotateAuth:
			rq, err := con.RotateAuthentication()
			if err != nil {
				con.pushFailureEventForHazelcastOpsReq(string(opsapi.HazelcastOpsRequestTypeRotateAuth), err.Error())
			} else if rq != DefaultDuration {
				c.reqQueue.GetQueue().AddAfter(key, rq)
			}
		case opsapi.HazelcastOpsRequestTypeReconfigureTLS:
			rq, err := con.reconfigureTLS()
			if err != nil {
				con.pushFailureEventForHazelcastOpsReq(string(opsapi.HazelcastOpsRequestTypeReconfigureTLS), err.Error())
			} else if rq != DefaultDuration {
				c.reqQueue.GetQueue().AddAfter(key, rq)
			}
		case opsapi.HazelcastOpsRequestTypeUpdateVersion:
			rq, err := con.UpdateVersion()
			if err != nil {
				con.pushFailureEventForHazelcastOpsReq(string(opsapi.HazelcastOpsRequestTypeUpdateVersion), err.Error())
			} else if rq != DefaultDuration {
				c.reqQueue.GetQueue().AddAfter(key, rq)
			}
		default:
			return fmt.Errorf("defined OpsRequestType %s is not supported, supported types for Hazelcast are %s", req.Spec.Type, strings.Join(opsapi.HazelcastOpsRequestTypeNames(), ", "))
		}
	}

	return nil
}

// returns data node pvc names

func (c *hzOpsReqController) getPVCNames() []string {
	var pvcNames []string
	pvcName := c.db.PVCName(kubedb.HazelcastVolumeData)
	pods := c.getPodsName()
	for _, pod := range pods {
		pvcNames = append(pvcNames, fmt.Sprintf("%s-%s", pvcName, pod))
	}
	return pvcNames
}
