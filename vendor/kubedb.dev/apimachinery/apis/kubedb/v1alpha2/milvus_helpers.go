package v1alpha2

import (
	"fmt"

	"gomodules.xyz/pointer"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ptr "k8s.io/utils/pointer"
	"kmodules.xyz/client-go/apiextensions"
	meta_util "kmodules.xyz/client-go/meta"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	"kubedb.dev/apimachinery/apis"
	"kubedb.dev/apimachinery/apis/kubedb"
	"kubedb.dev/apimachinery/crds"
)

type MilvusApp struct {
	*Milvus
}

func (_ Milvus) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourcePluralMySQL))
}

func (m *Milvus) ResourceKind() string {
	return ResourceKindMilvus
}

func (m *Milvus) ResourceSingular() string {
	return ResourceSingularMilvus
}

func (m *Milvus) Finalizer() string {
	return fmt.Sprintf("%s/%s", apis.Finalizer, m.ResourceSingular())
}

func (m *Milvus) Owner() *metav1.OwnerReference {
	return metav1.NewControllerRef(m, SchemeGroupVersion.WithKind(m.ResourceKind()))
}

func (r MilvusApp) Name() string {
	return r.Milvus.Name
}

func (m Milvus) Type() appcat.AppType {
	return appcat.AppType(fmt.Sprintf("%s/%s", kubedb.GroupName, ResourceSingularMilvus))
}

func (m *Milvus) AppBindingMeta() appcat.AppBindingMeta {
	return &MilvusApp{m}
}

func (m *Milvus) GetConnectionScheme() string {
	scheme := "localhost"
	return scheme
}

func (m *Milvus) OffshootName() string {
	return m.Name
}

func (m *Milvus) ServiceName() string {
	return m.OffshootName()
}

func (m *Milvus) DefaultPodRoleName() string {
	return meta_util.NameWithSuffix(m.OffshootName(), "role")
}

func (m *Milvus) DefaultPodRoleBindingName() string {
	return meta_util.NameWithSuffix(m.OffshootName(), "rolebinding")
}

func (m *Milvus) ServiceAccountName() string {
	return m.OffshootName()
}
func (m *Milvus) SetHealthCheckerDefaults() {
	if m.Spec.HealthChecker.PeriodSeconds == nil {
		m.Spec.HealthChecker.PeriodSeconds = pointer.Int32P(10)
	}
	if m.Spec.HealthChecker.TimeoutSeconds == nil {
		m.Spec.HealthChecker.TimeoutSeconds = pointer.Int32P(10)
	}
	if m.Spec.HealthChecker.FailureThreshold == nil {
		m.Spec.HealthChecker.FailureThreshold = pointer.Int32P(3)
	}
}

func (m *Milvus) GetAuthSecretName() string {
	if m.Spec.Standalone.AuthSecret != nil && m.Spec.Standalone.AuthSecret.Name != "" {
		return m.Spec.Standalone.AuthSecret.Name
	}
	return m.DefaultUserCredSecretName()
}

func (m *Milvus) GetPersistentSecrets() []string {
	var secrets []string
	if m.Spec.Standalone.AuthSecret != nil {
		secrets = append(secrets, m.GetAuthSecretName())
	}
	return secrets
}

func (m *Milvus) DefaultUserCredSecretName() string {
	return meta_util.NameWithSuffix(m.OffshootName(), "auth")
}

func (m *Milvus) ResourcePlural() string {
	return ResourcePluralMilvus
}

func (m *Milvus) ResourceFQN() string {
	return fmt.Sprintf("%s.%s", m.ResourcePlural(), kubedb.GroupName)
}

func (m *Milvus) offshootLabels(selector, override map[string]string) map[string]string {
	selector[meta_util.ComponentLabelKey] = kubedb.ComponentDatabase
	return meta_util.FilterKeys(kubedb.GroupName, selector, meta_util.OverwriteKeys(nil, m.Labels, override))
}

func (m *Milvus) OffshootSelectors(extraSelectors ...map[string]string) map[string]string {
	selector := map[string]string{
		meta_util.NameLabelKey:      m.ResourceFQN(),
		meta_util.InstanceLabelKey:  m.Name,
		meta_util.ManagedByLabelKey: kubedb.GroupName,
	}
	return meta_util.OverwriteKeys(selector, extraSelectors...)
}

func (m *Milvus) OffshootLabels() map[string]string {
	return m.offshootLabels(m.OffshootSelectors(), nil)
}

func (m *Milvus) GoverningServiceName() string {
	return meta_util.NameWithSuffix(m.ServiceName(), "pods")
}

func (m *Milvus) PodLabels(extraLabels ...map[string]string) map[string]string {
	var podTemplateLabels map[string]string
	if m.Spec.Standalone != nil && m.Spec.Standalone.PodTemplate.Labels != nil {
		podTemplateLabels = m.Spec.Standalone.PodTemplate.Labels
	}
	return m.offshootLabels(meta_util.OverwriteKeys(m.OffshootSelectors(), extraLabels...), podTemplateLabels)
}

func (m *Milvus) PetSetName() string {
	return m.OffshootName()
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
			"http://milvus-standalone-etcd-%d.milvus-standalone-etcd.milvus-standalone.svc.cluster.local:2379",
			i,
		)
	}

	return endpoints
}

func GetDefaultSecurityContext() *core.SecurityContext {
	return &core.SecurityContext{
		RunAsNonRoot:             ptr.Bool(true),
		RunAsUser:                ptr.Int64(10001),
		RunAsGroup:               ptr.Int64(10001),
		AllowPrivilegeEscalation: ptr.Bool(false),
		Capabilities: &core.Capabilities{
			Drop: []core.Capability{"ALL"},
		},
		SeccompProfile: &core.SeccompProfile{
			Type: core.SeccompProfileTypeRuntimeDefault,
		},
	}
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
