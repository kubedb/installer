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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindConfigSyncer = "ConfigSyncer"
	ResourceConfigSyncer     = "configsyncer"
	ResourceConfigSyncers    = "configsyncers"
)

// ConfigSyncer defines the schama for ConfigSyncer Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=configsyncers,singular=configsyncer,categories={kubeops,appscode}
type ConfigSyncer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigSyncerSpec `json:"spec,omitempty"`
}

// +kubebuilder:validation:Enum=oss;enterprise
// +kubebuilder:default:=oss
type LicenseMode string

// ConfigSyncerSpec is the schema for ConfigSyncer Operator values file
type ConfigSyncerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`
	ReplicaCount     int32  `json:"replicaCount"`
	RegistryFQDN     string `json:"registryFQDN"`
	// +optional
	License string `json:"license"`
	// +optional
	Mode            LicenseMode `json:"mode"`
	Image           Container   `json:"image"`
	ImagePullPolicy string      `json:"imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector" protobuf:"bytes,12,rep,name=nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations" protobuf:"bytes,13,rep,name=tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity" protobuf:"bytes,14,opt,name=affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext  `json:"podSecurityContext"`
	ServiceAccount     ServiceAccountSpec        `json:"serviceAccount"`
	Apiserver          AConfigSyncerpiserverSpec `json:"apiserver"`
	Config             ConfigSyncerConfig        `json:"config"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
}

type AConfigSyncerpiserverSpec struct {
	SecurePort                 string          `json:"securePort"`
	UseKubeapiserverFqdnForAks bool            `json:"useKubeapiserverFqdnForAks"`
	Healthcheck                HealthcheckSpec `json:"healthcheck"`
	ServingCerts               ServingCerts    `json:"servingCerts"`
}

type ConfigSyncerConfig struct {
	ClusterName           string   `json:"clusterName"`
	ConfigSourceNamespace string   `json:"configSourceNamespace"`
	KubeconfigContent     string   `json:"kubeconfigContent"`
	AdditionalOptions     []string `json:"additionalOptions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigSyncerList is a list of ConfigSyncers
type ConfigSyncerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigSyncer CRD objects
	Items []ConfigSyncer `json:"items,omitempty"`
}
