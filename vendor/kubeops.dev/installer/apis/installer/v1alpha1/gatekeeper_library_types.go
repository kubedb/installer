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
	ResourceKindGatekeeperLibrary = "GatekeeperLibrary"
	ResourceGatekeeperLibrary     = "gatekeeperlibrary"
	ResourceGatekeeperLibrarys    = "gatekeeperlibrarys"
)

// GatekeeperLibrary defines the schama for GatekeeperLibrary Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=gatekeeperlibrarys,singular=gatekeeperlibrary,categories={kubeops,appscode}
type GatekeeperLibrary struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatekeeperLibrarySpec `json:"spec,omitempty"`
}

// GatekeeperLibrarySpec is the schema for GatekeeperLibrary Operator values file
type GatekeeperLibrarySpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`
	// +kubebuilder:default:=templates
	Enable            GatekeeperResource         `json:"enable"`
	EnableConstraints map[string]map[string]bool `json:"enableConstraints"`
	// +kubebuilder:default:=warn
	EnforcementAction EnforcementAction `json:"enforcementAction"`
}

// +kubebuilder:validation:Enum=templates;constraints
type GatekeeperResource string

// +kubebuilder:validation:Enum=warn;deny;dryrun
type EnforcementAction string

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GatekeeperLibraryList is a list of GatekeeperLibrarys
type GatekeeperLibraryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GatekeeperLibrary CRD objects
	Items []GatekeeperLibrary `json:"items,omitempty"`
}
