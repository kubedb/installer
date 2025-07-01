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

$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-4.2.24.tar $IMAGE_REGISTRY/appscode-images/mongo:4.2.24
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-4.4.26.tar $IMAGE_REGISTRY/appscode-images/mongo:4.4.26
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-5.0.23.tar $IMAGE_REGISTRY/appscode-images/mongo:5.0.23
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-5.0.26.tar $IMAGE_REGISTRY/appscode-images/mongo:5.0.26
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-5.0.31.tar $IMAGE_REGISTRY/appscode-images/mongo:5.0.31
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-6.0.12.tar $IMAGE_REGISTRY/appscode-images/mongo:6.0.12
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-6.0.24.tar $IMAGE_REGISTRY/appscode-images/mongo:6.0.24
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-7.0.16.tar $IMAGE_REGISTRY/appscode-images/mongo:7.0.16
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-7.0.21.tar $IMAGE_REGISTRY/appscode-images/mongo:7.0.21
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-7.0.5.tar $IMAGE_REGISTRY/appscode-images/mongo:7.0.5
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-7.0.8.tar $IMAGE_REGISTRY/appscode-images/mongo:7.0.8
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-8.0.10.tar $IMAGE_REGISTRY/appscode-images/mongo:8.0.10
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-8.0.4.tar $IMAGE_REGISTRY/appscode-images/mongo:8.0.4
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-init-4.2-v9.tar $IMAGE_REGISTRY/kubedb/mongodb-init:4.2-v9
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-init-6.0-v11.tar $IMAGE_REGISTRY/kubedb/mongodb-init:6.0-v11
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb_exporter-v0.40.0.tar $IMAGE_REGISTRY/kubedb/mongodb_exporter:v0.40.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-wal-g-v2024.12.18_mongo.tar $IMAGE_REGISTRY/kubedb/wal-g:v2024.12.18_mongo
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-4.2.24.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:4.2.24
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-4.4.26.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:4.4.26
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-5.0.23.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:5.0.23
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-5.0.29.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:5.0.29
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-6.0.12.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:6.0.12
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-6.0.24.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:6.0.24
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-7.0.18.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:7.0.18
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-7.0.4.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:7.0.4
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-8.0.8.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:8.0.8
