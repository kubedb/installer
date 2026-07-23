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

if [ -z "${IMAGE_REGISTRY}" ]; then
    echo "IMAGE_REGISTRY is not set"
    exit 1
fi

TARBALL=${1:-}
tar -zxvf $TARBALL

CMD="./crane"

$CMD push --allow-nondistributable-artifacts --insecure images/tianon-toybox-0.8.11.tar $IMAGE_REGISTRY/tianon/toybox:0.8.11
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-kubectl-nonroot-1.34.tar $IMAGE_REGISTRY/appscode/kubectl-nonroot:1.34
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-petset-v0.1.0.tar $IMAGE_REGISTRY/appscode/petset:v0.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-sidekick-v0.0.15.tar $IMAGE_REGISTRY/appscode/sidekick:v0.0.15
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-cassandra-medusa-plugin-v0.13.0.tar $IMAGE_REGISTRY/kubedb/cassandra-medusa-plugin:v0.13.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-clickhouse-backup-plugin-v0.3.0.tar $IMAGE_REGISTRY/kubedb/clickhouse-backup-plugin:v0.3.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-dashboard-restic-plugin-v0.24.0.tar $IMAGE_REGISTRY/kubedb/dashboard-restic-plugin:v0.24.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-restic-plugin-v0.29.0.tar $IMAGE_REGISTRY/kubedb/elasticsearch-restic-plugin:v0.29.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-autoscaler-v0.51.0.tar $IMAGE_REGISTRY/kubedb/kubedb-autoscaler:v0.51.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-courier-v0.6.0.tar $IMAGE_REGISTRY/kubedb/kubedb-courier:v0.6.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-crd-manager-v0.21.0.tar $IMAGE_REGISTRY/kubedb/kubedb-crd-manager:v0.21.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-gitops-v0.14.0.tar $IMAGE_REGISTRY/kubedb/kubedb-gitops:v0.14.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-kibana-v0.42.0.tar $IMAGE_REGISTRY/kubedb/kubedb-kibana:v0.42.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-manifest-plugin-v0.29.0.tar $IMAGE_REGISTRY/kubedb/kubedb-manifest-plugin:v0.29.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-ops-manager-v0.53.0.tar $IMAGE_REGISTRY/kubedb/kubedb-ops-manager:v0.53.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-provisioner-v0.66.0.tar $IMAGE_REGISTRY/kubedb/kubedb-provisioner:v0.66.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-schema-manager-v0.42.0.tar $IMAGE_REGISTRY/kubedb/kubedb-schema-manager:v0.42.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-ui-server-v0.42.0.tar $IMAGE_REGISTRY/kubedb/kubedb-ui-server:v0.42.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-verifier-v0.17.0.tar $IMAGE_REGISTRY/kubedb/kubedb-verifier:v0.17.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-webhook-server-v0.42.0.tar $IMAGE_REGISTRY/kubedb/kubedb-webhook-server:v0.42.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mariadb-csi-snapshotter-plugin-v0.26.0.tar $IMAGE_REGISTRY/kubedb/mariadb-csi-snapshotter-plugin:v0.26.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mariadb-restic-plugin-v0.24.0_10.11.6-jammy.tar $IMAGE_REGISTRY/kubedb/mariadb-restic-plugin:v0.24.0_10.11.6-jammy
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mariadb-restic-plugin-v0.24.0_10.4.32-focal.tar $IMAGE_REGISTRY/kubedb/mariadb-restic-plugin:v0.24.0_10.4.32-focal
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mariadb-restic-plugin-v0.24.0_10.6.16-focal.tar $IMAGE_REGISTRY/kubedb/mariadb-restic-plugin:v0.24.0_10.6.16-focal
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mariadb-restic-plugin-v0.24.0_11.1.3-jammy.tar $IMAGE_REGISTRY/kubedb/mariadb-restic-plugin:v0.24.0_11.1.3-jammy
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-csi-snapshotter-plugin-v0.27.0.tar $IMAGE_REGISTRY/kubedb/mongodb-csi-snapshotter-plugin:v0.27.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-restic-plugin-v0.29.0_4.2.3.tar $IMAGE_REGISTRY/kubedb/mongodb-restic-plugin:v0.29.0_4.2.3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-restic-plugin-v0.29.0_4.4.6.tar $IMAGE_REGISTRY/kubedb/mongodb-restic-plugin:v0.29.0_4.4.6
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-restic-plugin-v0.29.0_5.0.15.tar $IMAGE_REGISTRY/kubedb/mongodb-restic-plugin:v0.29.0_5.0.15
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-restic-plugin-v0.29.0_5.0.3.tar $IMAGE_REGISTRY/kubedb/mongodb-restic-plugin:v0.29.0_5.0.3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-restic-plugin-v0.29.0_6.0.5.tar $IMAGE_REGISTRY/kubedb/mongodb-restic-plugin:v0.29.0_6.0.5
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-restic-plugin-v0.29.0_8.0.3.tar $IMAGE_REGISTRY/kubedb/mongodb-restic-plugin:v0.29.0_8.0.3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mssqlserver-walg-plugin-v0.20.0.tar $IMAGE_REGISTRY/kubedb/mssqlserver-walg-plugin:v0.20.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-csi-snapshotter-plugin-v0.27.0.tar $IMAGE_REGISTRY/kubedb/mysql-csi-snapshotter-plugin:v0.27.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-restic-plugin-v0.29.0_5.7.25.tar $IMAGE_REGISTRY/kubedb/mysql-restic-plugin:v0.29.0_5.7.25
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-restic-plugin-v0.29.0_8.0.21.tar $IMAGE_REGISTRY/kubedb/mysql-restic-plugin:v0.29.0_8.0.21
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-restic-plugin-v0.29.0_8.0.3.tar $IMAGE_REGISTRY/kubedb/mysql-restic-plugin:v0.29.0_8.0.3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-restic-plugin-v0.29.0_8.4.2.tar $IMAGE_REGISTRY/kubedb/mysql-restic-plugin:v0.29.0_8.4.2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-restic-plugin-v0.29.0_9.0.1.tar $IMAGE_REGISTRY/kubedb/mysql-restic-plugin:v0.29.0_9.0.1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-neo4j-backup-plugin-v0.2.0.tar $IMAGE_REGISTRY/kubedb/neo4j-backup-plugin:v0.2.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-csi-snapshotter-plugin-v0.27.0.tar $IMAGE_REGISTRY/kubedb/postgres-csi-snapshotter-plugin:v0.27.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-restic-plugin-v0.29.0_12.17.tar $IMAGE_REGISTRY/kubedb/postgres-restic-plugin:v0.29.0_12.17
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-restic-plugin-v0.29.0_14.10.tar $IMAGE_REGISTRY/kubedb/postgres-restic-plugin:v0.29.0_14.10
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-restic-plugin-v0.29.0_16.4.tar $IMAGE_REGISTRY/kubedb/postgres-restic-plugin:v0.29.0_16.4
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-restic-plugin-v0.29.0_17.2.tar $IMAGE_REGISTRY/kubedb/postgres-restic-plugin:v0.29.0_17.2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-restic-plugin-v0.29.0_18.2.tar $IMAGE_REGISTRY/kubedb/postgres-restic-plugin:v0.29.0_18.2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-provider-aws-v0.27.0.tar $IMAGE_REGISTRY/kubedb/provider-aws:v0.27.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-provider-azure-v0.27.0.tar $IMAGE_REGISTRY/kubedb/provider-azure:v0.27.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-provider-gcp-v0.27.0.tar $IMAGE_REGISTRY/kubedb/provider-gcp:v0.27.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-qdrant-restic-plugin-v0.2.0.tar $IMAGE_REGISTRY/kubedb/qdrant-restic-plugin:v0.2.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-redis-restic-plugin-v0.29.0.tar $IMAGE_REGISTRY/kubedb/redis-restic-plugin:v0.29.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-singlestore-restic-plugin-v0.24.0_alma-8.1.32-e3d3cde6da.tar $IMAGE_REGISTRY/kubedb/singlestore-restic-plugin:v0.24.0_alma-8.1.32-e3d3cde6da
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-singlestore-restic-plugin-v0.24.0_alma-8.5.7-bf633c1a54.tar $IMAGE_REGISTRY/kubedb/singlestore-restic-plugin:v0.24.0_alma-8.5.7-bf633c1a54
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-xtrabackup-restic-plugin-v0.14.0_2.4.29.tar $IMAGE_REGISTRY/kubedb/xtrabackup-restic-plugin:v0.14.0_2.4.29
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-xtrabackup-restic-plugin-v0.14.0_8.0.35.tar $IMAGE_REGISTRY/kubedb/xtrabackup-restic-plugin:v0.14.0_8.0.35
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-xtrabackup-restic-plugin-v0.14.0_8.1.0.tar $IMAGE_REGISTRY/kubedb/xtrabackup-restic-plugin:v0.14.0_8.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-xtrabackup-restic-plugin-v0.14.0_8.2.0.tar $IMAGE_REGISTRY/kubedb/xtrabackup-restic-plugin:v0.14.0_8.2.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-xtrabackup-restic-plugin-v0.14.0_8.4.0.tar $IMAGE_REGISTRY/kubedb/xtrabackup-restic-plugin:v0.14.0_8.4.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-zookeeper-restic-plugin-v0.21.0.tar $IMAGE_REGISTRY/kubedb/zookeeper-restic-plugin:v0.21.0
