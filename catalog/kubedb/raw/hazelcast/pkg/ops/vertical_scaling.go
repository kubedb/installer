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

	"kubedb.dev/apimachinery/apis/kubedb"
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"
	"kubedb.dev/apimachinery/pkg/lib"

	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	cu "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	coreutil "kmodules.xyz/client-go/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (c *hzOpsReqController) VerticalScale() (time.Duration, error) {
	// Updating Hazelcast ops-request phase to progressing and pause the rabbitmq operator reconcile
	rq, err := c.UpdateHazelcastPhaseProgressingAndPauseReconcile(string(opsapi.HazelcastOpsRequestTypeVerticalScaling),
		"Hazelcast ops-request has started to vertically scaling the Hazelcast nodes")

	if (err != nil) || (rq == RequeueDuration) {
		return rq, err
	}
	var changedPods []string
	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.UpdateStatefulSets) {
		changedPods, err = c.updateStatefulSetsResources()
		if err != nil {
			return DefaultDuration, err
		}
		if err := c.UpdateHazelcastOpsReqConditions(opsapi.UpdateStatefulSets, "Successfully updated StatefulSets Resources"); err != nil {
			return DefaultDuration, err
		}
		c.Recorder.Event(
			c.req,
			core.EventTypeNormal,
			opsapi.UpdateStatefulSets,
			"Successfully updated StatefulSets Resources",
		)
		c.log.Info("Successfully updated StatefulSets Resources")
	}

	if cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.UpdateStatefulSets) &&
		!cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.RestartPods) {
		c.RunParallel(opsapi.RestartPods,
			"Successfully Restarted Pods With Resources",
			c.newRestartFunc(
				changedPods,
				c.db),
		)
		return DefaultDuration, nil
	}

	c.log.Info("Vertical Scale Succeeded")

	verticalScaling := c.req.Spec.VerticalScaling

	if verticalScaling.Hazelcast != nil {
		err := c.patchHazelcastContainers(verticalScaling.Hazelcast)
		if err != nil {
			c.log.Error(err, "failed to patch Hazelcast")
			return DefaultDuration, err
		}
	}

	// resume and Change the opsapi request phase to "Successful".
	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.Successful) {
		if err := c.resumeHazelcast(); err != nil {
			return DefaultDuration, err
		}
		if err := c.updateHazelcastOpsRequestPhase(opsapi.Successful, "Successfully completed the vertical scaling for RabbitMQ", opsapi.OpsRequestPhaseSuccessful); err != nil {
			c.log.Error(err, "failed to update status")
			return DefaultDuration, err
		}
	}

	return DefaultDuration, nil
}

func (c *hzOpsReqController) getHazelcastContainers(hz *dbapi.Hazelcast) []core.Container {
	return hz.Spec.PodTemplate.Spec.Containers
}

func (c *hzOpsReqController) patchHazelcastContainers(node *opsapi.PodResources) error {
	_, err := cu.CreateOrPatch(context.TODO(), c.KBClient, c.db, func(obj client.Object, createOp bool) client.Object {
		ret := obj.(*dbapi.Hazelcast)
		lib.UpdateMachineProfileAnnotation(ret.Annotations, c.req.Annotations)
		dbContainer := coreutil.GetContainerByName(c.getHazelcastContainers(ret), kubedb.HazelcastContainerName)
		if dbContainer != nil {
			dbContainer.Resources = node.Resources
		}
		setTopologyInfo(ret, node)
		return ret
	})
	return err
}

func setTopologyInfo(hz *dbapi.Hazelcast, node *opsapi.PodResources) {
	hz.Spec.PodTemplate.Spec.NodeSelector = lib.SetNodeSelector(node.NodeSelectionPolicy, hz.Spec.PodTemplate.Spec.NodeSelector, node.Topology)
	hz.Spec.PodTemplate.Spec.Tolerations = lib.SetToleration(node.NodeSelectionPolicy, hz.Spec.PodTemplate.Spec.Tolerations, node.Topology)
}

func (c *hzOpsReqController) patchStatefulSet(node *opsapi.PodResources) error {
	_, err := cu.CreateOrPatch(context.TODO(), c.KBClient, &apps.StatefulSet{
		ObjectMeta: meta.ObjectMeta{
			Name:      c.db.StatefulSetName(),
			Namespace: c.db.Namespace,
		},
	}, func(obj client.Object, createOp bool) client.Object {
		ret := obj.(*apps.StatefulSet)
		for i, psContainer := range ret.Spec.Template.Spec.Containers {
			// Update main container Resources
			if psContainer.Name == kubedb.HazelcastContainerName {
				c.log.Info("Updating Hazelcast Container Vertically")
				ret.Spec.Template.Spec.Containers[i].Resources = node.Resources
			}
		}

		if node.Topology != nil {
			ret.Spec.Template.Spec.NodeSelector = lib.SetNodeSelector(node.NodeSelectionPolicy, ret.Spec.Template.Spec.NodeSelector, node.Topology)
			ret.Spec.Template.Spec.Tolerations = lib.SetToleration(node.NodeSelectionPolicy, ret.Spec.Template.Spec.Tolerations, node.Topology)
		}
		return ret
	})

	return err
}

func (c *hzOpsReqController) updateStatefulSetsResources() ([]string, error) {
	verticalScaling := c.req.Spec.VerticalScaling

	if c.req.Spec.VerticalScaling.Hazelcast != nil {
		err := c.patchStatefulSet(verticalScaling.Hazelcast)
		if err != nil {
			c.log.Error(err, fmt.Sprintf("failed to patch statefulset %s", c.db.StatefulSetName()))
			return []string{}, err
		}
	}

	return c.getPodsName(), nil
}
