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
	ResourceKindKubedbAutoscaler = "KubedbAutoscaler"
	ResourceKubedbAutoscaler     = "kubedbautoscaler"
	ResourceKubedbAutoscalers    = "kubedbautoscalers"
)

// KubedbAutoscaler defines the schama for KubeDB Enterprise Operator Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubedbautoscalers,singular=kubedbautoscaler,categories={kubedb,appscode}
type KubedbAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubedbAutoscalerSpec `json:"spec,omitempty"`
}

// KubedbAutoscalerSpec is the spec for redis version
type KubedbAutoscalerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string     `json:"fullnameOverride"`
	ReplicaCount     int32      `json:"replicaCount"`
	Operator         Container  `json:"operator"`
	Cleaner          CleanerRef `json:"cleaner"`
	ImagePullPolicy  string     `json:"imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
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
	EnableAnalytics bool       `json:"enableAnalytics"`
	Monitoring      Monitoring `json:"monitoring"`
	// +optional
	AdditionalPodSecurityPolicies []string `json:"additionalPodSecurityPolicies"`
	// +optional
	License string `json:"license"`
	// +optional
	UpdateInterval string `json:"updateInterval"`
	// +optional
	StorageAutoscaler StorageAutoscalerSpec `json:"storageAutoscaler"`
}

type StorageAutoscalerSpec struct {
	Prometheus PrometheusSpec `json:"prometheus"`
}

type PrometheusSpec struct {
	Address string `json:"address"`
	// +optional
	BearerToken string `json:"bearerToken"`
	// +optional
	CACert string `json:"caCert"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubedbAutoscalerList is a list of KubedbAutoscalers
type KubedbAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubedbAutoscaler CRD objects
	Items []KubedbAutoscaler `json:"items,omitempty"`
}
