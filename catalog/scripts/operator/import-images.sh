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

$CMD push --allow-nondistributable-artifacts --insecure images/appscode-kubectl-nonroot-1.34.tar $IMAGE_REGISTRY/appscode/kubectl-nonroot:1.34
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-petset-v0.0.16.tar $IMAGE_REGISTRY/appscode/petset:v0.0.16
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-sidekick-v0.0.13.tar $IMAGE_REGISTRY/appscode/sidekick:v0.0.13
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-cassandra-medusa-plugin-v0.8.0-rc.0.tar $IMAGE_REGISTRY/kubedb/cassandra-medusa-plugin:v0.8.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-dashboard-restic-plugin-v0.19.0-rc.0.tar $IMAGE_REGISTRY/kubedb/dashboard-restic-plugin:v0.19.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-restic-plugin-v0.24.0-rc.0.tar $IMAGE_REGISTRY/kubedb/elasticsearch-restic-plugin:v0.24.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-autoscaler-v0.46.0-rc.0.tar $IMAGE_REGISTRY/kubedb/kubedb-autoscaler:v0.46.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-crd-manager-v0.16.0-rc.0.tar $IMAGE_REGISTRY/kubedb/kubedb-crd-manager:v0.16.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-gitops-v0.9.0-rc.0.tar $IMAGE_REGISTRY/kubedb/kubedb-gitops:v0.9.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-kibana-v0.37.0-rc.0.tar $IMAGE_REGISTRY/kubedb/kubedb-kibana:v0.37.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-manifest-plugin-v0.24.0-rc.0.tar $IMAGE_REGISTRY/kubedb/kubedb-manifest-plugin:v0.24.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-ops-manager-v0.48.0-rc.0.tar $IMAGE_REGISTRY/kubedb/kubedb-ops-manager:v0.48.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-provisioner-v0.61.0-rc.0.tar $IMAGE_REGISTRY/kubedb/kubedb-provisioner:v0.61.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-schema-manager-v0.37.0-rc.0.tar $IMAGE_REGISTRY/kubedb/kubedb-schema-manager:v0.37.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-ui-server-v0.37.0-rc.0.tar $IMAGE_REGISTRY/kubedb/kubedb-ui-server:v0.37.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-verifier-v0.12.0-rc.0.tar $IMAGE_REGISTRY/kubedb/kubedb-verifier:v0.12.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-webhook-server-v0.37.0-rc.0.tar $IMAGE_REGISTRY/kubedb/kubedb-webhook-server:v0.37.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mariadb-csi-snapshotter-plugin-v0.21.0-rc.0.tar $IMAGE_REGISTRY/kubedb/mariadb-csi-snapshotter-plugin:v0.21.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-migrator-operator-v2026.2.16-rc.0.tar $IMAGE_REGISTRY/kubedb/migrator-operator:v2026.2.16-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-csi-snapshotter-plugin-v0.22.0-rc.0.tar $IMAGE_REGISTRY/kubedb/mongodb-csi-snapshotter-plugin:v0.22.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mssqlserver-walg-plugin-v0.15.0-rc.0.tar $IMAGE_REGISTRY/kubedb/mssqlserver-walg-plugin:v0.15.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-csi-snapshotter-plugin-v0.22.0-rc.0.tar $IMAGE_REGISTRY/kubedb/mysql-csi-snapshotter-plugin:v0.22.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-csi-snapshotter-plugin-v0.22.0-rc.0.tar $IMAGE_REGISTRY/kubedb/postgres-csi-snapshotter-plugin:v0.22.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-restic-plugin-v0.24.0-rc.0_16.1.tar $IMAGE_REGISTRY/kubedb/postgres-restic-plugin:v0.24.0-rc.0_16.1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-provider-aws-v0.22.0-rc.0.tar $IMAGE_REGISTRY/kubedb/provider-aws:v0.22.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-provider-azure-v0.22.0-rc.0.tar $IMAGE_REGISTRY/kubedb/provider-azure:v0.22.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-provider-gcp-v0.22.0-rc.0.tar $IMAGE_REGISTRY/kubedb/provider-gcp:v0.22.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-redis-restic-plugin-v0.24.0-rc.0.tar $IMAGE_REGISTRY/kubedb/redis-restic-plugin:v0.24.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-zookeeper-restic-plugin-v0.16.0-rc.0.tar $IMAGE_REGISTRY/kubedb/zookeeper-restic-plugin:v0.16.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/tianon-toybox-0.8.11.tar $IMAGE_REGISTRY/tianon/toybox:0.8.11
