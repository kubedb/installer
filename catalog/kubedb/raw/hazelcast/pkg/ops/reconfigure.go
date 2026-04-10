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
	"sync"
	"time"

	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"
	"kubedb.dev/apimachinery/pkg/lib"
	"kubedb.dev/hazelcast/util"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	cu "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (c *hzOpsReqController) Reconfigure() (time.Duration, error) {
	// Updating Hazelcast ops-request phase to progressing and pause the Hazelcast operator reconcile
	rq, err := c.UpdateHazelcastPhaseProgressingAndPauseReconcile(string(opsapi.HazelcastOpsRequestTypeReconfigure),
		"Hazelcast ops-request has started to reconfigure Hazelcast nodes")

	if (err != nil) || (rq == RequeueDuration) {
		return rq, err
	}

	dbCopy := c.db.DeepCopy()
	reqSpec := c.req.Spec.Configuration

	if reqSpec.RemoveCustomConfig {
		dbCopy.Spec.Configuration = nil
	}

	// Update for config secret
	if reqSpec.ConfigSecret != nil && len(reqSpec.ConfigSecret.Name) != 0 {
		if dbCopy.Spec.Configuration == nil {
			dbCopy.Spec.Configuration = &dbapi.ConfigurationSpec{}
		}
		dbCopy.Spec.Configuration.SecretName = reqSpec.ConfigSecret.Name
	}

	if len(reqSpec.ApplyConfig) != 0 {
		if dbCopy.Spec.Configuration == nil {
			dbCopy.Spec.Configuration = &dbapi.ConfigurationSpec{}
		}
		if dbCopy.Spec.Configuration.Inline == nil {
			dbCopy.Spec.Configuration.Inline = make(map[string]string)
		}

		// Update the ApplyConfig field in dbCopy.Spec.Configuration
		dbCopy.Spec.Configuration.Inline, err = util.GetMergedConfig(dbCopy.Spec.Configuration.Inline, reqSpec.ApplyConfig)
		if err != nil {
			return DefaultDuration, fmt.Errorf("failed to merge apply configs with error %s", err.Error())
		}

		if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.PrepareApplyConfig) {
			// Make PrepareApplyConfig condition true after successfully updating the ApplyConfig field
			err = c.UpdateHazelcastOpsReqConditions(opsapi.PrepareApplyConfig, "Successfully prepared user provided apply configs")
			if err != nil {
				return DefaultDuration, err
			}
		}
	}
	// Update petSet with new volumes/volumeMounts for the new config secrets
	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.UpdateStatefulSets) {
		// perform reconcile(dbCopy) to update petSets

		reconcile, err := c.NewHazelcastReconcile(dbCopy)
		if err != nil {
			c.log.Error(err, "failed to get hazelcast reconcile")
			return DefaultDuration, err
		}

		// Reconcile Hazelcast resources
		c.RunParallel(opsapi.UpdateStatefulSets, "successfully reconciled the Hazelcast with new configure",
			c.NewReconcileFunc(dbCopy, reconcile))
		// return from here, process rest in the next cycles

		return DefaultDuration, nil

	}
	needRestart := IsRestartNeeded(c.req.Spec.Configuration)
	if needRestart {
		// Restart pods to apply new configuration
		if cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.UpdateStatefulSets) && !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.RestartNodes) {
			pName := c.getPodsName()
			if len(pName) == 0 {
				return DefaultDuration, errors.New("Pod list is empty")
			}
			c.RunParallel(opsapi.RestartNodes, "Successfully restarted all nodes", c.newRestartFunc(pName, dbCopy))
			return DefaultDuration, nil
		}
	}

	if !needRestart || cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.RestartNodes) {
		_, err := cu.CreateOrPatch(context.TODO(), c.KBClient, c.db, func(obj client.Object, createOp bool) client.Object {
			ret := obj.(*dbapi.Hazelcast)
			ret.Spec = dbCopy.Spec
			return ret
		})
		if err != nil {
			c.log.Error(err, "failed to patch the Hazelcast object")
			return DefaultDuration, err
		}
	}
	// resume and Change the opsapi request phase to "Successful".
	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.Successful) {
		if err := c.resumeHazelcast(); err != nil {
			return DefaultDuration, err
		}
		if err := c.updateHazelcastOpsRequestPhase(opsapi.Successful, "Successfully completed reconfigure Hazelcast", opsapi.OpsRequestPhaseSuccessful); err != nil {
			c.log.Error(err, "failed to update status")
			return DefaultDuration, err
		}
	}
	return DefaultDuration, nil
}

// merge logic

func (c *Controller) mergeConfigurations(pendingReconfigureOps []opsapi.Accessor) (any, error) {
	ret := &opsapi.ReconfigurationSpec{}
	var err error
	for _, o := range pendingReconfigureOps {
		op := o.(*opsapi.HazelcastOpsRequest)
		if op.Spec.Configuration == nil {
			continue
		}
		ret, err = c.mergeOpsConfigurations(ret, op.Spec.Configuration)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func (c *Controller) mergeOpsConfigurations(merged, cfg *opsapi.ReconfigurationSpec) (*opsapi.ReconfigurationSpec, error) {
	if cfg == nil {
		return nil, nil
	}

	if merged == nil {
		merged = &opsapi.ReconfigurationSpec{
			ApplyConfig: map[string]string{},
		}
	}

	// Handle RemoveCustomConfig - if any request wants to remove config, apply it
	// But only if it's not overridden by a later request with actual config
	if cfg.RemoveCustomConfig {
		merged.ConfigSecret = nil
		merged.ApplyConfig = make(map[string]string)
		merged.RemoveCustomConfig = true
	}

	// If there's a ConfigSecret, use the latest one
	if cfg.ConfigSecret != nil {
		merged.ConfigSecret = cfg.ConfigSecret
	}

	// Merge ApplyConfig - later values override earlier ones
	if len(cfg.ApplyConfig) > 0 {
		if merged.ApplyConfig == nil {
			merged.ApplyConfig = make(map[string]string)
		}
		currentMergedConfig, err := util.GetMergedConfig(merged.ApplyConfig, cfg.ApplyConfig)
		if err != nil {
			return nil, err
		}
		merged.ApplyConfig = currentMergedConfig
	}

	if merged.ApplyConfig == nil && merged.ConfigSecret == nil && !merged.RemoveCustomConfig {
		return nil, nil
	}

	return merged, nil
}

func (c *Controller) convert(u *unstructured.Unstructured) (opsapi.Accessor, error) {
	var mergedOps opsapi.HazelcastOpsRequest
	if err := runtime.DefaultUnstructuredConverter.
		FromUnstructured(u.Object, &mergedOps); err != nil {
		return nil, err
	}

	return &mergedOps, nil
}

var muxForReconfigureMerger sync.Mutex

func (c *Controller) runSkippingLogicForReconfigure(req *opsapi.HazelcastOpsRequest, log logr.Logger) (int, error) {
	muxForReconfigureMerger.Lock()
	defer muxForReconfigureMerger.Unlock()
	merger, err := lib.NewReconfigureMerger(c.KBClient, opsapi.ResourceKindHazelcastOpsRequest, req,
		c.mergeConfigurations, c.convert, log)
	if err != nil {
		log.Error(err, "failed to construct NewReconfigureMerger")
	}
	return merger.Run()
}
