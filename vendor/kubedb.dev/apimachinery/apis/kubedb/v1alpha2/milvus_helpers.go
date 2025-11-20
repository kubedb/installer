package v1alpha2

import (
	"context"
	"fmt"

	"kubedb.dev/apimachinery/apis"
	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	"kubedb.dev/apimachinery/crds"

	"gomodules.xyz/pointer"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"kmodules.xyz/client-go/apiextensions"
	coreutil "kmodules.xyz/client-go/core/v1"
	meta_util "kmodules.xyz/client-go/meta"
	"kmodules.xyz/client-go/policy/secomp"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	ofstv2 "kmodules.xyz/offshoot-api/api/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type MilvusApp struct {
	*Milvus
}

func (_ Milvus) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourcePluralMilvus))
}

func (m *Milvus) ResourceKind() string {
	return ResourceKindMilvus
}

func (m *Milvus) ResourceSingular() string {
	return ResourceSingularMilvus
}

func (m *Milvus) ResourcePlural() string {
	return ResourcePluralMilvus
}

func (m *Milvus) ResourceFQN() string {
	return fmt.Sprintf("%s.%s", m.ResourcePlural(), kubedb.GroupName)
}

func (m *Milvus) AppBindingMeta() appcat.AppBindingMeta {
	return &MilvusApp{m}
}

func (r MilvusApp) Name() string {
	return r.Milvus.Name
}

func (m Milvus) Type() appcat.AppType {
	return appcat.AppType(fmt.Sprintf("%s/%s", kubedb.GroupName, m.ResourceSingular()))
}

func (m *Milvus) GetConnectionScheme() string {
	scheme := "http"
	return scheme
}

func (m *Milvus) Owner() *metav1.OwnerReference {
	return metav1.NewControllerRef(m, SchemeGroupVersion.WithKind(m.ResourceKind()))
}

func (m *Milvus) OffshootName() string {
	return m.Name
}

func (m *Milvus) ServiceName() string {
	return m.OffshootName()
}

func (m *Milvus) GoverningServiceName() string {
	return meta_util.NameWithSuffix(m.ServiceName(), "pods")
}

func (m *Milvus) PetSetName() string {
	return m.OffshootName()
}

func (m *Milvus) ServiceAccountName() string {
	return m.OffshootName()
}

func (m *Milvus) GetAuthSecretName() string {
	if m.Spec.AuthSecret != nil && m.Spec.AuthSecret.Name != "" {
		return m.Spec.AuthSecret.Name
	}
	return meta_util.NameWithSuffix(m.OffshootName(), "auth")
}

func (m *Milvus) ConfigSecretName() string {
	return meta_util.NameWithSuffix(m.OffshootName(), "config")
}

func (m *Milvus) GetPersistentSecrets() []string {
	var secrets []string
	if m.Spec.AuthSecret != nil {
		secrets = append(secrets, m.GetAuthSecretName())
	}
	secrets = append(secrets, m.ConfigSecretName())
	return secrets
}

func (m *Milvus) offshootLabels(selector, override map[string]string) map[string]string {
	selector[meta_util.ComponentLabelKey] = kubedb.ComponentDatabase
	return meta_util.FilterKeys(kubedb.GroupName, selector, meta_util.OverwriteKeys(nil, m.Labels, override))
}

func (m *Milvus) OffshootLabels() map[string]string {
	return m.offshootLabels(m.OffshootSelectors(), nil)
}

func (m *Milvus) OffshootSelectors(extraSelectors ...map[string]string) map[string]string {
	selector := map[string]string{
		meta_util.NameLabelKey:      m.ResourceFQN(),
		meta_util.InstanceLabelKey:  m.Name,
		meta_util.ManagedByLabelKey: kubedb.GroupName,
	}
	return meta_util.OverwriteKeys(selector, extraSelectors...)
}

func (m *Milvus) PodLabels(extraLabels ...map[string]string) map[string]string {
	var podTemplateLabels map[string]string
	if m.Spec.PodTemplate.Labels != nil {
		podTemplateLabels = m.Spec.PodTemplate.Labels
	}
	return m.offshootLabels(meta_util.OverwriteKeys(m.OffshootSelectors(), extraLabels...), podTemplateLabels)
}

func (m *Milvus) ServiceDNS() string {
	return fmt.Sprintf("%s.%s.svc.cluster.local:%d", m.ServiceName(), m.Namespace, kubedb.MilvusGrpcPort)
}

func (m *Milvus) getAuthSecret(ctx context.Context, kc client.Client) (*core.Secret, error) {
	secret := &core.Secret{}
	err := kc.Get(ctx, types.NamespacedName{
		Name:      m.Spec.AuthSecret.Name,
		Namespace: m.Namespace,
	}, secret)
	return secret, err
}

func (m *Milvus) GetUsername(ctx context.Context, kc client.Client) (string, error) {
	secret, _ := m.getAuthSecret(ctx, kc)
	data, ok := secret.Data[core.BasicAuthUsernameKey]
	if !ok || len(data) == 0 {
		return "", fmt.Errorf("username key %q missing in secret %s", core.BasicAuthUsernameKey, secret.Name)
	}
	return string(data), nil
}

func (m *Milvus) GetPassword(ctx context.Context, kc client.Client) (string, error) {
	secret, _ := m.getAuthSecret(ctx, kc)
	data, ok := secret.Data[core.BasicAuthPasswordKey]
	if !ok || len(data) == 0 {
		return "", fmt.Errorf("password key %q missing in secret %s", core.BasicAuthPasswordKey, secret.Name)
	}
	return string(data), nil
}

func (m *Milvus) SetHealthCheckerDefaults() {
	if m.Spec.HealthChecker.PeriodSeconds == nil {
		m.Spec.HealthChecker.PeriodSeconds = pointer.Int32P(10)
	}
	if m.Spec.HealthChecker.TimeoutSeconds == nil {
		m.Spec.HealthChecker.TimeoutSeconds = pointer.Int32P(10)
	}
	if m.Spec.HealthChecker.FailureThreshold == nil {
		m.Spec.HealthChecker.FailureThreshold = pointer.Int32P(1)
	}
}

func (m *Milvus) EtcdServiceName() string {
	return fmt.Sprintf("%s-%s", m.Namespace, kubedb.EtcdName)
}

func (m *Milvus) MetaStorageEndpoints() []string {
	if m.Spec.MetaStorage.ExternallyManaged {
		if len(m.Spec.MetaStorage.Endpoints) == 0 {
			klog.Errorf("metadata storage is externally managed but no endpoints were provided")
			return []string{}
		}
		return m.Spec.MetaStorage.Endpoints
	}

	size := m.Spec.MetaStorage.Size

	endpoints := make([]string, size)
	for i := 0; i < size; i++ {
		// Use pod DNS names for the etcd cluster
		endpoints[i] = fmt.Sprintf(
			"http://%s-%d.%s.%s.svc.cluster.local:%d",
			m.EtcdServiceName(), i,
			m.EtcdServiceName(), m.Namespace,
			2379,
		)
	}

	return endpoints
}

func (m *Milvus) SetDefaults(kc client.Client) {
	if m.Spec.Topology.Mode == nil {
		mode := MilvusMode("Standalone")
		m.Spec.Topology.Mode = &mode
	}

	if m.Spec.DeletionPolicy == "" {
		m.Spec.DeletionPolicy = DeletionPolicyDelete
	}

	if m.Spec.StorageType == "" {
		m.Spec.StorageType = StorageTypeDurable
	}

	if m.Spec.AuthSecret == nil {
		m.Spec.AuthSecret = &SecretReference{}
	}

	if m.Spec.AuthSecret.Kind == "" {
		m.Spec.AuthSecret.Kind = kubedb.ResourceKindSecret
	}

	if m.Spec.PodTemplate == nil {
		m.Spec.PodTemplate = &ofstv2.PodTemplateSpec{}
	}

	var mvVersion catalog.MilvusVersion
	err := kc.Get(context.TODO(), types.NamespacedName{
		Name: m.Spec.Version,
	}, &mvVersion)
	if err != nil {
		return
	}

	m.setDefaultContainerSecurityContext(&mvVersion, m.Spec.PodTemplate)

	m.SetHealthCheckerDefaults()

	m.setDefaultContainerResourceLimits(m.Spec.PodTemplate)
}

func (m *Milvus) setDefaultContainerSecurityContext(mvVersion *catalog.MilvusVersion, podTemplate *ofstv2.PodTemplateSpec) {
	if podTemplate == nil {
		return
	}
	if podTemplate.Spec.SecurityContext == nil {
		podTemplate.Spec.SecurityContext = &core.PodSecurityContext{}
	}
	if podTemplate.Spec.SecurityContext.FSGroup == nil {
		podTemplate.Spec.SecurityContext.FSGroup = mvVersion.Spec.SecurityContext.RunAsUser
	}

	container := coreutil.GetContainerByName(podTemplate.Spec.Containers, kubedb.MilvusContainerName)
	if container == nil {
		container = &core.Container{
			Name: kubedb.MilvusContainerName,
		}
	}
	if container.SecurityContext == nil {
		container.SecurityContext = &core.SecurityContext{}
	}
	m.AssignDefaultContainerSecurityContext(mvVersion, container.SecurityContext)
	podTemplate.Spec.Containers = coreutil.UpsertContainer(podTemplate.Spec.Containers, *container)
}

func (m *Milvus) AssignDefaultContainerSecurityContext(mvVersion *catalog.MilvusVersion, rc *core.SecurityContext) {
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
		rc.RunAsUser = mvVersion.Spec.SecurityContext.RunAsUser
	}
	if rc.RunAsGroup == nil {
		rc.RunAsGroup = mvVersion.Spec.SecurityContext.RunAsUser
	}
	if rc.SeccompProfile == nil {
		rc.SeccompProfile = secomp.DefaultSeccompProfile()
	}
}

func (m *Milvus) setDefaultContainerResourceLimits(podTemplate *ofstv2.PodTemplateSpec) {
	dbContainer := coreutil.GetContainerByName(podTemplate.Spec.Containers, kubedb.MilvusContainerName)
	if dbContainer != nil && (dbContainer.Resources.Requests == nil && dbContainer.Resources.Limits == nil) {
		apis.SetDefaultResourceLimits(&dbContainer.Resources, kubedb.DefaultResources)
	}
}
