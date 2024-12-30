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

$CMD push --allow-nondistributable-artifacts --insecure images/appscode-kube-rbac-proxy-v0.15.0.tar $IMAGE_REGISTRY/appscode/kube-rbac-proxy:v0.15.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-kubectl-nonroot-1.31.tar $IMAGE_REGISTRY/appscode/kubectl-nonroot:1.31
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-petset-v0.0.7.tar $IMAGE_REGISTRY/appscode/petset:v0.0.7
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-sidekick-v0.0.9.tar $IMAGE_REGISTRY/appscode/sidekick:v0.0.9
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-dashboard-restic-plugin-v0.1.0.tar $IMAGE_REGISTRY/kubedb/dashboard-restic-plugin:v0.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-dashboard-restic-plugin-v0.8.0.tar $IMAGE_REGISTRY/kubedb/dashboard-restic-plugin:v0.8.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-restic-plugin-v0.13.0.tar $IMAGE_REGISTRY/kubedb/elasticsearch-restic-plugin:v0.13.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-autoscaler-v0.35.0.tar $IMAGE_REGISTRY/kubedb/kubedb-autoscaler:v0.35.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-crd-manager-v0.5.0.tar $IMAGE_REGISTRY/kubedb/kubedb-crd-manager:v0.5.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-kibana-v0.26.0.tar $IMAGE_REGISTRY/kubedb/kubedb-kibana:v0.26.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-manifest-plugin-v0.13.0.tar $IMAGE_REGISTRY/kubedb/kubedb-manifest-plugin:v0.13.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-ops-manager-v0.37.2.tar $IMAGE_REGISTRY/kubedb/kubedb-ops-manager:v0.37.2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-provisioner-v0.50.1.tar $IMAGE_REGISTRY/kubedb/kubedb-provisioner:v0.50.1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-schema-manager-v0.26.0.tar $IMAGE_REGISTRY/kubedb/kubedb-schema-manager:v0.26.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-ui-server-v0.26.0.tar $IMAGE_REGISTRY/kubedb/kubedb-ui-server:v0.26.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-verifier-v0.1.0.tar $IMAGE_REGISTRY/kubedb/kubedb-verifier:v0.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-webhook-server-v0.26.1.tar $IMAGE_REGISTRY/kubedb/kubedb-webhook-server:v0.26.1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mariadb-csi-snapshotter-plugin-v0.10.0.tar $IMAGE_REGISTRY/kubedb/mariadb-csi-snapshotter-plugin:v0.10.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mariadb-restic-plugin-v0.8.0.tar $IMAGE_REGISTRY/kubedb/mariadb-restic-plugin:v0.8.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-csi-snapshotter-plugin-v0.11.0.tar $IMAGE_REGISTRY/kubedb/mongodb-csi-snapshotter-plugin:v0.11.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mssqlserver-walg-plugin-v0.0.1.tar $IMAGE_REGISTRY/kubedb/mssqlserver-walg-plugin:v0.0.1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-csi-snapshotter-plugin-v0.11.0.tar $IMAGE_REGISTRY/kubedb/mysql-csi-snapshotter-plugin:v0.11.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-csi-snapshotter-plugin-v0.11.0.tar $IMAGE_REGISTRY/kubedb/postgres-csi-snapshotter-plugin:v0.11.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-provider-aws-v0.12.0.tar $IMAGE_REGISTRY/kubedb/provider-aws:v0.12.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-provider-azure-v0.12.0.tar $IMAGE_REGISTRY/kubedb/provider-azure:v0.12.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-provider-gcp-v0.12.0.tar $IMAGE_REGISTRY/kubedb/provider-gcp:v0.12.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-redis-restic-plugin-v0.13.0.tar $IMAGE_REGISTRY/kubedb/redis-restic-plugin:v0.13.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-zookeeper-restic-plugin-v0.6.0.tar $IMAGE_REGISTRY/kubedb/zookeeper-restic-plugin:v0.6.0
$CMD push --allow-nondistributable-artifacts --insecure images/tianon-toybox-0.8.11.tar $IMAGE_REGISTRY/tianon/toybox:0.8.11
