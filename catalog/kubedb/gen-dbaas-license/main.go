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
	"fmt"
	"os"

	configdata "kubedb.dev/apimachinery/apis/config/v1alpha1"
	dbv1 "kubedb.dev/apimachinery/apis/kubedb/v1"
	dbv1a2 "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
)

func main() {
	r := configdata.LicenseRestrictions{
		dbv1a2.ResourceKindDruid: configdata.Restrictions{{
			VersionConstraint: "*",
		}},
		dbv1.ResourceKindElasticsearch: configdata.Restrictions{{
			VersionConstraint: "*",
			Distributions:     []string{"OpenSearch"},
		}},
		dbv1a2.ResourceKindFerretDB: configdata.Restrictions{{
			VersionConstraint: "*",
		}},
		dbv1.ResourceKindKafka: configdata.Restrictions{{
			VersionConstraint: "*",
		}},
		dbv1.ResourceKindMariaDB: configdata.Restrictions{{
			VersionConstraint: "*",
		}},
		dbv1.ResourceKindMemcached: configdata.Restrictions{{
			VersionConstraint: "*",
		}},
		dbv1.ResourceKindMySQL: configdata.Restrictions{{
			VersionConstraint: "*",
		}},
		dbv1.ResourceKindPerconaXtraDB: configdata.Restrictions{{
			VersionConstraint: "*",
		}},
		dbv1.ResourceKindPgBouncer: configdata.Restrictions{{
			VersionConstraint: "*",
		}},
		dbv1a2.ResourceKindPgpool: configdata.Restrictions{{
			VersionConstraint: "*",
		}},
		dbv1.ResourceKindPostgres: configdata.Restrictions{{
			VersionConstraint: "*",
			Distributions:     []string{"Official"},
		}},
		dbv1.ResourceKindProxySQL: configdata.Restrictions{{
			VersionConstraint: "*",
		}},
		dbv1a2.ResourceKindRabbitmq: configdata.Restrictions{{
			VersionConstraint: "*",
		}},
		dbv1.ResourceKindRedis: configdata.Restrictions{
			{
				VersionConstraint: "<= 7.2.4",
				Distributions:     []string{"Official"},
			},
			{
				VersionConstraint: ">= 7.2.5",
				Distributions:     []string{"Valkey"},
			},
		},
		dbv1a2.ResourceKindSolr: configdata.Restrictions{{
			VersionConstraint: "*",
		}},
		dbv1a2.ResourceKindZooKeeper: configdata.Restrictions{{
			VersionConstraint: "*",
		}},
	}

	fmt.Printf("License Restrictions:\n")
	je := json.NewEncoder(os.Stdout)
	je.SetEscapeHTML(false)
	err := je.Encode(r)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Printf("License Restrictions v1:\n")
	err = je.Encode(r.ToV1())
	if err != nil {
		panic(err)
	}
}
