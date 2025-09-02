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
	ResourceKindKubedbCrdManager = "KubedbCrdManager"
	ResourceKubedbCrdManager     = "kubedbcrdmanager"
	ResourceKubedbCrdManagers    = "kubedbcrdmanagers"
)

// KubedbCrdManager defines the schama for KubedbCrdManager operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubedbcrdmanagers,singular=kubedbcrdmanager,categories={kubeops,appscode}
type KubedbCrdManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubedbCrdManagerSpec `json:"spec,omitempty"`
}

// KubedbCrdManagerSpec is the schema for Identity Server values file
type KubedbCrdManagerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`
	//+optional
	RegistryFQDN    string   `json:"registryFQDN"`
	Image           ImageRef `json:"image"`
	ImagePullPolicy string   `json:"imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	//+optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity                *core.Affinity     `json:"affinity"`
	ServiceAccount          ServiceAccountSpec `json:"serviceAccount"`
	FeatureGates            map[string]bool    `json:"featureGates"`
	RemoveUnusedCRDs        bool               `json:"removeUnusedCRDs"`
	InstallGitOpsCRDs       bool               `json:"installGitOpsCRDs"`
	TTLSecondsAfterFinished int                `json:"ttlSecondsAfterFinished"`
	// +optional
	Distro DistroSpec `json:"distro"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubedbCrdManagerList is a list of KubedbCrdManagers
type KubedbCrdManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubedbCrdManager CRD objects
	Items []KubedbCrdManager `json:"items,omitempty"`
}
