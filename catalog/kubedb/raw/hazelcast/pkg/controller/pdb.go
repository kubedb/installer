/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"math"

	apps "k8s.io/api/apps/v1"
	policy "k8s.io/api/policy/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	core_util "kmodules.xyz/client-go/core/v1"
	policy_util "kmodules.xyz/client-go/policy"
)

func (rs *ReconcileState) SyncStatefulSetSetPodDisruptionBudget(sts *apps.StatefulSet) error {
	if sts == nil {
		return nil
	}
	// CleanUp PDB for petSet with replica 1
	if *sts.Spec.Replicas <= 1 {
		// pdb name & namespace is same as the corresponding statefulSet's name & namespace.
		err := policy_util.DeletePodDisruptionBudget(context.TODO(), rs.Client, types.NamespacedName{
			Namespace: sts.Namespace,
			Name:      sts.Name,
		})
		if !kerr.IsNotFound(err) {
			return err
		}
		return nil
	}
	return rs.SyncStatefulSetPDBWithCustomLabelSelectors(sts, *sts.Spec.Replicas, sts.Labels, sts.Spec.Selector.MatchLabels)
}

func (rs *ReconcileState) SyncStatefulSetPDBWithCustomLabelSelectors(sts *apps.StatefulSet, replicas int32, labels map[string]string, matchLabelSelectors map[string]string) error {
	if sts == nil {
		return nil
	}
	pdbRef := metav1.ObjectMeta{
		Name:      sts.Name,
		Namespace: sts.Namespace,
	}

	r := int32(math.Max(1, math.Floor(float64(replicas-1)/2.0)))
	maxUnavailable := &intstr.IntOrString{IntVal: r}

	owner := metav1.NewControllerRef(sts, apps.SchemeGroupVersion.WithKind("StatefulSet"))
	_, _, err := policy_util.CreateOrPatchPodDisruptionBudget(context.TODO(), rs.Client, pdbRef,
		func(in *policy.PodDisruptionBudget) *policy.PodDisruptionBudget {
			in.Labels = labels
			core_util.EnsureOwnerReference(&in.ObjectMeta, owner)
			in.Spec.Selector = &metav1.LabelSelector{
				MatchLabels: matchLabelSelectors,
			}
			in.Spec.MaxUnavailable = maxUnavailable
			in.Spec.MinAvailable = nil
			return in
		}, metav1.PatchOptions{})
	return err
}
