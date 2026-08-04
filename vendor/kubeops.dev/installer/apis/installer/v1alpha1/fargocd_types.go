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
	ResourceKindFargocd = "Fargocd"
	ResourceFargocd     = "fargocd"
	ResourceFargocds    = "fargocds"
)

// Fargocd defines the schama for Fargocd operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=fargocds,singular=fargocd,categories={kubeops,appscode}
type Fargocd struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FargocdSpec `json:"spec,omitempty"`
}

// FargocdSpec is the schema for Identity Server values file
type FargocdSpec struct {
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
	Affinity       *core.Affinity     `json:"affinity"`
	ServiceAccount ServiceAccountSpec `json:"serviceAccount"`
	Apiserver      FargocdApiserver   `json:"apiserver"`
	Monitoring     Monitoring         `json:"monitoring"`
	// +optional
	NetworkPolicy NetworkPolicySpec `json:"networkPolicy"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
	Argocd FargocdArgocd     `json:"argocd"`
}

// FargocdArgocd configures how the controller talks to Argo CD. It
// mirrors the `fargocd run` flags.
type FargocdArgocd struct {
	// Mode is one of "in-cluster", "autonomous", or "managed".
	// +kubebuilder:validation:Enum=in-cluster;autonomous;managed
	Mode string `json:"mode"`
	// Namespace overrides argocd-server namespace auto-discovery.
	// +optional
	Namespace string `json:"namespace"`
	// DestServer is written into Application.spec.destination.server.
	// +optional
	DestServer string `json:"destServer"`
	// DestName is written into Application.spec.destination.name; used
	// when Argo CD references the workload cluster by symbolic name.
	// +optional
	DestName string `json:"destName"`
	// Project is the Argo CD Project assigned to generated Applications.
	// +optional
	Project string `json:"project"`
	// ClusterName is the symbolic name of the workload cluster.
	// Required in managed mode.
	// +optional
	ClusterName string `json:"clusterName"`
	// KubeconfigSecret is the name of a Secret (in the release namespace)
	// whose key `kubeconfig` holds the kubeconfig for the Argo CD
	// principal cluster. Required in managed mode if Kubeconfig is empty.
	// +optional
	KubeconfigSecret string `json:"kubeconfigSecret"`
	// Kubeconfig is the raw kubeconfig content for the Argo CD principal
	// cluster. When set, the chart creates a Secret containing this
	// kubeconfig and mounts it into the operator pod. Either Kubeconfig
	// or KubeconfigSecret is required in managed mode.
	// +optional
	Kubeconfig string `json:"kubeconfig"`
}

type FargocdApiserver struct {
	EnableMutatingWebhook   bool                `json:"enableMutatingWebhook"`
	EnableValidatingWebhook bool                `json:"enableValidatingWebhook"`
	Healthcheck             HealthcheckSpec     `json:"healthcheck"`
	ServingCerts            FargocdServingCerts `json:"servingCerts"`
}

type FargocdServingCerts struct {
	Generate bool `json:"generate"`
	//+optional
	CertManager FargocdCertManagerCerts `json:"certManager"`
	//+optional
	CaCrt string `json:"caCrt"`
	//+optional
	ServerCrt string `json:"serverCrt"`
	//+optional
	ServerKey string `json:"serverKey"`
}

type FargocdCertManagerCerts struct {
	Enabled bool `json:"enabled"`
	//+optional
	IssuerRef FargocdCertManagerIssuerRef `json:"issuerRef"`
}

type FargocdCertManagerIssuerRef struct {
	//+optional
	Name string `json:"name"`
	//+optional
	Kind string `json:"kind"`
	//+optional
	Group string `json:"group"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FargocdList is a list of Fargocds
type FargocdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Fargocd CRD objects
	Items []Fargocd `json:"items,omitempty"`
}
