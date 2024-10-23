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

$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode/kubectl:v1.22 images/appscode-kubectl-v1.22.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/kubedump:0.1.0-v15 images/stashed-kubedump-0.1.0-v15.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-crd-installer:v0.36.0 images/stashed-stash-crd-installer-v0.36.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:5.6.4-v32 images/stashed-stash-elasticsearch-5.6.4-v32.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:6.2.4-v32 images/stashed-stash-elasticsearch-6.2.4-v32.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:6.3.0-v32 images/stashed-stash-elasticsearch-6.3.0-v32.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:6.4.0-v32 images/stashed-stash-elasticsearch-6.4.0-v32.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:6.5.3-v32 images/stashed-stash-elasticsearch-6.5.3-v32.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:6.8.0-v32 images/stashed-stash-elasticsearch-6.8.0-v32.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:7.14.0-v18 images/stashed-stash-elasticsearch-7.14.0-v18.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:7.2.0-v32 images/stashed-stash-elasticsearch-7.2.0-v32.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:7.3.2-v32 images/stashed-stash-elasticsearch-7.3.2-v32.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:8.2.0-v15 images/stashed-stash-elasticsearch-8.2.0-v15.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-enterprise:v0.36.0 images/stashed-stash-enterprise-v0.36.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-etcd:3.5.0-v19 images/stashed-stash-etcd-3.5.0-v19.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mariadb:10.5.8-v26 images/stashed-stash-mariadb-10.5.8-v26.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:3.4.17-v33 images/stashed-stash-mongodb-3.4.17-v33.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:3.4.22-v33 images/stashed-stash-mongodb-3.4.22-v33.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:3.6.13-v33 images/stashed-stash-mongodb-3.6.13-v33.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:3.6.8-v33 images/stashed-stash-mongodb-3.6.8-v33.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.0.11-v33 images/stashed-stash-mongodb-4.0.11-v33.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.0.3-v33 images/stashed-stash-mongodb-4.0.3-v33.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.0.5-v33 images/stashed-stash-mongodb-4.0.5-v33.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.1.13-v33 images/stashed-stash-mongodb-4.1.13-v33.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.1.4-v33 images/stashed-stash-mongodb-4.1.4-v33.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.1.7-v33 images/stashed-stash-mongodb-4.1.7-v33.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.2.3-v33 images/stashed-stash-mongodb-4.2.3-v33.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.4.6-v24 images/stashed-stash-mongodb-4.4.6-v24.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:5.0.15-v6 images/stashed-stash-mongodb-5.0.15-v6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:5.0.3-v21 images/stashed-stash-mongodb-5.0.3-v21.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:6.0.5-v9 images/stashed-stash-mongodb-6.0.5-v9.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mysql:5.7.25-v32 images/stashed-stash-mysql-5.7.25-v32.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mysql:8.0.14-v32 images/stashed-stash-mysql-8.0.14-v32.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mysql:8.0.21-v26 images/stashed-stash-mysql-8.0.21-v26.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mysql:8.0.3-v32 images/stashed-stash-mysql-8.0.3-v32.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-nats:2.6.1-v20 images/stashed-stash-nats-2.6.1-v20.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-nats:2.8.2-v15 images/stashed-stash-nats-2.8.2-v15.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-percona-xtradb:5.7-v26 images/stashed-stash-percona-xtradb-5.7-v26.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:10.14-v31 images/stashed-stash-postgres-10.14-v31.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:11.9-v31 images/stashed-stash-postgres-11.9-v31.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:12.4-v31 images/stashed-stash-postgres-12.4-v31.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:13.1-v28 images/stashed-stash-postgres-13.1-v28.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:14.0-v20 images/stashed-stash-postgres-14.0-v20.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:15.1-v12 images/stashed-stash-postgres-15.1-v12.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:16.1-v1 images/stashed-stash-postgres-16.1-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:9.6.19-v31 images/stashed-stash-postgres-9.6.19-v31.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-redis:5.0.13-v20 images/stashed-stash-redis-5.0.13-v20.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-redis:6.2.5-v20 images/stashed-stash-redis-6.2.5-v20.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-redis:7.0.5-v13 images/stashed-stash-redis-7.0.5-v13.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-ui-server:v0.17.0 images/stashed-stash-ui-server-v0.17.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-vault:1.10.3-v12 images/stashed-stash-vault-1.10.3-v12.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash:v0.36.0 images/stashed-stash-v0.36.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure prom/pushgateway:v1.4.2 images/prom-pushgateway-v1.4.2.tar

tar -czvf images.tar.gz images
