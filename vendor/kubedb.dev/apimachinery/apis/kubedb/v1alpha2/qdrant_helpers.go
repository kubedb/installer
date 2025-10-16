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

package v1alpha2

import (
	"context"
	"fmt"

	"gomodules.xyz/pointer"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"kmodules.xyz/client-go/apiextensions"
	coreutil "kmodules.xyz/client-go/core/v1"
	meta_util "kmodules.xyz/client-go/meta"
	"kmodules.xyz/client-go/policy/secomp"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	ofst "kmodules.xyz/offshoot-api/api/v2"
	"kubedb.dev/apimachinery/apis"
	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	"kubedb.dev/apimachinery/crds"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (q *Qdrant) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourcePluralQdrant))
}

func (q *Qdrant) AsOwner() *meta.OwnerReference {
	return meta.NewControllerRef(q, SchemeGroupVersion.WithKind(ResourceKindQdrant))
}

func (q *Qdrant) ResourceKind() string {
	return ResourceKindQdrant
}

func (q *Qdrant) ResourceSingular() string {
	return ResourceSingularQdrant
}

func (q *Qdrant) ResourcePlural() string {
	return ResourcePluralQdrant
}

func (q *Qdrant) Finalizer() string {
	return fmt.Sprintf("%s/%s", apis.Finalizer, q.ResourceSingular())
}

func (q *Qdrant) ResourceFQN() string {
	return fmt.Sprintf("%s.%s", q.ResourcePlural(), kubedb.GroupName)
}

func (q *Qdrant) OffshootSelectors(extraSelectors ...map[string]string) map[string]string {
	selector := map[string]string{
		meta_util.NameLabelKey:      q.ResourceFQN(),
		meta_util.InstanceLabelKey:  q.Name,
		meta_util.ManagedByLabelKey: kubedb.GroupName,
	}
	return meta_util.OverwriteKeys(selector, extraSelectors...)
}

func (q *Qdrant) OffshootName() string {
	return q.Name
}

func (q *Qdrant) GetAuthSecretName() string {
	if q.Spec.AuthSecret != nil && q.Spec.AuthSecret.Name != "" {
		return q.Spec.AuthSecret.Name
	}
	return q.DefaultAuthSecretName()
}

func (q *Qdrant) GetPersistentSecrets() []string {
	var secrets []string
	if q.Spec.AuthSecret != nil {
		secrets = append(secrets, q.GetAuthSecretName())
	}
	return secrets
}

// Owner returns owner reference to resources
func (q *Qdrant) Owner() *meta.OwnerReference {
	return meta.NewControllerRef(q, SchemeGroupVersion.WithKind(q.ResourceKind()))
}

func (q *Qdrant) SetDefaults(kc client.Client) {
	if q.Spec.Replicas == nil {
		q.Spec.Replicas = pointer.Int32P(1)
	}

	if q.Spec.DeletionPolicy == "" {
		q.Spec.DeletionPolicy = DeletionPolicyDelete
	}

	if q.Spec.StorageType == "" {
		q.Spec.StorageType = StorageTypeDurable
	}

	var qdVersion catalog.QdrantVersion
	err := kc.Get(context.TODO(), types.NamespacedName{
		Name: q.Spec.Version,
	}, &qdVersion)
	if err != nil {
		return
	}

	q.setDefaultContainerSecurityContext(&qdVersion, &q.Spec.PodTemplate)

	dbContainer := coreutil.GetContainerByName(q.Spec.PodTemplate.Spec.Containers, kubedb.QdrantContainerName)
	if dbContainer != nil && (dbContainer.Resources.Requests == nil || dbContainer.Resources.Limits == nil) {
		apis.SetDefaultResourceLimits(&dbContainer.Resources, kubedb.DefaultResources)
	}

	q.SetHealthCheckerDefaults()
}

func (q *Qdrant) SetHealthCheckerDefaults() {
	if q.Spec.HealthChecker.PeriodSeconds == nil {
		q.Spec.HealthChecker.PeriodSeconds = pointer.Int32P(10)
	}
	if q.Spec.HealthChecker.TimeoutSeconds == nil {
		q.Spec.HealthChecker.TimeoutSeconds = pointer.Int32P(10)
	}
	if q.Spec.HealthChecker.FailureThreshold == nil {
		q.Spec.HealthChecker.FailureThreshold = pointer.Int32P(3)
	}
}

func (q *Qdrant) setDefaultContainerSecurityContext(qdVersion *catalog.QdrantVersion, podTemplate *ofst.PodTemplateSpec) {
	if podTemplate == nil {
		return
	}
	if podTemplate.Spec.SecurityContext == nil {
		podTemplate.Spec.SecurityContext = &core.PodSecurityContext{}
	}
	if podTemplate.Spec.SecurityContext.FSGroup == nil {
		podTemplate.Spec.SecurityContext.FSGroup = qdVersion.Spec.SecurityContext.RunAsUser
	}

	container := coreutil.GetContainerByName(podTemplate.Spec.Containers, kubedb.QdrantContainerName)
	if container == nil {
		container = &core.Container{
			Name: kubedb.QdrantContainerName,
		}
		podTemplate.Spec.Containers = coreutil.UpsertContainer(podTemplate.Spec.Containers, *container)
	}

	if container.SecurityContext == nil {
		container.SecurityContext = &core.SecurityContext{}
	}
	q.assignDefaultContainerSecurityContext(qdVersion, container.SecurityContext)

	initContainer := coreutil.GetContainerByName(podTemplate.Spec.InitContainers, kubedb.QdrantInitContainerName)
	if initContainer == nil {
		initContainer = &core.Container{
			Name: kubedb.QdrantInitContainerName,
		}
		podTemplate.Spec.InitContainers = coreutil.UpsertContainer(podTemplate.Spec.InitContainers, *initContainer)
	}
	if initContainer.SecurityContext == nil {
		initContainer.SecurityContext = &core.SecurityContext{}
	}
	q.assignDefaultContainerSecurityContext(qdVersion, initContainer.SecurityContext)
}

func (q *Qdrant) assignDefaultContainerSecurityContext(qdVersion *catalog.QdrantVersion, rc *core.SecurityContext) {
	if rc.AllowPrivilegeEscalation == nil {
		rc.AllowPrivilegeEscalation = pointer.BoolP(false)
	}
	if rc.Capabilities == nil {
		rc.Capabilities = &core.Capabilities{
			Drop: []core.Capability{"ALL"},
		}
	}
	if rc.RunAsNonRoot == nil {
		rc.RunAsNonRoot = pointer.BoolP(true)
	}
	if rc.RunAsUser == nil {
		rc.RunAsUser = qdVersion.Spec.SecurityContext.RunAsUser
	}
	if rc.SeccompProfile == nil {
		rc.SeccompProfile = secomp.DefaultSeccompProfile()
	}
}

func (q *Qdrant) PetSetName() string {
	return q.OffshootName()
}

func (q *Qdrant) ServiceName() string { return q.OffshootName() }

func (q *Qdrant) offshootLabels(selector, override map[string]string) map[string]string {
	selector[meta_util.ComponentLabelKey] = kubedb.ComponentDatabase
	return meta_util.FilterKeys(kubedb.GroupName, selector, meta_util.OverwriteKeys(nil, q.Labels, override))
}

func (q *Qdrant) OffshootLabels() map[string]string {
	return q.offshootLabels(q.OffshootSelectors(), nil)
}

func (q *Qdrant) GoverningServiceName() string {
	return meta_util.NameWithSuffix(q.ServiceName(), "pods")
}

func (q *Qdrant) DefaultAuthSecretName() string {
	return meta_util.NameWithSuffix(q.OffshootName(), "auth")
}

func (q *Qdrant) ServiceAccountName() string {
	return q.OffshootName()
}

func (q *Qdrant) DefaultPodRoleName() string {
	return meta_util.NameWithSuffix(q.OffshootName(), "role")
}

func (q *Qdrant) DefaultPodRoleBindingName() string {
	return meta_util.NameWithSuffix(q.OffshootName(), "rolebinding")
}

type QdrantApp struct {
	*Qdrant
}

func (q *QdrantApp) Name() string {
	return q.Qdrant.Name
}

func (q QdrantApp) Type() appcat.AppType {
	return appcat.AppType(fmt.Sprintf("%s/%s", kubedb.GroupName, ResourceSingularQdrant))
}

func (q *Qdrant) AppBindingMeta() appcat.AppBindingMeta {
	return &QdrantApp{q}
}

func (q *Qdrant) GetConnectionScheme() string {
	scheme := "http"
	return scheme
}

func (q *Qdrant) PodLabels(extraLabels ...map[string]string) map[string]string {
	return q.offshootLabels(meta_util.OverwriteKeys(q.OffshootSelectors(), extraLabels...), q.Spec.PodTemplate.Labels)
}

func (q *Qdrant) ConfigSecretName() string {
	return meta_util.NameWithSuffix(q.OffshootName(), "config")
}

func (q *Qdrant) PVCName(alias string) string {
	return alias
}

func (q *Qdrant) Address() string {
	return fmt.Sprintf("%v.%v.svc", q.Name, q.Namespace)
}
