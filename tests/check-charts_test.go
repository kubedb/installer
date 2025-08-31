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
	"errors"
	"path/filepath"
	"runtime"
	"testing"

	"kmodules.xyz/image-packer/pkg/lib"
)

func Test_CheckImageArchitectures(t *testing.T) {
	dir, err := rootDir()
	if err != nil {
		t.Error(err)
	}

	if err := lib.CheckImageArchitectures([]string{
		filepath.Join(dir, "catalog", "imagelist.yaml"),
	}, []string{
		"floragunncom/sg-elasticsearch:7.9.3-oss-47.1.0",
		"ghcr.io/appscode-images/druid:28.0.1",
		"ghcr.io/appscode-images/druid:30.0.1",
		"ghcr.io/appscode-images/druid:31.0.0",
		"ghcr.io/appscode-images/elastic:6.8.23",
		"ghcr.io/appscode-images/kibana:6.8.23",
		"ghcr.io/appscode-images/mysql:5.7.42-debian",
		"ghcr.io/appscode-images/mysql:5.7.44-oracle",
		"ghcr.io/appscode-images/mysql:8.0.36-debian",
		"ghcr.io/appscode-images/percona-xtradb-cluster:5.7.44",
		"ghcr.io/appscode-images/percona-xtradb-cluster:8.0.40",
		"ghcr.io/appscode-images/percona-xtradb-cluster:8.4.3",
		"ghcr.io/appscode-images/postgres-documentdb:15-0.102.0-ferretdb-2.0.0",
		"ghcr.io/appscode-images/postgres-documentdb:16-0.102.0-ferretdb-2.0.0",
		"ghcr.io/appscode-images/postgres-documentdb:17-0.102.0-ferretdb-2.0.0",
		"ghcr.io/appscode-images/singlestore-node:alma-8.1.32-e3d3cde6da",
		"ghcr.io/appscode-images/singlestore-node:alma-8.5.30-4f46ab16a5",
		"ghcr.io/appscode-images/singlestore-node:alma-8.5.7-bf633c1a54",
		"ghcr.io/appscode-images/singlestore-node:alma-8.7.10-95e2357384",
		"ghcr.io/appscode-images/singlestore-node:alma-8.7.21-f0b8de04d5",
		"ghcr.io/appscode-images/singlestore-node:alma-8.9.3-bfa36a984a",
		"ghcr.io/kubedb/mysql-archiver:v0.19.0_5.7.44",
		"ghcr.io/kubedb/mysql-init:5.7-v7",
		"ghcr.io/kubedb/oracle-ee:21.3.0",
		"ghcr.io/kubedb/proxysql-exporter:v1.1.0",
		"mcr.microsoft.com/mssql/server:2022-CU12-ubuntu-22.04",
		"mcr.microsoft.com/mssql/server:2022-CU14-ubuntu-22.04",
		"mcr.microsoft.com/mssql/server:2022-CU16-ubuntu-22.04",
		"mcr.microsoft.com/mssql/server:2022-CU19-ubuntu-22.04",
		"mysql/mysql-router:8.0.31",
		"percona/percona-server-mongodb:4.2.24",
		"percona/percona-server-mongodb:4.4.26",
		"percona/percona-server-mongodb:5.0.23",
		"percona/percona-server-mongodb:5.0.29",
		"percona/percona-server-mongodb:6.0.12",
		"percona/percona-server-mongodb:7.0.4",
		"postgis/postgis:11-3.3",
		"postgis/postgis:12-3.4",
		"postgis/postgis:13-3.4",
		"postgis/postgis:14-3.4",
		"postgis/postgis:15-3.4",
		"postgis/postgis:16-3.4",
		"singlestore/cluster-in-a-box:alma-8.1.32-e3d3cde6da-4.0.16-1.17.6",
		"singlestore/cluster-in-a-box:alma-8.5.22-fe61f40cd1-4.1.0-1.17.11",
		"singlestore/cluster-in-a-box:alma-8.5.7-bf633c1a54-4.0.17-1.17.8",
		"singlestore/cluster-in-a-box:alma-8.7.10-95e2357384-4.1.0-1.17.14",
	}, nil); err != nil {
		t.Errorf("CheckImageArchitectures() error = %v", err)
	}
}

func rootDir() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("failed to locate root dir")
	}

	return filepath.Clean(filepath.Join(filepath.Dir(file), "..")), nil
}
