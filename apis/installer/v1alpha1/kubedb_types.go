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
	KubedbProvisioner KubedbProvisionerValues `json:"kubedb-provisioner"`

	//+optional
	KubedbCatalog KubedbCatalogValues `json:"kubedb-catalog"`

	//+optional
	KubedbWebhookServer KubedbWebhookServerValues `json:"kubedb-webhook-server"`

	//+optional
	KubedbOpsManager KubedbOpsManagerValues `json:"kubedb-ops-manager"`

	//+optional
	KubedbAutoscaler KubedbAutoscalerValues `json:"kubedb-autoscaler"`

	//+optional
	KubedbDashboard KubedbDashboardValues `json:"kubedb-dashboard"`

	//+optional
	KubedbSchemaManager KubedbSchemaManagerValues `json:"kubedb-schema-manager"`

	//+optional
	KubedbMetrics KubedbMetricsValues `json:"kubedb-metrics"`
}

type KubedbProvisionerValues struct {
	Enabled                *bool `json:"enabled"`
	*KubedbProvisionerSpec `json:",inline,omitempty"`
}

type KubedbCatalogValues struct {
	Enabled            *bool `json:"enabled"`
	*KubedbCatalogSpec `json:",inline,omitempty"`
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

type KubedbDashboardValues struct {
	Enabled              bool `json:"enabled"`
	*KubedbDashboardSpec `json:",inline,omitempty"`
}

type KubedbSchemaManagerValues struct {
	Enabled                  bool `json:"enabled"`
	*KubedbSchemaManagerSpec `json:",inline,omitempty"`
}

type KubedbMetricsValues struct {
	Enabled bool `json:"enabled"`
}

type GlobalValues struct {
	License            string   `json:"license"`
	Registry           string   `json:"registry"`
	RegistryFQDN       string   `json:"registryFQDN"`
	InsecureRegistries []string `json:"insecureRegistries"`
	//+optional
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets"`
	Monitoring       EASMonitoring               `json:"monitoring"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubedbList is a list of Kubedbs
type KubedbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Kubedb CRD objects
	Items []Kubedb `json:"items,omitempty"`
}
