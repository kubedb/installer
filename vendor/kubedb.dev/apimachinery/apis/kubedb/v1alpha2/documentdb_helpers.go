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
	"net/url"

	"kubedb.dev/apimachinery/apis"
	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	"kubedb.dev/apimachinery/crds"

	"github.com/Masterminds/semver/v3"
	"github.com/fatih/structs"
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

func (f *DocumentDB) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourcePluralDocumentDB))
}

func (f *DocumentDB) ResourcePlural() string {
	return ResourcePluralDocumentDB
}

func (f *DocumentDB) ResourceFQN() string {
	return fmt.Sprintf("%s.%s", f.ResourcePlural(), "kubedb.com")
}

func (f *DocumentDB) ServiceName() string {
	return f.Name
}

type DocumentDBApp struct {
	*DocumentDB
}

func (f DocumentDBApp) Name() string {
	return f.DocumentDB.Name
}

func (f DocumentDBApp) Type() appcat.AppType {
	return appcat.AppType(fmt.Sprintf("%s/%s", kubedb.GroupName, ResourceSingularDocumentDB))
}

func (f *DocumentDB) AppBindingMeta() appcat.AppBindingMeta {
	return &DocumentDBApp{f}
}

func (f *DocumentDB) OffshootName() string {
	return f.Name
}

func (f *DocumentDB) OffshootSelectors() map[string]string {
	selector := map[string]string{
		meta_util.NameLabelKey:      f.ResourceFQN(),
		meta_util.InstanceLabelKey:  f.Name,
		meta_util.ManagedByLabelKey: "kubedb.com",
	}
	return selector
}

func (f *DocumentDB) PodControllerLabels(podControllerLabels map[string]string, extraLabels ...map[string]string) map[string]string {
	return f.offshootLabels(meta_util.OverwriteKeys(f.OffshootSelectors(), extraLabels...), podControllerLabels)
}

func (f *DocumentDB) OffshootLabels() map[string]string {
	return f.offshootLabels(f.OffshootSelectors(), nil)
}

func (f *DocumentDB) PrimaryServerSelectors() map[string]string {
	return meta_util.OverwriteKeys(f.OffshootSelectors(), map[string]string{
		kubedb.DocumentDBPrimaryLabelKey: f.OffshootName(),
	})
}

func (f *DocumentDB) PrimaryServerLabels() map[string]string {
	return meta_util.OverwriteKeys(f.OffshootLabels(), f.PrimaryServerSelectors())
}

func (f *DocumentDB) offshootLabels(selector, override map[string]string) map[string]string {
	selector[meta_util.ComponentLabelKey] = kubedb.ComponentDatabase
	return meta_util.FilterKeys("kubedb.com", selector, meta_util.OverwriteKeys(nil, f.Labels, override))
}

func (f *DocumentDB) GetAuthSecretName() string {
	if f.Spec.AuthSecret != nil && f.Spec.AuthSecret.Name != "" {
		return f.Spec.AuthSecret.Name
	}
	return meta_util.NameWithSuffix(f.OffshootName(), "auth")
}

func (f *DocumentDB) GetPersistentSecrets() []string {
	var secrets []string
	if f.Spec.AuthSecret != nil {
		secrets = append(secrets, f.GetAuthSecretName())
	}
	return secrets
}

// AsOwner returns owner reference to resources
func (f *DocumentDB) AsOwner() *meta.OwnerReference {
	return meta.NewControllerRef(f, SchemeGroupVersion.WithKind(f.ResourceKind()))
}

func (f *DocumentDB) ResourceKind() string {
	return ResourceKindDocumentDB
}

func (f *DocumentDB) PgBackendName() string {
	return f.OffshootName() + "-pg-backend"
}

func (f *DocumentDB) PodLabels(podTemplateLabels map[string]string, extraLabels ...map[string]string) map[string]string {
	return f.offshootLabels(meta_util.OverwriteKeys(f.OffshootSelectors(), extraLabels...), podTemplateLabels)
}

func (f *DocumentDB) CertificateName(alias DocumentDBCertificateAlias) string {
	return meta_util.NameWithSuffix(f.Name, fmt.Sprintf("%s-cert", string(alias)))
}

func (f *DocumentDB) GetCertSecretName(alias DocumentDBCertificateAlias) string {
	name, ok := kmapi.GetCertificateSecretName(f.Spec.TLS.Certificates, string(alias))
	if ok {
		return name
	}

	return f.CertificateName(alias)
}

func (f *DocumentDB) GetExternalBackendClientSecretName() string {
	return f.Name + "-ext-pg-client-cert"
}

func (f *DocumentDB) GetBackendConnectionSecretName() string {
	return f.OffshootName() + "-backend-connection"
}

func (f *DocumentDB) GetSecretVolumeName(secretName string) string {
	return secretName + "-vol"
}

func (f *DocumentDB) SetHealthCheckerDefaults() {
	if f.Spec.HealthChecker.PeriodSeconds == nil {
		f.Spec.HealthChecker.PeriodSeconds = pointer.Int32P(10)
	}
	if f.Spec.HealthChecker.TimeoutSeconds == nil {
		f.Spec.HealthChecker.TimeoutSeconds = pointer.Int32P(10)
	}
	if f.Spec.HealthChecker.FailureThreshold == nil {
		f.Spec.HealthChecker.FailureThreshold = pointer.Int32P(2)
	}
}

func (f *DocumentDB) SetDefaults(kc client.Client) {
	if f == nil {
		return
	}

	if f.Spec.Backend == nil {
		f.Spec.Backend = &DocumentDBBackendSpec{}
	}

	if f.Spec.Backend.Replicas == nil {
		f.Spec.Backend.Replicas = pointer.Int32P(3)
	}

	if f.Spec.Backend.PodTemplate == nil {
		f.Spec.Backend.PodTemplate = &ofst.PodTemplateSpec{}
	}

	if f.Spec.Backend.StorageType == "" {
		f.Spec.Backend.StorageType = StorageTypeDurable
	}

	if f.Spec.DeletionPolicy == "" {
		f.Spec.DeletionPolicy = DeletionPolicyWipeOut
	}

	if f.Spec.SSLMode == "" {
		f.Spec.SSLMode = SSLModeDisabled
	}

	if f.Spec.AuthSecret == nil {
		f.Spec.AuthSecret = &SecretReference{}
	}
	if f.Spec.AuthSecret.Kind == "" {
		f.Spec.AuthSecret.Kind = kubedb.ResourceKindSecret
	}

	var dcVersion catalog.DocumentDBVersion
	err := kc.Get(context.TODO(), types.NamespacedName{
		Name: f.Spec.Version,
	}, &dcVersion)
	if err != nil {
		klog.Errorf("can't get the DocumentDB version object %s for %s \n", err.Error(), f.Spec.Version)
		return
	}

	if f.Spec.Server == nil {
		f.Spec.Server = &DocumentDBServer{
			Primary: &DocumentDBServerSpec{
				Replicas:    pointer.Int32P(1),
				PodTemplate: &ofst.PodTemplateSpec{},
			},
		}
	}

	if f.Spec.Server.Primary != nil {
		if f.Spec.Server.Primary.Replicas == nil {
			f.Spec.Server.Primary.Replicas = pointer.Int32P(1)
		}
		if f.Spec.Server.Primary.PodTemplate == nil {
			f.Spec.Server.Primary.PodTemplate = &ofst.PodTemplateSpec{}
		}
		f.setDefaultPodTemplateValues(f.Spec.Server.Primary.PodTemplate, &dcVersion)
	}

	if f.Spec.Server.Secondary != nil {
		if f.Spec.Server.Secondary.Replicas == nil {
			f.Spec.Server.Secondary.Replicas = pointer.Int32P(1)
		}
		if f.Spec.Server.Secondary.PodTemplate == nil {
			f.Spec.Server.Secondary.PodTemplate = &ofst.PodTemplateSpec{}
		}
		f.setDefaultPodTemplateValues(f.Spec.Server.Secondary.PodTemplate, &dcVersion)
	}

	if f.Spec.AuthSecret == nil {
		f.Spec.AuthSecret = &SecretReference{
			ExternallyManaged: false,
		}
	}

	f.Spec.Monitor.SetDefaults()
	if f.Spec.Monitor != nil && f.Spec.Monitor.Prometheus != nil {
		if f.Spec.Monitor.Prometheus.Exporter.SecurityContext.RunAsUser == nil {
			f.Spec.Monitor.Prometheus.Exporter.SecurityContext.RunAsUser = dcVersion.Spec.SecurityContext.RunAsUser
		}
		if f.Spec.Monitor.Prometheus.Exporter.SecurityContext.RunAsGroup == nil {
			f.Spec.Monitor.Prometheus.Exporter.SecurityContext.RunAsGroup = dcVersion.Spec.SecurityContext.RunAsUser
		}
	}

	f.SetTLSDefaults()
	f.SetHealthCheckerDefaults()
}

func (f *DocumentDB) setDefaultPodTemplateValues(podTemplate *ofst.PodTemplateSpec, dcVersion *catalog.DocumentDBVersion) {
	dbContainer := coreutil.GetContainerByName(podTemplate.Spec.Containers, kubedb.DocumentDBContainerName)
	if dbContainer == nil {
		dbContainer = &core.Container{
			Name: kubedb.DocumentDBContainerName,
		}
		podTemplate.Spec.Containers = append(podTemplate.Spec.Containers, *dbContainer)
	}
	if structs.IsZero(dbContainer.Resources) {
		apis.SetDefaultResourceLimits(&dbContainer.Resources, kubedb.DefaultResources)
	}
	if dbContainer.SecurityContext == nil {
		dbContainer.SecurityContext = &core.SecurityContext{}
	}
	f.setDefaultContainerSecurityContext(dcVersion, dbContainer.SecurityContext)
	f.setDefaultPodTemplateSecurityContext(dcVersion, podTemplate)
}

func (f *DocumentDB) IsVersionAtLeast(dcVersion *catalog.DocumentDBVersion, version uint64) bool {
	v, _ := semver.NewVersion(dcVersion.Spec.Version)
	return v.Major() >= version
}

func (f *DocumentDB) setDefaultPodTemplateSecurityContext(dcVersion *catalog.DocumentDBVersion, podTemplate *ofst.PodTemplateSpec) {
	if podTemplate == nil {
		return
	}
	if podTemplate.Spec.SecurityContext == nil {
		podTemplate.Spec.SecurityContext = &core.PodSecurityContext{}
	}
	if podTemplate.Spec.SecurityContext.FSGroup == nil {
		podTemplate.Spec.SecurityContext.FSGroup = dcVersion.Spec.SecurityContext.RunAsUser
	}
}

func (f *DocumentDB) setDefaultContainerSecurityContext(dcVersion *catalog.DocumentDBVersion, sc *core.SecurityContext) {
	if sc.AllowPrivilegeEscalation == nil {
		sc.AllowPrivilegeEscalation = pointer.BoolP(false)
	}
	if sc.Capabilities == nil {
		sc.Capabilities = &core.Capabilities{
			Drop: []core.Capability{"ALL"},
		}
	}
	if sc.RunAsNonRoot == nil {
		sc.RunAsNonRoot = pointer.BoolP(true)
	}
	if sc.RunAsUser == nil {
		sc.RunAsUser = dcVersion.Spec.SecurityContext.RunAsUser
	}
	if sc.RunAsGroup == nil {
		sc.RunAsGroup = dcVersion.Spec.SecurityContext.RunAsUser
	}
	if sc.SeccompProfile == nil {
		sc.SeccompProfile = secomp.DefaultSeccompProfile()
	}
}

func (f *DocumentDB) SetTLSDefaults() {
	if f.Spec.TLS == nil || f.Spec.TLS.IssuerRef == nil {
		return
	}

	defaultServerOrg := []string{kubedb.KubeDBOrganization}
	defaultServerOrgUnit := []string{string(DocumentDBServerCert)}

	_, cert := kmapi.GetCertificate(f.Spec.TLS.Certificates, string(DocumentDBServerCert))
	if cert != nil && cert.Subject != nil {
		if cert.Subject.Organizations != nil {
			defaultServerOrg = cert.Subject.Organizations
		}
		if cert.Subject.OrganizationalUnits != nil {
			defaultServerOrgUnit = cert.Subject.OrganizationalUnits
		}
	}
	f.Spec.TLS.Certificates = kmapi.SetMissingSpecForCertificate(f.Spec.TLS.Certificates, kmapi.CertificateSpec{
		Alias:      string(DocumentDBServerCert),
		SecretName: f.GetCertSecretName(DocumentDBServerCert),
		Subject: &kmapi.X509Subject{
			Organizations:       defaultServerOrg,
			OrganizationalUnits: defaultServerOrgUnit,
		},
	})

	// Client-cert
	defaultClientOrg := []string{kubedb.KubeDBOrganization}
	defaultClientOrgUnit := []string{string(DocumentDBClientCert)}
	_, cert = kmapi.GetCertificate(f.Spec.TLS.Certificates, string(DocumentDBClientCert))
	if cert != nil && cert.Subject != nil {
		if cert.Subject.Organizations != nil {
			defaultClientOrg = cert.Subject.Organizations
		}
		if cert.Subject.OrganizationalUnits != nil {
			defaultClientOrgUnit = cert.Subject.OrganizationalUnits
		}
	}
	f.Spec.TLS.Certificates = kmapi.SetMissingSpecForCertificate(f.Spec.TLS.Certificates, kmapi.CertificateSpec{
		Alias:      string(DocumentDBClientCert),
		SecretName: f.GetCertSecretName(DocumentDBClientCert),
		Subject: &kmapi.X509Subject{
			Organizations:       defaultClientOrg,
			OrganizationalUnits: defaultClientOrgUnit,
		},
	})
}

type DocumentDBStatsService struct {
	*DocumentDB
}

func (fs DocumentDBStatsService) ServiceMonitorName() string {
	return fs.OffshootName() + "-stats"
}

func (fs DocumentDBStatsService) ServiceMonitorAdditionalLabels() map[string]string {
	return fs.OffshootLabels()
}

func (fs DocumentDBStatsService) Path() string {
	return kubedb.DocumentDBMetricsPath
}

func (fs DocumentDBStatsService) Scheme() string {
	sc := promapi.SchemeHTTP
	return sc.String()
}

func (fs DocumentDBStatsService) TLSConfig() *promapi.TLSConfig {
	return nil
}

func (fs DocumentDBStatsService) ServiceName() string {
	return fs.OffshootName() + "-stats"
}

func (f *DocumentDB) StatsService() mona.StatsAccessor {
	return &DocumentDBStatsService{f}
}

func (f *DocumentDB) GoverningServiceName() string {
	return f.OffshootName() + "-pods"
}

func (f *DocumentDB) ServiceLabels(alias ServiceAlias, extraLabels ...map[string]string) map[string]string {
	svcTemplate := GetServiceTemplate(f.Spec.ServiceTemplates, alias)
	return f.offshootLabels(meta_util.OverwriteKeys(f.OffshootSelectors(), extraLabels...), svcTemplate.Labels)
}

func (f *DocumentDB) StatsServiceLabels() map[string]string {
	return f.ServiceLabels(StatsServiceAlias, map[string]string{kubedb.LabelRole: kubedb.RoleStats})
}

func (f *DocumentDB) ReplicasAreReady(lister pslister.PetSetLister) (bool, string, error) {
	// Desire number of petSets
	expectedItems := 1
	return checkReplicasOfPetSet(lister.PetSets(f.Namespace), labels.SelectorFromSet(f.OffshootLabels()), expectedItems)
}

func (f *DocumentDB) GetSSLModeFromAppBinding(apb *appcat.AppBinding) (PostgresSSLMode, error) {
	var sslMode string
	if apb.Spec.ClientConfig.URL != nil {
		parsedURL, err := url.Parse(*apb.Spec.ClientConfig.URL)
		if err != nil {
			return "", fmt.Errorf("parse error in appbinding %s/%s 'spec.clientConfig.url'. error: %v", apb.Namespace, apb.Name, err)
		}
		if parsedURL.Scheme != "postgres" && parsedURL.Scheme != "postgresql" {
			return "", fmt.Errorf("invalid scheme provided in URL in provided appbinding %s/%s", apb.Namespace, apb.Name)
		}
		sslMode = parsedURL.Query().Get("sslmode")
	}
	if apb.Spec.ClientConfig.Service != nil {
		values, err := url.ParseQuery(apb.Spec.ClientConfig.Service.Query)
		if err != nil {
			return "", fmt.Errorf("parse error in appbinding %s/%s 'spec.clientConfig.service.query'. error: %v", apb.Namespace, apb.Name, err)
		}
		if sslMode != "" && sslMode != values.Get("sslmode") {
			return "", fmt.Errorf("sslMode is not same in 'spec.clientConfig.service.query' and 'spec.clientConfig.url' of appbinding %s/%s", apb.Namespace, apb.Name)
		}
		sslMode = values.Get("sslmode")
	}
	// If sslMode is not specified anywhere, it will be disabled
	if sslMode == "" {
		sslMode = "disable"
	}
	return PostgresSSLMode(sslMode), nil
}
