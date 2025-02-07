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

$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-redis-5.0.14-bullseye.tar $IMAGE_REGISTRY/appscode-images/redis:5.0.14-bullseye
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-redis-6.0.20-bookworm.tar $IMAGE_REGISTRY/appscode-images/redis:6.0.20-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-redis-6.2.14-bookworm.tar $IMAGE_REGISTRY/appscode-images/redis:6.2.14-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-redis-6.2.16-bookworm.tar $IMAGE_REGISTRY/appscode-images/redis:6.2.16-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-redis-7.0.14-bookworm.tar $IMAGE_REGISTRY/appscode-images/redis:7.0.14-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-redis-7.0.15-bookworm.tar $IMAGE_REGISTRY/appscode-images/redis:7.0.15-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-redis-7.2.3-bookworm.tar $IMAGE_REGISTRY/appscode-images/redis:7.2.3-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-redis-7.2.4-bookworm.tar $IMAGE_REGISTRY/appscode-images/redis:7.2.4-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-redis-7.2.6-bookworm.tar $IMAGE_REGISTRY/appscode-images/redis:7.2.6-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-redis-7.4.0-bookworm.tar $IMAGE_REGISTRY/appscode-images/redis:7.4.0-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-redis-7.4.1-bookworm.tar $IMAGE_REGISTRY/appscode-images/redis:7.4.1-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-redis-coordinator-v0.31.0-rc.0.tar $IMAGE_REGISTRY/kubedb/redis-coordinator:v0.31.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-redis-init-0.9.0.tar $IMAGE_REGISTRY/kubedb/redis-init:0.9.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-redis-4.0.11.tar $IMAGE_REGISTRY/kubedb/redis:4.0.11
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-redis_exporter-1.66.0.tar $IMAGE_REGISTRY/kubedb/redis_exporter:1.66.0
