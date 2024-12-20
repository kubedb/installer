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

$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:5.7.42-debian images/appscode-images-mysql-5.7.42-debian.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:5.7.44-oracle images/appscode-images-mysql-5.7.44-oracle.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:8.0.31-oracle images/appscode-images-mysql-8.0.31-oracle.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:8.0.35-oracle images/appscode-images-mysql-8.0.35-oracle.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:8.0.36-debian images/appscode-images-mysql-8.0.36-debian.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:8.1.0-oracle images/appscode-images-mysql-8.1.0-oracle.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:8.2.0-oracle images/appscode-images-mysql-8.2.0-oracle.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:8.4.2-oracle images/appscode-images-mysql-8.4.2-oracle.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:8.4.3-oracle images/appscode-images-mysql-8.4.3-oracle.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:9.0.1-oracle images/appscode-images-mysql-9.0.1-oracle.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mysql:9.1.0-oracle images/appscode-images-mysql-9.1.0-oracle.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-archiver:v0.11.0_5.7.44 images/kubedb-mysql-archiver-v0.11.0_5.7.44.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-archiver:v0.11.0_8.0.35 images/kubedb-mysql-archiver-v0.11.0_8.0.35.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-archiver:v0.11.0_8.1.0 images/kubedb-mysql-archiver-v0.11.0_8.1.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-archiver:v0.11.0_8.2.0 images/kubedb-mysql-archiver-v0.11.0_8.2.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-archiver:v0.11.0_8.4.3 images/kubedb-mysql-archiver-v0.11.0_8.4.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-archiver:v0.11.0_9.1.0 images/kubedb-mysql-archiver-v0.11.0_9.1.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-coordinator:v0.28.0 images/kubedb-mysql-coordinator-v0.28.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-init:5.7-v5 images/kubedb-mysql-init-5.7-v5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-init:8.0.31-v4 images/kubedb-mysql-init-8.0.31-v4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-init:8.4.2-v3 images/kubedb-mysql-init-8.4.2-v3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-init:8.4.3-v3 images/kubedb-mysql-init-8.4.3-v3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-init:9.0.1-v1 images/kubedb-mysql-init-9.0.1-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-init:9.1.0-v1 images/kubedb-mysql-init-9.1.0-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysql-router-init:v0.28.0 images/kubedb-mysql-router-init-v0.28.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mysqld-exporter:v0.13.1 images/kubedb-mysqld-exporter-v0.13.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/replication-mode-detector:v0.37.0 images/kubedb-replication-mode-detector-v0.37.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure mysql/mysql-router:8.0.31 images/mysql-mysql-router-8.0.31.tar
$CMD pull --allow-nondistributable-artifacts --insecure registry.k8s.io/git-sync/git-sync:v4.2.1 images/git-sync-git-sync-v4.2.1.tar

tar -czvf images.tar.gz images
