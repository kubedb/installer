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
	"os"

	configdata "kubedb.dev/apimachinery/apis/config/v1alpha1"
	dbv1 "kubedb.dev/apimachinery/apis/kubedb/v1"
	dbv1a2 "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
)

func main() {
	r := configdata.LicenseRestrictions{
		dbv1a2.ResourceKindDruid: {
			VersionConstraint: "*",
		},
		dbv1.ResourceKindElasticsearch: {
			VersionConstraint: "*",
			Distributions:     []string{"OpenSearch"},
		},
		dbv1a2.ResourceKindFerretDB: {
			VersionConstraint: "*",
		},
		dbv1.ResourceKindKafka: {
			VersionConstraint: "*",
		},
		dbv1.ResourceKindMariaDB: {
			VersionConstraint: "*",
		},
		dbv1.ResourceKindMemcached: {
			VersionConstraint: "*",
		},
		dbv1.ResourceKindMySQL: {
			VersionConstraint: "*",
		},
		dbv1.ResourceKindPerconaXtraDB: {
			VersionConstraint: "*",
		},
		dbv1.ResourceKindPgBouncer: {
			VersionConstraint: "*",
		},
		dbv1a2.ResourceKindPgpool: {
			VersionConstraint: "*",
		},
		dbv1.ResourceKindPostgres: {
			VersionConstraint: "*",
			Distributions:     []string{"Official"},
		},
		dbv1.ResourceKindProxySQL: {
			VersionConstraint: "*",
		},
		dbv1a2.ResourceKindRabbitmq: {
			VersionConstraint: "*",
		},
		dbv1.ResourceKindRedis: {
			VersionConstraint: "<= 7.40",
		},
		dbv1a2.ResourceKindSolr: {
			VersionConstraint: "*",
		},
		dbv1a2.ResourceKindZooKeeper: {
			VersionConstraint: "*",
		},
	}
	je := json.NewEncoder(os.Stdout)
	je.SetEscapeHTML(false)
	err := je.Encode(r)
	if err != nil {
		panic(err)
	}
}
