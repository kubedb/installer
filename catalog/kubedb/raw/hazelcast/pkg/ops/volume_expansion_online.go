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
	"errors"
	"fmt"

	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"

	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	cutil "kmodules.xyz/client-go/conditions"
	core_util "kmodules.xyz/client-go/core/v1"
)

type volumeExpansionOnlineFunc struct {
	*hzOpsReqController
	storageReq *resource.Quantity
	retries    volumeExpansionOnlineFuncRetries
	pvcsName   []string
	infantPVC  string
}
type volumeExpansionOnlineFuncRetries struct {
	getPVCRetries         *Retries
	compareStorageRetries *Retries
	patchPVCRetries       *Retries
}

func (c *hzOpsReqController) newVolumeExpansionOnlineFunc(storageReq *resource.Quantity, pvcsName []string) func() (bool, error) {
	opts := &volumeExpansionOnlineFunc{
		hzOpsReqController: c,
		storageReq:         storageReq,
		retries:            volumeExpansionOnlineFuncRetries{},
		pvcsName:           pvcsName,
	}

	opts.retries.getPVCRetries = c.newRetries("GetPvc")
	opts.retries.compareStorageRetries = c.newRetries("CompareStorage")
	opts.retries.patchPVCRetries = c.newRetries("PatchPvc")

	return opts.run
}

func (v *volumeExpansionOnlineFunc) run() (bool, error) {
	if v.pvcsName == nil {
		return false, nil
	}

	if v.infantPVC != "" {
		// Get PVC
		pvc, err := v.Client.CoreV1().PersistentVolumeClaims(v.db.Namespace).Get(context.TODO(), v.infantPVC, meta.GetOptions{})
		if err != nil {
			// retry
			return v.retries.getPVCRetries.Wait(), err
		}
		v.retries.getPVCRetries.Initialize()

		// quantity.Cmp(y) returns 0 if the quantity is equal to y, -1 if the quantity is less than y, or 1 if the
		// quantity is greater than y.
		ck := pvc.Status.Capacity.Storage().Cmp(*v.storageReq)
		if ck == -1 {
			return v.retries.compareStorageRetries.Wait(), fmt.Errorf("pvc: %s is not yet expanded", v.infantPVC)
		}
		v.retries.compareStorageRetries.Initialize()

		v.log.Info("PVC successfully expanded", "PVCName", v.infantPVC)
		// reset infantPVC
		v.infantPVC = ""
		// delete first item of the list
		if len(v.pvcsName) > 1 {
			v.pvcsName = v.pvcsName[1:]
		} else {
			v.pvcsName = nil
		}

		// retry
		return true, nil
	} else {

		pvcName := v.pvcsName[0]
		// Get PVC
		pvc, err := v.Client.CoreV1().PersistentVolumeClaims(v.db.Namespace).Get(context.TODO(), pvcName, meta.GetOptions{})
		if err != nil {
			if kerr.IsNotFound(err) {
				return false, fmt.Errorf("pvc: %s doesn't exist", pvcName)
			}
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
			}, meta.PatchOptions{})
			if err != nil {
				return v.retries.patchPVCRetries.Wait(), err
			}
			v.retries.patchPVCRetries.Initialize()
		}

		// set infant PVC
		v.infantPVC = pvcName
	}

	// process the next pvc
	return true, errors.New("process the next pvc of the list")
}

// VolumeExpansionOnline
// - List PVC
// - Patch PVC with the updated volume size and wait for it to become updated!
func (c *hzOpsReqController) VolumeExpansionOnline() error {
	// Update
	if pvcNames := c.getPVCNames(); len(pvcNames) > 0 {
		if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.VolumeExpansionSucceeded) {
			c.RunParallel(opsapi.VolumeExpansionSucceeded, "successfully updated PVC sizes",
				c.newVolumeExpansionOnlineFunc(c.req.Spec.VolumeExpansion.Hazelcast, pvcNames))
			// return from here, process rest in the next cycles
			return nil
		}
	}
	return nil
}
