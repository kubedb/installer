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

$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/redis:5.0.14-bullseye $IMAGE_REGISTRY/appscode-images/redis:5.0.14-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/redis:6.0.20-bookworm $IMAGE_REGISTRY/appscode-images/redis:6.0.20-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/redis:6.2.14-bookworm $IMAGE_REGISTRY/appscode-images/redis:6.2.14-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/redis:6.2.16-bookworm $IMAGE_REGISTRY/appscode-images/redis:6.2.16-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/redis:7.0.14-bookworm $IMAGE_REGISTRY/appscode-images/redis:7.0.14-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/redis:7.0.15-bookworm $IMAGE_REGISTRY/appscode-images/redis:7.0.15-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/redis:7.2.3-bookworm $IMAGE_REGISTRY/appscode-images/redis:7.2.3-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/redis:7.2.4-bookworm $IMAGE_REGISTRY/appscode-images/redis:7.2.4-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/redis:7.2.6-bookworm $IMAGE_REGISTRY/appscode-images/redis:7.2.6-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/redis:7.4.0-bookworm $IMAGE_REGISTRY/appscode-images/redis:7.4.0-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/redis:7.4.1-bookworm $IMAGE_REGISTRY/appscode-images/redis:7.4.1-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/valkey:7.2.5 $IMAGE_REGISTRY/appscode-images/valkey:7.2.5
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/valkey:7.2.9 $IMAGE_REGISTRY/appscode-images/valkey:7.2.9
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/valkey:8.0.3 $IMAGE_REGISTRY/appscode-images/valkey:8.0.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/valkey:8.1.1 $IMAGE_REGISTRY/appscode-images/valkey:8.1.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/redis-coordinator:v0.36.0 $IMAGE_REGISTRY/kubedb/redis-coordinator:v0.36.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/redis-init:0.12.0 $IMAGE_REGISTRY/kubedb/redis-init:0.12.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/redis:4.0.11 $IMAGE_REGISTRY/kubedb/redis:4.0.11
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/redis_exporter:1.66.0 $IMAGE_REGISTRY/kubedb/redis_exporter:1.66.0
$CMD cp --allow-nondistributable-artifacts --insecure registry.k8s.io/git-sync/git-sync:v4.4.2 $IMAGE_REGISTRY/git-sync/git-sync:v4.4.2
