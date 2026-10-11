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

package v1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindKubedbProvisioner = "KubedbProvisioner"
	ResourceKubedbProvisioner     = "kubedbprovisioner"
	ResourceKubedbProvisioners    = "kubedbprovisioners"
)

// KubedbProvisioner defines the schama for KubeDB Operator Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubedbprovisioners,singular=kubedbprovisioner,categories={kubedb,appscode}
type KubedbProvisioner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubedbProvisionerSpec `json:"spec,omitempty"`
}

// KubedbProvisionerSpec is the schema for kubedb-provisioner chart values file
type KubedbProvisionerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride   string    `json:"fullnameOverride"`
	ReplicaCount       int32     `json:"replicaCount"`
	RegistryFQDN       string    `json:"registryFQDN"`
	InsecureRegistries []string  `json:"insecureRegistries"`
	Operator           Container `json:"operator"`
	Waitfor            ImageRef  `json:"waitfor"`
	ImagePullPolicy    string    `json:"imagePullPolicy"`
	//+optional
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets"`
	// +optional
	CriticalAddon bool `json:"criticalAddon"`
	// +optional
	LogLevel int32 `json:"logLevel"`
	// +optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	// +optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	Apiserver          WebHookSpec              `json:"apiserver"`
	// +optional
	EnforceTerminationPolicy bool       `json:"enforceTerminationPolicy"`
	Monitoring               Monitoring `json:"monitoring"`
	// +optional
	AdditionalPodSecurityPolicies []string `json:"additionalPodSecurityPolicies"`
	// +optional
	License string `json:"license"`
	// +optional
	LicenseSecretName string  `json:"licenseSecretName"`
	Psp               PSPSpec `json:"psp"`
	// +optional
	MaxConcurrentReconciles int `json:"maxConcurrentReconciles"`
	// List of sources to populate environment variables in the container.
	// The keys defined within a source must be a C_IDENTIFIER. All invalid keys
	// will be reported as an event when the container is starting. When a key exists in multiple
	// sources, the value associated with the last source will take precedence.
	// Values defined by an Env with a duplicate key will take precedence.
	// Cannot be updated.
	// +optional
	// +listType=atomic
	EnvFrom []core.EnvFromSource `json:"envFrom"`
	// List of environment variables to set in the container.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	Env           []core.EnvVar     `json:"env"`
	NetworkPolicy NetworkPolicySpec `json:"networkPolicy"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
	// DcDR enables the cross data center (DC-DR) failover orchestrator. Unset (the
	// default) runs a plain, non DC-DR provisioner.
	// +optional
	DcDR *DcDRSpec `json:"dcDR,omitempty"`
}

// DcDRSpec configures the DC-DR failover orchestrator in the provisioner.
type DcDRSpec struct {
	// Enable the DC-DR failover orchestrator.
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// The data center (OCM spoke cluster) name this operator instance runs in; it is the
	// marker activeDC, the Lease holderIdentity, and the PlacementPolicy clusterName.
	// +optional
	LocalDC string `json:"localDC,omitempty"`
	// Name of a Secret in the operator namespace whose "kubeconfig" key is the
	// coordination control plane kubeconfig (where the primary-DC Leases live), mounted at
	// /etc/dr/coord. Leave empty when this cluster is itself the coordination plane.
	// +optional
	CoordKubeconfigSecret string `json:"coordKubeconfigSecret,omitempty"`
	// QPS for the client talking to the coordination control plane apiserver. Unset, the
	// operator binary's default applies.
	// +optional
	CoordQPS *int32 `json:"coordQPS,omitempty"`
	// Burst for the client talking to the coordination control plane apiserver. Unset, the
	// operator binary's default applies.
	// +optional
	CoordBurst *int32 `json:"coordBurst,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubedbProvisionerList is a list of KubedbProvisioner-s
type KubedbProvisionerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubedbProvisioner CRD objects
	Items []KubedbProvisioner `json:"items,omitempty"`
}
