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
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode/sidekick:v0.0.14 images/appscode-sidekick-v0.0.14.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/cassandra-medusa-plugin:v0.12.0 images/kubedb-cassandra-medusa-plugin-v0.12.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/clickhouse-backup-plugin:v0.2.0 images/kubedb-clickhouse-backup-plugin-v0.2.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/dashboard-restic-plugin:v0.23.0 images/kubedb-dashboard-restic-plugin-v0.23.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch-restic-plugin:v0.28.0 images/kubedb-elasticsearch-restic-plugin-v0.28.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-autoscaler:v0.50.0 images/kubedb-kubedb-autoscaler-v0.50.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-crd-manager:v0.20.0 images/kubedb-kubedb-crd-manager-v0.20.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-gitops:v0.13.0 images/kubedb-kubedb-gitops-v0.13.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-kibana:v0.41.0 images/kubedb-kubedb-kibana-v0.41.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-manifest-plugin:v0.28.0 images/kubedb-kubedb-manifest-plugin-v0.28.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-ops-manager:v0.52.0 images/kubedb-kubedb-ops-manager-v0.52.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-provisioner:v0.65.0 images/kubedb-kubedb-provisioner-v0.65.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-schema-manager:v0.41.0 images/kubedb-kubedb-schema-manager-v0.41.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-ui-server:v0.41.0 images/kubedb-kubedb-ui-server-v0.41.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-verifier:v0.16.0 images/kubedb-kubedb-verifier-v0.16.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-webhook-server:v0.41.0 images/kubedb-kubedb-webhook-server-v0.41.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-csi-snapshotter-plugin:v0.25.0 images/kubedb-mariadb-csi-snapshotter-plugin-v0.25.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/migrator-operator:v0.5.0 images/kubedb-migrator-operator-v0.5.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-csi-snapshotter-plugin:v0.26.0 images/kubedb-mongodb-csi-snapshotter-plugin-v0.26.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mssqlserver-walg-plugin:v0.19.0 images/kubedb-mssqlserver-walg-plugin-v0.19.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-csi-snapshotter-plugin:v0.26.0 images/kubedb-mysql-csi-snapshotter-plugin-v0.26.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/neo4j-backup-plugin:v0.1.0 images/kubedb-neo4j-backup-plugin-v0.1.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-csi-snapshotter-plugin:v0.26.0 images/kubedb-postgres-csi-snapshotter-plugin-v0.26.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-restic-plugin:v0.28.0_16.4 images/kubedb-postgres-restic-plugin-v0.28.0_16.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/provider-aws:v0.26.0 images/kubedb-provider-aws-v0.26.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/provider-azure:v0.26.0 images/kubedb-provider-azure-v0.26.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/provider-gcp:v0.26.0 images/kubedb-provider-gcp-v0.26.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/redis-restic-plugin:v0.28.0 images/kubedb-redis-restic-plugin-v0.28.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/zookeeper-restic-plugin:v0.20.0 images/kubedb-zookeeper-restic-plugin-v0.20.0.tar

tar -czvf images.tar.gz images
