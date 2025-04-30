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

$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres-documentdb:15-0.102.0-ferretdb-2.0.0 images/appscode-images-postgres-documentdb-15-0.102.0-ferretdb-2.0.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres-documentdb:16-0.102.0-ferretdb-2.0.0 images/appscode-images-postgres-documentdb-16-0.102.0-ferretdb-2.0.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres-documentdb:17-0.102.0-ferretdb-2.0.0 images/appscode-images-postgres-documentdb-17-0.102.0-ferretdb-2.0.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:10.23-alpine images/appscode-images-postgres-10.23-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:10.23-bullseye images/appscode-images-postgres-10.23-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:11.22-alpine images/appscode-images-postgres-11.22-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:11.22-bookworm images/appscode-images-postgres-11.22-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:12.17-alpine images/appscode-images-postgres-12.17-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:12.17-bookworm images/appscode-images-postgres-12.17-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:12.22-alpine images/appscode-images-postgres-12.22-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:12.22-bookworm images/appscode-images-postgres-12.22-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.13-alpine images/appscode-images-postgres-13.13-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.13-bookworm images/appscode-images-postgres-13.13-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.18-alpine images/appscode-images-postgres-13.18-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.18-bookworm images/appscode-images-postgres-13.18-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.20-alpine images/appscode-images-postgres-13.20-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.20-bookworm images/appscode-images-postgres-13.20-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.10-alpine images/appscode-images-postgres-14.10-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.10-bookworm images/appscode-images-postgres-14.10-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.13-alpine images/appscode-images-postgres-14.13-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.13-bookworm images/appscode-images-postgres-14.13-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.15-alpine images/appscode-images-postgres-14.15-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.15-bookworm images/appscode-images-postgres-14.15-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.17-alpine images/appscode-images-postgres-14.17-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.17-bookworm images/appscode-images-postgres-14.17-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.10-alpine images/appscode-images-postgres-15.10-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.10-bookworm images/appscode-images-postgres-15.10-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.12-alpine images/appscode-images-postgres-15.12-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.12-bookworm images/appscode-images-postgres-15.12-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.5-alpine images/appscode-images-postgres-15.5-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.5-bookworm images/appscode-images-postgres-15.5-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.8-alpine images/appscode-images-postgres-15.8-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.8-bookworm images/appscode-images-postgres-15.8-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.1-alpine images/appscode-images-postgres-16.1-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.1-bookworm images/appscode-images-postgres-16.1-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.4-alpine images/appscode-images-postgres-16.4-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.4-bookworm images/appscode-images-postgres-16.4-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.6-alpine images/appscode-images-postgres-16.6-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.6-bookworm images/appscode-images-postgres-16.6-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.8-alpine images/appscode-images-postgres-16.8-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.8-bookworm images/appscode-images-postgres-16.8-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.2-alpine images/appscode-images-postgres-17.2-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.2-bookworm images/appscode-images-postgres-17.2-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.4-alpine images/appscode-images-postgres-17.4-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.4-bookworm images/appscode-images-postgres-17.4-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/pg-coordinator:v0.38.0 images/kubedb-pg-coordinator-v0.38.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_11.22-alpine images/kubedb-postgres-archiver-v0.15.0_11.22-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_11.22-bookworm images/kubedb-postgres-archiver-v0.15.0_11.22-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_12.17-alpine images/kubedb-postgres-archiver-v0.15.0_12.17-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_12.17-bookworm images/kubedb-postgres-archiver-v0.15.0_12.17-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_13.13-alpine images/kubedb-postgres-archiver-v0.15.0_13.13-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_13.13-bookworm images/kubedb-postgres-archiver-v0.15.0_13.13-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_14.10-alpine images/kubedb-postgres-archiver-v0.15.0_14.10-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_14.10-bookworm images/kubedb-postgres-archiver-v0.15.0_14.10-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_15.5-alpine images/kubedb-postgres-archiver-v0.15.0_15.5-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_15.5-bookworm images/kubedb-postgres-archiver-v0.15.0_15.5-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_16.1-alpine images/kubedb-postgres-archiver-v0.15.0_16.1-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_16.1-bookworm images/kubedb-postgres-archiver-v0.15.0_16.1-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_17.2-alpine images/kubedb-postgres-archiver-v0.15.0_17.2-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.15.0_17.2-bookworm images/kubedb-postgres-archiver-v0.15.0_17.2-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-init:0.17.1 images/kubedb-postgres-init-0.17.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure postgis/postgis:11-3.3 images/postgis-postgis-11-3.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure postgis/postgis:12-3.4 images/postgis-postgis-12-3.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure postgis/postgis:13-3.4 images/postgis-postgis-13-3.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure postgis/postgis:14-3.4 images/postgis-postgis-14-3.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure postgis/postgis:15-3.4 images/postgis-postgis-15-3.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure postgis/postgis:16-3.4 images/postgis-postgis-16-3.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure prometheuscommunity/postgres-exporter:v0.15.0 images/prometheuscommunity-postgres-exporter-v0.15.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure timescale/timescaledb:2.14.2-pg13-oss images/timescale-timescaledb-2.14.2-pg13-oss.tar
$CMD pull --allow-nondistributable-artifacts --insecure timescale/timescaledb:2.14.2-pg14-oss images/timescale-timescaledb-2.14.2-pg14-oss.tar
$CMD pull --allow-nondistributable-artifacts --insecure timescale/timescaledb:2.14.2-pg15-oss images/timescale-timescaledb-2.14.2-pg15-oss.tar
$CMD pull --allow-nondistributable-artifacts --insecure timescale/timescaledb:2.14.2-pg16-oss images/timescale-timescaledb-2.14.2-pg16-oss.tar

tar -czvf images.tar.gz images
