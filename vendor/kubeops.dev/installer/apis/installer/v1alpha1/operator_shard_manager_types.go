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
	ResourceKindOperatorShardManager = "OperatorShardManager"
	ResourceOperatorShardManager     = "operatorshardmanager"
	ResourceOperatorShardManagers    = "operatorshardmanagers"
)

// OperatorShardManager defines the schama for Operator Shard Manager installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=operatorshardmanagers,singular=operatorshardmanager,categories={kubeops,appscode}
type OperatorShardManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OperatorShardManagerSpec `json:"spec,omitempty"`
}

// OperatorShardManagerSpec is the schema for Identity Server values file
type OperatorShardManagerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string    `json:"fullnameOverride"`
	ReplicaCount     int       `json:"replicaCount"`
	RegistryFQDN     string    `json:"registryFQDN"`
	Image            Container `json:"image"`
	ImagePullPolicy  string    `json:"imagePullPolicy"`
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
	PodLabels map[string]string `json:"podLabels"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity       *core.Affinity      `json:"affinity"`
	ServiceAccount ServiceAccountSpec  `json:"serviceAccount"`
	Apiserver      SupervisorApiserver `json:"apiserver"`
	Monitoring     Monitoring          `json:"monitoring"`
	// +optional
	NetworkPolicy NetworkPolicySpec `json:"networkPolicy"`
	// +optional
	Distro DistroSpec `json:"distro"`
}

type ImageReference struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	PullPolicy string `json:"pullPolicy"`
}

type ServiceSpec struct {
	Type string `json:"type"`
	Port int    `json:"port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OperatorShardManagerList is a list of OperatorShardManagers
type OperatorShardManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OperatorShardManager CRD objects
	Items []OperatorShardManager `json:"items,omitempty"`
}
