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
	ResourceKindFalcoUiServer = "FalcoUiServer"
	ResourceFalcoUiServer     = "falcouiserver"
	ResourceFalcoUiServers    = "falcouiservers"
)

// FalcoUiServer defines the schama for FalcoUiServer Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=falcouiservers,singular=falcouiserver,categories={kubeops,appscode}
type FalcoUiServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FalcoUiServerSpec `json:"spec,omitempty"`
}

// FalcoUiServerSpec is the schema for FalcoUiServer Operator values file
type FalcoUiServerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string        `json:"fullnameOverride"`
	ReplicaCount     int32         `json:"replicaCount"`
	RegistryFQDN     string        `json:"registryFQDN"`
	App              Container     `json:"app"`
	Etcd             EtcdContainer `json:"etcd"`
	Kine             Container     `json:"kine"`
	ImagePullPolicy  string        `json:"imagePullPolicy"`
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
	NodeSelector map[string]string `json:"nodeSelector" protobuf:"bytes,12,rep,name=nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations" protobuf:"bytes,13,rep,name=tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity" protobuf:"bytes,14,opt,name=affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	StorageClass       LocalObjectReference     `json:"storageClass"`
	Persistence        Persistence              `json:"persistence"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	Apiserver          FalcoUiserverSpec        `json:"apiserver"`
	Monitoring         Monitoring               `json:"monitoring"`
	Dashboard          GrafanaDashboard         `json:"dashboard"`
	Grafana            ObjectReference          `json:"grafana"`
	EventTTL           metav1.Duration          `json:"eventTTL"`
}

type FalcoUiserverSpec struct {
	ApiserverSpec `json:",inline"`
	DB            ApiserverDB `json:"db"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FalcoUiServerList is a list of FalcoUiServers
type FalcoUiServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FalcoUiServer CRD objects
	Items []FalcoUiServer `json:"items,omitempty"`
}
