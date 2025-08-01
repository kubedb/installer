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
	"path/filepath"
	"strconv"
	"strings"

	"kubedb.dev/apimachinery/apis"
	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	"kubedb.dev/apimachinery/crds"

	promapi "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"gomodules.xyz/pointer"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	kmapi "kmodules.xyz/client-go/api/v1"
	"kmodules.xyz/client-go/apiextensions"
	coreutil "kmodules.xyz/client-go/core/v1"
	meta_util "kmodules.xyz/client-go/meta"
	"kmodules.xyz/client-go/policy/secomp"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	mona "kmodules.xyz/monitoring-agent-api/api/v1"
	ofst "kmodules.xyz/offshoot-api/api/v2"
	pslister "kubeops.dev/petset/client/listers/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CassandraApp struct {
	*Cassandra
}

func (r *Cassandra) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourcePluralCassandra))
}

func (r *Cassandra) AppBindingMeta() appcat.AppBindingMeta {
	return &CassandraApp{r}
}

func (r CassandraApp) Name() string {
	return r.Cassandra.Name
}

func (r CassandraApp) Type() appcat.AppType {
	return appcat.AppType(fmt.Sprintf("%s/%s", kubedb.GroupName, ResourceSingularCassandra))
}

func (c Cassandra) SidekickLabels(skName string) map[string]string {
	return meta_util.OverwriteKeys(nil, kubedb.CommonSidekickLabels(), map[string]string{
		meta_util.InstanceLabelKey: skName,
		kubedb.SidekickOwnerName:   c.Name,
		kubedb.SidekickOwnerKind:   c.ResourceFQN(),
	})
}

// Owner returns owner reference to resources
func (r *Cassandra) Owner() *meta.OwnerReference {
	return meta.NewControllerRef(r, SchemeGroupVersion.WithKind(r.ResourceKind()))
}

func (r *Cassandra) ResourceKind() string {
	return ResourceKindCassandra
}

func (r *Cassandra) OffshootName() string {
	return r.Name
}

func (r *Cassandra) OffshootRackName(value string) string {
	return meta_util.NameWithSuffix(r.OffshootName(), value)
}

func (r *Cassandra) OffshootRackPetSetName(rackName string) string {
	rack := meta_util.NameWithSuffix("rack", rackName)
	return meta_util.NameWithSuffix(r.OffshootName(), rack)
}

func (r *Cassandra) OffshootLabels() map[string]string {
	return r.offshootLabels(r.OffshootSelectors(), nil)
}

func (r *Cassandra) OffshootRackLabels(petSetName string) map[string]string {
	return r.offshootLabels(r.OffshootRackSelectors(petSetName), nil)
}

func (r *Cassandra) offshootLabels(selector, override map[string]string) map[string]string {
	selector[meta_util.ComponentLabelKey] = kubedb.ComponentDatabase
	return meta_util.FilterKeys(kubedb.GroupName, selector, meta_util.OverwriteKeys(nil, r.Labels, override))
}

func (r *Cassandra) OffshootSelectors(extraSelectors ...map[string]string) map[string]string {
	selector := map[string]string{
		meta_util.NameLabelKey:      r.ResourceFQN(),
		meta_util.InstanceLabelKey:  r.Name,
		meta_util.ManagedByLabelKey: kubedb.GroupName,
	}
	return meta_util.OverwriteKeys(selector, extraSelectors...)
}

func (r *Cassandra) OffshootRackSelectors(petSetName string, extraSelectors ...map[string]string) map[string]string {
	selector := map[string]string{
		meta_util.NameLabelKey:      r.ResourceFQN(),
		meta_util.InstanceLabelKey:  r.Name,
		meta_util.ManagedByLabelKey: kubedb.GroupName,
		meta_util.PartOfLabelKey:    petSetName,
	}
	return meta_util.OverwriteKeys(selector, extraSelectors...)
}

func (r *Cassandra) ResourceFQN() string {
	return fmt.Sprintf("%s.%s", r.ResourcePlural(), kubedb.GroupName)
}

func (r *Cassandra) ResourcePlural() string {
	return ResourcePluralCassandra
}

func (r *Cassandra) ServiceName() string {
	return r.OffshootName()
}

func (r *Cassandra) PrimaryServiceDNS() string {
	return fmt.Sprintf("%s.%s.svc", r.ServiceName(), r.Namespace)
}

func (r *Cassandra) GoverningServiceName() string {
	return meta_util.NameWithSuffix(r.ServiceName(), "pods")
}

func (r *Cassandra) RackGoverningServiceName(name string) string {
	return meta_util.NameWithSuffix(name, "pods")
}

func (r *Cassandra) RackGoverningServiceDNS(petSetName string, replicaNo int) string {
	return fmt.Sprintf("%s-%d.%s.%s.svc", petSetName, replicaNo, r.RackGoverningServiceName(petSetName), r.GetNamespace())
}

func (r *Cassandra) GetAuthSecretName() string {
	if r.Spec.AuthSecret != nil && r.Spec.AuthSecret.Name != "" {
		return r.Spec.AuthSecret.Name
	}
	return meta_util.NameWithSuffix(r.OffshootName(), "auth")
}

func (r *Cassandra) ConfigSecretName() string {
	return meta_util.NameWithSuffix(r.OffshootName(), "config")
}

func (r *Cassandra) DefaultUserCredSecretName(username string) string {
	return meta_util.NameWithSuffix(r.Name, strings.ReplaceAll(fmt.Sprintf("%s-cred", username), "_", "-"))
}

func (r *Cassandra) CassandraKeystoreCredSecretName() string {
	return meta_util.NameWithSuffix(r.OffshootName(), kubedb.CassandraKeystoreSecretKey)
}

func (r *Cassandra) PVCName(alias string) string {
	return alias
}

func (r *Cassandra) PetSetName() string {
	return r.OffshootName()
}

func (r *Cassandra) PodLabels(extraLabels ...map[string]string) map[string]string {
	return r.offshootLabels(meta_util.OverwriteKeys(r.OffshootSelectors(), extraLabels...), r.Spec.PodTemplate.Labels)
}

func (r *Cassandra) RackPodLabels(petSetName string, labels map[string]string, extraLabels ...map[string]string) map[string]string {
	return r.offshootLabels(meta_util.OverwriteKeys(r.OffshootRackSelectors(petSetName), extraLabels...), labels)
}

func (r *Cassandra) GetConnectionScheme() string {
	scheme := "http"
	return scheme
}

func (r *Cassandra) OffShootSelectors(extraSelectors ...map[string]string) map[string]string {
	selector := map[string]string{
		meta_util.NameLabelKey:      r.ResourceFQN(),
		meta_util.InstanceLabelKey:  r.Name,
		meta_util.ManagedByLabelKey: kubedb.GroupName,
	}
	return meta_util.OverwriteKeys(selector, extraSelectors...)
}

func (r *Cassandra) offShootLabels(selector, override map[string]string) map[string]string {
	selector[meta_util.ComponentLabelKey] = kubedb.ComponentDatabase
	return meta_util.FilterKeys(kubedb.GroupName, selector, meta_util.OverwriteKeys(nil, r.Labels, override))
}

func (r *Cassandra) OffShootLabels() map[string]string {
	return r.offShootLabels(r.OffShootSelectors(), nil)
}

func (r *Cassandra) ServiceLabels(alias ServiceAlias, extraLabels ...map[string]string) map[string]string {
	svcTemplate := GetServiceTemplate(r.Spec.ServiceTemplates, alias)
	return r.offShootLabels(meta_util.OverwriteKeys(r.OffShootSelectors(), extraLabels...), svcTemplate.Labels)
}

func (r *Cassandra) OffShootName() string {
	return r.Name
}

type CassandraStatsService struct {
	*Cassandra
}

func (ks CassandraStatsService) TLSConfig() *promapi.TLSConfig {
	return nil
}

func (ks CassandraStatsService) GetNamespace() string {
	return ks.Cassandra.GetNamespace()
}

func (ks CassandraStatsService) ServiceName() string {
	return ks.OffShootName() + "-stats"
}

func (ks CassandraStatsService) ServiceMonitorName() string {
	return ks.ServiceName()
}

func (ks CassandraStatsService) ServiceMonitorAdditionalLabels() map[string]string {
	return ks.OffshootLabels()
}

func (ks CassandraStatsService) Path() string {
	return kubedb.DefaultStatsPath
}

func (ks CassandraStatsService) Scheme() string {
	return ""
}

func (r *Cassandra) StatsService() mona.StatsAccessor {
	return &CassandraStatsService{r}
}

func (r *Cassandra) StatsServiceLabels() map[string]string {
	return r.ServiceLabels(StatsServiceAlias, map[string]string{kubedb.LabelRole: kubedb.RoleStats})
}

func (r *Cassandra) SetHealthCheckerDefaults() {
	if r.Spec.HealthChecker.PeriodSeconds == nil {
		r.Spec.HealthChecker.PeriodSeconds = pointer.Int32P(30)
	}
	if r.Spec.HealthChecker.TimeoutSeconds == nil {
		r.Spec.HealthChecker.TimeoutSeconds = pointer.Int32P(10)
	}
	if r.Spec.HealthChecker.FailureThreshold == nil {
		r.Spec.HealthChecker.FailureThreshold = pointer.Int32P(3)
	}
}

func (r *Cassandra) Finalizer() string {
	return fmt.Sprintf("%s/%s", apis.Finalizer, r.ResourceSingular())
}

func (r *Cassandra) ResourceSingular() string {
	return ResourceSingularCassandra
}

func (r *Cassandra) SetTLSDefaults() {
	if r.Spec.TLS == nil || r.Spec.TLS.IssuerRef == nil {
		return
	}
	r.Spec.TLS.Certificates = kmapi.SetMissingSecretNameForCertificate(r.Spec.TLS.Certificates, string(CassandraServerCert), r.CertificateName(CassandraServerCert))
	r.Spec.TLS.Certificates = kmapi.SetMissingSecretNameForCertificate(r.Spec.TLS.Certificates, string(CassandraClientCert), r.CertificateName(CassandraClientCert))
}

func (r *Cassandra) SetDefaults(kc client.Client) {
	if r.Spec.DeletionPolicy == "" {
		r.Spec.DeletionPolicy = DeletionPolicyDelete
	}

	if r.Spec.EnableSSL {
		if r.Spec.KeystoreCredSecret == nil {
			r.Spec.KeystoreCredSecret = &SecretReference{
				LocalObjectReference: core.LocalObjectReference{
					Name: r.CassandraKeystoreCredSecretName(),
				},
			}
		}
	}

	var casVersion catalog.CassandraVersion
	err := kc.Get(context.TODO(), types.NamespacedName{
		Name: r.Spec.Version,
	}, &casVersion)
	if err != nil {
		klog.Errorf("can't get the cassandra version object %s for %s \n", err.Error(), r.Spec.Version)
		return
	}
	if r.Spec.Topology != nil {
		rackName := map[string]bool{}
		racks := r.Spec.Topology.Rack
		for index, rack := range racks {
			if rack.Replicas == nil {
				rack.Replicas = pointer.Int32P(1)
			}
			if rack.Name == "" {
				for i := 1; ; i += 1 {
					rack.Name = r.OffshootRackName(strconv.Itoa(i))
					if !rackName[rack.Name] {
						rackName[rack.Name] = true
						break
					}
				}
			} else {
				rackName[rack.Name] = true
			}
			if rack.StorageType == "" {
				rack.StorageType = StorageTypeDurable
			}

			if rack.PodTemplate == nil {
				rack.PodTemplate = &ofst.PodTemplateSpec{}
			}

			dbContainer := coreutil.GetContainerByName(rack.PodTemplate.Spec.Containers, kubedb.CassandraContainerName)
			if dbContainer != nil && (dbContainer.Resources.Requests == nil && dbContainer.Resources.Limits == nil) {
				apis.SetDefaultResourceLimits(&dbContainer.Resources, kubedb.DefaultResources)
			}
			r.setDefaultContainerSecurityContext(&casVersion, rack.PodTemplate)
			racks[index] = rack
		}
		r.Spec.Topology.Rack = racks
	} else {
		if r.Spec.Replicas == nil {
			r.Spec.Replicas = pointer.Int32P(1)
		}
		if r.Spec.StorageType == "" {
			r.Spec.StorageType = StorageTypeDurable
		}

		if r.Spec.PodTemplate == nil {
			r.Spec.PodTemplate = &ofst.PodTemplateSpec{}
		}
		r.setDefaultContainerSecurityContext(&casVersion, r.Spec.PodTemplate)
		dbContainer := coreutil.GetContainerByName(r.Spec.PodTemplate.Spec.Containers, kubedb.CassandraContainerName)
		if dbContainer != nil && (dbContainer.Resources.Requests == nil && dbContainer.Resources.Limits == nil) {
			apis.SetDefaultResourceLimits(&dbContainer.Resources, kubedb.DefaultResources)
		}
		r.SetHealthCheckerDefaults()
	}
	r.SetTLSDefaults()
	r.Spec.Monitor.SetDefaults()

	if r.Spec.Monitor != nil && r.Spec.Monitor.Prometheus != nil {
		if r.Spec.Monitor.Prometheus.Exporter.SecurityContext.RunAsUser == nil {
			r.Spec.Monitor.Prometheus.Exporter.SecurityContext.RunAsUser = casVersion.Spec.SecurityContext.RunAsUser
		}
		if r.Spec.Monitor.Prometheus.Exporter.SecurityContext.RunAsGroup == nil {
			r.Spec.Monitor.Prometheus.Exporter.SecurityContext.RunAsGroup = casVersion.Spec.SecurityContext.RunAsUser
		}
	}
}

func (r *Cassandra) setDefaultContainerSecurityContext(csVersion *catalog.CassandraVersion, podTemplate *ofst.PodTemplateSpec) {
	if podTemplate == nil {
		return
	}
	if podTemplate.Spec.SecurityContext == nil {
		podTemplate.Spec.SecurityContext = &core.PodSecurityContext{}
	}
	if podTemplate.Spec.SecurityContext.FSGroup == nil {
		podTemplate.Spec.SecurityContext.FSGroup = csVersion.Spec.SecurityContext.RunAsUser
	}

	container := coreutil.GetContainerByName(podTemplate.Spec.Containers, kubedb.CassandraContainerName)
	if container == nil {
		container = &core.Container{
			Name: kubedb.CassandraContainerName,
		}
		podTemplate.Spec.Containers = coreutil.UpsertContainer(podTemplate.Spec.Containers, *container)
	}
	if container.SecurityContext == nil {
		container.SecurityContext = &core.SecurityContext{}
	}
	r.assignDefaultContainerSecurityContext(csVersion, container.SecurityContext)

	initContainer := coreutil.GetContainerByName(podTemplate.Spec.InitContainers, kubedb.CassandraInitContainerName)
	if initContainer == nil {
		initContainer = &core.Container{
			Name: kubedb.CassandraInitContainerName,
		}
		podTemplate.Spec.InitContainers = coreutil.UpsertContainer(podTemplate.Spec.InitContainers, *initContainer)
	}
	if initContainer.SecurityContext == nil {
		initContainer.SecurityContext = &core.SecurityContext{}
	}
	r.assignDefaultContainerSecurityContext(csVersion, initContainer.SecurityContext)
}

func (r *Cassandra) assignDefaultContainerSecurityContext(csVersion *catalog.CassandraVersion, rc *core.SecurityContext) {
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
		rc.RunAsUser = csVersion.Spec.SecurityContext.RunAsUser
	}
	if rc.SeccompProfile == nil {
		rc.SeccompProfile = secomp.DefaultSeccompProfile()
	}
}

func (r *Cassandra) GetSeed() string {
	seed := " "
	namespace := r.Namespace
	name := r.Name
	if r.Spec.Topology == nil {
		seed = fmt.Sprintf("%s-0.%s-pods.%s.svc.cluster.local", name, name, namespace)
		seed = seed + " , "
		return seed
	}
	for _, rack := range r.Spec.Topology.Rack {
		rackCount := min(*rack.Replicas, 3)
		for i := int32(0); i < rackCount; i++ {
			current_seed := fmt.Sprintf("%s-rack-%s-%d.%s-rack-%s-pods.%s.svc.cluster.local", name, rack.Name, i, name, rack.Name, namespace)
			seed += current_seed + " , "
		}
	}
	return seed
}

func (c *Cassandra) ReplicasAreReady(lister pslister.PetSetLister) (bool, string, error) {
	// Desire number of petSets
	expectedItems := 1
	if c.Spec.Topology != nil {
		expectedItems = len(c.Spec.Topology.Rack)
	}
	return checkReplicasOfPetSet(lister.PetSets(c.Namespace), labels.SelectorFromSet(c.OffshootLabels()), expectedItems)
}

// CertificateName returns the default certificate name and/or certificate secret name for a certificate alias
func (m *Cassandra) CertificateName(alias CassandraCertificateAlias) string {
	return meta_util.NameWithSuffix(m.Name, fmt.Sprintf("%s-cert", string(alias)))
}

// GetCertSecretName returns the secret name for a certificate alias if any provide,
// otherwise returns default certificate secret name for the given alias.
func (m *Cassandra) GetCertSecretName(alias CassandraCertificateAlias) string {
	if m.Spec.TLS != nil {
		name, ok := kmapi.GetCertificateSecretName(m.Spec.TLS.Certificates, string(alias))
		if ok {
			return name
		}
	}
	return m.CertificateName(alias)
}

// CertSecretVolumeName returns the CertSecretVolumeName
// Values will be like: client-certs, server-certs etc.
func (c *Cassandra) CertSecretVolumeName(alias CassandraCertificateAlias) string {
	return string(alias) + "-certs"
}

// CertSecretVolumeMountPath returns the CertSecretVolumeMountPath
// if configDir is "/var/cassandra/ssl",
// mountPath will be, "/var/cassandra/ssl/<alias>".
func (c *Cassandra) CertSecretVolumeMountPath(configDir string, cert string) string {
	return filepath.Join(configDir, cert)
}
