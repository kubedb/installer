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
	ResourceKindKubedbOpsManager = "KubedbOpsManager"
	ResourceKubedbOpsManager     = "kubedbopsmanager"
	ResourceKubedbOpsManagers    = "kubedbopsmanagers"
)

// KubedbOpsManager defines the schama for KubeDB Ops Manager Operator Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubedbopsmanagers,singular=kubedbopsmanager,categories={kubedb,appscode}
type KubedbOpsManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubedbOpsManagerSpec `json:"spec,omitempty"`
}

// KubedbOpsManagerSpec is the schema for kubedb-ops-manager chart values file
type KubedbOpsManagerSpec struct {
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
	Monitoring         Monitoring               `json:"monitoring"`
	// +optional
	License string `json:"license"`
	// +optional
	LicenseSecretName string `json:"licenseSecretName"`
	// +optional
	RecommendationEngine RecommendationEngineConfig `json:"recommendationEngine"`
	Psp                  PSPSpec                    `json:"psp"`
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
	Env []core.EnvVar `json:"env"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubedbOpsManagerList is a list of KubedbOpsManagers
type KubedbOpsManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubedbOpsManager CRD objects
	Items []KubedbOpsManager `json:"items,omitempty"`
}

type RecommendationEngineConfig struct {
	RecommendationResyncPeriod                  metav1.Duration `json:"recommendationResyncPeriod"`
	GenRotateTLSRecommendationBeforeExpiryYear  int             `json:"genRotateTLSRecommendationBeforeExpiryYear"`
	GenRotateTLSRecommendationBeforeExpiryMonth int             `json:"genRotateTLSRecommendationBeforeExpiryMonth"`
	GenRotateTLSRecommendationBeforeExpiryDay   int             `json:"genRotateTLSRecommendationBeforeExpiryDay"`
}
