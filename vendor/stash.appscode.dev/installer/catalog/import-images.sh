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
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-kubedump-0.2.0-v4.tar $IMAGE_REGISTRY/stashed/kubedump:0.2.0-v4
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-crd-installer-v0.40.0.tar $IMAGE_REGISTRY/stashed/stash-crd-installer:v0.40.0
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-5.6.4-v36.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:5.6.4-v36
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-6.2.4-v36.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.2.4-v36
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-6.3.0-v36.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.3.0-v36
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-6.4.0-v36.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.4.0-v36
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-6.5.3-v36.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.5.3-v36
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-6.8.0-v36.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:6.8.0-v36
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-7.14.0-v22.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:7.14.0-v22
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-7.2.0-v36.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:7.2.0-v36
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-7.3.2-v36.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:7.3.2-v36
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-elasticsearch-8.2.0-v19.tar $IMAGE_REGISTRY/stashed/stash-elasticsearch:8.2.0-v19
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-enterprise-v0.40.0.tar $IMAGE_REGISTRY/stashed/stash-enterprise:v0.40.0
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-etcd-3.5.0-v23.tar $IMAGE_REGISTRY/stashed/stash-etcd:3.5.0-v23
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mariadb-10.5.8-v30.tar $IMAGE_REGISTRY/stashed/stash-mariadb:10.5.8-v30
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-3.4.17-v37.tar $IMAGE_REGISTRY/stashed/stash-mongodb:3.4.17-v37
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-3.4.22-v37.tar $IMAGE_REGISTRY/stashed/stash-mongodb:3.4.22-v37
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-3.6.13-v37.tar $IMAGE_REGISTRY/stashed/stash-mongodb:3.6.13-v37
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-3.6.8-v37.tar $IMAGE_REGISTRY/stashed/stash-mongodb:3.6.8-v37
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.0.11-v37.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.0.11-v37
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.0.3-v37.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.0.3-v37
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.0.5-v37.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.0.5-v37
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.1.13-v37.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.1.13-v37
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.1.4-v37.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.1.4-v37
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.1.7-v37.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.1.7-v37
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.2.3-v37.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.2.3-v37
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-4.4.6-v28.tar $IMAGE_REGISTRY/stashed/stash-mongodb:4.4.6-v28
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-5.0.15-v10.tar $IMAGE_REGISTRY/stashed/stash-mongodb:5.0.15-v10
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-5.0.3-v25.tar $IMAGE_REGISTRY/stashed/stash-mongodb:5.0.3-v25
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mongodb-6.0.5-v13.tar $IMAGE_REGISTRY/stashed/stash-mongodb:6.0.5-v13
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mysql-5.7.25-v37.tar $IMAGE_REGISTRY/stashed/stash-mysql:5.7.25-v37
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mysql-8.0.14-v36.tar $IMAGE_REGISTRY/stashed/stash-mysql:8.0.14-v36
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mysql-8.0.21-v30.tar $IMAGE_REGISTRY/stashed/stash-mysql:8.0.21-v30
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-mysql-8.0.3-v36.tar $IMAGE_REGISTRY/stashed/stash-mysql:8.0.3-v36
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-nats-2.6.1-v24.tar $IMAGE_REGISTRY/stashed/stash-nats:2.6.1-v24
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-nats-2.8.2-v19.tar $IMAGE_REGISTRY/stashed/stash-nats:2.8.2-v19
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-percona-xtradb-5.7-v26.tar $IMAGE_REGISTRY/stashed/stash-percona-xtradb:5.7-v26
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-percona-xtradb-8.0.tar $IMAGE_REGISTRY/stashed/stash-percona-xtradb:8.0
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-percona-xtradb-8.4.tar $IMAGE_REGISTRY/stashed/stash-percona-xtradb:8.4
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-10.14-v35.tar $IMAGE_REGISTRY/stashed/stash-postgres:10.14-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-11.9-v35.tar $IMAGE_REGISTRY/stashed/stash-postgres:11.9-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-12.4-v35.tar $IMAGE_REGISTRY/stashed/stash-postgres:12.4-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-13.1-v32.tar $IMAGE_REGISTRY/stashed/stash-postgres:13.1-v32
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-14.0-v24.tar $IMAGE_REGISTRY/stashed/stash-postgres:14.0-v24
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-15.1-v16.tar $IMAGE_REGISTRY/stashed/stash-postgres:15.1-v16
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-16.1-v5.tar $IMAGE_REGISTRY/stashed/stash-postgres:16.1-v5
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-17.2-v3.tar $IMAGE_REGISTRY/stashed/stash-postgres:17.2-v3
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-postgres-9.6.19-v35.tar $IMAGE_REGISTRY/stashed/stash-postgres:9.6.19-v35
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-redis-5.0.13-v24.tar $IMAGE_REGISTRY/stashed/stash-redis:5.0.13-v24
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-redis-6.2.5-v24.tar $IMAGE_REGISTRY/stashed/stash-redis:6.2.5-v24
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-redis-7.0.5-v17.tar $IMAGE_REGISTRY/stashed/stash-redis:7.0.5-v17
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-ui-server-v0.21.0.tar $IMAGE_REGISTRY/stashed/stash-ui-server:v0.21.0
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-vault-1.10.3-v16.tar $IMAGE_REGISTRY/stashed/stash-vault:1.10.3-v16
$CMD push --allow-nondistributable-artifacts --insecure images/stashed-stash-v0.40.0.tar $IMAGE_REGISTRY/stashed/stash:v0.40.0
$CMD push --allow-nondistributable-artifacts --insecure images/prom-pushgateway-v1.4.2.tar $IMAGE_REGISTRY/prom/pushgateway:v1.4.2
