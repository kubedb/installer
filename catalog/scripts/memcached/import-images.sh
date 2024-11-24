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

$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-memcached-1.5.22-alpine.tar $IMAGE_REGISTRY/appscode-images/memcached:1.5.22-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-memcached-1.6.22-alpine.tar $IMAGE_REGISTRY/appscode-images/memcached:1.6.22-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-memcached-1.6.29-alpine.tar $IMAGE_REGISTRY/appscode-images/memcached:1.6.29-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-memcached_exporter-v0.14.3-ac.tar $IMAGE_REGISTRY/appscode-images/memcached_exporter:v0.14.3-ac
$CMD push --allow-nondistributable-artifacts --insecure images/prom-memcached-exporter-v0.14.2.tar $IMAGE_REGISTRY/prom/memcached-exporter:v0.14.2
