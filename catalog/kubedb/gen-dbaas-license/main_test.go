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

package main

import (
	"encoding/json"
	"reflect"
	"slices"
	"testing"

	configdata "kubedb.dev/apimachinery/apis/config/v1alpha1"

	"github.com/Masterminds/semver/v3"
)

func TestOpenSourceRestrictions(t *testing.T) {
	restrictions := openSourceRestrictions()
	for _, kind := range []string{
		"Aerospike", "Cassandra", "ClickHouse", "DocumentDB", "Druid", "Etcd",
		"Ignite", "Kafka", "MariaDB", "Memcached", "Milvus", "MySQL",
		"PerconaXtraDB", "PgBouncer", "Pgpool", "ProxySQL", "Qdrant",
		"RabbitMQ", "Solr", "Weaviate", "ZooKeeper",
	} {
		t.Run(kind, func(t *testing.T) {
			if !allowed(t, restrictions, kind, "1.2.3", "") {
				t.Errorf("%s is not enabled", kind)
			}
		})
	}

	for _, tt := range []struct {
		kind, distribution, version string
		want                        bool
	}{
		{"Elasticsearch", "OpenSearch", "3.8.0", true},
		{"Elasticsearch", "ElasticStack", "9.5.3", false},
		{"Elasticsearch", "SearchGuard", "7.9.3", false},
		{"Postgres", "Official", "18.3", true},
		{"Postgres", "PostGIS", "16.2", true},
		{"Postgres", "DocumentDB", "17.4", true},
		{"Postgres", "Percona", "17.9", true},
		{"Postgres", "TimescaleDB", "11.11", true},
		{"Postgres", "TimescaleDB", "12.6", true},
		{"Postgres", "TimescaleDB", "13.14", true},
		{"Postgres", "TimescaleDB", "14.11", true},
		{"Postgres", "TimescaleDB", "17.9", false},
		{"Postgres", "AppsCode", "18.6", false},
		{"Postgres", "KubeDB", "18.6", false},
		{"Redis", "Official", "7.2.4", true},
		{"Redis", "Official", "7.2.6", true},
		{"Redis", "Official", "7.3.0", false},
		{"Redis", "Official", "7.4.0", false},
		{"Redis", "Official", "7.4.6", false},
		{"Redis", "Official", "8.0.0", true},
		{"Redis", "Official", "8.2.2", true},
		{"Redis", "Valkey", "7.2.4", true},
		{"Redis", "Valkey", "8.1.1", true},
		{"MongoDB", "Official", "8.0.17", false},
		{"MongoDB", "Percona", "8.0.17", false},
		{"Hazelcast", "", "5.5.6", false},
		{"Neo4j", "", "5.26.28", false},
		{"MSSQLServer", "", "2022.0.0", false},
	} {
		t.Run(tt.kind+"/"+tt.distribution+"/"+tt.version, func(t *testing.T) {
			if got := allowed(t, restrictions, tt.kind, tt.version, tt.distribution); got != tt.want {
				t.Errorf("allowed = %t, want %t", got, tt.want)
			}
		})
	}
}

func allowed(t *testing.T, restrictions configdata.LicenseRestrictions, kind, version, distribution string) bool {
	t.Helper()
	v, err := semver.NewVersion(version)
	if err != nil {
		t.Fatal(err)
	}
	for _, restriction := range restrictions[kind] {
		constraint, err := semver.NewConstraint(restriction.VersionConstraint)
		if err != nil {
			t.Fatalf("%s: invalid constraint %q: %v", kind, restriction.VersionConstraint, err)
		}
		if constraint.Check(v) && (len(restriction.Distributions) == 0 || slices.Contains(restriction.Distributions, distribution)) {
			return true
		}
	}
	return false
}

func TestLegacyOutput(t *testing.T) {
	restrictions := openSourceRestrictions()
	legacy := restrictions.ToV1()
	for kind, entries := range restrictions {
		if !reflect.DeepEqual(legacy[kind], entries[0]) {
			t.Errorf("%s: legacy output must contain the first restriction", kind)
		}
	}
	data, err := json.Marshal(legacy)
	if err != nil {
		t.Fatal(err)
	}
	var decoded configdata.LicenseRestrictions
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	for kind, entry := range legacy {
		if !reflect.DeepEqual(decoded[kind], configdata.Restrictions{entry}) {
			t.Errorf("%s: legacy JSON did not round-trip", kind)
		}
	}
}
