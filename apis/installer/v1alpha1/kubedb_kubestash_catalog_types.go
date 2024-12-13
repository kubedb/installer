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
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindKubedbKubestashCatalog = "KubedbKubestashCatalog"
	ResourceKubedbKubestashCatalog     = "kubedbkubestashcatalog"
	ResourceKubedbKubestashCatalogs    = "kubedbkubestashcatalogs"
)

// KubedbKubestashCatalog defines the schema for KubeStash Catalog chart.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubedbkubestashcatalogs,singular=kubedbkubestashcatalog,categories={stash,appscode}
type KubedbKubestashCatalog struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubedbKubestashCatalogSpec `json:"spec,omitempty"`
}

// KubedbKubestashCatalogSpec is the schema for Stash Postgres values file
type KubedbKubestashCatalogSpec struct {
	//+optional
	Proxies        shared.RegistryProxies      `json:"proxies"`
	FeatureGates   map[string]bool             `json:"featureGates"`
	WaitTimeout    int64                       `json:"waitTimeout"`
	Druid          KubestashDatabaseSpec       `json:"druid"`
	Elasticsearch  KubestashDatabaseSpec       `json:"elasticsearch"`
	Opensearch     KubestashDatabaseSpec       `json:"opensearch"`
	Kubedbmanifest KubestashKubedbmanifestSpec `json:"kubedbmanifest"`
	Mariadb        KubestashDatabaseSpec       `json:"mariadb"`
	Mongodb        KubestashMongodbSpec        `json:"mongodb"`
	MSSQLServer    KubestashMongodbSpec        `json:"mssqlserver"`
	Mysql          KubestashDatabaseSpec       `json:"mysql"`
	Redis          KubestashDatabaseSpec       `json:"redis"`
	Postgres       KubestashPostgresSpec       `json:"postgres"`
	Singlestore    KubestashDatabaseSpec       `json:"singlestore"`
	ZooKeeper      KubestashDatabaseSpec       `json:"zookeeper"`
	Kubedbverifier KubestashVerifierSpec       `json:"kubedbverifier"`
}

// KubestashDatabaseSpec is the schema for DB values file
type KubestashDatabaseSpec struct {
	Backup  DatabaseBackup  `json:"backup"`
	Restore DatabaseRestore `json:"restore"`
}

type DatabaseBackup struct {
	//+optional
	Args string `json:"args"`
}

type DatabaseRestore struct {
	//+optional
	Args string `json:"args"`
}

type KubestashKubedbmanifestSpec struct {
	Enabled bool `json:"enabled"`
}

type KubeDumpSpec struct {
	Enabled bool           `json:"enabled"`
	Backup  KubeDumpBackup `json:"backup"`
}

type KubeDumpBackup struct {
	Sanitize          bool   `json:"sanitize"`
	LabelSelector     string `json:"labelSelector"`
	IncludeDependants bool   `json:"includeDependants"`
}

// KubestashMongodbSpec is the schema for KubeStash MongoDB values file
type KubestashMongodbSpec struct {
	MaxConcurrency int32          `json:"maxConcurrency"`
	Backup         MongoDBBackup  `json:"backup"`
	Restore        MongoDBRestore `json:"restore"`
}

type MongoDBBackup struct {
	// +optional
	Args string `json:"args"`
}

type MongoDBRestore struct {
	// +optional
	Args string `json:"args"`
}

// KubestashPostgresSpec is the schema for KubeStash Postgres values file
type KubestashPostgresSpec struct {
	Backup  PostgresBackup  `json:"backup"`
	Restore PostgresRestore `json:"restore"`
}

type PostgresBackup struct {
	// +optional
	CMD string `json:"cmd"`
	// +optional
	Args string `json:"args"`
}

type PostgresRestore struct {
	// +optional
	Args string `json:"args"`
}

type KubestashVerifierSpec struct {
	Enabled bool `json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubedbKubestashCatalogList is a list of KubedbKubestashCatalog
type KubedbKubestashCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StashPostgres CRD objects
	Items []KubedbKubestashCatalog `json:"items,omitempty"`
}
