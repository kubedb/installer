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

	hazelcastgo "kubedb.dev/db-client-go/hazelcast"

	"github.com/pkg/errors"
	apps "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientutil "kmodules.xyz/client-go/client"
	"kubestash.dev/apimachinery/apis"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type scaleUpNode struct {
	*hzOpsReqController
	statefulSet     *apps.StatefulSet
	retries         scaleUpNodeRetries
	desiredReplicas int32
	infantNode      string
}

type scaleUpNodeRetries struct {
	isNodeInClusterRetries  *Retries
	patchStatefulSetRetries *Retries
}

func (c *hzOpsReqController) NewScaleUpFunc(sts *apps.StatefulSet, desiredReplicas int32) func() (bool, error) {
	up := &scaleUpNode{
		hzOpsReqController: c,
		statefulSet:        sts,
		retries:            scaleUpNodeRetries{},
		desiredReplicas:    desiredReplicas,
	}
	up.retries.isNodeInClusterRetries = c.newRetries("IsNodeInCluster")
	up.retries.patchStatefulSetRetries = c.newRetries("PatchStatefulSet")
	return up.run
}

// function returns error with a boolean which specifies
// whether we should retry the process or fails the opsReq.
func (s *scaleUpNode) run() (bool, error) {
	if s.statefulSet == nil {
		return false, errors.New("statefulSet is empty")
	}

	// Check whether the new node joined the cluster or not.
	// If the node successfully joined the cluster, set "infantNode" to empty string.
	if s.infantNode != "" {
		joined, err := s.isNodeInCluster(s.infantNode)
		if err != nil {
			return s.retries.isNodeInClusterRetries.Wait(), err
		}
		// If the node hasn't joined the cluster yet.
		if !joined {
			return s.retries.isNodeInClusterRetries.Wait(), errors.Wrapf(err, "Hazelcast node: %s hasn't joined the cluster yet.", s.infantNode)
		}

		s.retries.isNodeInClusterRetries.Initialize()

		// Node joined the cluster,
		// set infant node to empty.
		s.log.Info("A new node has joined the cluster.", "NodeName", s.infantNode)

		s.infantNode = ""

		// return with retry
		return true, nil
	} else if *s.statefulSet.Spec.Replicas < s.desiredReplicas {
		// Increase current replica count
		sts, err := s.UpdateStatefulSetReplicas(*s.statefulSet.Spec.Replicas+1, s.statefulSet.Name)
		if err != nil {
			s.log.Error(err, "failed to update statefulSet replica", "StatefulSet", s.statefulSet.Name)
			return s.retries.patchStatefulSetRetries.Wait(), err
		}

		s.retries.patchStatefulSetRetries.Initialize()

		// update statefulSet
		s.statefulSet = sts

		// set infant node
		s.infantNode = fmt.Sprintf("%s-%d", sts.Name, *sts.Spec.Replicas-1)

		// return with retry
		// wait for this infant node to join the cluster
		return true, nil
	}

	// process is completed, no need to retry
	return false, nil
}

// Returns true if the given node joined other nodes in the cluster.
func (c *hzOpsReqController) isNodeInCluster(podName string) (bool, error) {
	ctx := context.TODO()
	hzClient, err := hazelcastgo.NewKubeDBClientBuilder(c.KBClient, c.db).WithPod(podName).WithLog(c.log).WithContext(ctx).GetHazelcastClient()
	if err != nil {
		return false, err
	}

	defer func() {
		err := hzClient.Shutdown(ctx)
		if err != nil {
			c.log.Error(err, "Failed to shutdown client")
		}
	}()

	return true, nil
}

func (c *hzOpsReqController) UpdateStatefulSetReplicas(replicas int32, statefulSetName string) (*apps.StatefulSet, error) {
	meta := metav1.ObjectMeta{
		Name:      statefulSetName,
		Namespace: c.db.Namespace,
	}
	statefulSet := &apps.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       apis.KindStatefulSet,
			APIVersion: apps.SchemeGroupVersion.String(),
		},
		ObjectMeta: meta,
	}

	_, err := clientutil.CreateOrPatch(context.TODO(), c.KBClient, statefulSet, func(obj client.Object, createOp bool) client.Object {
		in := obj.(*apps.StatefulSet)
		in.Spec.Replicas = &replicas
		in.Spec.UpdateStrategy.Type = apps.OnDeleteStatefulSetStrategyType
		return in
	})
	if err != nil {
		return nil, err
	}

	return statefulSet, nil
}
