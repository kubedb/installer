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

$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-10.23-alpine.tar $IMAGE_REGISTRY/appscode-images/postgres:10.23-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-10.23-bullseye.tar $IMAGE_REGISTRY/appscode-images/postgres:10.23-bullseye
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-11.22-alpine.tar $IMAGE_REGISTRY/appscode-images/postgres:11.22-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-11.22-bookworm.tar $IMAGE_REGISTRY/appscode-images/postgres:11.22-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-12.17-alpine.tar $IMAGE_REGISTRY/appscode-images/postgres:12.17-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-12.17-bookworm.tar $IMAGE_REGISTRY/appscode-images/postgres:12.17-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-13.13-alpine.tar $IMAGE_REGISTRY/appscode-images/postgres:13.13-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-13.13-bookworm.tar $IMAGE_REGISTRY/appscode-images/postgres:13.13-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-14.10-alpine.tar $IMAGE_REGISTRY/appscode-images/postgres:14.10-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-14.10-bookworm.tar $IMAGE_REGISTRY/appscode-images/postgres:14.10-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-14.13-alpine.tar $IMAGE_REGISTRY/appscode-images/postgres:14.13-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-14.13-bookworm.tar $IMAGE_REGISTRY/appscode-images/postgres:14.13-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-15.5-alpine.tar $IMAGE_REGISTRY/appscode-images/postgres:15.5-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-15.5-bookworm.tar $IMAGE_REGISTRY/appscode-images/postgres:15.5-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-15.8-alpine.tar $IMAGE_REGISTRY/appscode-images/postgres:15.8-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-15.8-bookworm.tar $IMAGE_REGISTRY/appscode-images/postgres:15.8-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-16.1-alpine.tar $IMAGE_REGISTRY/appscode-images/postgres:16.1-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-16.1-bookworm.tar $IMAGE_REGISTRY/appscode-images/postgres:16.1-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-16.4-alpine.tar $IMAGE_REGISTRY/appscode-images/postgres:16.4-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-postgres-16.4-bookworm.tar $IMAGE_REGISTRY/appscode-images/postgres:16.4-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-pg-coordinator-v0.33.0.tar $IMAGE_REGISTRY/kubedb/pg-coordinator:v0.33.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-archiver-v0.10.0_11.22-alpine.tar $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.10.0_11.22-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-archiver-v0.10.0_11.22-bookworm.tar $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.10.0_11.22-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-archiver-v0.10.0_12.17-alpine.tar $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.10.0_12.17-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-archiver-v0.10.0_12.17-bookworm.tar $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.10.0_12.17-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-archiver-v0.10.0_13.13-alpine.tar $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.10.0_13.13-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-archiver-v0.10.0_13.13-bookworm.tar $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.10.0_13.13-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-archiver-v0.10.0_14.10-alpine.tar $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.10.0_14.10-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-archiver-v0.10.0_14.10-bookworm.tar $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.10.0_14.10-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-archiver-v0.10.0_15.5-alpine.tar $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.10.0_15.5-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-archiver-v0.10.0_15.5-bookworm.tar $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.10.0_15.5-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-archiver-v0.10.0_16.1-alpine.tar $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.10.0_16.1-alpine
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-archiver-v0.10.0_16.1-bookworm.tar $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.10.0_16.1-bookworm
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-postgres-init-0.16.0.tar $IMAGE_REGISTRY/kubedb/postgres-init:0.16.0
$CMD push --allow-nondistributable-artifacts --insecure images/postgis-postgis-11-3.3.tar $IMAGE_REGISTRY/postgis/postgis:11-3.3
$CMD push --allow-nondistributable-artifacts --insecure images/postgis-postgis-12-3.4.tar $IMAGE_REGISTRY/postgis/postgis:12-3.4
$CMD push --allow-nondistributable-artifacts --insecure images/postgis-postgis-13-3.4.tar $IMAGE_REGISTRY/postgis/postgis:13-3.4
$CMD push --allow-nondistributable-artifacts --insecure images/postgis-postgis-14-3.4.tar $IMAGE_REGISTRY/postgis/postgis:14-3.4
$CMD push --allow-nondistributable-artifacts --insecure images/postgis-postgis-15-3.4.tar $IMAGE_REGISTRY/postgis/postgis:15-3.4
$CMD push --allow-nondistributable-artifacts --insecure images/postgis-postgis-16-3.4.tar $IMAGE_REGISTRY/postgis/postgis:16-3.4
$CMD push --allow-nondistributable-artifacts --insecure images/prometheuscommunity-postgres-exporter-v0.15.0.tar $IMAGE_REGISTRY/prometheuscommunity/postgres-exporter:v0.15.0
$CMD push --allow-nondistributable-artifacts --insecure images/timescale-timescaledb-2.14.2-pg13-oss.tar $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg13-oss
$CMD push --allow-nondistributable-artifacts --insecure images/timescale-timescaledb-2.14.2-pg14-oss.tar $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg14-oss
$CMD push --allow-nondistributable-artifacts --insecure images/timescale-timescaledb-2.14.2-pg15-oss.tar $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg15-oss
$CMD push --allow-nondistributable-artifacts --insecure images/timescale-timescaledb-2.14.2-pg16-oss.tar $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg16-oss