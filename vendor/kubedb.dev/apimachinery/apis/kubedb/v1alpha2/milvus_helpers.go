package v1alpha2

import (
	"context"
	"fmt"

	"gomodules.xyz/pointer"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"kmodules.xyz/client-go/apiextensions"
	coreutil "kmodules.xyz/client-go/core/v1"
	meta_util "kmodules.xyz/client-go/meta"
	"kmodules.xyz/client-go/policy/secomp"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	ofstv2 "kmodules.xyz/offshoot-api/api/v2"
	"kubedb.dev/apimachinery/apis"
	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	"kubedb.dev/apimachinery/crds"
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
	scheme := "localhost"
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
	if m.Spec.Topology.Standalone.AuthSecret != nil && m.Spec.Topology.Standalone.AuthSecret.Name != "" {
		return m.Spec.Topology.Standalone.AuthSecret.Name
	}
	return m.DefaultUserCredSecretName()
}

func (m *Milvus) ConfigSecretName() string {
	return meta_util.NameWithSuffix(m.OffshootName(), "config")
}

func (m *Milvus) GetPersistentSecrets() []string {
	var secrets []string
	if m.Spec.Topology.Standalone.AuthSecret != nil {
		secrets = append(secrets, m.GetAuthSecretName())
		secrets = append(secrets, m.ConfigSecretName())
	}
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
	if m.Spec.Topology.Standalone != nil && m.Spec.Topology.Standalone.PodTemplate.Labels != nil {
		podTemplateLabels = m.Spec.Topology.Standalone.PodTemplate.Labels
	}
	return m.offshootLabels(meta_util.OverwriteKeys(m.OffshootSelectors(), extraLabels...), podTemplateLabels)
}

func (m *Milvus) GetGRPCAddress() string {
	port := kubedb.MilvusGrpcPort
	if m.Spec.Topology.Standalone.GRPCPort != nil {
		port = *m.Spec.Topology.Standalone.GRPCPort
	}
	return fmt.Sprintf("localhost:%d", port)
}

func (m *Milvus) getAuthSecret(ctx context.Context, kc client.Client) (*core.Secret, error) {
	if m.Spec.Topology.Standalone.DisableSecurity || m.Spec.Topology.Standalone.AuthSecret == nil {
		return nil, nil
	}
	if m.Spec.Topology.Standalone.AuthSecret.Name == "" {
		return nil, nil
	}

	secret := &core.Secret{}
	err := kc.Get(ctx, types.NamespacedName{
		Name:      m.Spec.Topology.Standalone.AuthSecret.Name,
		Namespace: m.Namespace,
	}, secret)
	return secret, err
}

func (m *Milvus) GetUsername(ctx context.Context, kc client.Client) (string, error) {
	secret, err := m.getAuthSecret(ctx, kc)
	if err != nil {
		return "", err
	}
	if secret == nil {
		return "", nil // security disabled
	}

	data, ok := secret.Data[kubedb.MilvusUsernameKey]
	if !ok || len(data) == 0 {
		return "", fmt.Errorf("username key %q missing in secret %s", kubedb.MilvusUsernameKey, secret.Name)
	}
	return string(data), nil
}

func (m *Milvus) GetPassword(ctx context.Context, kc client.Client) (string, error) {
	secret, err := m.getAuthSecret(ctx, kc)
	if err != nil {
		return "", err
	}
	if secret == nil {
		return "", nil
	}

	data, ok := secret.Data[kubedb.MilvusPasswordKey]
	if !ok || len(data) == 0 {
		return "", fmt.Errorf("password key %q missing in secret %s", kubedb.MilvusPasswordKey, secret.Name)
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

func (m *Milvus) DefaultUserCredSecretName() string {
	return meta_util.NameWithSuffix(m.OffshootName(), "auth")
}

func (m *Milvus) EtcdEndpoints() []string {
	if m.Spec.Etcd == nil {
		return nil
	}

	if m.Spec.Etcd.ExternallyManaged {
		if len(m.Spec.Etcd.Endpoints) == 0 {
			fmt.Println("Warning: Etcd is externally managed but no endpoints are provided")
			return nil
		}
		// Return user-provided endpoints
		return m.Spec.Etcd.Endpoints
	}

	size := 3
	if m.Spec.Etcd.Size > 0 {
		size = m.Spec.Etcd.Size
	}

	endpoints := make([]string, size)
	for i := 0; i < size; i++ {
		// Use pod DNS names for the etcd cluster
		endpoints[i] = fmt.Sprintf(
			"http://%s-%d.%s.%s.svc.cluster.local:%d",
			kubedb.EtcdPodTemplateSuffix, i,
			kubedb.EtcdServiceSuffix, m.Namespace,
			kubedb.EtcdPort,
		)
	}

	return endpoints
}

func (m *Milvus) SetDefaults(kc client.Client) {
	if m.Spec.Topology.Standalone.Replicas == nil {
		m.Spec.Topology.Standalone.Replicas = pointer.Int32P(1)
	}

	if m.Spec.DeletionPolicy == "" {
		m.Spec.DeletionPolicy = DeletionPolicyDelete
	}

	if m.Spec.Topology.Standalone.StorageType == "" {
		m.Spec.Topology.Standalone.StorageType = StorageTypeDurable
	}

	if m.Spec.Topology.Standalone.AuthSecret == nil {
		m.Spec.Topology.Standalone.AuthSecret = &SecretReference{}
	}

	if m.Spec.Topology.Standalone.AuthSecret.Kind == "" {
		m.Spec.Topology.Standalone.AuthSecret.Kind = kubedb.ResourceKindSecret
	}

	if m.Spec.Topology.Standalone.PodTemplate == nil {
		m.Spec.Topology.Standalone.PodTemplate = &ofstv2.PodTemplateSpec{}
	}

	var mlvVersion catalog.MilvusVersion
	err := kc.Get(context.TODO(), types.NamespacedName{
		Name: m.Spec.Version,
	}, &mlvVersion)
	if err != nil {
		return
	}

	m.setDefaultContainerSecurityContext(&mlvVersion, m.Spec.Topology.Standalone.PodTemplate)

	dbContainer := coreutil.GetContainerByName(m.Spec.Topology.Standalone.PodTemplate.Spec.Containers, kubedb.MilvusContainerName)
	if dbContainer != nil && (dbContainer.Resources.Requests == nil || dbContainer.Resources.Limits == nil) {
		apis.SetDefaultResourceLimits(&dbContainer.Resources, kubedb.DefaultResources)
	}

	m.SetHealthCheckerDefaults()

	m.setDefaultContainerResourceLimits(m.Spec.Topology.Standalone.PodTemplate)
}

func GetDefaultReadinessProbe() *core.Probe {
	return &core.Probe{
		ProbeHandler: core.ProbeHandler{
			GRPC: &core.GRPCAction{Port: 19530},
		},
		InitialDelaySeconds: 60,
		PeriodSeconds:       10,
		TimeoutSeconds:      5,
		FailureThreshold:    18,
	}
}

func (m *Milvus) setDefaultContainerSecurityContext(mlvVersion *catalog.MilvusVersion, podTemplate *ofstv2.PodTemplateSpec) {
	if podTemplate == nil {
		return
	}
	if podTemplate.Spec.SecurityContext == nil {
		podTemplate.Spec.SecurityContext = &core.PodSecurityContext{}
	}
	if podTemplate.Spec.SecurityContext.FSGroup == nil {
		podTemplate.Spec.SecurityContext.FSGroup = mlvVersion.Spec.SecurityContext.RunAsUser
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
	m.AssignDefaultContainerSecurityContext(mlvVersion, container.SecurityContext)
	podTemplate.Spec.Containers = coreutil.UpsertContainer(podTemplate.Spec.Containers, *container)

	initContainer := coreutil.GetContainerByName(podTemplate.Spec.InitContainers, kubedb.MilvusInitContainerName)
	if initContainer == nil {
		initContainer = &core.Container{
			Name: kubedb.MilvusInitContainerName,
		}
	}
	if initContainer.SecurityContext == nil {
		initContainer.SecurityContext = &core.SecurityContext{}
	}
	m.AssignDefaultInitContainerSecurityContext(mlvVersion, initContainer.SecurityContext)
	podTemplate.Spec.InitContainers = coreutil.UpsertContainer(podTemplate.Spec.InitContainers, *initContainer)
}

func (m *Milvus) AssignDefaultInitContainerSecurityContext(mlvVersion *catalog.MilvusVersion, rc *core.SecurityContext) {
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
		rc.RunAsUser = mlvVersion.Spec.SecurityContext.RunAsUser
	}
	if rc.SeccompProfile == nil {
		rc.SeccompProfile = secomp.DefaultSeccompProfile()
	}
	rc.RunAsUser = pointer.Int64P(0)
	rc.RunAsNonRoot = pointer.BoolP(false)
	rc.RunAsGroup = pointer.Int64P(0)
}

func (m *Milvus) AssignDefaultContainerSecurityContext(mlvVersion *catalog.MilvusVersion, rc *core.SecurityContext) {
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
		rc.RunAsUser = mlvVersion.Spec.SecurityContext.RunAsUser
	}
	if rc.RunAsGroup == nil {
		rc.RunAsGroup = mlvVersion.Spec.SecurityContext.RunAsUser
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
