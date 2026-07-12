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

$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:10.16 $IMAGE_REGISTRY/postgres:10.16
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:10.16-alpine $IMAGE_REGISTRY/postgres:10.16-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:10.19-bullseye $IMAGE_REGISTRY/postgres:10.19-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:10.20-bullseye $IMAGE_REGISTRY/postgres:10.20-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.11 $IMAGE_REGISTRY/postgres:11.11
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.11-alpine $IMAGE_REGISTRY/postgres:11.11-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.14-alpine $IMAGE_REGISTRY/postgres:11.14-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.14-bullseye $IMAGE_REGISTRY/postgres:11.14-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.15-alpine $IMAGE_REGISTRY/postgres:11.15-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.15-bullseye $IMAGE_REGISTRY/postgres:11.15-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.19-alpine $IMAGE_REGISTRY/postgres:11.19-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.19-bullseye $IMAGE_REGISTRY/postgres:11.19-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.20-alpine $IMAGE_REGISTRY/postgres:11.20-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.20-bullseye $IMAGE_REGISTRY/postgres:11.20-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.10-alpine $IMAGE_REGISTRY/postgres:12.10-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.10-bullseye $IMAGE_REGISTRY/postgres:12.10-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.13-alpine $IMAGE_REGISTRY/postgres:12.13-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.13-bullseye $IMAGE_REGISTRY/postgres:12.13-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.14-alpine $IMAGE_REGISTRY/postgres:12.14-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.14-bullseye $IMAGE_REGISTRY/postgres:12.14-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.15-alpine $IMAGE_REGISTRY/postgres:12.15-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.15-bullseye $IMAGE_REGISTRY/postgres:12.15-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.6 $IMAGE_REGISTRY/postgres:12.6
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.6-alpine $IMAGE_REGISTRY/postgres:12.6-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.9-alpine $IMAGE_REGISTRY/postgres:12.9-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.9-bullseye $IMAGE_REGISTRY/postgres:12.9-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.10-alpine $IMAGE_REGISTRY/postgres:13.10-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.10-bullseye $IMAGE_REGISTRY/postgres:13.10-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.11-alpine $IMAGE_REGISTRY/postgres:13.11-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.11-bullseye $IMAGE_REGISTRY/postgres:13.11-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.2 $IMAGE_REGISTRY/postgres:13.2
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.2-alpine $IMAGE_REGISTRY/postgres:13.2-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.5-alpine $IMAGE_REGISTRY/postgres:13.5-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.5-bullseye $IMAGE_REGISTRY/postgres:13.5-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.6-alpine $IMAGE_REGISTRY/postgres:13.6-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.6-bullseye $IMAGE_REGISTRY/postgres:13.6-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.9-alpine $IMAGE_REGISTRY/postgres:13.9-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.9-bullseye $IMAGE_REGISTRY/postgres:13.9-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.1-alpine $IMAGE_REGISTRY/postgres:14.1-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.1-bullseye $IMAGE_REGISTRY/postgres:14.1-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.2-alpine $IMAGE_REGISTRY/postgres:14.2-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.2-bullseye $IMAGE_REGISTRY/postgres:14.2-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.6-alpine $IMAGE_REGISTRY/postgres:14.6-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.6-bullseye $IMAGE_REGISTRY/postgres:14.6-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.7-alpine $IMAGE_REGISTRY/postgres:14.7-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.7-bullseye $IMAGE_REGISTRY/postgres:14.7-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.8-alpine $IMAGE_REGISTRY/postgres:14.8-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.8-bullseye $IMAGE_REGISTRY/postgres:14.8-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:15.1-alpine $IMAGE_REGISTRY/postgres:15.1-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:15.1-bullseye $IMAGE_REGISTRY/postgres:15.1-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:15.2-alpine $IMAGE_REGISTRY/postgres:15.2-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:15.2-bullseye $IMAGE_REGISTRY/postgres:15.2-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:15.3-alpine $IMAGE_REGISTRY/postgres:15.3-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:15.3-bullseye $IMAGE_REGISTRY/postgres:15.3-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:9.6.21 $IMAGE_REGISTRY/postgres:9.6.21
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:9.6.21-alpine $IMAGE_REGISTRY/postgres:9.6.21-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:9.6.24-alpine $IMAGE_REGISTRY/postgres:9.6.24-alpine
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/library/postgres:9.6.24-bullseye $IMAGE_REGISTRY/postgres:9.6.24-bullseye
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/postgis/postgis:11-3.3 $IMAGE_REGISTRY/postgis/postgis:11-3.3
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/postgis/postgis:12-3.4 $IMAGE_REGISTRY/postgis/postgis:12-3.4
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/postgis/postgis:13-3.4 $IMAGE_REGISTRY/postgis/postgis:13-3.4
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/postgis/postgis:14-3.4 $IMAGE_REGISTRY/postgis/postgis:14-3.4
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/postgis/postgis:15-3.4 $IMAGE_REGISTRY/postgis/postgis:15-3.4
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/postgis/postgis:16-3.4 $IMAGE_REGISTRY/postgis/postgis:16-3.4
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/prometheuscommunity/postgres-exporter:v0.18.1 $IMAGE_REGISTRY/prometheuscommunity/postgres-exporter:v0.18.1
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/timescale/timescaledb:2.1.0-pg11-oss $IMAGE_REGISTRY/timescale/timescaledb:2.1.0-pg11-oss
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/timescale/timescaledb:2.1.0-pg12-oss $IMAGE_REGISTRY/timescale/timescaledb:2.1.0-pg12-oss
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/timescale/timescaledb:2.14.2-pg13-oss $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg13-oss
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/timescale/timescaledb:2.14.2-pg14-oss $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg14-oss
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/timescale/timescaledb:2.14.2-pg15-oss $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg15-oss
$CMD cp --allow-nondistributable-artifacts --insecure docker.io/timescale/timescaledb:2.14.2-pg16-oss $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg16-oss
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres-documentdb:15-0.102.0-ferretdb-2.0.0 $IMAGE_REGISTRY/appscode-images/postgres-documentdb:15-0.102.0-ferretdb-2.0.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres-documentdb:16-0.102.0-ferretdb-2.0.0 $IMAGE_REGISTRY/appscode-images/postgres-documentdb:16-0.102.0-ferretdb-2.0.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres-documentdb:17-0.102.0-ferretdb-2.0.0 $IMAGE_REGISTRY/appscode-images/postgres-documentdb:17-0.102.0-ferretdb-2.0.0
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
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.20-alpine $IMAGE_REGISTRY/appscode-images/postgres:13.20-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.20-bookworm $IMAGE_REGISTRY/appscode-images/postgres:13.20-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.21-alpine $IMAGE_REGISTRY/appscode-images/postgres:13.21-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.21-bookworm $IMAGE_REGISTRY/appscode-images/postgres:13.21-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.10-alpine $IMAGE_REGISTRY/appscode-images/postgres:14.10-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.10-bookworm $IMAGE_REGISTRY/appscode-images/postgres:14.10-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.13-alpine $IMAGE_REGISTRY/appscode-images/postgres:14.13-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.13-bookworm $IMAGE_REGISTRY/appscode-images/postgres:14.13-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.15-alpine $IMAGE_REGISTRY/appscode-images/postgres:14.15-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.15-bookworm $IMAGE_REGISTRY/appscode-images/postgres:14.15-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.17-alpine $IMAGE_REGISTRY/appscode-images/postgres:14.17-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.17-bookworm $IMAGE_REGISTRY/appscode-images/postgres:14.17-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.18-alpine $IMAGE_REGISTRY/appscode-images/postgres:14.18-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.18-bookworm $IMAGE_REGISTRY/appscode-images/postgres:14.18-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.21-alpine $IMAGE_REGISTRY/appscode-images/postgres:14.21-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.21-bookworm $IMAGE_REGISTRY/appscode-images/postgres:14.21-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.22-alpine $IMAGE_REGISTRY/appscode-images/postgres:14.22-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.22-bookworm $IMAGE_REGISTRY/appscode-images/postgres:14.22-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.10-alpine $IMAGE_REGISTRY/appscode-images/postgres:15.10-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.10-bookworm $IMAGE_REGISTRY/appscode-images/postgres:15.10-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.12-alpine $IMAGE_REGISTRY/appscode-images/postgres:15.12-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.12-bookworm $IMAGE_REGISTRY/appscode-images/postgres:15.12-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.13-alpine $IMAGE_REGISTRY/appscode-images/postgres:15.13-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.13-bookworm $IMAGE_REGISTRY/appscode-images/postgres:15.13-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.16-alpine $IMAGE_REGISTRY/appscode-images/postgres:15.16-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.16-bookworm $IMAGE_REGISTRY/appscode-images/postgres:15.16-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.17-alpine $IMAGE_REGISTRY/appscode-images/postgres:15.17-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.17-bookworm $IMAGE_REGISTRY/appscode-images/postgres:15.17-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.5-alpine $IMAGE_REGISTRY/appscode-images/postgres:15.5-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.5-bookworm $IMAGE_REGISTRY/appscode-images/postgres:15.5-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.8-alpine $IMAGE_REGISTRY/appscode-images/postgres:15.8-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.8-bookworm $IMAGE_REGISTRY/appscode-images/postgres:15.8-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.1-alpine $IMAGE_REGISTRY/appscode-images/postgres:16.1-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.1-bookworm $IMAGE_REGISTRY/appscode-images/postgres:16.1-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.10-alpine $IMAGE_REGISTRY/appscode-images/postgres:16.10-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.10-bookworm $IMAGE_REGISTRY/appscode-images/postgres:16.10-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.12-alpine $IMAGE_REGISTRY/appscode-images/postgres:16.12-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.12-bookworm $IMAGE_REGISTRY/appscode-images/postgres:16.12-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.13-alpine $IMAGE_REGISTRY/appscode-images/postgres:16.13-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.13-alpine-ext $IMAGE_REGISTRY/appscode-images/postgres:16.13-alpine-ext
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.13-bookworm $IMAGE_REGISTRY/appscode-images/postgres:16.13-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.13-bookworm-ext $IMAGE_REGISTRY/appscode-images/postgres:16.13-bookworm-ext
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.4-alpine $IMAGE_REGISTRY/appscode-images/postgres:16.4-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.4-bookworm $IMAGE_REGISTRY/appscode-images/postgres:16.4-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.6-alpine $IMAGE_REGISTRY/appscode-images/postgres:16.6-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.6-bookworm $IMAGE_REGISTRY/appscode-images/postgres:16.6-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.8-alpine $IMAGE_REGISTRY/appscode-images/postgres:16.8-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.8-bookworm $IMAGE_REGISTRY/appscode-images/postgres:16.8-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.9-alpine $IMAGE_REGISTRY/appscode-images/postgres:16.9-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.9-bookworm $IMAGE_REGISTRY/appscode-images/postgres:16.9-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.2-alpine $IMAGE_REGISTRY/appscode-images/postgres:17.2-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.2-bookworm $IMAGE_REGISTRY/appscode-images/postgres:17.2-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.4-alpine $IMAGE_REGISTRY/appscode-images/postgres:17.4-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.4-bookworm $IMAGE_REGISTRY/appscode-images/postgres:17.4-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.5-alpine $IMAGE_REGISTRY/appscode-images/postgres:17.5-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.5-bookworm $IMAGE_REGISTRY/appscode-images/postgres:17.5-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.8-alpine $IMAGE_REGISTRY/appscode-images/postgres:17.8-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.8-bookworm $IMAGE_REGISTRY/appscode-images/postgres:17.8-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.9-alpine $IMAGE_REGISTRY/appscode-images/postgres:17.9-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.9-alpine-ext $IMAGE_REGISTRY/appscode-images/postgres:17.9-alpine-ext
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.9-bookworm $IMAGE_REGISTRY/appscode-images/postgres:17.9-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.9-bookworm-ext $IMAGE_REGISTRY/appscode-images/postgres:17.9-bookworm-ext
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:18.2-alpine $IMAGE_REGISTRY/appscode-images/postgres:18.2-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:18.2-bookworm $IMAGE_REGISTRY/appscode-images/postgres:18.2-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:18.3-alpine $IMAGE_REGISTRY/appscode-images/postgres:18.3-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:18.3-alpine-ext $IMAGE_REGISTRY/appscode-images/postgres:18.3-alpine-ext
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:18.3-bookworm $IMAGE_REGISTRY/appscode-images/postgres:18.3-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:18.3-bookworm-ext $IMAGE_REGISTRY/appscode-images/postgres:18.3-bookworm-ext
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-migrator-postgresql:v0.6.0 $IMAGE_REGISTRY/kubedb/kubedb-migrator-postgresql:v0.6.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/pg-coordinator:v0.50.0 $IMAGE_REGISTRY/kubedb/pg-coordinator:v0.50.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_11.22-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_11.22-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_11.22-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_11.22-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_12.17-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_12.17-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_12.17-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_12.17-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_13.13-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_13.13-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_13.13-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_13.13-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_14.10-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_14.10-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_14.10-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_14.10-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_15.5-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_15.5-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_15.5-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_15.5-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_16.4-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_16.4-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_16.4-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_16.4-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_17.2-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_17.2-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_17.2-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_17.2-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_18.2-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_18.2-alpine
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_18.2-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.27.0_18.2-bookworm
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-init:0.19.0 $IMAGE_REGISTRY/kubedb/postgres-init:0.19.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-init:0.20.0 $IMAGE_REGISTRY/kubedb/postgres-init:0.20.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.2 $IMAGE_REGISTRY/kubedb/postgres:10.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.2-v2 $IMAGE_REGISTRY/kubedb/postgres:10.2-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.2-v3 $IMAGE_REGISTRY/kubedb/postgres:10.2-v3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.2-v4 $IMAGE_REGISTRY/kubedb/postgres:10.2-v4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.2-v5 $IMAGE_REGISTRY/kubedb/postgres:10.2-v5
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.2-v6 $IMAGE_REGISTRY/kubedb/postgres:10.2-v6
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.6 $IMAGE_REGISTRY/kubedb/postgres:10.6
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.6-v1 $IMAGE_REGISTRY/kubedb/postgres:10.6-v1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.6-v2 $IMAGE_REGISTRY/kubedb/postgres:10.6-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.6-v3 $IMAGE_REGISTRY/kubedb/postgres:10.6-v3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:11.1 $IMAGE_REGISTRY/kubedb/postgres:11.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:11.1-v1 $IMAGE_REGISTRY/kubedb/postgres:11.1-v1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:11.1-v2 $IMAGE_REGISTRY/kubedb/postgres:11.1-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:11.1-v3 $IMAGE_REGISTRY/kubedb/postgres:11.1-v3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:11.2 $IMAGE_REGISTRY/kubedb/postgres:11.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:11.2-v1 $IMAGE_REGISTRY/kubedb/postgres:11.2-v1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6 $IMAGE_REGISTRY/kubedb/postgres:9.6
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6-v2 $IMAGE_REGISTRY/kubedb/postgres:9.6-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6-v3 $IMAGE_REGISTRY/kubedb/postgres:9.6-v3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6-v4 $IMAGE_REGISTRY/kubedb/postgres:9.6-v4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6-v5 $IMAGE_REGISTRY/kubedb/postgres:9.6-v5
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6-v6 $IMAGE_REGISTRY/kubedb/postgres:9.6-v6
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6.7 $IMAGE_REGISTRY/kubedb/postgres:9.6.7
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6.7-v2 $IMAGE_REGISTRY/kubedb/postgres:9.6.7-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6.7-v3 $IMAGE_REGISTRY/kubedb/postgres:9.6.7-v3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6.7-v4 $IMAGE_REGISTRY/kubedb/postgres:9.6.7-v4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6.7-v5 $IMAGE_REGISTRY/kubedb/postgres:9.6.7-v5
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6.7-v6 $IMAGE_REGISTRY/kubedb/postgres:9.6.7-v6
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres_exporter:v0.4.6 $IMAGE_REGISTRY/kubedb/postgres_exporter:v0.4.6
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres_exporter:v0.4.7 $IMAGE_REGISTRY/kubedb/postgres_exporter:v0.4.7
