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

$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:10.10.7-jammy $IMAGE_REGISTRY/appscode-images/mariadb:10.10.7-jammy
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:10.11.6-jammy $IMAGE_REGISTRY/appscode-images/mariadb:10.11.6-jammy
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:10.4.32-focal $IMAGE_REGISTRY/appscode-images/mariadb:10.4.32-focal
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:10.5.23-focal $IMAGE_REGISTRY/appscode-images/mariadb:10.5.23-focal
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:10.6.16-focal $IMAGE_REGISTRY/appscode-images/mariadb:10.6.16-focal
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:11.0.4-jammy $IMAGE_REGISTRY/appscode-images/mariadb:11.0.4-jammy
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:11.1.3-jammy $IMAGE_REGISTRY/appscode-images/mariadb:11.1.3-jammy
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:11.2.2-jammy $IMAGE_REGISTRY/appscode-images/mariadb:11.2.2-jammy
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:11.3.2-jammy $IMAGE_REGISTRY/appscode-images/mariadb:11.3.2-jammy
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:11.4.3-noble $IMAGE_REGISTRY/appscode-images/mariadb:11.4.3-noble
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:11.5.2-noble $IMAGE_REGISTRY/appscode-images/mariadb:11.5.2-noble
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.9.0_10.10.7-jammy $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.9.0_10.10.7-jammy
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.9.0_10.11.6-jammy $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.9.0_10.11.6-jammy
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.9.0_10.4.32-focal $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.9.0_10.4.32-focal
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.9.0_10.5.23-focal $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.9.0_10.5.23-focal
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.9.0_10.6.16-focal $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.9.0_10.6.16-focal
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.9.0_11.0.4-jammy $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.9.0_11.0.4-jammy
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.9.0_11.1.3-jammy $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.9.0_11.1.3-jammy
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.9.0_11.2.2-jammy $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.9.0_11.2.2-jammy
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-coordinator:v0.29.0 $IMAGE_REGISTRY/kubedb/mariadb-coordinator:v0.29.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-init:0.5.2 $IMAGE_REGISTRY/kubedb/mariadb-init:0.5.2
