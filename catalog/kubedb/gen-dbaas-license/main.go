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
		dbv1a2.ResourceKindDruid: configdata.Restriction{
			VersionConstraint: "*",
		},
		dbv1.ResourceKindElasticsearch: configdata.Restriction{
			VersionConstraint: "*",
			Distributions:     []string{"OpenSearch"},
		},
		dbv1a2.ResourceKindFerretDB: configdata.Restriction{
			VersionConstraint: "*",
		},
		dbv1.ResourceKindKafka: configdata.Restriction{
			VersionConstraint: "*",
		},
		dbv1.ResourceKindMariaDB: configdata.Restriction{
			VersionConstraint: "*",
		},
		dbv1.ResourceKindMemcached: configdata.Restriction{
			VersionConstraint: "*",
		},
		dbv1.ResourceKindMySQL: configdata.Restriction{
			VersionConstraint: "*",
		},
		dbv1.ResourceKindPerconaXtraDB: configdata.Restriction{
			VersionConstraint: "*",
		},
		dbv1.ResourceKindPgBouncer: configdata.Restriction{
			VersionConstraint: "*",
		},
		dbv1a2.ResourceKindPgpool: configdata.Restriction{
			VersionConstraint: "*",
		},
		dbv1.ResourceKindPostgres: configdata.Restriction{
			VersionConstraint: "*",
			Distributions:     []string{"Official"},
		},
		dbv1.ResourceKindProxySQL: configdata.Restriction{
			VersionConstraint: "*",
		},
		dbv1a2.ResourceKindRabbitmq: configdata.Restriction{
			VersionConstraint: "*",
		},
		dbv1.ResourceKindRedis: configdata.Restriction{
			VersionConstraint: "<= 7.40",
		},
		dbv1a2.ResourceKindSolr: configdata.Restriction{
			VersionConstraint: "*",
		},
		dbv1a2.ResourceKindZooKeeper: configdata.Restriction{
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
