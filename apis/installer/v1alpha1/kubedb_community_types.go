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
	ResourceKindKubedbCommunity = "KubedbCommunity"
	ResourceKubedbCommunity     = "kubedbcommunity"
	ResourceKubedbCommunitys    = "kubedbcommunitys"
)

// KubedbCommunity defines the schama for KubeDB Operator Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubedbcommunitys,singular=kubedbcommunity,categories={kubedb,appscode}
type KubedbCommunity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubedbCommunitySpec `json:"spec,omitempty"`
}

// KubedbCommunitySpec is the schema for kubedb-community chart values file
type KubedbCommunitySpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string     `json:"fullnameOverride"`
	ReplicaCount     int32      `json:"replicaCount"`
	Operator         Container  `json:"operator"`
	Cleaner          CleanerRef `json:"cleaner"`
	ImagePullPolicy  string     `json:"imagePullPolicy"`
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
	EnforceTerminationPolicy bool `json:"enforceTerminationPolicy"`
	// +optional
	EnableAnalytics bool       `json:"enableAnalytics"`
	Monitoring      Monitoring `json:"monitoring"`
	// +optional
	AdditionalPodSecurityPolicies []string `json:"additionalPodSecurityPolicies"`
	// +optional
	License string `json:"license"`
}

type ServiceAccountSpec struct {
	Create bool `json:"create"`
	//+optional
	Name *string `json:"name"`
	//+optional
	Annotations map[string]string `json:"annotations"`
}

type WebHookSpec struct {
	GroupPriorityMinimum    int32  `json:"groupPriorityMinimum"`
	VersionPriority         int32  `json:"versionPriority"`
	EnableMutatingWebhook   bool   `json:"enableMutatingWebhook"`
	EnableValidatingWebhook bool   `json:"enableValidatingWebhook"`
	CA                      string `json:"ca"`
	// +optional
	BypassValidatingWebhookXray bool            `json:"bypassValidatingWebhookXray"`
	UseKubeapiserverFqdnForAks  bool            `json:"useKubeapiserverFqdnForAks"`
	Healthcheck                 HealthcheckSpec `json:"healthcheck"`
	Port                        int32           `json:"port"`
	ServingCerts                ServingCerts    `json:"servingCerts"`
}

type ServingCerts struct {
	Generate bool `json:"generate"`
	// +optional
	CaCrt string `json:"caCrt"`
	// +optional
	ServerCrt string `json:"serverCrt"`
	// +optional
	ServerKey string `json:"serverKey"`
}

type HealthcheckSpec struct {
	// +optional
	Enabled bool `json:"enabled"`
}

type Monitoring struct {
	// +optional
	Enabled        bool                  `json:"enabled"`
	Agent          string                `json:"agent"`
	ServiceMonitor *ServiceMonitorLabels `json:"serviceMonitor"`
}

type ServiceMonitorLabels struct {
	// +optional
	Labels map[string]string `json:"labels"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubedbCommunityList is a list of KubedbCommunity-s
type KubedbCommunityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubedbCommunity CRD objects
	Items []KubedbCommunity `json:"items,omitempty"`
}
