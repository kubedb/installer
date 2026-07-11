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

$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/mongo:3.4.17 $IMAGE_REGISTRY/mongo:3.4.17
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/mongo:3.4.22 $IMAGE_REGISTRY/mongo:3.4.22
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/mongo:3.6.13 $IMAGE_REGISTRY/mongo:3.6.13
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/mongo:3.6.8 $IMAGE_REGISTRY/mongo:3.6.8
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/mongo:4.0.11 $IMAGE_REGISTRY/mongo:4.0.11
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/mongo:4.0.3 $IMAGE_REGISTRY/mongo:4.0.3
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/mongo:4.0.5 $IMAGE_REGISTRY/mongo:4.0.5
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/mongo:4.1.13 $IMAGE_REGISTRY/mongo:4.1.13
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/mongo:4.1.4 $IMAGE_REGISTRY/mongo:4.1.4
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/mongo:4.1.7 $IMAGE_REGISTRY/mongo:4.1.7
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/mongo:5.0.2 $IMAGE_REGISTRY/mongo:5.0.2
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/mongo:5.0.3 $IMAGE_REGISTRY/mongo:5.0.3
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:3.6.18 $IMAGE_REGISTRY/percona/percona-server-mongodb:3.6.18
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:4.0.10 $IMAGE_REGISTRY/percona/percona-server-mongodb:4.0.10
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:4.2.24 $IMAGE_REGISTRY/percona/percona-server-mongodb:4.2.24
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:4.2.7-7 $IMAGE_REGISTRY/percona/percona-server-mongodb:4.2.7-7
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:4.4.10 $IMAGE_REGISTRY/percona/percona-server-mongodb:4.4.10
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:4.4.26 $IMAGE_REGISTRY/percona/percona-server-mongodb:4.4.26
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:5.0.23 $IMAGE_REGISTRY/percona/percona-server-mongodb:5.0.23
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:5.0.29 $IMAGE_REGISTRY/percona/percona-server-mongodb:5.0.29
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:6.0.12 $IMAGE_REGISTRY/percona/percona-server-mongodb:6.0.12
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:6.0.24 $IMAGE_REGISTRY/percona/percona-server-mongodb:6.0.24
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:7.0.18 $IMAGE_REGISTRY/percona/percona-server-mongodb:7.0.18
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:7.0.28 $IMAGE_REGISTRY/percona/percona-server-mongodb:7.0.28
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:7.0.4 $IMAGE_REGISTRY/percona/percona-server-mongodb:7.0.4
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:8.0.17 $IMAGE_REGISTRY/percona/percona-server-mongodb:8.0.17
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:8.0.8 $IMAGE_REGISTRY/percona/percona-server-mongodb:8.0.8
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:4.2.24 $IMAGE_REGISTRY/appscode-images/mongo:4.2.24
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:4.2.3 $IMAGE_REGISTRY/appscode-images/mongo:4.2.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:4.4.26 $IMAGE_REGISTRY/appscode-images/mongo:4.4.26
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:4.4.6 $IMAGE_REGISTRY/appscode-images/mongo:4.4.6
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:5.0.15 $IMAGE_REGISTRY/appscode-images/mongo:5.0.15
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:5.0.23 $IMAGE_REGISTRY/appscode-images/mongo:5.0.23
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:5.0.26 $IMAGE_REGISTRY/appscode-images/mongo:5.0.26
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:5.0.31 $IMAGE_REGISTRY/appscode-images/mongo:5.0.31
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:6.0.12 $IMAGE_REGISTRY/appscode-images/mongo:6.0.12
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:6.0.24 $IMAGE_REGISTRY/appscode-images/mongo:6.0.24
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:6.0.5 $IMAGE_REGISTRY/appscode-images/mongo:6.0.5
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:7.0.16 $IMAGE_REGISTRY/appscode-images/mongo:7.0.16
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:7.0.21 $IMAGE_REGISTRY/appscode-images/mongo:7.0.21
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:7.0.28 $IMAGE_REGISTRY/appscode-images/mongo:7.0.28
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:7.0.5 $IMAGE_REGISTRY/appscode-images/mongo:7.0.5
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:7.0.8 $IMAGE_REGISTRY/appscode-images/mongo:7.0.8
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:8.0.10 $IMAGE_REGISTRY/appscode-images/mongo:8.0.10
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:8.0.17 $IMAGE_REGISTRY/appscode-images/mongo:8.0.17
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:8.0.3 $IMAGE_REGISTRY/appscode-images/mongo:8.0.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:8.0.4 $IMAGE_REGISTRY/appscode-images/mongo:8.0.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-migrator-mongodb:v0.6.0 $IMAGE_REGISTRY/kubedb/kubedb-migrator-mongodb:v0.6.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4 $IMAGE_REGISTRY/kubedb/mongo:3.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4-v1 $IMAGE_REGISTRY/kubedb/mongo:3.4-v1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4-v2 $IMAGE_REGISTRY/kubedb/mongo:3.4-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4-v3 $IMAGE_REGISTRY/kubedb/mongo:3.4-v3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4-v4 $IMAGE_REGISTRY/kubedb/mongo:3.4-v4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4-v5 $IMAGE_REGISTRY/kubedb/mongo:3.4-v5
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4.17 $IMAGE_REGISTRY/kubedb/mongo:3.4.17
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4.22 $IMAGE_REGISTRY/kubedb/mongo:3.4.22
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6 $IMAGE_REGISTRY/kubedb/mongo:3.6
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6-v1 $IMAGE_REGISTRY/kubedb/mongo:3.6-v1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6-v2 $IMAGE_REGISTRY/kubedb/mongo:3.6-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6-v3 $IMAGE_REGISTRY/kubedb/mongo:3.6-v3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6-v4 $IMAGE_REGISTRY/kubedb/mongo:3.6-v4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6-v5 $IMAGE_REGISTRY/kubedb/mongo:3.6-v5
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6.13 $IMAGE_REGISTRY/kubedb/mongo:3.6.13
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6.8 $IMAGE_REGISTRY/kubedb/mongo:3.6.8
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0 $IMAGE_REGISTRY/kubedb/mongo:4.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0-v1 $IMAGE_REGISTRY/kubedb/mongo:4.0-v1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0-v2 $IMAGE_REGISTRY/kubedb/mongo:4.0-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0-v3 $IMAGE_REGISTRY/kubedb/mongo:4.0-v3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0.11 $IMAGE_REGISTRY/kubedb/mongo:4.0.11
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0.3 $IMAGE_REGISTRY/kubedb/mongo:4.0.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0.5 $IMAGE_REGISTRY/kubedb/mongo:4.0.5
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0.5-v1 $IMAGE_REGISTRY/kubedb/mongo:4.0.5-v1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0.5-v2 $IMAGE_REGISTRY/kubedb/mongo:4.0.5-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1 $IMAGE_REGISTRY/kubedb/mongo:4.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1-v1 $IMAGE_REGISTRY/kubedb/mongo:4.1-v1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1.13 $IMAGE_REGISTRY/kubedb/mongo:4.1.13
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1.4 $IMAGE_REGISTRY/kubedb/mongo:4.1.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1.7 $IMAGE_REGISTRY/kubedb/mongo:4.1.7
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1.7-v1 $IMAGE_REGISTRY/kubedb/mongo:4.1.7-v1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1.7-v2 $IMAGE_REGISTRY/kubedb/mongo:4.1.7-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.2 $IMAGE_REGISTRY/kubedb/mongo:4.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:0.1.0 $IMAGE_REGISTRY/kubedb/mongodb-init:0.1.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:0.3.0 $IMAGE_REGISTRY/kubedb/mongodb-init:0.3.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:4.1-v9 $IMAGE_REGISTRY/kubedb/mongodb-init:4.1-v9
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:4.1.4-v9 $IMAGE_REGISTRY/kubedb/mongodb-init:4.1.4-v9
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:4.2-v9 $IMAGE_REGISTRY/kubedb/mongodb-init:4.2-v9
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:6.0-v11 $IMAGE_REGISTRY/kubedb/mongodb-init:6.0-v11
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb_exporter:v0.20.4 $IMAGE_REGISTRY/kubedb/mongodb_exporter:v0.20.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb_exporter:v0.47.2 $IMAGE_REGISTRY/kubedb/mongodb_exporter:v0.47.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/wal-g:v2026.3.30_mongo $IMAGE_REGISTRY/kubedb/wal-g:v2026.3.30_mongo
