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

$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mysql-5.7.42-debian.tar $IMAGE_REGISTRY/appscode-images/mysql:5.7.42-debian
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mysql-5.7.44-oracle.tar $IMAGE_REGISTRY/appscode-images/mysql:5.7.44-oracle
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mysql-8.0.31-oracle.tar $IMAGE_REGISTRY/appscode-images/mysql:8.0.31-oracle
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mysql-8.0.35-oracle.tar $IMAGE_REGISTRY/appscode-images/mysql:8.0.35-oracle
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mysql-8.0.36-debian.tar $IMAGE_REGISTRY/appscode-images/mysql:8.0.36-debian
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mysql-8.1.0-oracle.tar $IMAGE_REGISTRY/appscode-images/mysql:8.1.0-oracle
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mysql-8.2.0-oracle.tar $IMAGE_REGISTRY/appscode-images/mysql:8.2.0-oracle
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mysql-8.4.2-oracle.tar $IMAGE_REGISTRY/appscode-images/mysql:8.4.2-oracle
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mysql-8.4.3-oracle.tar $IMAGE_REGISTRY/appscode-images/mysql:8.4.3-oracle
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mysql-9.0.1-oracle.tar $IMAGE_REGISTRY/appscode-images/mysql:9.0.1-oracle
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mysql-9.1.0-oracle.tar $IMAGE_REGISTRY/appscode-images/mysql:9.1.0-oracle
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-archiver-v0.18.0-rc.0_5.7.44.tar $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.18.0-rc.0_5.7.44
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-archiver-v0.18.0-rc.0_8.0.35.tar $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.18.0-rc.0_8.0.35
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-archiver-v0.18.0-rc.0_8.1.0.tar $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.18.0-rc.0_8.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-archiver-v0.18.0-rc.0_8.2.0.tar $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.18.0-rc.0_8.2.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-archiver-v0.18.0-rc.0_8.4.3.tar $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.18.0-rc.0_8.4.3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-archiver-v0.18.0-rc.0_9.1.0.tar $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.18.0-rc.0_9.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-coordinator-v0.35.0-rc.0.tar $IMAGE_REGISTRY/kubedb/mysql-coordinator:v0.35.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-init-5.7-v7.tar $IMAGE_REGISTRY/kubedb/mysql-init:5.7-v7
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-init-8.0.31-v6.tar $IMAGE_REGISTRY/kubedb/mysql-init:8.0.31-v6
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-init-8.4.2-v5.tar $IMAGE_REGISTRY/kubedb/mysql-init:8.4.2-v5
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-init-8.4.3-v5.tar $IMAGE_REGISTRY/kubedb/mysql-init:8.4.3-v5
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-init-9.0.1-v3.tar $IMAGE_REGISTRY/kubedb/mysql-init:9.0.1-v3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-init-9.1.0-v3.tar $IMAGE_REGISTRY/kubedb/mysql-init:9.1.0-v3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysql-router-init-v0.35.0-rc.0.tar $IMAGE_REGISTRY/kubedb/mysql-router-init:v0.35.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mysqld-exporter-v0.13.1.tar $IMAGE_REGISTRY/kubedb/mysqld-exporter:v0.13.1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-replication-mode-detector-v0.44.0-rc.0.tar $IMAGE_REGISTRY/kubedb/replication-mode-detector:v0.44.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/mysql-mysql-router-8.0.31.tar $IMAGE_REGISTRY/mysql/mysql-router:8.0.31
$CMD push --allow-nondistributable-artifacts --insecure images/git-sync-git-sync-v4.2.1.tar $IMAGE_REGISTRY/git-sync/git-sync:v4.2.1
