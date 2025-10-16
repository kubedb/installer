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
	"fmt"

	"kubedb.dev/apimachinery/apis"
	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	"kubedb.dev/apimachinery/crds"

	promapi "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"gomodules.xyz/pointer"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apimachinery/pkg/runtime/schema"
	_ "k8s.io/klog/v2"
	_ "k8s.io/utils/ptr"
	"kmodules.xyz/client-go/apiextensions"
	metautil "kmodules.xyz/client-go/meta"
	_ "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	mona "kmodules.xyz/monitoring-agent-api/api/v1"
	ofst "kmodules.xyz/offshoot-api/api/v2"
	ofstutil "kmodules.xyz/offshoot-api/util"
	"sigs.k8s.io/controller-runtime/pkg/client"
	_ "sigs.k8s.io/controller-runtime/pkg/client"
)

func (_ HanaDB) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourcePluralHanaDB))
}

func (o *HanaDB) ResourceKind() string {
	return ResourceKindHanaDB
}

func (o *HanaDB) ResourcePlural() string {
	return ResourcePluralHanaDB
}

func (o *HanaDB) ResourceFQN() string {
	return fmt.Sprintf("%s.%s", o.ResourcePlural(), SchemeGroupVersion.Group)
}

func (o *HanaDB) ResourceShortCode() string {
	return ResourceCodeHanaDB
}

func (o *HanaDB) OffshootName() string {
	return o.Name
}

func (o *HanaDB) ServiceName() string {
	return o.OffshootName()
}

func (o *HanaDB) ObserverServiceName() string {
	return o.OffshootName() + kubedb.OracleDatabaseRoleObserver
}

func (o *HanaDB) GoverningServiceName() string {
	return metautil.NameWithSuffix(o.ServiceName(), "pods")
}

func (o *HanaDB) StandbyServiceName() string {
	return metautil.NameWithPrefix(o.ServiceName(), kubedb.OracleStandbyServiceSuffix)
}

func (o *HanaDB) offshootLabels(selector, override map[string]string) map[string]string {
	selector[metautil.ComponentLabelKey] = kubedb.ComponentDatabase
	return metautil.FilterKeys(SchemeGroupVersion.Group, selector, metautil.OverwriteKeys(nil, o.Labels, override))
}

//func (o *HanaDB) ServiceLabels(alias ServiceAlias, extraLabels ...map[string]string) map[string]string {
//	svcTemplate := GetServiceTemplate(o.Spec.ServiceTemplates, alias)
//	return o.offshootLabels(metautil.OverwriteKeys(o.OffshootSelectors(), extraLabels...), svcTemplate.Labels)
//}

func (o *HanaDB) OffshootLabels() map[string]string {
	return o.offshootLabels(o.OffshootSelectors(), nil)
}

func (o *HanaDB) OffshootSelectors(extraSelectors ...map[string]string) map[string]string {
	selector := map[string]string{
		metautil.NameLabelKey:      o.ResourceFQN(),
		metautil.InstanceLabelKey:  o.Name,
		metautil.ManagedByLabelKey: SchemeGroupVersion.Group,
	}
	return metautil.OverwriteKeys(selector, extraSelectors...)
}

func (o *HanaDB) OffshootPodSelectors(extraSelectors ...map[string]string) map[string]string {
	selector := map[string]string{
		metautil.NameLabelKey:        o.ResourceFQN(),
		metautil.InstanceLabelKey:    o.Name,
		metautil.ManagedByLabelKey:   SchemeGroupVersion.Group,
		kubedb.OracleDatabaseRoleKey: kubedb.OracleDatabaseRoleInstance,
	}
	return metautil.OverwriteKeys(selector, extraSelectors...)
}

func (o *HanaDB) ObserverSelectors(extraSelectors ...map[string]string) map[string]string {
	selector := o.OffshootSelectors()
	selector[kubedb.OracleDatabaseRoleKey] = kubedb.OracleDatabaseRoleObserver
	return metautil.OverwriteKeys(selector, extraSelectors...)
}

func (o *HanaDB) PodControllerLabels(podTemplate *ofst.PodTemplateSpec, extraLabels ...map[string]string) map[string]string {
	if podTemplate != nil && podTemplate.Controller.Labels != nil {
		return o.offshootLabels(metautil.OverwriteKeys(o.OffshootSelectors(), extraLabels...), podTemplate.Controller.Labels)
	}
	return o.offshootLabels(metautil.OverwriteKeys(o.OffshootSelectors(), extraLabels...), nil)
}

func (o *HanaDB) ObserverPodControllerLabels(podTemplate *ofst.PodTemplateSpec, extraLabels ...map[string]string) map[string]string {
	if podTemplate != nil && podTemplate.Controller.Labels != nil {
		return o.offshootLabels(metautil.OverwriteKeys(o.OffshootSelectors(), extraLabels...), podTemplate.Controller.Labels)
	}
	labels := o.offshootLabels(metautil.OverwriteKeys(o.OffshootSelectors(), extraLabels...), nil)
	labels[kubedb.OracleDatabaseRoleKey] = kubedb.OracleDatabaseRoleObserver
	return labels
}

func (o *HanaDB) PodLabels(podTemplate *ofst.PodTemplateSpec, extraLabels ...map[string]string) map[string]string {
	if podTemplate != nil && podTemplate.Labels != nil {
		return o.offshootLabels(metautil.OverwriteKeys(o.OffshootSelectors(), extraLabels...), podTemplate.Labels)
	}
	return o.offshootLabels(metautil.OverwriteKeys(o.OffshootSelectors(), extraLabels...), nil)
}

func (o *HanaDB) ObserverPodLabels(podTemplate *ofst.PodTemplateSpec, extraLabels ...map[string]string) map[string]string {
	labels := make(map[string]string)
	labels[kubedb.OracleDatabaseRoleKey] = kubedb.OracleDatabaseRoleObserver
	extraLabels = append(extraLabels, labels)
	return o.PodLabels(podTemplate, extraLabels...)
}

func (o *HanaDB) ServiceAccountName() string {
	return o.OffshootName()
}

// Owner returns owner reference to resources
func (o *HanaDB) Owner() *metav1.OwnerReference {
	return metav1.NewControllerRef(o, SchemeGroupVersion.WithKind(o.ResourceKind()))
}

func (o *HanaDB) GetAuthSecretName() string {
	if o.Spec.AuthSecret != nil && o.Spec.AuthSecret.Name != "" {
		return o.Spec.AuthSecret.Name
	}
	return metautil.NameWithSuffix(o.OffshootName(), "auth")
}

func (o *HanaDB) GetPersistentSecrets() []string {
	var secrets []string
	secrets = append(secrets, o.GetAuthSecretName())
	return secrets
}

func (m *HanaDB) GetNameSpacedName() string {
	return m.Namespace + "/" + m.Name
}

func (o *HanaDB) DefaultPodRoleName() string {
	return metautil.NameWithSuffix(o.OffshootName(), "role")
}

func (o *HanaDB) DefaultPodRoleBindingName() string {
	return metautil.NameWithSuffix(o.OffshootName(), "rolebinding")
}

func (r *HanaDB) Finalizer() string {
	return fmt.Sprintf("%s/%s", apis.Finalizer, r.ResourceSingular())
}

func (r *HanaDB) ResourceSingular() string {
	return ResourceSingularHanaDB
}

type hanadbStatsService struct {
	*HanaDB
}

func (os hanadbStatsService) TLSConfig() *promapi.TLSConfig {
	return nil
}

func (os hanadbStatsService) GetNamespace() string {
	return os.HanaDB.GetNamespace()
}

func (os hanadbStatsService) ServiceName() string {
	return os.OffshootName() + "-stats"
}

func (os hanadbStatsService) ServiceMonitorName() string {
	return os.ServiceName()
}

func (os hanadbStatsService) ServiceMonitorAdditionalLabels() map[string]string {
	return os.OffshootLabels()
}

func (os hanadbStatsService) Path() string {
	return kubedb.DefaultStatsPath
}

func (os hanadbStatsService) Scheme() string {
	return ""
}

func (o *HanaDB) StatsService() mona.StatsAccessor {
	return &hanadbStatsService{o}
}

type hanadbApp struct {
	*HanaDB
}

func (r hanadbApp) Name() string {
	return r.HanaDB.Name
}

func (r hanadbApp) Type() appcat.AppType {
	return appcat.AppType(fmt.Sprintf("%s/%s", SchemeGroupVersion.Group, ResourceSingularHanaDB))
}

//
//func (o HanaDB) AppBindingMeta() appcat.AppBindingMeta {
//	return &oracleApp{&o}
//}

//func (o *HanaDB) StatsServiceLabels() map[string]string {
//	return o.ServiceLabels(StatsServiceAlias, map[string]string{kubedb.LabelRole: kubedb.RoleStats})
//}

func (o *HanaDB) PetSetName() string {
	return o.OffshootName()
}

func (o *HanaDB) ObserverPetSetName() string {
	return fmt.Sprintf("%s-observer", o.PetSetName())
}

func (o *HanaDB) ConfigSecretName() string {
	return metautil.NameWithSuffix(o.OffshootName(), "config")
}

//func (o *HanaDB) IsStandalone() bool {
//	return o.Spec.Mode == OracleModeStandalone
//}
//
//func (o *HanaDB) IsDataGuardEnabled() bool {
//	return o.Spec.Mode == OracleModeDataGuard
//}

func (o *HanaDB) SetHealthCheckerDefaults() {
	if o.Spec.HealthChecker.PeriodSeconds == nil {
		o.Spec.HealthChecker.PeriodSeconds = pointer.Int32P(10)
	}
	if o.Spec.HealthChecker.TimeoutSeconds == nil {
		o.Spec.HealthChecker.TimeoutSeconds = pointer.Int32P(10)
	}
	if o.Spec.HealthChecker.FailureThreshold == nil {
		o.Spec.HealthChecker.FailureThreshold = pointer.Int32P(3)
	}
}

func (o *HanaDB) SetDefaults(kc client.Client) {
	//if o.Spec.Halted {
	//	if o.Spec.DeletionPolicy == DeletionPolicyDoNotTerminate {
	//		klog.Errorf(`Can't halt, since deletion policy is 'DoNotTerminate'`)
	//		return
	//	}
	//	o.Spec.DeletionPolicy = DeletionPolicyHalt
	//}

	//if o.Spec.DeletionPolicy == "" {
	//	o.Spec.DeletionPolicy = DeletionPolicyDelete
	//}
	//
	//if o.Spec.StorageType == "" {
	//	o.Spec.StorageType = StorageTypeDurable
	//}
	//
	//o.SetListenerDefaults()
	//o.initializePodTemplates()
	//
	//oraVersion := &catalog.OracleVersion{}
	//err := kc.Get(context.Background(), types.NamespacedName{Name: o.Spec.Version}, oraVersion)
	//if err != nil {
	//	klog.Errorf("can't get the oracle version object %s for %s \n", err.Error(), o.Spec.Version)
	//	return
	//}
	//
	//if o.Spec.PodTemplate.Spec.ServiceAccountName == "" {
	//	o.Spec.PodTemplate.Spec.ServiceAccountName = o.OffshootName()
	//}
	//
	//if o.Spec.Mode == OracleModeDataGuard {
	//	o.SetDataGuardDefaults()
	//	o.SetObserverInitContainerDefaults(o.Spec.DataGuard.Observer.PodTemplate, oraVersion)
	//	o.SetHanaDBObserverContainerDefaults(o.Spec.DataGuard.Observer.PodTemplate, oraVersion)
	//}
	//
	//o.SetDefaultPodSecurityContext(o.Spec.PodTemplate, oraVersion)
	//o.SetHanaDBContainerDefaults(o.Spec.PodTemplate, oraVersion)
	//o.SetCoordinatorContainerDefaults(o.Spec.PodTemplate, oraVersion)
	//o.SetInitContainerDefaults(o.Spec.PodTemplate, oraVersion)
	//o.SetHealthCheckerDefaults()
	//o.Spec.Monitor.SetDefaults()
	//if o.Spec.Monitor != nil && o.Spec.Monitor.Prometheus != nil {
	//	if o.Spec.Monitor.Prometheus.Exporter.SecurityContext.RunAsUser == nil {
	//		o.Spec.Monitor.Prometheus.Exporter.SecurityContext.RunAsUser = oraVersion.Spec.SecurityContext.RunAsUser
	//	}
	//	if o.Spec.Monitor.Prometheus.Exporter.SecurityContext.RunAsGroup == nil {
	//		o.Spec.Monitor.Prometheus.Exporter.SecurityContext.RunAsGroup = oraVersion.Spec.SecurityContext.RunAsUser
	//	}
	//}
}

//func (o *HanaDB) SetListenerDefaults() {
//	if o.Spec.Listener == nil {
//		o.Spec.Listener = &ListenerSpec{}
//	}
//	o.Spec.Listener.Port = ptr.To(int32(kubedb.OracleDatabasePort))
//	o.Spec.Listener.Protocol = OracleListenerProtocolTCP
//	o.Spec.Listener.Service = ptr.To(kubedb.OracleDatabaseServiceName)
//}

//func (o *HanaDB) initializePodTemplates() {
//	if o.Spec.Mode == OracleModeDataGuard {
//		if o.Spec.DataGuard == nil {
//			o.Spec.DataGuard = &DataGuardSpec{}
//		}
//		if o.Spec.DataGuard.Observer == nil {
//			o.Spec.DataGuard.Observer = &ObserverSpec{}
//		}
//		if o.Spec.DataGuard.Observer.PodTemplate == nil {
//			o.Spec.DataGuard.Observer.PodTemplate = new(ofst.PodTemplateSpec)
//		}
//	}
//	if o.Spec.PodTemplate == nil {
//		o.Spec.PodTemplate = new(ofst.PodTemplateSpec)
//	}
//}

func (o *HanaDB) SetDefaultPodSecurityContext(podTemplate *ofst.PodTemplateSpec, oraVersion *catalog.OracleVersion) {
	if podTemplate == nil {
		return
	}

	if podTemplate.Spec.SecurityContext == nil {
		podTemplate.Spec.SecurityContext = &core.PodSecurityContext{}
	}
	if podTemplate.Spec.SecurityContext.FSGroup == nil {
		podTemplate.Spec.SecurityContext.FSGroup = oraVersion.Spec.SecurityContext.RunAsUser
	}
	if podTemplate.Spec.SecurityContext.RunAsUser == nil {
		podTemplate.Spec.SecurityContext.RunAsUser = oraVersion.Spec.SecurityContext.RunAsUser
	}
	if podTemplate.Spec.SecurityContext.RunAsGroup == nil {
		podTemplate.Spec.SecurityContext.RunAsGroup = oraVersion.Spec.SecurityContext.RunAsUser
	}
}

func (o *HanaDB) SetInitContainerDefaults(podTemplate *ofst.PodTemplateSpec, oraVersion *catalog.OracleVersion) {
	if podTemplate == nil {
		return
	}
	container := ofstutil.EnsureInitContainerExists(podTemplate, kubedb.OracleInitContainerName)
	o.setContainerDefaultSecurityContext(container, oraVersion)
	o.setContainerDefaultResources(container, *kubedb.DefaultInitContainerResource.DeepCopy())
}

func (o *HanaDB) SetObserverInitContainerDefaults(podTemplate *ofst.PodTemplateSpec, oraVersion *catalog.OracleVersion) {
	if podTemplate == nil {
		return
	}
	container := ofstutil.EnsureInitContainerExists(podTemplate, kubedb.OracleObserverInitContainerName)
	o.setContainerDefaultSecurityContext(container, oraVersion)
	o.setContainerDefaultResources(container, *kubedb.DefaultInitContainerResource.DeepCopy())
}

func (o *HanaDB) SetHanaDBContainerDefaults(podTemplate *ofst.PodTemplateSpec, oraVersion *catalog.OracleVersion) {
	if podTemplate == nil {
		return
	}
	container := ofstutil.EnsureContainerExists(podTemplate, kubedb.OracleContainerName)
	o.setContainerDefaultSecurityContext(container, oraVersion)
	o.setContainerDefaultResources(container, *kubedb.DefaultResourcesCoreAndMemoryIntensiveOracle.DeepCopy())
}

func (o *HanaDB) SetHanaDBObserverContainerDefaults(podTemplate *ofst.PodTemplateSpec, oraVersion *catalog.OracleVersion) {
	if podTemplate == nil {
		return
	}
	container := ofstutil.EnsureContainerExists(podTemplate, kubedb.OracleObserverContainerName)
	o.setContainerDefaultSecurityContext(container, oraVersion)
	o.setContainerDefaultResources(container, *kubedb.DefaultResourcesCoreAndMemoryIntensiveOracleObserver.DeepCopy())
}

func (o *HanaDB) SetCoordinatorContainerDefaults(podTemplate *ofst.PodTemplateSpec, oraVersion *catalog.OracleVersion) {
	if podTemplate == nil {
		return
	}
	container := ofstutil.EnsureContainerExists(podTemplate, kubedb.OracleCoordinatorContainerName)
	o.setContainerDefaultSecurityContext(container, oraVersion)
	o.setContainerDefaultResources(container, *kubedb.CoordinatorDefaultResources.DeepCopy())
}

func (o *HanaDB) setContainerDefaultSecurityContext(container *core.Container, _ *catalog.OracleVersion) {
	if container.SecurityContext == nil {
		container.SecurityContext = &core.SecurityContext{}
	}
	// TODO: Check what part of security context make hanadb fail to run
	// o.assignDefaultContainerSecurityContext(container.SecurityContext, oraVersion)
}

func (o *HanaDB) setContainerDefaultResources(container *core.Container, defaultResources core.ResourceRequirements) {
	if container.Resources.Requests == nil && container.Resources.Limits == nil {
		apis.SetDefaultResourceLimits(&container.Resources, defaultResources)
	}
}
