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
	ResourceKindFargocdManager = "FargocdManager"
	ResourceFargocdManager     = "fargocdmanager"
	ResourceFargocdManagers    = "fargocdmanagers"
)

// FargocdManager defines the schema for the fargocd OCM AddOn manager
// installer chart. The manager runs on an Open Cluster Management hub
// and ships the fargocd installer chart to every selected spoke via
// the addon-framework Helm agent.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=fargocdmanagers,singular=fargocdmanager,categories={kubeops,appscode}
type FargocdManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FargocdManagerSpec `json:"spec,omitempty"`
}

// FargocdManagerSpec is the schema for the fargocd-manager values file.
type FargocdManagerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`

	// RegistryFQDN is the docker registry fqdn used to pull fargocd
	// docker images on spoke clusters (overrides the spoke chart default).
	RegistryFQDN string `json:"registryFQDN"`
	// Image is the fully qualified manager image name (registry/repository).
	Image string `json:"image"`
	// Tag overrides the manager image tag (defaults to the chart appVersion).
	//+optional
	Tag             string `json:"tag"`
	ImagePullPolicy string `json:"imagePullPolicy"`

	// Placement controls the OCM Placement that selects the spoke
	// clusters which receive the fargocd addon.
	Placement FargocdManagerPlacement `json:"placement"`

	// Argocd is the fargocd configuration propagated to every spoke
	// addon installation. Per-spoke clusterName is injected automatically
	// from the ManagedCluster name.
	Argocd FargocdManagerArgocd `json:"argocd"`

	// KubeconfigSecretName is the name of a Secret holding a kubeconfig
	// for the OCM hub itself. When set, the manager uses it (via
	// --kubeconfig) instead of the in-cluster ServiceAccount.
	//+optional
	KubeconfigSecretName string `json:"kubeconfigSecretName"`

	SecurityContext *core.SecurityContext `json:"securityContext"`
	//+optional
	EnvFrom []core.EnvFromSource `json:"envFrom"`
	//+optional
	Env []core.EnvVar `json:"env"`
	//+optional
	Distro shared.DistroSpec `json:"distro"`
}

// FargocdManagerPlacement controls the OCM Placement that selects the
// spoke clusters which receive the fargocd addon.
type FargocdManagerPlacement struct {
	Create bool   `json:"create"`
	Name   string `json:"name"`
}

// FargocdManagerArgocd holds the hub-wide configuration that the manager
// pushes to each spoke's fargocd installation. Field semantics match the
// `argocd.*` block of the fargocd installer chart.
type FargocdManagerArgocd struct {
	// Mode is one of "in-cluster", "autonomous", or "managed".
	// +kubebuilder:validation:Enum=in-cluster;autonomous;managed
	Mode string `json:"mode"`
	// Namespace overrides argocd-server namespace auto-discovery on each spoke.
	// +optional
	Namespace string `json:"namespace"`
	// DestServer is written into Application.spec.destination.server.
	// +optional
	DestServer string `json:"destServer"`
	// DestName is written into Application.spec.destination.name.
	// +optional
	DestName string `json:"destName"`
	// Project is the Argo CD Project assigned to generated Applications.
	// +optional
	Project string `json:"project"`
	// Kubeconfig is the raw kubeconfig content for the Argo CD principal
	// cluster. When set, the chart creates a Secret on the hub holding
	// this kubeconfig and mounts it into the manager pod; the manager
	// then propagates the kubeconfig to every spoke. Required in managed
	// mode unless KubeconfigSecret is set.
	// +optional
	Kubeconfig string `json:"kubeconfig"`
	// KubeconfigSecret is the name of an existing Secret in the release
	// namespace holding the Argo CD principal kubeconfig. Mutually
	// exclusive with Kubeconfig.
	// +optional
	KubeconfigSecret string `json:"kubeconfigSecret"`
	// KubeconfigSpokeSecret is the name of a pre-created Secret on each
	// spoke holding the principal kubeconfig. When set, the manager
	// propagates this name to spokes instead of pushing inline kubeconfig.
	// +optional
	KubeconfigSpokeSecret string `json:"kubeconfigSpokeSecret"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FargocdManagerList is a list of FargocdManagers
type FargocdManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FargocdManager CRD objects
	Items []FargocdManager `json:"items,omitempty"`
}
