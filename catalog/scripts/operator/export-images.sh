#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the AppsCode Community License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -x

mkdir -p images

OS=$(uname -o)
if [ "${OS}" = "GNU/Linux" ]; then
    OS=Linux
fi
ARCH=$(uname -m)
if [ "${ARCH}" = "aarch64" ]; then
    ARCH=arm64
fi
curl -sL "https://github.com/google/go-containerregistry/releases/latest/download/go-containerregistry_${OS}_${ARCH}.tar.gz" >/tmp/go-containerregistry.tar.gz
tar -zxvf /tmp/go-containerregistry.tar.gz -C /tmp/
mv /tmp/crane images

CMD="./images/crane"

$CMD pull --allow-nondistributable-artifacts --insecure docker.io/tianon/toybox:0.8.11 images/tianon-toybox-0.8.11.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode/kubectl-nonroot:1.34 images/appscode-kubectl-nonroot-1.34.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode/petset:v0.1.0 images/appscode-petset-v0.1.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode/sidekick:v0.0.15 images/appscode-sidekick-v0.0.15.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/cassandra-medusa-plugin:v0.13.0 images/kubedb-cassandra-medusa-plugin-v0.13.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/clickhouse-backup-plugin:v0.3.0 images/kubedb-clickhouse-backup-plugin-v0.3.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/dashboard-restic-plugin:v0.24.0 images/kubedb-dashboard-restic-plugin-v0.24.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch-restic-plugin:v0.29.0 images/kubedb-elasticsearch-restic-plugin-v0.29.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-autoscaler:v0.51.0 images/kubedb-kubedb-autoscaler-v0.51.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-courier:v0.6.0 images/kubedb-kubedb-courier-v0.6.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-crd-manager:v0.21.0 images/kubedb-kubedb-crd-manager-v0.21.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-gitops:v0.14.0 images/kubedb-kubedb-gitops-v0.14.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-kibana:v0.42.0 images/kubedb-kubedb-kibana-v0.42.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-manifest-plugin:v0.29.0 images/kubedb-kubedb-manifest-plugin-v0.29.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-ops-manager:v0.53.0 images/kubedb-kubedb-ops-manager-v0.53.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-provisioner:v0.66.0 images/kubedb-kubedb-provisioner-v0.66.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-schema-manager:v0.42.0 images/kubedb-kubedb-schema-manager-v0.42.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-ui-server:v0.42.0 images/kubedb-kubedb-ui-server-v0.42.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-verifier:v0.17.0 images/kubedb-kubedb-verifier-v0.17.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-webhook-server:v0.42.0 images/kubedb-kubedb-webhook-server-v0.42.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-csi-snapshotter-plugin:v0.26.0 images/kubedb-mariadb-csi-snapshotter-plugin-v0.26.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-restic-plugin:v0.24.0_10.11.6-jammy images/kubedb-mariadb-restic-plugin-v0.24.0_10.11.6-jammy.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-restic-plugin:v0.24.0_10.4.32-focal images/kubedb-mariadb-restic-plugin-v0.24.0_10.4.32-focal.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-restic-plugin:v0.24.0_10.6.16-focal images/kubedb-mariadb-restic-plugin-v0.24.0_10.6.16-focal.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-restic-plugin:v0.24.0_11.1.3-jammy images/kubedb-mariadb-restic-plugin-v0.24.0_11.1.3-jammy.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-csi-snapshotter-plugin:v0.27.0 images/kubedb-mongodb-csi-snapshotter-plugin-v0.27.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-restic-plugin:v0.29.0_4.2.3 images/kubedb-mongodb-restic-plugin-v0.29.0_4.2.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-restic-plugin:v0.29.0_4.4.6 images/kubedb-mongodb-restic-plugin-v0.29.0_4.4.6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-restic-plugin:v0.29.0_5.0.15 images/kubedb-mongodb-restic-plugin-v0.29.0_5.0.15.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-restic-plugin:v0.29.0_5.0.3 images/kubedb-mongodb-restic-plugin-v0.29.0_5.0.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-restic-plugin:v0.29.0_6.0.5 images/kubedb-mongodb-restic-plugin-v0.29.0_6.0.5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-restic-plugin:v0.29.0_8.0.3 images/kubedb-mongodb-restic-plugin-v0.29.0_8.0.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mssqlserver-walg-plugin:v0.20.0 images/kubedb-mssqlserver-walg-plugin-v0.20.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-csi-snapshotter-plugin:v0.27.0 images/kubedb-mysql-csi-snapshotter-plugin-v0.27.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-restic-plugin:v0.29.0_5.7.25 images/kubedb-mysql-restic-plugin-v0.29.0_5.7.25.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-restic-plugin:v0.29.0_8.0.21 images/kubedb-mysql-restic-plugin-v0.29.0_8.0.21.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-restic-plugin:v0.29.0_8.0.3 images/kubedb-mysql-restic-plugin-v0.29.0_8.0.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-restic-plugin:v0.29.0_8.4.2 images/kubedb-mysql-restic-plugin-v0.29.0_8.4.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-restic-plugin:v0.29.0_9.0.1 images/kubedb-mysql-restic-plugin-v0.29.0_9.0.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/neo4j-backup-plugin:v0.2.0 images/kubedb-neo4j-backup-plugin-v0.2.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-csi-snapshotter-plugin:v0.27.0 images/kubedb-postgres-csi-snapshotter-plugin-v0.27.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-restic-plugin:v0.29.0_12.17 images/kubedb-postgres-restic-plugin-v0.29.0_12.17.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-restic-plugin:v0.29.0_14.10 images/kubedb-postgres-restic-plugin-v0.29.0_14.10.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-restic-plugin:v0.29.0_16.4 images/kubedb-postgres-restic-plugin-v0.29.0_16.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-restic-plugin:v0.29.0_17.2 images/kubedb-postgres-restic-plugin-v0.29.0_17.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-restic-plugin:v0.29.0_18.2 images/kubedb-postgres-restic-plugin-v0.29.0_18.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/provider-aws:v0.27.0 images/kubedb-provider-aws-v0.27.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/provider-azure:v0.27.0 images/kubedb-provider-azure-v0.27.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/provider-gcp:v0.27.0 images/kubedb-provider-gcp-v0.27.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/qdrant-restic-plugin:v0.2.0 images/kubedb-qdrant-restic-plugin-v0.2.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/redis-restic-plugin:v0.29.0 images/kubedb-redis-restic-plugin-v0.29.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/singlestore-restic-plugin:v0.24.0_alma-8.1.32-e3d3cde6da images/kubedb-singlestore-restic-plugin-v0.24.0_alma-8.1.32-e3d3cde6da.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/singlestore-restic-plugin:v0.24.0_alma-8.5.7-bf633c1a54 images/kubedb-singlestore-restic-plugin-v0.24.0_alma-8.5.7-bf633c1a54.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/xtrabackup-restic-plugin:v0.14.0_2.4.29 images/kubedb-xtrabackup-restic-plugin-v0.14.0_2.4.29.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/xtrabackup-restic-plugin:v0.14.0_8.0.35 images/kubedb-xtrabackup-restic-plugin-v0.14.0_8.0.35.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/xtrabackup-restic-plugin:v0.14.0_8.1.0 images/kubedb-xtrabackup-restic-plugin-v0.14.0_8.1.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/xtrabackup-restic-plugin:v0.14.0_8.2.0 images/kubedb-xtrabackup-restic-plugin-v0.14.0_8.2.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/xtrabackup-restic-plugin:v0.14.0_8.4.0 images/kubedb-xtrabackup-restic-plugin-v0.14.0_8.4.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/zookeeper-restic-plugin:v0.21.0 images/kubedb-zookeeper-restic-plugin-v0.21.0.tar

tar -czvf images.tar.gz images
