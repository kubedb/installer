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

package v1alpha1

import (
	core "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindStorageMetricsServer = "StorageMetricsServer"
	ResourceStorageMetricsServer     = "storagemetricsserver"
	ResourceStorageMetricsServers    = "storagemetricsservers"
)

// StorageMetricsServer defines the schema for the Storage Metrics API server installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
type StorageMetricsServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageMetricsServerSpec `json:"spec,omitempty"`
}

// StorageMetricsServerSpec is the schema for the storage-metrics-server chart values file.
type StorageMetricsServerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string   `json:"fullnameOverride"`
	RegistryFQDN     string   `json:"registryFQDN"`
	Image            ImageRef `json:"image"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	ImagePullPolicy  string   `json:"imagePullPolicy"`
	ReplicaCount     int32    `json:"replicaCount"`
	// Extra args appended to the container command line.
	//+optional
	ExtraArgs []string `json:"extraArgs"`
	// Extra environment variables passed to the container.
	//+optional
	ExtraEnv []core.EnvVar `json:"extraEnv"`
	// Compute Resources required by the apiserver container.
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
	// Security options the apiserver container should run with.
	//+optional
	SecurityContext *core.SecurityContext        `json:"securityContext"`
	Service         StorageMetricsServiceSpec    `json:"service"`
	ApiService      StorageMetricsAPIServiceSpec `json:"apiService"`
	//+optional
	PriorityClassName string `json:"priorityClassName"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	//+optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints.
	//+optional
	Affinity            *core.Affinity                    `json:"affinity"`
	PodDisruptionBudget StorageMetricsPodDisruptionBudget `json:"podDisruptionBudget"`
	// Additional subjects bound to the metrics-reader ClusterRole.
	//+optional
	MetricsReaderSubjects []rbac.Subject `json:"metricsReaderSubjects"`
	// Distro-specific overrides (OpenShift, UBI image variant).
	//+optional
	Distro shared.DistroSpec `json:"distro"`
}

// StorageMetricsServiceSpec configures the metrics Service fronting the aggregated apiserver.
type StorageMetricsServiceSpec struct {
	Type       string `json:"type"`
	Port       int32  `json:"port"`
	TargetPort int32  `json:"targetPort"`
}

// StorageMetricsAPIServiceSpec configures the APIService registered with kube-aggregator.
type StorageMetricsAPIServiceSpec struct {
	Create                bool  `json:"create"`
	InsecureSkipTLSVerify bool  `json:"insecureSkipTLSVerify"`
	GroupPriorityMinimum  int32 `json:"groupPriorityMinimum"`
	VersionPriority       int32 `json:"versionPriority"`
}

// StorageMetricsPodDisruptionBudget configures the PodDisruptionBudget for the apiserver Deployment.
type StorageMetricsPodDisruptionBudget struct {
	Enabled      bool  `json:"enabled"`
	MinAvailable int32 `json:"minAvailable"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageMetricsServerList is a list of StorageMetricsServers
type StorageMetricsServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageMetricsServer CRD objects
	Items []StorageMetricsServer `json:"items,omitempty"`
}
