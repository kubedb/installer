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

$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/mongo:3.4.17 images/library-mongo-3.4.17.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/mongo:3.4.22 images/library-mongo-3.4.22.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/mongo:3.6.13 images/library-mongo-3.6.13.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/mongo:3.6.8 images/library-mongo-3.6.8.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/mongo:4.0.11 images/library-mongo-4.0.11.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/mongo:4.0.3 images/library-mongo-4.0.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/mongo:4.0.5 images/library-mongo-4.0.5.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/mongo:4.1.13 images/library-mongo-4.1.13.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/mongo:4.1.4 images/library-mongo-4.1.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/mongo:4.1.7 images/library-mongo-4.1.7.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/mongo:5.0.2 images/library-mongo-5.0.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/mongo:5.0.3 images/library-mongo-5.0.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:3.6.18 images/percona-percona-server-mongodb-3.6.18.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:4.0.10 images/percona-percona-server-mongodb-4.0.10.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:4.2.24 images/percona-percona-server-mongodb-4.2.24.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:4.2.7-7 images/percona-percona-server-mongodb-4.2.7-7.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:4.4.10 images/percona-percona-server-mongodb-4.4.10.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:4.4.26 images/percona-percona-server-mongodb-4.4.26.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:5.0.23 images/percona-percona-server-mongodb-5.0.23.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:5.0.29 images/percona-percona-server-mongodb-5.0.29.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:6.0.12 images/percona-percona-server-mongodb-6.0.12.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:6.0.24 images/percona-percona-server-mongodb-6.0.24.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:7.0.18 images/percona-percona-server-mongodb-7.0.18.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:7.0.28 images/percona-percona-server-mongodb-7.0.28.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:7.0.4 images/percona-percona-server-mongodb-7.0.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:8.0.17 images/percona-percona-server-mongodb-8.0.17.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/percona/percona-server-mongodb:8.0.8 images/percona-percona-server-mongodb-8.0.8.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:4.2.24 images/appscode-images-mongo-4.2.24.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:4.2.3 images/appscode-images-mongo-4.2.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:4.4.26 images/appscode-images-mongo-4.4.26.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:4.4.6 images/appscode-images-mongo-4.4.6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:5.0.15 images/appscode-images-mongo-5.0.15.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:5.0.23 images/appscode-images-mongo-5.0.23.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:5.0.26 images/appscode-images-mongo-5.0.26.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:5.0.31 images/appscode-images-mongo-5.0.31.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:6.0.12 images/appscode-images-mongo-6.0.12.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:6.0.24 images/appscode-images-mongo-6.0.24.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:6.0.5 images/appscode-images-mongo-6.0.5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:7.0.16 images/appscode-images-mongo-7.0.16.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:7.0.21 images/appscode-images-mongo-7.0.21.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:7.0.28 images/appscode-images-mongo-7.0.28.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:7.0.5 images/appscode-images-mongo-7.0.5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:7.0.8 images/appscode-images-mongo-7.0.8.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:8.0.10 images/appscode-images-mongo-8.0.10.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:8.0.17 images/appscode-images-mongo-8.0.17.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:8.0.3 images/appscode-images-mongo-8.0.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:8.0.4 images/appscode-images-mongo-8.0.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-migrator-mongodb:v0.6.0 images/kubedb-kubedb-migrator-mongodb-v0.6.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4 images/kubedb-mongo-3.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4-v1 images/kubedb-mongo-3.4-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4-v2 images/kubedb-mongo-3.4-v2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4-v3 images/kubedb-mongo-3.4-v3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4-v4 images/kubedb-mongo-3.4-v4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4-v5 images/kubedb-mongo-3.4-v5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4.17 images/kubedb-mongo-3.4.17.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.4.22 images/kubedb-mongo-3.4.22.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6 images/kubedb-mongo-3.6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6-v1 images/kubedb-mongo-3.6-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6-v2 images/kubedb-mongo-3.6-v2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6-v3 images/kubedb-mongo-3.6-v3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6-v4 images/kubedb-mongo-3.6-v4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6-v5 images/kubedb-mongo-3.6-v5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6.13 images/kubedb-mongo-3.6.13.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:3.6.8 images/kubedb-mongo-3.6.8.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0 images/kubedb-mongo-4.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0-v1 images/kubedb-mongo-4.0-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0-v2 images/kubedb-mongo-4.0-v2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0-v3 images/kubedb-mongo-4.0-v3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0.11 images/kubedb-mongo-4.0.11.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0.3 images/kubedb-mongo-4.0.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0.5 images/kubedb-mongo-4.0.5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0.5-v1 images/kubedb-mongo-4.0.5-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.0.5-v2 images/kubedb-mongo-4.0.5-v2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1 images/kubedb-mongo-4.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1-v1 images/kubedb-mongo-4.1-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1.13 images/kubedb-mongo-4.1.13.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1.4 images/kubedb-mongo-4.1.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1.7 images/kubedb-mongo-4.1.7.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1.7-v1 images/kubedb-mongo-4.1.7-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.1.7-v2 images/kubedb-mongo-4.1.7-v2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongo:4.2 images/kubedb-mongo-4.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:0.1.0 images/kubedb-mongodb-init-0.1.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:0.3.0 images/kubedb-mongodb-init-0.3.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:4.1-v9 images/kubedb-mongodb-init-4.1-v9.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:4.1.4-v9 images/kubedb-mongodb-init-4.1.4-v9.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:4.2-v9 images/kubedb-mongodb-init-4.2-v9.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:6.0-v11 images/kubedb-mongodb-init-6.0-v11.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb_exporter:v0.20.4 images/kubedb-mongodb_exporter-v0.20.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb_exporter:v0.47.2 images/kubedb-mongodb_exporter-v0.47.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/wal-g:v2026.3.30_mongo images/kubedb-wal-g-v2026.3.30_mongo.tar

tar -czvf images.tar.gz images
