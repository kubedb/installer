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

$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:5.7.42-debian $IMAGE_REGISTRY/appscode-images/mysql:5.7.42-debian
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:5.7.44-oracle $IMAGE_REGISTRY/appscode-images/mysql:5.7.44-oracle
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:8.0.31-oracle $IMAGE_REGISTRY/appscode-images/mysql:8.0.31-oracle
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:8.0.35-oracle $IMAGE_REGISTRY/appscode-images/mysql:8.0.35-oracle
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:8.0.36-debian $IMAGE_REGISTRY/appscode-images/mysql:8.0.36-debian
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:8.1.0-oracle $IMAGE_REGISTRY/appscode-images/mysql:8.1.0-oracle
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:8.2.0-oracle $IMAGE_REGISTRY/appscode-images/mysql:8.2.0-oracle
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:8.4.2-oracle $IMAGE_REGISTRY/appscode-images/mysql:8.4.2-oracle
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-archiver:v0.10.0_5.7.44 $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.10.0_5.7.44
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-archiver:v0.10.0_8.0.35 $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.10.0_8.0.35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-archiver:v0.10.0_8.1.0 $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.10.0_8.1.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-archiver:v0.10.0_8.2.0 $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.10.0_8.2.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-coordinator:v0.27.0 $IMAGE_REGISTRY/kubedb/mysql-coordinator:v0.27.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-init:5.7-v4 $IMAGE_REGISTRY/kubedb/mysql-init:5.7-v4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-init:8.0.31-v3 $IMAGE_REGISTRY/kubedb/mysql-init:8.0.31-v3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-init:8.4.2-v2 $IMAGE_REGISTRY/kubedb/mysql-init:8.4.2-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-router-init:v0.27.0 $IMAGE_REGISTRY/kubedb/mysql-router-init:v0.27.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysqld-exporter:v0.13.1 $IMAGE_REGISTRY/kubedb/mysqld-exporter:v0.13.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/replication-mode-detector:v0.36.0 $IMAGE_REGISTRY/kubedb/replication-mode-detector:v0.36.0
$CMD cp --allow-nondistributable-artifacts --insecure mysql/mysql-router:8.0.31 $IMAGE_REGISTRY/mysql/mysql-router:8.0.31
$CMD cp --allow-nondistributable-artifacts --insecure registry.k8s.io/git-sync/git-sync:v4.2.1 $IMAGE_REGISTRY/git-sync/git-sync:v4.2.1