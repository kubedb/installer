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

$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode/kubectl-nonroot:1.31 $IMAGE_REGISTRY/appscode/kubectl-nonroot:1.31
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/kubedump:0.2.0-v2 $IMAGE_REGISTRY/stashed/kubedump:0.2.0-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-crd-installer:v0.38.0 $IMAGE_REGISTRY/stashed/stash-crd-installer:v0.38.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:5.6.4-v34 $IMAGE_REGISTRY/stashed/stash-elasticsearch:5.6.4-v34
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:6.2.4-v34 $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.2.4-v34
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:6.3.0-v34 $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.3.0-v34
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:6.4.0-v34 $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.4.0-v34
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:6.5.3-v34 $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.5.3-v34
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:6.8.0-v34 $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.8.0-v34
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:7.14.0-v20 $IMAGE_REGISTRY/stashed/stash-elasticsearch:7.14.0-v20
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:7.2.0-v34 $IMAGE_REGISTRY/stashed/stash-elasticsearch:7.2.0-v34
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:7.3.2-v34 $IMAGE_REGISTRY/stashed/stash-elasticsearch:7.3.2-v34
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-elasticsearch:8.2.0-v17 $IMAGE_REGISTRY/stashed/stash-elasticsearch:8.2.0-v17
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-enterprise:v0.38.0 $IMAGE_REGISTRY/stashed/stash-enterprise:v0.38.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-etcd:3.5.0-v21 $IMAGE_REGISTRY/stashed/stash-etcd:3.5.0-v21
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mariadb:10.5.8-v28 $IMAGE_REGISTRY/stashed/stash-mariadb:10.5.8-v28
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:3.4.17-v35 $IMAGE_REGISTRY/stashed/stash-mongodb:3.4.17-v35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:3.4.22-v35 $IMAGE_REGISTRY/stashed/stash-mongodb:3.4.22-v35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:3.6.13-v35 $IMAGE_REGISTRY/stashed/stash-mongodb:3.6.13-v35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:3.6.8-v35 $IMAGE_REGISTRY/stashed/stash-mongodb:3.6.8-v35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.0.11-v35 $IMAGE_REGISTRY/stashed/stash-mongodb:4.0.11-v35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.0.3-v35 $IMAGE_REGISTRY/stashed/stash-mongodb:4.0.3-v35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.0.5-v35 $IMAGE_REGISTRY/stashed/stash-mongodb:4.0.5-v35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.1.13-v35 $IMAGE_REGISTRY/stashed/stash-mongodb:4.1.13-v35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.1.4-v35 $IMAGE_REGISTRY/stashed/stash-mongodb:4.1.4-v35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.1.7-v35 $IMAGE_REGISTRY/stashed/stash-mongodb:4.1.7-v35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.2.3-v35 $IMAGE_REGISTRY/stashed/stash-mongodb:4.2.3-v35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:4.4.6-v26 $IMAGE_REGISTRY/stashed/stash-mongodb:4.4.6-v26
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:5.0.15-v8 $IMAGE_REGISTRY/stashed/stash-mongodb:5.0.15-v8
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:5.0.3-v23 $IMAGE_REGISTRY/stashed/stash-mongodb:5.0.3-v23
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mongodb:6.0.5-v11 $IMAGE_REGISTRY/stashed/stash-mongodb:6.0.5-v11
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mysql:5.7.25-v35 $IMAGE_REGISTRY/stashed/stash-mysql:5.7.25-v35
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mysql:8.0.14-v34 $IMAGE_REGISTRY/stashed/stash-mysql:8.0.14-v34
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mysql:8.0.21-v28 $IMAGE_REGISTRY/stashed/stash-mysql:8.0.21-v28
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-mysql:8.0.3-v34 $IMAGE_REGISTRY/stashed/stash-mysql:8.0.3-v34
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-nats:2.6.1-v22 $IMAGE_REGISTRY/stashed/stash-nats:2.6.1-v22
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-nats:2.8.2-v17 $IMAGE_REGISTRY/stashed/stash-nats:2.8.2-v17
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-percona-xtradb:5.7-v26 $IMAGE_REGISTRY/stashed/stash-percona-xtradb:5.7-v26
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:10.14-v33 $IMAGE_REGISTRY/stashed/stash-postgres:10.14-v33
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:11.9-v33 $IMAGE_REGISTRY/stashed/stash-postgres:11.9-v33
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:12.4-v33 $IMAGE_REGISTRY/stashed/stash-postgres:12.4-v33
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:13.1-v30 $IMAGE_REGISTRY/stashed/stash-postgres:13.1-v30
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:14.0-v22 $IMAGE_REGISTRY/stashed/stash-postgres:14.0-v22
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:15.1-v14 $IMAGE_REGISTRY/stashed/stash-postgres:15.1-v14
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:16.1-v3 $IMAGE_REGISTRY/stashed/stash-postgres:16.1-v3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:17.2-v1 $IMAGE_REGISTRY/stashed/stash-postgres:17.2-v1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-postgres:9.6.19-v33 $IMAGE_REGISTRY/stashed/stash-postgres:9.6.19-v33
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-redis:5.0.13-v22 $IMAGE_REGISTRY/stashed/stash-redis:5.0.13-v22
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-redis:6.2.5-v22 $IMAGE_REGISTRY/stashed/stash-redis:6.2.5-v22
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-redis:7.0.5-v15 $IMAGE_REGISTRY/stashed/stash-redis:7.0.5-v15
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-ui-server:v0.19.0 $IMAGE_REGISTRY/stashed/stash-ui-server:v0.19.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash-vault:1.10.3-v14 $IMAGE_REGISTRY/stashed/stash-vault:1.10.3-v14
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/stashed/stash:v0.38.0 $IMAGE_REGISTRY/stashed/stash:v0.38.0
$CMD cp --allow-nondistributable-artifacts --insecure prom/pushgateway:v1.4.2 $IMAGE_REGISTRY/prom/pushgateway:v1.4.2
