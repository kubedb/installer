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
	kubeopsinstaller "kubeops.dev/installer/apis/installer/v1alpha1"
)

const (
	ResourceKindKubedb = "Kubedb"
	ResourceKubedb     = "kubedb"
	ResourceKubedbs    = "kubedbs"
)

// Kubedb defines the schama for KubeDB combined Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubedbs,singular=kubedb,categories={kubedb,appscode}
type Kubedb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubedbSpec `json:"spec,omitempty"`
}

// KubedbSpec is the schema for kubedb chart values file
type KubedbSpec struct {
	Global GlobalValues `json:"global"`

	//+optional
	Petset PetsetValues `json:"petset"`
	//+optional
	OperatorShardManager OperatorShardManagerValues `json:"operator-shard-manager"`
	//+optional
	Sidekick SidekickValues `json:"sidekick"`
	//+optional
	Supervisor SupervisorValues `json:"supervisor"`

	//+optional
	KubedbCrdManager KubedbCrdManagerValues `json:"kubedb-crd-manager"`

	//+optional
	KubedbProvisioner KubedbProvisionerValues `json:"kubedb-provisioner"`

	//+optional
	KubedbCatalog KubedbCatalogValues `json:"kubedb-catalog"`

	//+optional
	KubedbKubestashCatalog KubedbKubestashCatalogValues `json:"kubedb-kubestash-catalog"`

	//+optional
	KubedbWebhookServer KubedbWebhookServerValues `json:"kubedb-webhook-server"`

	//+optional
	KubedbOpsManager KubedbOpsManagerValues `json:"kubedb-ops-manager"`

	//+optional
	KubedbAutoscaler KubedbAutoscalerValues `json:"kubedb-autoscaler"`

	//+optional
	KubedbSchemaManager KubedbSchemaManagerValues `json:"kubedb-schema-manager"`

	//+optional
	KubedbMetrics KubedbMetricsValues `json:"kubedb-metrics"`

	//+optional
	KubedbGitops KubedbGitopsValues `json:"kubedb-gitops"`

	//+optional
	AceUserRoles AceUserRolesValues `json:"ace-user-roles"`
}

type PetsetValues struct {
	Enabled bool `json:"enabled"`
}

type OperatorShardManagerValues struct {
	Enabled bool `json:"enabled"`
}

type SidekickValues struct {
	Enabled bool `json:"enabled"`
}

type SupervisorValues struct {
	Enabled bool `json:"enabled"`
}

type KubedbCrdManagerValues struct {
	Enabled               *bool `json:"enabled"`
	*KubedbCrdManagerSpec `json:",inline,omitempty"`
}

type KubedbProvisionerValues struct {
	Enabled                *bool `json:"enabled"`
	*KubedbProvisionerSpec `json:",inline,omitempty"`
}

type KubedbCatalogValues struct {
	Enabled            *bool `json:"enabled"`
	*KubedbCatalogSpec `json:",inline,omitempty"`
}

type KubedbKubestashCatalogValues struct {
	Enabled                     bool `json:"enabled"`
	*KubedbKubestashCatalogSpec `json:",inline,omitempty"`
}

type KubedbWebhookServerValues struct {
	Enabled                  bool `json:"enabled"`
	*KubedbWebhookServerSpec `json:",inline,omitempty"`
}

type KubedbOpsManagerValues struct {
	Enabled               *bool `json:"enabled"`
	*KubedbOpsManagerSpec `json:",inline,omitempty"`
}

type KubedbAutoscalerValues struct {
	Enabled               *bool `json:"enabled"`
	*KubedbAutoscalerSpec `json:",inline,omitempty"`
}

type KubedbSchemaManagerValues struct {
	Enabled                  bool `json:"enabled"`
	*KubedbSchemaManagerSpec `json:",inline,omitempty"`
}

type KubedbMetricsValues struct {
	Enabled bool `json:"enabled"`
}

type KubedbGitopsValues struct {
	Enabled           bool `json:"enabled"`
	*KubedbGitopsSpec `json:",inline,omitempty"`
}

type AceUserRolesValues struct {
	Enabled            bool                               `json:"enabled"`
	EnableClusterRoles *kubeopsinstaller.UserClusterRoles `json:"enableClusterRoles,omitempty"`
}

type GlobalValues struct {
	License            string   `json:"license"`
	LicenseSecretName  string   `json:"licenseSecretName"`
	Registry           string   `json:"registry"`
	RegistryFQDN       string   `json:"registryFQDN"`
	InsecureRegistries []string `json:"insecureRegistries"`
	//+optional
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets"`
	FeatureGates     map[string]bool             `json:"featureGates"`
	Monitoring       EASMonitoring               `json:"monitoring"`
	// +optional
	MaxConcurrentReconciles int `json:"maxConcurrentReconciles"`
	// +optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity       *core.Affinity `json:"affinity"`
	WaitForWebhook bool           `json:"waitForWebhook"`

	// +optional
	NetworkPolicy NetworkPolicy `json:"networkPolicy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubedbList is a list of Kubedbs
type KubedbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Kubedb CRD objects
	Items []Kubedb `json:"items,omitempty"`
}
