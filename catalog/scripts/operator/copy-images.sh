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
mv /tmp/crane .

CMD="./crane"

$CMD cp --allow-nondistributable-artifacts --insecure docker.io/tianon/toybox:0.8.11 $IMAGE_REGISTRY/tianon/toybox:0.8.11
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode/kubectl-nonroot:1.34 $IMAGE_REGISTRY/appscode/kubectl-nonroot:1.34
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode/petset:v0.1.0 $IMAGE_REGISTRY/appscode/petset:v0.1.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode/sidekick:v0.0.15 $IMAGE_REGISTRY/appscode/sidekick:v0.0.15
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/cassandra-medusa-plugin:v0.14.0-rc.2 $IMAGE_REGISTRY/kubedb/cassandra-medusa-plugin:v0.14.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/clickhouse-backup-plugin:v0.3.0 $IMAGE_REGISTRY/kubedb/clickhouse-backup-plugin:v0.3.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/clickhouse-backup-plugin:v0.4.0-rc.2 $IMAGE_REGISTRY/kubedb/clickhouse-backup-plugin:v0.4.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/dashboard-restic-plugin:v0.25.0-rc.2 $IMAGE_REGISTRY/kubedb/dashboard-restic-plugin:v0.25.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch-restic-plugin:v0.30.0-rc.2 $IMAGE_REGISTRY/kubedb/elasticsearch-restic-plugin:v0.30.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/etcd-restic-plugin:v0.1.0-rc.2_3.5.21 $IMAGE_REGISTRY/kubedb/etcd-restic-plugin:v0.1.0-rc.2_3.5.21
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/etcd-restic-plugin:v0.1.0-rc.2_3.6.4 $IMAGE_REGISTRY/kubedb/etcd-restic-plugin:v0.1.0-rc.2_3.6.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-autoscaler:v0.52.0-rc.2 $IMAGE_REGISTRY/kubedb/kubedb-autoscaler:v0.52.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-courier:v0.7.0-rc.2 $IMAGE_REGISTRY/kubedb/kubedb-courier:v0.7.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-crd-manager:v0.22.0-rc.2 $IMAGE_REGISTRY/kubedb/kubedb-crd-manager:v0.22.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-gitops:v0.15.0-rc.2 $IMAGE_REGISTRY/kubedb/kubedb-gitops:v0.15.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-kibana:v0.43.0-rc.2 $IMAGE_REGISTRY/kubedb/kubedb-kibana:v0.43.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-manifest-plugin:v0.30.0-rc.2 $IMAGE_REGISTRY/kubedb/kubedb-manifest-plugin:v0.30.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-ops-manager:v0.54.0-rc.2 $IMAGE_REGISTRY/kubedb/kubedb-ops-manager:v0.54.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-provisioner:v0.67.0-rc.2 $IMAGE_REGISTRY/kubedb/kubedb-provisioner:v0.67.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-schema-manager:v0.43.0-rc.2 $IMAGE_REGISTRY/kubedb/kubedb-schema-manager:v0.43.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-ui-server:v0.43.0-rc.2 $IMAGE_REGISTRY/kubedb/kubedb-ui-server:v0.43.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-verifier:v0.18.0-rc.2 $IMAGE_REGISTRY/kubedb/kubedb-verifier:v0.18.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-webhook-server:v0.43.0-rc.2 $IMAGE_REGISTRY/kubedb/kubedb-webhook-server:v0.43.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-csi-snapshotter-plugin:v0.27.0-rc.2 $IMAGE_REGISTRY/kubedb/mariadb-csi-snapshotter-plugin:v0.27.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-restic-plugin:v0.25.0-rc.2_10.11.6-jammy $IMAGE_REGISTRY/kubedb/mariadb-restic-plugin:v0.25.0-rc.2_10.11.6-jammy
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-restic-plugin:v0.25.0-rc.2_10.4.32-focal $IMAGE_REGISTRY/kubedb/mariadb-restic-plugin:v0.25.0-rc.2_10.4.32-focal
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-restic-plugin:v0.25.0-rc.2_10.6.16-focal $IMAGE_REGISTRY/kubedb/mariadb-restic-plugin:v0.25.0-rc.2_10.6.16-focal
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-restic-plugin:v0.25.0-rc.2_11.1.3-jammy $IMAGE_REGISTRY/kubedb/mariadb-restic-plugin:v0.25.0-rc.2_11.1.3-jammy
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-csi-snapshotter-plugin:v0.28.0-rc.2 $IMAGE_REGISTRY/kubedb/mongodb-csi-snapshotter-plugin:v0.28.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-restic-plugin:v0.30.0-rc.2_4.2.3 $IMAGE_REGISTRY/kubedb/mongodb-restic-plugin:v0.30.0-rc.2_4.2.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-restic-plugin:v0.30.0-rc.2_4.4.6 $IMAGE_REGISTRY/kubedb/mongodb-restic-plugin:v0.30.0-rc.2_4.4.6
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-restic-plugin:v0.30.0-rc.2_5.0.15 $IMAGE_REGISTRY/kubedb/mongodb-restic-plugin:v0.30.0-rc.2_5.0.15
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-restic-plugin:v0.30.0-rc.2_5.0.3 $IMAGE_REGISTRY/kubedb/mongodb-restic-plugin:v0.30.0-rc.2_5.0.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-restic-plugin:v0.30.0-rc.2_6.0.5 $IMAGE_REGISTRY/kubedb/mongodb-restic-plugin:v0.30.0-rc.2_6.0.5
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-restic-plugin:v0.30.0-rc.2_8.0.3 $IMAGE_REGISTRY/kubedb/mongodb-restic-plugin:v0.30.0-rc.2_8.0.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mssqlserver-walg-plugin:v0.21.0-rc.2 $IMAGE_REGISTRY/kubedb/mssqlserver-walg-plugin:v0.21.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-csi-snapshotter-plugin:v0.28.0-rc.2 $IMAGE_REGISTRY/kubedb/mysql-csi-snapshotter-plugin:v0.28.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-restic-plugin:v0.30.0-rc.2_5.7.25 $IMAGE_REGISTRY/kubedb/mysql-restic-plugin:v0.30.0-rc.2_5.7.25
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-restic-plugin:v0.30.0-rc.2_8.0.21 $IMAGE_REGISTRY/kubedb/mysql-restic-plugin:v0.30.0-rc.2_8.0.21
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-restic-plugin:v0.30.0-rc.2_8.0.3 $IMAGE_REGISTRY/kubedb/mysql-restic-plugin:v0.30.0-rc.2_8.0.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-restic-plugin:v0.30.0-rc.2_8.4.2 $IMAGE_REGISTRY/kubedb/mysql-restic-plugin:v0.30.0-rc.2_8.4.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-restic-plugin:v0.30.0-rc.2_9.0.1 $IMAGE_REGISTRY/kubedb/mysql-restic-plugin:v0.30.0-rc.2_9.0.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/neo4j-backup-plugin:v0.3.0-rc.2 $IMAGE_REGISTRY/kubedb/neo4j-backup-plugin:v0.3.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-csi-snapshotter-plugin:v0.28.0-rc.2 $IMAGE_REGISTRY/kubedb/postgres-csi-snapshotter-plugin:v0.28.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-restic-plugin:v0.30.0-rc.2_12.17 $IMAGE_REGISTRY/kubedb/postgres-restic-plugin:v0.30.0-rc.2_12.17
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-restic-plugin:v0.30.0-rc.2_14.10 $IMAGE_REGISTRY/kubedb/postgres-restic-plugin:v0.30.0-rc.2_14.10
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-restic-plugin:v0.30.0-rc.2_16.4 $IMAGE_REGISTRY/kubedb/postgres-restic-plugin:v0.30.0-rc.2_16.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-restic-plugin:v0.30.0-rc.2_17.2 $IMAGE_REGISTRY/kubedb/postgres-restic-plugin:v0.30.0-rc.2_17.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-restic-plugin:v0.30.0-rc.2_18.2 $IMAGE_REGISTRY/kubedb/postgres-restic-plugin:v0.30.0-rc.2_18.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/provider-aws:v0.28.0-rc.2 $IMAGE_REGISTRY/kubedb/provider-aws:v0.28.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/provider-azure:v0.28.0-rc.2 $IMAGE_REGISTRY/kubedb/provider-azure:v0.28.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/provider-gcp:v0.28.0-rc.2 $IMAGE_REGISTRY/kubedb/provider-gcp:v0.28.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/qdrant-restic-plugin:v0.3.0-rc.2 $IMAGE_REGISTRY/kubedb/qdrant-restic-plugin:v0.3.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/redis-restic-plugin:v0.30.0-rc.2 $IMAGE_REGISTRY/kubedb/redis-restic-plugin:v0.30.0-rc.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/singlestore-restic-plugin:v0.25.0-rc.2_alma-8.1.32-e3d3cde6da $IMAGE_REGISTRY/kubedb/singlestore-restic-plugin:v0.25.0-rc.2_alma-8.1.32-e3d3cde6da
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/singlestore-restic-plugin:v0.25.0-rc.2_alma-8.5.7-bf633c1a54 $IMAGE_REGISTRY/kubedb/singlestore-restic-plugin:v0.25.0-rc.2_alma-8.5.7-bf633c1a54
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/xtrabackup-restic-plugin:v0.15.0-rc.2_2.4.29 $IMAGE_REGISTRY/kubedb/xtrabackup-restic-plugin:v0.15.0-rc.2_2.4.29
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/xtrabackup-restic-plugin:v0.15.0-rc.2_8.0.35 $IMAGE_REGISTRY/kubedb/xtrabackup-restic-plugin:v0.15.0-rc.2_8.0.35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/xtrabackup-restic-plugin:v0.15.0-rc.2_8.1.0 $IMAGE_REGISTRY/kubedb/xtrabackup-restic-plugin:v0.15.0-rc.2_8.1.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/xtrabackup-restic-plugin:v0.15.0-rc.2_8.2.0 $IMAGE_REGISTRY/kubedb/xtrabackup-restic-plugin:v0.15.0-rc.2_8.2.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/xtrabackup-restic-plugin:v0.15.0-rc.2_8.4.0 $IMAGE_REGISTRY/kubedb/xtrabackup-restic-plugin:v0.15.0-rc.2_8.4.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/xtrabackup-restic-plugin:v0.15.0-rc.2_9.7.1 $IMAGE_REGISTRY/kubedb/xtrabackup-restic-plugin:v0.15.0-rc.2_9.7.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/zookeeper-restic-plugin:v0.22.0-rc.2 $IMAGE_REGISTRY/kubedb/zookeeper-restic-plugin:v0.22.0-rc.2
