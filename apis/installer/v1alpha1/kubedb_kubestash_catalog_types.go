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
	ResourceKindKubedbKubestashCatalog = "KubedbKubestashCatalog"
	ResourceKubedbKubestashCatalog     = "kubedbkubestashcatalog"
	ResourceKubedbKubestashCatalogs    = "kubedbkubestashcatalogs"
)

// KubedbKubestashCatalog defines the schema for Stash Catalog chart.

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
	Proxies        RegistryProxies         `json:"proxies"`
	WaitTimeout    int64                   `json:"waitTimeout"`
	Elasticsearch  StashElasticsearchSpec  `json:"elasticsearch"`
	Opensearch     StashOpensearchSpec     `json:"opensearch"`
	Kubedbmanifest StashKubedbmanifestSpec `json:"kubedbmanifest"`
	KubeDump       KubeDumpSpec            `json:"kubedump"`
	Mongodb        StashMongodbSpec        `json:"mongodb"`
	Mysql          StashMysqlSpec          `json:"mysql"`
	Pvc            StashPvcSpec            `json:"pvc"`
	Redis          StashRedisSpec          `json:"redis"`
	Volumesnapshot StashVolumesnapshotSpec `json:"volumesnapshot"`
	Workload       StashWorkloadSpec       `json:"workload"`
}

// StashElasticsearchSpec is the schema for Stash Elasticsearch values file
type StashElasticsearchSpec struct {
	Enabled bool                 `json:"enabled"`
	Backup  ElasticsearchBackup  `json:"backup"`
	Restore ElasticsearchRestore `json:"restore"`
}

type ElasticsearchBackup struct {
	//+optional
	Args string `json:"args"`
}

type ElasticsearchRestore struct {
	//+optional
	Args string `json:"args"`
}

// StashOpensearchSpec is the schema for Stash Opensearch values file
type StashOpensearchSpec struct {
	Enabled bool              `json:"enabled"`
	Backup  OpensearchBackup  `json:"backup"`
	Restore OpensearchRestore `json:"restore"`
}

type OpensearchBackup struct {
	//+optional
	Args string `json:"args"`
}

type OpensearchRestore struct {
	//+optional
	Args string `json:"args"`
}

type StashKubedbmanifestSpec struct {
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

// StashMongodbSpec is the schema for Stash MongoDB values file
type StashMongodbSpec struct {
	Enabled        bool           `json:"enabled"`
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

// StashMysqlSpec is the schema for Stash MySQL values file
type StashMysqlSpec struct {
	Enabled bool         `json:"enabled"`
	Backup  MySQLBackup  `json:"backup"`
	Restore MySQLRestore `json:"restore"`
}

type MySQLBackup struct {
	// +optional
	Args string `json:"args"`
}

type MySQLRestore struct {
	// +optional
	Args string `json:"args"`
}

type StashPvcSpec struct {
	Enabled bool `json:"enabled"`
}

// StashRedisSpec is the schema for Stash Redis values file
type StashRedisSpec struct {
	Enabled bool         `json:"enabled"`
	Backup  RedisBackup  `json:"backup"`
	Restore RedisRestore `json:"restore"`
}

type RedisBackup struct {
	// +optional
	Args string `json:"args"`
}

type RedisRestore struct {
	// +optional
	Args string `json:"args"`
}

type StashVolumesnapshotSpec struct {
	Enabled bool `json:"enabled"`
}

type StashWorkloadSpec struct {
	Enabled bool `json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubedbKubestashCatalogList is a list of KubedbKubestashCatalogs
type KubedbKubestashCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StashPostgres CRD objects
	Items []KubedbKubestashCatalog `json:"items,omitempty"`
}