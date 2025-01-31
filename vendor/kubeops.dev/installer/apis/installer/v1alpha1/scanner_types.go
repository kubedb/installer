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
	ResourceKindScanner = "Scanner"
	ResourceScanner     = "scanner"
	ResourceScanners    = "scanners"
)

// Scanner defines the schama for Scanner Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=scanners,singular=scanner,categories={kubeops,appscode}
type Scanner struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScannerSpec `json:"spec,omitempty"`
}

// ScannerSpec is the schema for Scanner Operator values file
type ScannerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string          `json:"fullnameOverride"`
	ReplicaCount     int32           `json:"replicaCount"`
	RegistryFQDN     string          `json:"registryFQDN"`
	App              Container       `json:"app"`
	Etcd             EtcdContainer   `json:"etcd"`
	Kine             Container       `json:"kine"`
	Cacher           CacherContainer `json:"cacher"`
	ImagePullPolicy  string          `json:"imagePullPolicy"`
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
	Apiserver          ScannerserverSpec        `json:"apiserver"`
	Monitoring         Monitoring               `json:"monitoring"`
	Dashboard          GrafanaDashboard         `json:"dashboard"`
	Grafana            ObjectReference          `json:"grafana"`
	Nats               ScannerNATS              `json:"nats"`
	// +optional
	License string `json:"license"`

	// +optional
	ScanRequestTTLAfterFinished metav1.Duration  `json:"scanRequestTTLAfterFinished"`
	ScanReportTTLAfterOutdated  metav1.Duration  `json:"scanReportTTLAfterOutdated"`
	Workspace                   ScannerWorkspace `json:"workspace"`
}

type ScannerserverSpec struct {
	ApiserverSpec `json:",inline"`
	DB            ApiserverDB `json:"db"`
}

type GrafanaDashboard struct {
	Enabled    bool                `json:"enabled"`
	FolderID   int                 `json:"folderID"`
	Overwrite  bool                `json:"overwrite"`
	Templatize DashboardTemplatize `json:"templatize"`
}

type DashboardTemplatize struct {
	Title      bool `json:"title"`
	Datasource bool `json:"datasource"`
}

type LocalObjectReference struct {
	Name string `json:"name"`
}

type ObjectReference struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type EtcdContainer struct {
	Container    `json:",inline,omitempty"`
	ServingCerts `json:"servingCerts"`
}

type CacherContainer struct {
	Container `json:",inline,omitempty"`
	Enable    bool   `json:"enable"`
	Schedule  string `json:"schedule"`
}

type Persistence struct {
	Size string `json:"size"`
}

type NatsAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ScannerNATS struct {
	Addr string   `json:"addr"`
	Auth NatsAuth `json:"auth"`
}

type ScannerWorkspace struct {
	Namespace string `json:"namespace"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScannerList is a list of Scanners
type ScannerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Scanner CRD objects
	Items []Scanner `json:"items,omitempty"`
}
