/*
Copyright The KubeDB Authors.

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

package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindKubeDBOperator = "KubeDBOperator"
	ResourceKubeDBOperator     = "kubedboperator"
	ResourceKubeDBOperators    = "kubedboperators"
)

// KubeDBOperator defines the schama for KubeDB Operator Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubedboperators,singular=kubedboperator,categories={kubedb,appscode}
type KubeDBOperator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeDBOperatorSpec `json:"spec,omitempty"`
}

type ImageRef struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

// KubeDBOperatorSpec is the spec for redis version
type KubeDBOperatorSpec struct {
	ReplicaCount    int               `json:"replicaCount"`
	KubeDB          ImageRef          `json:"kubedb"`
	Cleaner         ImageRef          `json:"cleaner"`
	ImagePullPolicy string            `json:"imagePullPolicy"`
	CriticalAddon   bool              `json:"criticalAddon"`
	LogLevel        int               `json:"logLevel"`
	Annotations     map[string]string `json:"annotations"`
	NodeSelector    map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations,omitempty"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity                      *core.Affinity     `json:"affinity,omitempty"`
	ServiceAccount                ServiceAccountSpec `json:"serviceAccount"`
	Apiserver                     WebHookSpec        `json:"apiserver"`
	EnableAnalytics               bool               `json:"enableAnalytics"`
	Monitoring                    Monitoring         `json:"monitoring"`
	AdditionalPodSecurityPolicies []string           `json:"additionalPodSecurityPolicies"`
	// Compute Resources required by the sidecar container.
	Resources core.ResourceRequirements `json:"resources,omitempty"`
}

type ServiceAccountSpec struct {
	Create bool   `json:"create"`
	Name   string `json:"name"`
}

type WebHookSpec struct {
	GroupPriorityMinimum        int             `json:"groupPriorityMinimum"`
	VersionPriority             int             `json:"versionPriority"`
	EnableMutatingWebhook       bool            `json:"enableMutatingWebhook"`
	EnableValidatingWebhook     bool            `json:"enableValidatingWebhook"`
	Ca                          string          `json:"ca"`
	BypassValidatingWebhookXray bool            `json:"bypassValidatingWebhookXray"`
	UseKubeapiserverFqdnForAks  bool            `json:"useKubeapiserverFqdnForAks"`
	Healthcheck                 HealthcheckSpec `json:"healthcheck"`
	Port int `json:"port"`
}

type HealthcheckSpec struct {
	Enabled bool `json:"enabled"`
}

type Monitoring struct {
	Enabled    bool   `json:"enabled"`
	Agent          string               `json:"agent"`
	Prometheus     PrometheusSpec       `json:"prometheus"`
	ServiceMonitor ServiceMonitorLabels `json:"serviceMonitor"`
}

type PrometheusSpec struct {
	Namespace string `json:"namespace"`
}

type ServiceMonitorLabels struct {
	Labels map[string]string `json:"labels"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeDBOperatorList is a list of KubeDBOperators
type KubeDBOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubeDBOperator CRD objects
	Items []KubeDBOperator `json:"items,omitempty"`
}
