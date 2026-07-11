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

$CMD push --allow-nondistributable-artifacts --insecure images/library-mongo-3.4.17.tar $IMAGE_REGISTRY/mongo:3.4.17
$CMD push --allow-nondistributable-artifacts --insecure images/library-mongo-3.4.22.tar $IMAGE_REGISTRY/mongo:3.4.22
$CMD push --allow-nondistributable-artifacts --insecure images/library-mongo-3.6.13.tar $IMAGE_REGISTRY/mongo:3.6.13
$CMD push --allow-nondistributable-artifacts --insecure images/library-mongo-3.6.8.tar $IMAGE_REGISTRY/mongo:3.6.8
$CMD push --allow-nondistributable-artifacts --insecure images/library-mongo-4.0.11.tar $IMAGE_REGISTRY/mongo:4.0.11
$CMD push --allow-nondistributable-artifacts --insecure images/library-mongo-4.0.3.tar $IMAGE_REGISTRY/mongo:4.0.3
$CMD push --allow-nondistributable-artifacts --insecure images/library-mongo-4.0.5.tar $IMAGE_REGISTRY/mongo:4.0.5
$CMD push --allow-nondistributable-artifacts --insecure images/library-mongo-4.1.13.tar $IMAGE_REGISTRY/mongo:4.1.13
$CMD push --allow-nondistributable-artifacts --insecure images/library-mongo-4.1.4.tar $IMAGE_REGISTRY/mongo:4.1.4
$CMD push --allow-nondistributable-artifacts --insecure images/library-mongo-4.1.7.tar $IMAGE_REGISTRY/mongo:4.1.7
$CMD push --allow-nondistributable-artifacts --insecure images/library-mongo-5.0.2.tar $IMAGE_REGISTRY/mongo:5.0.2
$CMD push --allow-nondistributable-artifacts --insecure images/library-mongo-5.0.3.tar $IMAGE_REGISTRY/mongo:5.0.3
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-3.6.18.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:3.6.18
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-4.0.10.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:4.0.10
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-4.2.24.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:4.2.24
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-4.2.7-7.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:4.2.7-7
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-4.4.10.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:4.4.10
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-4.4.26.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:4.4.26
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-5.0.23.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:5.0.23
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-5.0.29.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:5.0.29
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-6.0.12.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:6.0.12
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-6.0.24.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:6.0.24
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-7.0.18.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:7.0.18
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-7.0.28.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:7.0.28
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-7.0.4.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:7.0.4
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-8.0.17.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:8.0.17
$CMD push --allow-nondistributable-artifacts --insecure images/percona-percona-server-mongodb-8.0.8.tar $IMAGE_REGISTRY/percona/percona-server-mongodb:8.0.8
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-4.2.24.tar $IMAGE_REGISTRY/appscode-images/mongo:4.2.24
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-4.2.3.tar $IMAGE_REGISTRY/appscode-images/mongo:4.2.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-4.4.26.tar $IMAGE_REGISTRY/appscode-images/mongo:4.4.26
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-4.4.6.tar $IMAGE_REGISTRY/appscode-images/mongo:4.4.6
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-5.0.15.tar $IMAGE_REGISTRY/appscode-images/mongo:5.0.15
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-5.0.23.tar $IMAGE_REGISTRY/appscode-images/mongo:5.0.23
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-5.0.26.tar $IMAGE_REGISTRY/appscode-images/mongo:5.0.26
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-5.0.31.tar $IMAGE_REGISTRY/appscode-images/mongo:5.0.31
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-6.0.12.tar $IMAGE_REGISTRY/appscode-images/mongo:6.0.12
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-6.0.24.tar $IMAGE_REGISTRY/appscode-images/mongo:6.0.24
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-6.0.5.tar $IMAGE_REGISTRY/appscode-images/mongo:6.0.5
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-7.0.16.tar $IMAGE_REGISTRY/appscode-images/mongo:7.0.16
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-7.0.21.tar $IMAGE_REGISTRY/appscode-images/mongo:7.0.21
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-7.0.28.tar $IMAGE_REGISTRY/appscode-images/mongo:7.0.28
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-7.0.5.tar $IMAGE_REGISTRY/appscode-images/mongo:7.0.5
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-7.0.8.tar $IMAGE_REGISTRY/appscode-images/mongo:7.0.8
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-8.0.10.tar $IMAGE_REGISTRY/appscode-images/mongo:8.0.10
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-8.0.17.tar $IMAGE_REGISTRY/appscode-images/mongo:8.0.17
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-8.0.3.tar $IMAGE_REGISTRY/appscode-images/mongo:8.0.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-mongo-8.0.4.tar $IMAGE_REGISTRY/appscode-images/mongo:8.0.4
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-kubedb-migrator-mongodb-v0.6.0.tar $IMAGE_REGISTRY/kubedb/kubedb-migrator-mongodb:v0.6.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.4.tar $IMAGE_REGISTRY/kubedb/mongo:3.4
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.4-v1.tar $IMAGE_REGISTRY/kubedb/mongo:3.4-v1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.4-v2.tar $IMAGE_REGISTRY/kubedb/mongo:3.4-v2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.4-v3.tar $IMAGE_REGISTRY/kubedb/mongo:3.4-v3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.4-v4.tar $IMAGE_REGISTRY/kubedb/mongo:3.4-v4
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.4-v5.tar $IMAGE_REGISTRY/kubedb/mongo:3.4-v5
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.4.17.tar $IMAGE_REGISTRY/kubedb/mongo:3.4.17
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.4.22.tar $IMAGE_REGISTRY/kubedb/mongo:3.4.22
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.6.tar $IMAGE_REGISTRY/kubedb/mongo:3.6
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.6-v1.tar $IMAGE_REGISTRY/kubedb/mongo:3.6-v1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.6-v2.tar $IMAGE_REGISTRY/kubedb/mongo:3.6-v2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.6-v3.tar $IMAGE_REGISTRY/kubedb/mongo:3.6-v3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.6-v4.tar $IMAGE_REGISTRY/kubedb/mongo:3.6-v4
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.6-v5.tar $IMAGE_REGISTRY/kubedb/mongo:3.6-v5
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.6.13.tar $IMAGE_REGISTRY/kubedb/mongo:3.6.13
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-3.6.8.tar $IMAGE_REGISTRY/kubedb/mongo:3.6.8
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.0.tar $IMAGE_REGISTRY/kubedb/mongo:4.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.0-v1.tar $IMAGE_REGISTRY/kubedb/mongo:4.0-v1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.0-v2.tar $IMAGE_REGISTRY/kubedb/mongo:4.0-v2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.0-v3.tar $IMAGE_REGISTRY/kubedb/mongo:4.0-v3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.0.11.tar $IMAGE_REGISTRY/kubedb/mongo:4.0.11
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.0.3.tar $IMAGE_REGISTRY/kubedb/mongo:4.0.3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.0.5.tar $IMAGE_REGISTRY/kubedb/mongo:4.0.5
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.0.5-v1.tar $IMAGE_REGISTRY/kubedb/mongo:4.0.5-v1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.0.5-v2.tar $IMAGE_REGISTRY/kubedb/mongo:4.0.5-v2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.1.tar $IMAGE_REGISTRY/kubedb/mongo:4.1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.1-v1.tar $IMAGE_REGISTRY/kubedb/mongo:4.1-v1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.1.13.tar $IMAGE_REGISTRY/kubedb/mongo:4.1.13
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.1.4.tar $IMAGE_REGISTRY/kubedb/mongo:4.1.4
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.1.7.tar $IMAGE_REGISTRY/kubedb/mongo:4.1.7
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.1.7-v1.tar $IMAGE_REGISTRY/kubedb/mongo:4.1.7-v1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.1.7-v2.tar $IMAGE_REGISTRY/kubedb/mongo:4.1.7-v2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongo-4.2.tar $IMAGE_REGISTRY/kubedb/mongo:4.2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-init-0.1.0.tar $IMAGE_REGISTRY/kubedb/mongodb-init:0.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-init-0.3.0.tar $IMAGE_REGISTRY/kubedb/mongodb-init:0.3.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-init-4.1-v9.tar $IMAGE_REGISTRY/kubedb/mongodb-init:4.1-v9
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-init-4.1.4-v9.tar $IMAGE_REGISTRY/kubedb/mongodb-init:4.1.4-v9
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-init-4.2-v9.tar $IMAGE_REGISTRY/kubedb/mongodb-init:4.2-v9
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb-init-6.0-v11.tar $IMAGE_REGISTRY/kubedb/mongodb-init:6.0-v11
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb_exporter-v0.20.4.tar $IMAGE_REGISTRY/kubedb/mongodb_exporter:v0.20.4
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mongodb_exporter-v0.47.2.tar $IMAGE_REGISTRY/kubedb/mongodb_exporter:v0.47.2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-wal-g-v2026.3.30_mongo.tar $IMAGE_REGISTRY/kubedb/wal-g:v2026.3.30_mongo
