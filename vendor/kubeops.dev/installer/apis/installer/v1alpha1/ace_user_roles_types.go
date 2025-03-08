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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindAceUserRoles = "AceUserRoles"
	ResourceAceUserRoles     = "aceuserroles"
	ResourceAceUserRoless    = "aceuserroless"
)

// AceUserRoles defines the schama for ui server installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=aceuserroless,singular=aceuserroles,categories={kubeops,appscode}
type AceUserRoles struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AceUserRolesSpec `json:"spec,omitempty"`
}

// AceUserRolesSpec is the schema for Identity Server values file
type AceUserRolesSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride   string           `json:"fullnameOverride"`
	EnableClusterRoles UserClusterRoles `json:"enableClusterRoles"`
	//+optional
	Annotations map[string]string `json:"annotations"`
}

type UserClusterRoles struct {
	Ace                bool `json:"ace"`
	Appcatalog         bool `json:"appcatalog"`
	Catalog            bool `json:"catalog"`
	CertManager        bool `json:"cert-manager"`
	Kubedb             bool `json:"kubedb"`
	KubedbUI           bool `json:"kubedb-ui"`
	Kubestash          bool `json:"kubestash"`
	Kubevault          bool `json:"kubevault"`
	LicenseProxyserver bool `json:"license-proxyserver"`
	Metrics            bool `json:"metrics"`
	Prometheus         bool `json:"prometheus"`
	SecretsStore       bool `json:"secrets-store"`
	Stash              bool `json:"stash"`
	VirtualSecrets     bool `json:"virtual-secrets"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AceUserRolesList is a list of AceUserRoless
type AceUserRolesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AceUserRoles CRD objects
	Items []AceUserRoles `json:"items,omitempty"`
}
