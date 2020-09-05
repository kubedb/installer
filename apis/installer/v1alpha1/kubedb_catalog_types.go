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
	ResourceKindKubeDBCatalog = "KubeDBCatalog"
	ResourceKubeDBCatalog     = "kubedbcatalog"
	ResourceKubeDBCatalogs    = "kubedbcatalogs"
)

// KubeDBCatalog defines the schama for KubeDB Operator Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubedbcatalogs,singular=kubedbcatalog,categories={kubedb,appscode}
type KubeDBCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeDBCatalogSpec `json:"spec,omitempty"`
}

// KubeDBCatalogSpec is the spec for redis version
type KubeDBCatalogSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string      `json:"fullnameOverride"`
	Image            RegistryRef `json:"image"`
	Catalog          Catalog     `json:"catalog"`
	SkipDeprecated   bool        `json:"skipDeprecated"`
}

type RegistryRef struct {
	Registry string `json:"registry"`
}

type Catalog struct {
	//+optional
	Elasticsearch bool `json:"elasticsearch"`
	//+optional
	Etcd bool `json:"etcd"`
	//+optional
	Memcached bool `json:"memcached"`
	//+optional
	Mongo bool `json:"mongo"`
	//+optional
	Mysql bool `json:"mysql"`
	//+optional
	Perconaxtradb bool `json:"perconaxtradb"`
	//+optional
	Pgbouncer bool `json:"pgbouncer"`
	//+optional
	Postgres bool `json:"postgres"`
	//+optional
	Proxysql bool `json:"proxysql"`
	//+optional
	Redis bool `json:"redis"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeDBCatalogList is a list of KubeDBCatalogs
type KubeDBCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubeDBCatalog CRD objects
	Items []KubeDBCatalog `json:"items,omitempty"`
}
