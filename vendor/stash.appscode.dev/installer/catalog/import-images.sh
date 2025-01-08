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

$CMD push --allow-nondistributable-artifacts --insecure images/appscode-kubectl-nonroot-1.31.tar $IMAGE_REGISTRY/appscode/kubectl-nonroot:1.31
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-kubedump-0.2.0-v2.tar $IMAGE_REGISTRY/stashed/kubedump:0.2.0-v2
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-crd-installer-v0.38.0.tar $IMAGE_REGISTRY/stashed/stash-crd-installer:v0.38.0
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-5.6.4-v34.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:5.6.4-v34
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-6.2.4-v34.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.2.4-v34
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-6.3.0-v34.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.3.0-v34
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-6.4.0-v34.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.4.0-v34
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-6.5.3-v34.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.5.3-v34
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-6.8.0-v34.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.8.0-v34
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-7.14.0-v20.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:7.14.0-v20
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-7.2.0-v34.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:7.2.0-v34
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-7.3.2-v34.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:7.3.2-v34
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-8.2.0-v17.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:8.2.0-v17
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-enterprise-v0.38.0.tar $IMAGE_REGISTRY/stashed/stash-enterprise:v0.38.0
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-etcd-3.5.0-v21.tar $IMAGE_REGISTRY/stashed/stash-etcd:3.5.0-v21
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mariadb-10.5.8-v28.tar $IMAGE_REGISTRY/stashed/stash-mariadb:10.5.8-v28
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-3.4.17-v35.tar $IMAGE_REGISTRY/stashed/stash-mongodb:3.4.17-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-3.4.22-v35.tar $IMAGE_REGISTRY/stashed/stash-mongodb:3.4.22-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-3.6.13-v35.tar $IMAGE_REGISTRY/stashed/stash-mongodb:3.6.13-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-3.6.8-v35.tar $IMAGE_REGISTRY/stashed/stash-mongodb:3.6.8-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.0.11-v35.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.0.11-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.0.3-v35.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.0.3-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.0.5-v35.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.0.5-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.1.13-v35.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.1.13-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.1.4-v35.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.1.4-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.1.7-v35.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.1.7-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.2.3-v35.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.2.3-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.4.6-v26.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.4.6-v26
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-5.0.15-v8.tar $IMAGE_REGISTRY/stashed/stash-mongodb:5.0.15-v8
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-5.0.3-v23.tar $IMAGE_REGISTRY/stashed/stash-mongodb:5.0.3-v23
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-6.0.5-v11.tar $IMAGE_REGISTRY/stashed/stash-mongodb:6.0.5-v11
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mysql-5.7.25-v35.tar $IMAGE_REGISTRY/stashed/stash-mysql:5.7.25-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mysql-8.0.14-v34.tar $IMAGE_REGISTRY/stashed/stash-mysql:8.0.14-v34
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mysql-8.0.21-v28.tar $IMAGE_REGISTRY/stashed/stash-mysql:8.0.21-v28
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mysql-8.0.3-v34.tar $IMAGE_REGISTRY/stashed/stash-mysql:8.0.3-v34
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-nats-2.6.1-v22.tar $IMAGE_REGISTRY/stashed/stash-nats:2.6.1-v22
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-nats-2.8.2-v17.tar $IMAGE_REGISTRY/stashed/stash-nats:2.8.2-v17
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-percona-xtradb-5.7-v26.tar $IMAGE_REGISTRY/stashed/stash-percona-xtradb:5.7-v26
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-10.14-v33.tar $IMAGE_REGISTRY/stashed/stash-postgres:10.14-v33
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-11.9-v33.tar $IMAGE_REGISTRY/stashed/stash-postgres:11.9-v33
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-12.4-v33.tar $IMAGE_REGISTRY/stashed/stash-postgres:12.4-v33
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-13.1-v30.tar $IMAGE_REGISTRY/stashed/stash-postgres:13.1-v30
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-14.0-v22.tar $IMAGE_REGISTRY/stashed/stash-postgres:14.0-v22
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-15.1-v14.tar $IMAGE_REGISTRY/stashed/stash-postgres:15.1-v14
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-16.1-v3.tar $IMAGE_REGISTRY/stashed/stash-postgres:16.1-v3
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-17.2-v1.tar $IMAGE_REGISTRY/stashed/stash-postgres:17.2-v1
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-9.6.19-v33.tar $IMAGE_REGISTRY/stashed/stash-postgres:9.6.19-v33
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-redis-5.0.13-v22.tar $IMAGE_REGISTRY/stashed/stash-redis:5.0.13-v22
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-redis-6.2.5-v22.tar $IMAGE_REGISTRY/stashed/stash-redis:6.2.5-v22
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-redis-7.0.5-v15.tar $IMAGE_REGISTRY/stashed/stash-redis:7.0.5-v15
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-ui-server-v0.19.0.tar $IMAGE_REGISTRY/stashed/stash-ui-server:v0.19.0
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-vault-1.10.3-v14.tar $IMAGE_REGISTRY/stashed/stash-vault:1.10.3-v14
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-v0.38.0.tar $IMAGE_REGISTRY/stashed/stash:v0.38.0
$CMD push --allow-nondistributable-artifacts --insecure images/prom-pushgateway-v1.4.2.tar $IMAGE_REGISTRY/prom/pushgateway:v1.4.2
