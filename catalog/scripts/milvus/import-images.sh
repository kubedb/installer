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

$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-milvus-2.6.11.tar $IMAGE_REGISTRY/appscode-images/milvus:2.6.11
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-milvus-2.6.15.tar $IMAGE_REGISTRY/appscode-images/milvus:2.6.15
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-milvus-2.6.15-gpu.tar $IMAGE_REGISTRY/appscode-images/milvus:2.6.15-gpu
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-milvus-2.6.7.tar $IMAGE_REGISTRY/appscode-images/milvus:2.6.7
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-milvus-2.6.9.tar $IMAGE_REGISTRY/appscode-images/milvus:2.6.9
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-milvus-3.0.0.tar $IMAGE_REGISTRY/appscode-images/milvus:3.0.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-milvus-3.0.0-gpu.tar $IMAGE_REGISTRY/appscode-images/milvus:3.0.0-gpu
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-milvus-archiver-v0.2.1.tar $IMAGE_REGISTRY/kubedb/milvus-archiver:v0.2.1
