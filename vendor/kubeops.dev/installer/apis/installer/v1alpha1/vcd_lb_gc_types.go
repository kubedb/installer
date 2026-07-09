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
)

const (
	ResourceKindVcdLbGc = "VcdLbGc"
	ResourceVcdLbGc     = "vcdlbgc"
	ResourceVcdLbGcs    = "vcdlbgcs"
)

// VcdLbGc defines the schema for vcd-lb-gc Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=vcdlbgcs,singular=vcdlbgc,categories={kubeops,appscode}
type VcdLbGc struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VcdLbGcSpec `json:"spec,omitempty"`
}

// VcdLbGcSpec is the schema for vcd-lb-gc values file
type VcdLbGcSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string    `json:"fullnameOverride"`
	ReplicaCount     int32     `json:"replicaCount"`
	RegistryFQDN     string    `json:"registryFQDN"`
	Image            Container `json:"image"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	ImagePullPolicy  string   `json:"imagePullPolicy"`
	//+optional
	LogLevel int32 `json:"logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	Config             VcdLbGcConfig            `json:"config"`
	Vcd                VcdLbGcCredentials       `json:"vcd"`
}

// VcdLbGcConfig holds the controller runtime configuration.
type VcdLbGcConfig struct {
	// CPI cluster ID, e.g. "capvcdCluster:<uuid>"
	ClusterID string `json:"clusterID"`
	// URN of the edge gateway hosting the LB, e.g. "urn:vcloud:gateway:<uuid>"
	EdgeGatewayID string `json:"edgeGatewayID"`
	// Reconcile interval
	Interval string `json:"interval"`
	// If true, logs orphans without deleting them.
	DryRun bool `json:"dryRun"`
	// Skip DNAT cleanup (set when enableVirtualServiceSharedIP is on)
	SkipDNAT bool `json:"skipDNAT"`
}

// VcdLbGcCredentials holds the VMware Cloud Director tenant credentials.
type VcdLbGcCredentials struct {
	// VCD URL, e.g. https://vcd.example.com
	Endpoint string `json:"endpoint"`
	// VCD tenant org name
	Org string `json:"org"`
	// VCD tenant user
	User string `json:"user"`
	// VCD tenant password
	Password string `json:"password"`
	// Skip TLS verification when talking to VCD
	Insecure bool `json:"insecure"`
	// Use an existing Secret holding VCD_* keys instead of creating one.
	//+optional
	ExistingSecret string `json:"existingSecret"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VcdLbGcList is a list of VcdLbGcs
type VcdLbGcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VcdLbGc CRD objects
	Items []VcdLbGc `json:"items,omitempty"`
}
