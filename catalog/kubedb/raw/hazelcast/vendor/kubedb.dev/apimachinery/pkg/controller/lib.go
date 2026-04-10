/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

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

	_ "gomodules.xyz/stow/azure"
	_ "gomodules.xyz/stow/google"
	_ "gomodules.xyz/stow/s3"
	apps "k8s.io/api/apps/v1"
	policy "k8s.io/api/policy/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	core_util "kmodules.xyz/client-go/core/v1"
	policy_util "kmodules.xyz/client-go/policy"
	psapi "kubeops.dev/petset/apis/apps/v1"
)

// SyncPetSetPodDisruptionBudget syncs the PDB with the current state of the petSet.
// The maxUnavailable is calculated based petSet replica count, maxUnavailable = (replicas-1)/2.
// Also cleanup the PDB, when replica count is 1 or less.
func (c *Controller) SyncPetSetPodDisruptionBudget(ps *psapi.PetSet) error {
	if ps == nil {
		return nil
	}
	// CleanUp PDB for petSet with replica 1
	if *ps.Spec.Replicas <= 1 {
		// pdb name & namespace is same as the corresponding statefulSet's name & namespace.
		err := policy_util.DeletePodDisruptionBudget(context.TODO(), c.Client, types.NamespacedName{
			Namespace: ps.Namespace,
			Name:      ps.Name,
		})
		if !kerr.IsNotFound(err) {
			return err
		}
		return nil
	}
	return c.SyncPetSetPDBWithCustomLabelSelectors(ps, *ps.Spec.Replicas, ps.Labels, ps.Spec.Selector.MatchLabels)
}

func (c *Controller) SyncPetSetPDBWithCustomLabelSelectors(ps *psapi.PetSet, replicas int32, labels map[string]string, matchLabelSelectors map[string]string) error {
	if ps == nil {
		return nil
	}
	pdbRef := metav1.ObjectMeta{
		Name:      ps.Name,
		Namespace: ps.Namespace,
	}

	r := int32(math.Max(1, math.Floor(float64(replicas-1)/2.0)))
	maxUnavailable := &intstr.IntOrString{IntVal: r}

	owner := metav1.NewControllerRef(ps, psapi.SchemeGroupVersion.WithKind("PetSet"))
	_, _, err := policy_util.CreateOrPatchPodDisruptionBudget(context.TODO(), c.Client, pdbRef,
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

func (c *Controller) CreateDeploymentPodDisruptionBudget(deployment *apps.Deployment) error {
	owner := metav1.NewControllerRef(deployment, apps.SchemeGroupVersion.WithKind("Deployment"))

	m := metav1.ObjectMeta{
		Name:      deployment.Name,
		Namespace: deployment.Namespace,
	}

	_, _, err := policy_util.CreateOrPatchPodDisruptionBudget(context.TODO(), c.Client, m,
		func(in *policy.PodDisruptionBudget) *policy.PodDisruptionBudget {
			in.Labels = deployment.Labels
			core_util.EnsureOwnerReference(&in.ObjectMeta, owner)

			in.Spec.Selector = &metav1.LabelSelector{
				MatchLabels: deployment.Spec.Template.Labels,
			}

			in.Spec.MaxUnavailable = nil

			in.Spec.MinAvailable = &intstr.IntOrString{IntVal: 1}
			return in
		}, metav1.PatchOptions{})
	return err
}
