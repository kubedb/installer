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

	"github.com/pkg/errors"
	"gomodules.xyz/pointer"
	apps "k8s.io/api/apps/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/apis/core"
	cu "kmodules.xyz/client-go/client"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ScaleDownRetries struct {
	getPodRetries             *Retries
	patchStatefulSetRetries   *Retries
	deletePVCRetries          *Retries
	getPVCRetries             *Retries
	reassignPartitionsRetries *Retries
}
type hazelcastScaleDown struct {
	*hzOpsReqController
	statefulSet    *apps.StatefulSet
	newDB          *dbapi.Hazelcast
	deceasedNode   string
	targetReplicas int32
	nodeRole       string
	retries        ScaleDownRetries
}

func (c *hzOpsReqController) NewScaleDownFunc(statefulSet *apps.StatefulSet, db *dbapi.Hazelcast, targetReplicas int32, nodeType string) func() (bool, error) {
	opts := &hazelcastScaleDown{
		hzOpsReqController: c,
		statefulSet:        statefulSet,
		newDB:              db,
		targetReplicas:     targetReplicas,
		nodeRole:           nodeType,
		retries:            ScaleDownRetries{},
	}
	opts.retries.getPodRetries = c.newRetries("GetPod")
	opts.retries.deletePVCRetries = c.newRetries("DeletePvc")
	opts.retries.getPVCRetries = c.newRetries("GetPvc")
	opts.retries.patchStatefulSetRetries = c.newRetries("IsStatefulSetPatched")
	opts.retries.reassignPartitionsRetries = c.newRetries("ReassignPartitions")
	return opts.scaleDown
}

func (c *hazelcastScaleDown) scaleDown() (bool, error) {
	// Check whether the deceased node successfully deleted or not.
	// If deleted, set "deceasedNode" to empty.
	if c.deceasedNode != "" {
		pod := &core.Pod{}
		err := c.KBClient.Get(context.TODO(), types.NamespacedName{
			Name:      c.deceasedNode,
			Namespace: c.db.Namespace,
		}, pod)

		// keep returning until, err == nil
		if err == nil {
			return c.retries.getPodRetries.Wait(), errors.Wrapf(err, "Node: %s is not deleted yet", c.deceasedNode)
		}

		c.retries.getPodRetries.Initialize()

		// Delete PVC
		nodePVC := fmt.Sprintf("%s-%s", c.db.PVCName(kubedb.HazelcastVolumeData), c.deceasedNode)
		err = c.Client.CoreV1().PersistentVolumeClaims(c.db.Namespace).Delete(context.TODO(), nodePVC, metav1.DeleteOptions{})
		if err != nil && !kerr.IsNotFound(err) {
			return c.retries.deletePVCRetries.Wait(), err
		}
		c.retries.deletePVCRetries.Initialize()
		_, err = c.Client.CoreV1().PersistentVolumeClaims(c.db.Namespace).Get(context.TODO(), nodePVC, metav1.GetOptions{})
		// keep returning until, err != nil && IsNotFound(err)==true
		if err == nil || !kerr.IsNotFound(err) {
			return c.retries.getPVCRetries.Wait(), errors.Wrapf(err, "PVC: %s/%s is not deleted yet", c.db.Namespace, nodePVC)
		}
		c.retries.getPVCRetries.Initialize()
		c.log.Info("A node just left the cluster.", "NodeName", c.deceasedNode)
		c.deceasedNode = ""

		// return with retry
		return true, nil
	} else if *c.statefulSet.Spec.Replicas > c.targetReplicas {

		targetPodName := fmt.Sprintf("%s-%d", c.statefulSet.Name, *c.statefulSet.Spec.Replicas-1)

		// Reduce current replica count
		_, err := cu.CreateOrPatch(context.TODO(), c.KBClient, c.statefulSet, func(obj client.Object, createOp bool) client.Object {
			ret := obj.(*apps.StatefulSet)
			ret.Spec.Replicas = pointer.Int32P(*ret.Spec.Replicas - 1)
			return ret
		})
		if err != nil {
			c.log.Error(err, "Failed to patch StatefulSet", "StatefulSet", c.statefulSet.Name)
			return c.retries.patchStatefulSetRetries.Wait(), err
		}
		c.retries.patchStatefulSetRetries.Initialize()

		c.deceasedNode = targetPodName

		// return with retry
		// wait for this deceased node to discard from the cluster
		return true, nil
	}
	return false, nil
}
