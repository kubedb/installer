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
	ResourceKindOpencostGrafanaDashboards = "OpencostGrafanaDashboards"
	ResourceOpencostGrafanaDashboards     = "opencostgrafanadashboards"
	ResourceOpencostGrafanaDashboardss    = "opencostgrafanadashboardss"
)

// OpencostGrafanaDashboards defines the schama for ui server installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=opencostgrafanadashboardss,singular=opencostgrafanadashboards,categories={kubeops,appscode}
type OpencostGrafanaDashboards struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpencostGrafanaDashboardsSpec `json:"spec,omitempty"`
}

// OpencostGrafanaDashboardsSpec is the schema for Identity Server values file
type OpencostGrafanaDashboardsSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string           `json:"fullnameOverride"`
	Dashboard        GrafanaDashboard `json:"dashboard"`
	Grafana          ObjectReference  `json:"grafana"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpencostGrafanaDashboardsList is a list of OpencostGrafanaDashboardss
type OpencostGrafanaDashboardsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpencostGrafanaDashboards CRD objects
	Items []OpencostGrafanaDashboards `json:"items,omitempty"`
}
