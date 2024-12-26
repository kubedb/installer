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

$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode/kube-rbac-proxy:v0.15.0 $IMAGE_REGISTRY/appscode/kube-rbac-proxy:v0.15.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode/kubectl-nonroot:1.31 $IMAGE_REGISTRY/appscode/kubectl-nonroot:1.31
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode/petset:v0.0.7 $IMAGE_REGISTRY/appscode/petset:v0.0.7
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode/sidekick:v0.0.9 $IMAGE_REGISTRY/appscode/sidekick:v0.0.9
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/dashboard-restic-plugin:v0.1.0 $IMAGE_REGISTRY/kubedb/dashboard-restic-plugin:v0.1.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/dashboard-restic-plugin:v0.8.0 $IMAGE_REGISTRY/kubedb/dashboard-restic-plugin:v0.8.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch-restic-plugin:v0.13.0 $IMAGE_REGISTRY/kubedb/elasticsearch-restic-plugin:v0.13.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-autoscaler:v0.35.0 $IMAGE_REGISTRY/kubedb/kubedb-autoscaler:v0.35.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-crd-manager:v0.5.0 $IMAGE_REGISTRY/kubedb/kubedb-crd-manager:v0.5.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-kibana:v0.26.0 $IMAGE_REGISTRY/kubedb/kubedb-kibana:v0.26.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-manifest-plugin:v0.13.0 $IMAGE_REGISTRY/kubedb/kubedb-manifest-plugin:v0.13.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-ops-manager:v0.37.1 $IMAGE_REGISTRY/kubedb/kubedb-ops-manager:v0.37.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-provisioner:v0.50.1 $IMAGE_REGISTRY/kubedb/kubedb-provisioner:v0.50.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-schema-manager:v0.26.0 $IMAGE_REGISTRY/kubedb/kubedb-schema-manager:v0.26.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-ui-server:v0.26.0 $IMAGE_REGISTRY/kubedb/kubedb-ui-server:v0.26.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-verifier:v0.1.0 $IMAGE_REGISTRY/kubedb/kubedb-verifier:v0.1.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-webhook-server:v0.26.1 $IMAGE_REGISTRY/kubedb/kubedb-webhook-server:v0.26.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-csi-snapshotter-plugin:v0.10.0 $IMAGE_REGISTRY/kubedb/mariadb-csi-snapshotter-plugin:v0.10.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-restic-plugin:v0.8.0 $IMAGE_REGISTRY/kubedb/mariadb-restic-plugin:v0.8.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-csi-snapshotter-plugin:v0.11.0 $IMAGE_REGISTRY/kubedb/mongodb-csi-snapshotter-plugin:v0.11.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mssqlserver-walg-plugin:v0.0.1 $IMAGE_REGISTRY/kubedb/mssqlserver-walg-plugin:v0.0.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-csi-snapshotter-plugin:v0.11.0 $IMAGE_REGISTRY/kubedb/mysql-csi-snapshotter-plugin:v0.11.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-csi-snapshotter-plugin:v0.11.0 $IMAGE_REGISTRY/kubedb/postgres-csi-snapshotter-plugin:v0.11.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-restic-plugin:v0.13.0_16.1 $IMAGE_REGISTRY/kubedb/postgres-restic-plugin:v0.13.0_16.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/provider-aws:v0.12.0 $IMAGE_REGISTRY/kubedb/provider-aws:v0.12.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/provider-azure:v0.12.0 $IMAGE_REGISTRY/kubedb/provider-azure:v0.12.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/provider-gcp:v0.12.0 $IMAGE_REGISTRY/kubedb/provider-gcp:v0.12.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/redis-restic-plugin:v0.13.0 $IMAGE_REGISTRY/kubedb/redis-restic-plugin:v0.13.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/zookeeper-restic-plugin:v0.6.0 $IMAGE_REGISTRY/kubedb/zookeeper-restic-plugin:v0.6.0
$CMD cp --allow-nondistributable-artifacts --insecure tianon/toybox:0.8.11 $IMAGE_REGISTRY/tianon/toybox:0.8.11
