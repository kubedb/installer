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
	ResourceKindKubedbCatalog = "KubedbCatalog"
	ResourceKubedbCatalog     = "kubedbcatalog"
	ResourceKubedbCatalogs    = "kubedbcatalogs"
)

// KubedbCatalog defines the schama for KubeDB Operator Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubedbcatalogs,singular=kubedbcatalog,categories={kubedb,appscode}
type KubedbCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubedbCatalogSpec `json:"spec,omitempty"`
}

// KubedbCatalogSpec is the schema for kubedb-catalog chart values file
type KubedbCatalogSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`
	//+optional
	Proxies        RegistryProxies     `json:"proxies"`
	FeatureGates   map[string]bool     `json:"featureGates"`
	Psp            PSP                 `json:"psp"`
	SkipDeprecated bool                `json:"skipDeprecated"`
	EnableVersions map[string][]string `json:"enableVersions"`
}

type RegistryProxies struct {
	// company/bin:1.23
	//+optional
	DockerHub string `json:"dockerHub"`
	// alpine, nginx etc.
	//+optional
	DockerLibrary string `json:"dockerLibrary"`
	// ghcr.io
	//+optional
	GHCR string `json:"ghcr"`
	// registry.k8s.io
	//+optional
	Kubernetes string `json:"kubernetes"`
	// mcr.microsoft.com
	//+optional
	Microsoft string `json:"microsoft"`
	// r.appscode.com
	//+optional
	AppsCode string `json:"appscode"`
}

type PSP struct {
	Enabled bool `json:"enabled"`
	//+optional
	Elasticsearch PSPElasticsearch `json:"elasticsearch"`
	//+optional
	Mariadb PSPMariadb `json:"mariadb"`
	//+optional
	Memcached PSPMemcached `json:"memcached"`
	//+optional
	Mongodb PSPMongodb `json:"mongodb"`
	//+optional
	Mysql PSPMysql `json:"mysql"`
	//+optional
	Perconaxtradb PSPPerconaxtradb `json:"perconaxtradb"`
	//+optional
	Postgres PSPPostgres `json:"postgres"`
	//+optional
	Proxysql PSPProxysql `json:"proxysql"`
	//+optional
	Redis PSPRedis `json:"redis"`
	//+optional
	Kafka PSPKafka `json:"kafka"`
}

type PSPElasticsearch struct {
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation"`
	Privileged               bool `json:"privileged"`
}

type PSPMariadb struct {
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation"`
	Privileged               bool `json:"privileged"`
}

type PSPMemcached struct {
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation"`
	Privileged               bool `json:"privileged"`
}

type PSPMongodb struct {
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation"`
	Privileged               bool `json:"privileged"`
}

type PSPMysql struct {
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation"`
	Privileged               bool `json:"privileged"`
}

type PSPPerconaxtradb struct {
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation"`
	Privileged               bool `json:"privileged"`
}

type PSPPostgres struct {
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation"`
	Privileged               bool `json:"privileged"`
}

type PSPProxysql struct {
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation"`
	Privileged               bool `json:"privileged"`
}

type PSPRedis struct {
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation"`
	Privileged               bool `json:"privileged"`
}

type PSPKafka struct {
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation"`
	Privileged               bool `json:"privileged"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubedbCatalogList is a list of KubedbCatalogs
type KubedbCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubedbCatalog CRD objects
	Items []KubedbCatalog `json:"items,omitempty"`
}
