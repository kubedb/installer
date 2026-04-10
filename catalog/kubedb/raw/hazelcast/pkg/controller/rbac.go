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
	"fmt"

	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kutil "kmodules.xyz/client-go"
	coreutil "kmodules.xyz/client-go/core/v1"
	rbac_util "kmodules.xyz/client-go/rbac/v1"
)

func (rs *ReconcileState) createRoleBinding() error {
	// Ensure new RoleBindings
	_, _, err := rbac_util.CreateOrPatchRoleBinding(
		context.TODO(),
		rs.Client,
		metav1.ObjectMeta{
			Name:      rs.db.Name,
			Namespace: rs.db.Namespace,
		},
		func(in *rbac.RoleBinding) *rbac.RoleBinding {
			coreutil.EnsureOwnerReference(&in.ObjectMeta, rs.db.Owner())
			in.Labels = rs.db.OffshootLabels()
			in.RoleRef = rbac.RoleRef{
				APIGroup: rbac.GroupName,
				Kind:     "Role",
				Name:     rs.db.Name,
			}
			in.Subjects = []rbac.Subject{
				{
					Kind:      rbac.ServiceAccountKind,
					Name:      rs.db.Name,
					Namespace: rs.db.Namespace,
				},
			}
			return in
		},
		metav1.PatchOptions{},
	)

	return err
}

func (rs *ReconcileState) ensureRole() error {
	// Create new Roles
	_, v, err := rbac_util.CreateOrPatchRole(
		context.TODO(),
		rs.Client,
		metav1.ObjectMeta{
			Name:      rs.db.Name,
			Namespace: rs.db.Namespace,
		},
		func(in *rbac.Role) *rbac.Role {
			coreutil.EnsureOwnerReference(&in.ObjectMeta, rs.db.Owner())
			in.Labels = rs.db.OffshootLabels()
			in.Rules = []rbac.PolicyRule{}
			stsRule := rbac.PolicyRule{
				APIGroups: []string{apps.SchemeGroupVersion.Group},
				Resources: []string{"statefulsets"},
				Verbs:     []string{"get", "list", "watch"},
			}
			endRule := rbac.PolicyRule{
				APIGroups: []string{core.SchemeGroupVersion.Group},
				Resources: []string{"endpoints"},
				Verbs:     []string{"get"},
			}
			in.Rules = append(in.Rules, stsRule, endRule)

			return in
		},
		metav1.PatchOptions{},
	)
	if err != nil {
		return err
	}
	if v == kutil.VerbCreated {
		rs.log.Info(fmt.Sprintf("role: %s/%s created", rs.db.Namespace, rs.db.Name))
	}
	return nil
}

func (rs *ReconcileState) createServiceAccount() error {
	// Create new ServiceAccount
	_, _, err := coreutil.CreateOrPatchServiceAccount(
		context.TODO(),
		rs.Client,
		metav1.ObjectMeta{
			Name:      rs.db.Name,
			Namespace: rs.db.Namespace,
		},
		func(in *core.ServiceAccount) *core.ServiceAccount {
			coreutil.EnsureOwnerReference(&in.ObjectMeta, rs.db.Owner())
			in.Labels = rs.db.OffshootLabels()
			return in
		},
		metav1.PatchOptions{},
	)
	return err
}

func (rs *ReconcileState) ensureDatabaseRBAC() error {
	saName := rs.db.Spec.PodTemplate.Spec.ServiceAccountName
	if saName == "" {
		saName = rs.db.Name
		rs.db.Spec.PodTemplate.Spec.ServiceAccountName = saName
	}

	sa, err := rs.Client.CoreV1().ServiceAccounts(rs.db.Namespace).Get(context.TODO(), saName, metav1.GetOptions{})
	if kerr.IsNotFound(err) {
		// create service account, since it does not exist
		if err = rs.createServiceAccount(); err != nil {
			if !kerr.IsAlreadyExists(err) {
				return err
			}
		}
	} else if err != nil {
		return err
	} else if owned, _ := coreutil.IsOwnedBy(sa, rs.db); !owned {
		// user provided the service account, so do nothing.
		return nil
	}

	// Create New Role
	if err := rs.ensureRole(); err != nil {
		return err
	}

	// Create New RoleBinding
	if err := rs.createRoleBinding(); err != nil {
		return err
	}

	return nil
}
