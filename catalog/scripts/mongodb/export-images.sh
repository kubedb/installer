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

$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:4.2.24 images/appscode-images-mongo-4.2.24.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:4.4.26 images/appscode-images-mongo-4.4.26.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:5.0.23 images/appscode-images-mongo-5.0.23.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:5.0.26 images/appscode-images-mongo-5.0.26.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:6.0.12 images/appscode-images-mongo-6.0.12.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:7.0.16 images/appscode-images-mongo-7.0.16.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:7.0.5 images/appscode-images-mongo-7.0.5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:7.0.8 images/appscode-images-mongo-7.0.8.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mongo:8.0.4 images/appscode-images-mongo-8.0.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:4.2-v9 images/kubedb-mongodb-init-4.2-v9.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb-init:6.0-v10 images/kubedb-mongodb-init-6.0-v10.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mongodb_exporter:v0.40.0 images/kubedb-mongodb_exporter-v0.40.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/wal-g:v2024.12.18_mongo images/kubedb-wal-g-v2024.12.18_mongo.tar
$CMD pull --allow-nondistributable-artifacts --insecure percona/percona-server-mongodb:4.2.24 images/percona-percona-server-mongodb-4.2.24.tar
$CMD pull --allow-nondistributable-artifacts --insecure percona/percona-server-mongodb:4.4.26 images/percona-percona-server-mongodb-4.4.26.tar
$CMD pull --allow-nondistributable-artifacts --insecure percona/percona-server-mongodb:5.0.23 images/percona-percona-server-mongodb-5.0.23.tar
$CMD pull --allow-nondistributable-artifacts --insecure percona/percona-server-mongodb:6.0.12 images/percona-percona-server-mongodb-6.0.12.tar
$CMD pull --allow-nondistributable-artifacts --insecure percona/percona-server-mongodb:7.0.4 images/percona-percona-server-mongodb-7.0.4.tar

tar -czvf images.tar.gz images
