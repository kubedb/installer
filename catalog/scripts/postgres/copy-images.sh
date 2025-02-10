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

$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:10.23-alpine $IMAGE_REGISTRY/appscode-images/postgres:10.23-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:10.23-bullseye $IMAGE_REGISTRY/appscode-images/postgres:10.23-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:11.22-alpine $IMAGE_REGISTRY/appscode-images/postgres:11.22-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:11.22-bookworm $IMAGE_REGISTRY/appscode-images/postgres:11.22-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:12.17-alpine $IMAGE_REGISTRY/appscode-images/postgres:12.17-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:12.17-bookworm $IMAGE_REGISTRY/appscode-images/postgres:12.17-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:12.22-alpine $IMAGE_REGISTRY/appscode-images/postgres:12.22-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:12.22-bookworm $IMAGE_REGISTRY/appscode-images/postgres:12.22-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.13-alpine $IMAGE_REGISTRY/appscode-images/postgres:13.13-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.13-bookworm $IMAGE_REGISTRY/appscode-images/postgres:13.13-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.18-alpine $IMAGE_REGISTRY/appscode-images/postgres:13.18-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.18-bookworm $IMAGE_REGISTRY/appscode-images/postgres:13.18-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.10-alpine $IMAGE_REGISTRY/appscode-images/postgres:14.10-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.10-bookworm $IMAGE_REGISTRY/appscode-images/postgres:14.10-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.13-alpine $IMAGE_REGISTRY/appscode-images/postgres:14.13-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.13-bookworm $IMAGE_REGISTRY/appscode-images/postgres:14.13-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.15-alpine $IMAGE_REGISTRY/appscode-images/postgres:14.15-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.15-bookworm $IMAGE_REGISTRY/appscode-images/postgres:14.15-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.10-alpine $IMAGE_REGISTRY/appscode-images/postgres:15.10-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.10-bookworm $IMAGE_REGISTRY/appscode-images/postgres:15.10-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.5-alpine $IMAGE_REGISTRY/appscode-images/postgres:15.5-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.5-bookworm $IMAGE_REGISTRY/appscode-images/postgres:15.5-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.8-alpine $IMAGE_REGISTRY/appscode-images/postgres:15.8-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.8-bookworm $IMAGE_REGISTRY/appscode-images/postgres:15.8-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.1-alpine $IMAGE_REGISTRY/appscode-images/postgres:16.1-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.1-bookworm $IMAGE_REGISTRY/appscode-images/postgres:16.1-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.4-alpine $IMAGE_REGISTRY/appscode-images/postgres:16.4-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.4-bookworm $IMAGE_REGISTRY/appscode-images/postgres:16.4-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.6-alpine $IMAGE_REGISTRY/appscode-images/postgres:16.6-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.6-bookworm $IMAGE_REGISTRY/appscode-images/postgres:16.6-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.2-alpine $IMAGE_REGISTRY/appscode-images/postgres:17.2-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.2-bookworm $IMAGE_REGISTRY/appscode-images/postgres:17.2-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/pg-coordinator:v0.36.0-rc.0 $IMAGE_REGISTRY/kubedb/pg-coordinator:v0.36.0-rc.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_11.22-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_11.22-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_11.22-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_11.22-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_12.17-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_12.17-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_12.17-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_12.17-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_13.13-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_13.13-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_13.13-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_13.13-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_14.10-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_14.10-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_14.10-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_14.10-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_15.5-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_15.5-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_15.5-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_15.5-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_16.1-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_16.1-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_16.1-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_16.1-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_17.2-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_17.2-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.13.0-rc.0_17.2-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.13.0-rc.0_17.2-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-init:0.17.1 $IMAGE_REGISTRY/kubedb/postgres-init:0.17.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-init:0.17.1 $IMAGE_REGISTRY/kubedb/postgres-init:0.17.1
$CMD cp --allow-nondistributable-artifacts --insecure postgis/postgis:11-3.3 $IMAGE_REGISTRY/postgis/postgis:11-3.3
$CMD cp --allow-nondistributable-artifacts --insecure postgis/postgis:12-3.4 $IMAGE_REGISTRY/postgis/postgis:12-3.4
$CMD cp --allow-nondistributable-artifacts --insecure postgis/postgis:13-3.4 $IMAGE_REGISTRY/postgis/postgis:13-3.4
$CMD cp --allow-nondistributable-artifacts --insecure postgis/postgis:14-3.4 $IMAGE_REGISTRY/postgis/postgis:14-3.4
$CMD cp --allow-nondistributable-artifacts --insecure postgis/postgis:15-3.4 $IMAGE_REGISTRY/postgis/postgis:15-3.4
$CMD cp --allow-nondistributable-artifacts --insecure postgis/postgis:16-3.4 $IMAGE_REGISTRY/postgis/postgis:16-3.4
$CMD cp --allow-nondistributable-artifacts --insecure prometheuscommunity/postgres-exporter:v0.15.0 $IMAGE_REGISTRY/prometheuscommunity/postgres-exporter:v0.15.0
$CMD cp --allow-nondistributable-artifacts --insecure timescale/timescaledb:2.14.2-pg13-oss $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg13-oss
$CMD cp --allow-nondistributable-artifacts --insecure timescale/timescaledb:2.14.2-pg14-oss $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg14-oss
$CMD cp --allow-nondistributable-artifacts --insecure timescale/timescaledb:2.14.2-pg15-oss $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg15-oss
$CMD cp --allow-nondistributable-artifacts --insecure timescale/timescaledb:2.14.2-pg16-oss $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg16-oss
