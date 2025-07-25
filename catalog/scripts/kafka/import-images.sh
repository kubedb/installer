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

$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-connect-cluster-3.5.2.tar $IMAGE_REGISTRY/appscode-images/kafka-connect-cluster:3.5.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-connect-cluster-3.6.1.tar $IMAGE_REGISTRY/appscode-images/kafka-connect-cluster:3.6.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-connect-cluster-3.7.2.tar $IMAGE_REGISTRY/appscode-images/kafka-connect-cluster:3.7.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-connect-cluster-3.8.1.tar $IMAGE_REGISTRY/appscode-images/kafka-connect-cluster:3.8.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-connect-cluster-3.9.0.tar $IMAGE_REGISTRY/appscode-images/kafka-connect-cluster:3.9.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-connect-cluster-4.0.0.tar $IMAGE_REGISTRY/appscode-images/kafka-connect-cluster:4.0.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-cruise-control-3.5.2.tar $IMAGE_REGISTRY/appscode-images/kafka-cruise-control:3.5.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-cruise-control-3.6.1.tar $IMAGE_REGISTRY/appscode-images/kafka-cruise-control:3.6.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-cruise-control-3.7.2.tar $IMAGE_REGISTRY/appscode-images/kafka-cruise-control:3.7.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-cruise-control-3.8.1.tar $IMAGE_REGISTRY/appscode-images/kafka-cruise-control:3.8.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-cruise-control-3.9.0.tar $IMAGE_REGISTRY/appscode-images/kafka-cruise-control:3.9.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-cruise-control-4.0.0.tar $IMAGE_REGISTRY/appscode-images/kafka-cruise-control:4.0.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-kraft-3.5.2.tar $IMAGE_REGISTRY/appscode-images/kafka-kraft:3.5.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-kraft-3.6.1.tar $IMAGE_REGISTRY/appscode-images/kafka-kraft:3.6.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-kraft-3.7.2.tar $IMAGE_REGISTRY/appscode-images/kafka-kraft:3.7.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-kraft-3.8.1.tar $IMAGE_REGISTRY/appscode-images/kafka-kraft:3.8.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-kraft-3.9.0.tar $IMAGE_REGISTRY/appscode-images/kafka-kraft:3.9.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kafka-4.0.0.tar $IMAGE_REGISTRY/appscode-images/kafka:4.0.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kafka-init-4.0-v1.tar $IMAGE_REGISTRY/kubedb/kafka-init:4.0-v1
