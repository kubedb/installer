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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"kubedb.dev/apimachinery/apis/kubedb"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"
	"kubedb.dev/apimachinery/pkg/lib"
	hazelcastgo "kubedb.dev/db-client-go/hazelcast"

	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	cu "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	core_util "kmodules.xyz/client-go/core/v1"
	"kmodules.xyz/client-go/meta"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type volumeExpansionOfflineFunc struct {
	*hzOpsReqController
	retries     volumeExpansionOfflineFuncRetries
	storageReq  *resource.Quantity
	podsName    []string
	infantPod   string
	deceasedPod string
}

type volumeExpansionOfflineFuncRetries struct {
	getPodRetries           *Retries
	createPodRetries        *Retries
	deletePodRetries        *Retries
	getPVCRetries           *Retries
	patchPVCRetries         *Retries
	patchOpsReqRetries      *Retries
	compareStorageRetries   *Retries
	runningHazelcastRetries *Retries
}

func (c *hzOpsReqController) newVolumeExpansionOfflineFunc(podNames []string, storageReq *resource.Quantity) func() (bool, error) {
	opts := &volumeExpansionOfflineFunc{
		hzOpsReqController: c,
		retries:            volumeExpansionOfflineFuncRetries{},
		podsName:           podNames,
		storageReq:         storageReq,
	}
	opts.retries.getPodRetries = c.newRetries("GetPod")
	opts.retries.createPodRetries = c.newRetries("CreatePod")
	opts.retries.deletePodRetries = c.newRetries("DeletePod")
	opts.retries.getPVCRetries = c.newRetries("GetPvc")
	opts.retries.patchPVCRetries = c.newRetries("PatchPvc")
	opts.retries.patchOpsReqRetries = c.newRetries("PatchOpsRequest")
	opts.retries.compareStorageRetries = c.newRetries("CompareStorage")
	opts.retries.runningHazelcastRetries = c.newRetries("RunningHazelcast")
	return opts.run
}

// function returns error with a boolean which specifies
// whether we should retry the process or fails the opsReq.
// If the function completes its given task, it return "false, nil"
// i.e. No need to retry, no error occurred.
func (v *volumeExpansionOfflineFunc) run() (bool, error) {
	if v.podsName == nil {
		return false, nil
	}

	// Here a pod has 3 phases:
	// 3. Initial:
	//		- Get the Pod
	//		- Store Pod's YAML to OpsReq's annotation: CleanUp unwanted metadata
	//		- Delete the Pod
	//		- Set the Pod as DeceasedPod
	// 2. Deceased:
	//		- Get the Pod to make sure that the pod is deleted; otherwise retry
	//		- Get the associate PVC
	//		- Patch the PVC with new volume spec
	//		- Wait for the PVC to get expanded
	//		- Recreate the Pod with the YAML from OpsReq's annotation
	//		- Set the Pod as InfantPod
	// 1. Infant:
	//		- Get the Pod to make sure that the pod is recreated
	//		- Drop the Pod from process queue.

	if v.infantPod != "" {
		// Check whether the pod is deleted or not
		pod := &core.Pod{}
		err := v.KBClient.Get(context.TODO(), types.NamespacedName{
			Name:      v.infantPod,
			Namespace: v.db.Namespace,
		}, pod)
		if err != nil {
			v.log.Info("pod is waiting to be created", "pod", v.infantPod, "error", err.Error())
			return v.retries.getPodRetries.Wait(), err
		}
		v.retries.getPodRetries.Initialize()

		// creating client, we are checking Hazelcast cluster working fine or not
		isHazelcastRunning, err := v.isHazelcastRunning()
		if !isHazelcastRunning {
			return v.retries.runningHazelcastRetries.Wait(), err
		}

		// delete first item of the list
		if len(v.podsName) > 1 {
			v.podsName = v.podsName[1:]
		} else {
			v.podsName = nil
		}

		// reset infant pod
		v.infantPod = ""
		// retry
		return true, nil
	} else if v.deceasedPod != "" {
		// Check whether the pod is deleted or not
		pod := &core.Pod{}
		err := v.KBClient.Get(context.TODO(), types.NamespacedName{
			Name:      v.deceasedPod,
			Namespace: v.db.Namespace,
		}, pod)
		if err == nil || (err != nil && !kerr.IsNotFound(err)) {
			v.log.Info("pod is waiting to be deleted", "pod", v.deceasedPod)
			return v.retries.getPodRetries.Wait(), err
		}
		v.retries.getPodRetries.Initialize()

		// Now that the pod is deleted!
		// Time to patch the PVC
		// Get PVC
		pvcName := fmt.Sprintf("%s-%s", v.db.PVCName(kubedb.HazelcastVolumeData), v.deceasedPod)
		pvc, err := v.Client.CoreV1().PersistentVolumeClaims(v.db.Namespace).Get(context.TODO(), pvcName, metav1.GetOptions{})
		if err != nil {
			// retry
			return v.retries.getPVCRetries.Wait(), err
		}
		v.retries.getPVCRetries.Initialize()

		// quantity.Cmp(y) returns 0 if the quantity is equal to y, -1 if the quantity is less than y, or 1 if the
		// quantity is greater than y.
		ck := pvc.Spec.Resources.Requests.Storage().Cmp(*v.storageReq)
		if ck == -1 {
			v.log.Info("expanding PVC", "pvcName", pvcName)
			_, _, err = core_util.CreateOrPatchPVC(context.TODO(), v.Client, pvc.ObjectMeta, func(in *core.PersistentVolumeClaim) *core.PersistentVolumeClaim {
				in.Spec.Resources = core.VolumeResourceRequirements{
					Requests: map[core.ResourceName]resource.Quantity{
						core.ResourceStorage: *v.storageReq,
					},
				}
				return in
			}, metav1.PatchOptions{})
			if err != nil {
				return v.retries.patchPVCRetries.Wait(), err
			}
			v.retries.patchPVCRetries.Initialize()
		}

		// quantity.Cmp(y) returns 0 if the quantity is equal to y, -1 if the quantity is less than y, or 1 if the
		// quantity is greater than y.
		ck = pvc.Status.Capacity.Storage().Cmp(*v.storageReq)
		if ck == -1 {
			if !IsPVCConditionTrue(pvc.Status.Conditions, core.PersistentVolumeClaimFileSystemResizePending) {
				return v.retries.compareStorageRetries.Wait(), fmt.Errorf("FileSystemResizePending status of PVC is not true")
			}
		}
		v.retries.compareStorageRetries.Initialize()

		v.log.Info("PVC successfully expanded", "PVCName", pvc.Name)

		// Time to Recreate the pod
		if val, ok := v.req.Annotations[lib.VolumeExpansionAnnotationKey(v.deceasedPod)]; ok && val != "" {
			pod := &core.Pod{}
			podJson := val
			err = json.Unmarshal([]byte(podJson), pod)
			if err != nil {
				v.log.Info("failed to unmarshal pod from annotation", "Pod", v.deceasedPod, "error", err)
				return false, err
			}

			_, err = cu.CreateOrPatch(context.TODO(), v.KBClient, pod, func(obj client.Object, createOp bool) client.Object {
				ret := obj.(*core.Pod)
				return ret
			})
			if err != nil {
				return v.retries.createPodRetries.Wait(), err
			}
			v.retries.createPodRetries.Initialize()

			// Delete the pod YAML from annotation
			_, err := cu.CreateOrPatch(context.TODO(), v.KBClient, v.req, func(obj client.Object, createOp bool) client.Object {
				ret := obj.(*opsapi.HazelcastOpsRequest)
				delete(ret.Annotations, lib.VolumeExpansionAnnotationKey(v.deceasedPod))
				return ret
			})
			if err != nil {
				v.log.Info("failed to patch opsapi request", "error", err)
				return v.retries.patchOpsReqRetries.Wait(), err
			}
			v.retries.patchOpsReqRetries.Initialize()
		} else {
			// Since the Pod YAML is missing in annotation,
			// Check whether the POD is running or not.
			// Return error without retry, if the pod is missing and No YAML to recreate the Pod.
			pod := &core.Pod{}
			err := v.KBClient.Get(context.TODO(), types.NamespacedName{
				Namespace: v.db.Namespace,
				Name:      v.deceasedPod,
			}, pod)
			if err != nil {
				if kerr.IsNotFound(err) {
					return false, err
				}
				return v.retries.getPodRetries.Wait(), err
			}
			v.retries.getPodRetries.Initialize()
		}

		ck = pvc.Status.Capacity.Storage().Cmp(*v.storageReq)
		if ck == -1 {
			return v.retries.compareStorageRetries.Wait(), fmt.Errorf("pvc: %s is not yet expanded", pvc.Name)
		}

		// set the infant pod
		v.infantPod = v.deceasedPod

		// reset deceasedPod
		v.deceasedPod = ""
	} else {
		podName := v.podsName[0]
		klog.Info("podname now", podName)
		pod := &core.Pod{}
		err := v.KBClient.Get(context.TODO(), types.NamespacedName{
			Name:      podName,
			Namespace: v.db.Namespace,
		}, pod)
		if err != nil {
			if kerr.IsNotFound(err) {
				// Case: opsapi-manager operator got restarted during running opsapi request
				// Scenario:
				//		- annotation was stored in opsRequest
				//		- One pod was deleted (say pod-0)
				//		- operator restarted!!!

				// Since the pod doesn't exist, check whether the YAML is stored in annotation or not.
				if podYML, exists := v.req.Annotations[lib.VolumeExpansionAnnotationKey(podName)]; exists && podYML != "" {
					// No need to do the rest for this pod, skip to next step
					// set deceased Pod
					v.deceasedPod = podName
					return true, nil
				}
				return false, fmt.Errorf("pod: %s doesn't exist", podName)
			}
			// retry
			return v.retries.getPodRetries.Wait(), err
		}
		v.retries.getPodRetries.Initialize()
		// Drop unwanted metadata
		pod.ObjectMeta = metav1.ObjectMeta{
			Name:            pod.Name,
			Namespace:       pod.Namespace,
			Labels:          pod.Labels,
			Annotations:     pod.Annotations,
			Finalizers:      pod.Finalizers,
			OwnerReferences: pod.OwnerReferences,
		}
		// Drop whole status section
		pod.Status = core.PodStatus{}

		// Store pod YAML to opsapi request annotations
		podJson, err := json.Marshal(pod)
		if err != nil {
			v.log.Info("failed to marshal pod", "Pod", podName, "error", err)
			return false, err
		}
		buffer := new(bytes.Buffer)
		err = json.Compact(buffer, podJson)
		if err != nil {
			v.log.Info("failed to compact pod json", "Pod", podName, "error", err)
			return false, err
		}

		_, err = cu.CreateOrPatch(context.TODO(), v.KBClient, v.req, func(obj client.Object, createOp bool) client.Object {
			ret := obj.(*opsapi.HazelcastOpsRequest)
			annotation := map[string]string{
				lib.VolumeExpansionAnnotationKey(podName): buffer.String(),
			}
			ret.Annotations = meta.OverwriteKeys(ret.Annotations, annotation)
			return ret
		})
		if err != nil {
			v.log.Info("failed to patch opsapi request", "error", err)
			return v.retries.patchOpsReqRetries.Wait(), err
		}
		v.retries.patchOpsReqRetries.Initialize()

		v.log.Info("Deleting pod", "pod", podName)
		err = v.KBClient.Delete(context.TODO(), &core.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      podName,
				Namespace: v.db.Namespace,
			},
		})
		if err != nil {
			v.log.Info("failed to delete pod", "Pod", podName, "error", err)
			return v.retries.deletePodRetries.Wait(), err
		}
		v.retries.deletePodRetries.Initialize()

		// set deceased Pod
		v.deceasedPod = podName
	}

	// process the next pvc
	return true, errors.New("process the next pvc of the list")
}

func (c *hzOpsReqController) VolumeExpansionOffline() error {
	if podNames := c.getPodsName(); len(podNames) > 0 {
		if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.VolumeExpansionSucceeded) {
			c.RunParallel(opsapi.VolumeExpansionSucceeded, "successfully updated PVC sizes",
				c.newVolumeExpansionOfflineFunc(podNames, c.req.Spec.VolumeExpansion.Hazelcast))
			// return from here, process rest in the next cycles
			return nil
		}
	}

	return nil
}

func (v *volumeExpansionOfflineFunc) isHazelcastRunning() (bool, error) {
	_, err := hazelcastgo.NewKubeDBClientBuilder(v.KBClient, v.db).WithContext(context.TODO()).GetHazelcastClient()
	if err != nil {
		v.log.Error(err, "Failed to Create hazelcast")
		return false, err
	}

	return true, nil
}

// IsPVCConditionTrue returns "true" if the desired condition is in true state.
// It returns "false" if the desired condition is not in "true" state or is not in the condition list.
func IsPVCConditionTrue(conditions []core.PersistentVolumeClaimCondition, condType core.PersistentVolumeClaimConditionType) bool {
	for i := range conditions {
		if conditions[i].Type == condType && conditions[i].Status == core.ConditionTrue {
			return true
		}
	}
	return false
}
