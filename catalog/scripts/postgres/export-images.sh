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

$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:10.16 images/library-postgres-10.16.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:10.16-alpine images/library-postgres-10.16-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:10.19-bullseye images/library-postgres-10.19-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:10.20-bullseye images/library-postgres-10.20-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.11 images/library-postgres-11.11.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.11-alpine images/library-postgres-11.11-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.14-alpine images/library-postgres-11.14-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.14-bullseye images/library-postgres-11.14-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.15-alpine images/library-postgres-11.15-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.15-bullseye images/library-postgres-11.15-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.19-alpine images/library-postgres-11.19-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.19-bullseye images/library-postgres-11.19-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.20-alpine images/library-postgres-11.20-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:11.20-bullseye images/library-postgres-11.20-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.10-alpine images/library-postgres-12.10-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.10-bullseye images/library-postgres-12.10-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.13-alpine images/library-postgres-12.13-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.13-bullseye images/library-postgres-12.13-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.14-alpine images/library-postgres-12.14-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.14-bullseye images/library-postgres-12.14-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.15-alpine images/library-postgres-12.15-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.15-bullseye images/library-postgres-12.15-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.6 images/library-postgres-12.6.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.6-alpine images/library-postgres-12.6-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.9-alpine images/library-postgres-12.9-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:12.9-bullseye images/library-postgres-12.9-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.10-alpine images/library-postgres-13.10-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.10-bullseye images/library-postgres-13.10-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.11-alpine images/library-postgres-13.11-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.11-bullseye images/library-postgres-13.11-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.2 images/library-postgres-13.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.2-alpine images/library-postgres-13.2-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.5-alpine images/library-postgres-13.5-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.5-bullseye images/library-postgres-13.5-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.6-alpine images/library-postgres-13.6-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.6-bullseye images/library-postgres-13.6-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.9-alpine images/library-postgres-13.9-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:13.9-bullseye images/library-postgres-13.9-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.1-alpine images/library-postgres-14.1-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.1-bullseye images/library-postgres-14.1-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.2-alpine images/library-postgres-14.2-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.2-bullseye images/library-postgres-14.2-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.6-alpine images/library-postgres-14.6-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.6-bullseye images/library-postgres-14.6-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.7-alpine images/library-postgres-14.7-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.7-bullseye images/library-postgres-14.7-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.8-alpine images/library-postgres-14.8-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:14.8-bullseye images/library-postgres-14.8-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:15.1-alpine images/library-postgres-15.1-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:15.1-bullseye images/library-postgres-15.1-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:15.2-alpine images/library-postgres-15.2-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:15.2-bullseye images/library-postgres-15.2-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:15.3-alpine images/library-postgres-15.3-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:15.3-bullseye images/library-postgres-15.3-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:9.6.21 images/library-postgres-9.6.21.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:9.6.21-alpine images/library-postgres-9.6.21-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:9.6.24-alpine images/library-postgres-9.6.24-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/postgres:9.6.24-bullseye images/library-postgres-9.6.24-bullseye.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/postgis/postgis:11-3.3 images/postgis-postgis-11-3.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/postgis/postgis:12-3.4 images/postgis-postgis-12-3.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/postgis/postgis:13-3.4 images/postgis-postgis-13-3.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/postgis/postgis:14-3.4 images/postgis-postgis-14-3.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/postgis/postgis:15-3.4 images/postgis-postgis-15-3.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/postgis/postgis:16-3.4 images/postgis-postgis-16-3.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/prometheuscommunity/postgres-exporter:v0.18.1 images/prometheuscommunity-postgres-exporter-v0.18.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/timescale/timescaledb:2.1.0-pg11-oss images/timescale-timescaledb-2.1.0-pg11-oss.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/timescale/timescaledb:2.1.0-pg12-oss images/timescale-timescaledb-2.1.0-pg12-oss.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/timescale/timescaledb:2.14.2-pg13-oss images/timescale-timescaledb-2.14.2-pg13-oss.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/timescale/timescaledb:2.14.2-pg14-oss images/timescale-timescaledb-2.14.2-pg14-oss.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/timescale/timescaledb:2.14.2-pg15-oss images/timescale-timescaledb-2.14.2-pg15-oss.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/timescale/timescaledb:2.14.2-pg16-oss images/timescale-timescaledb-2.14.2-pg16-oss.tar
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
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.21-alpine images/appscode-images-postgres-13.21-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:13.21-bookworm images/appscode-images-postgres-13.21-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.10-alpine images/appscode-images-postgres-14.10-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.10-bookworm images/appscode-images-postgres-14.10-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.13-alpine images/appscode-images-postgres-14.13-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.13-bookworm images/appscode-images-postgres-14.13-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.15-alpine images/appscode-images-postgres-14.15-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.15-bookworm images/appscode-images-postgres-14.15-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.17-alpine images/appscode-images-postgres-14.17-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.17-bookworm images/appscode-images-postgres-14.17-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.18-alpine images/appscode-images-postgres-14.18-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.18-bookworm images/appscode-images-postgres-14.18-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.21-alpine images/appscode-images-postgres-14.21-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.21-bookworm images/appscode-images-postgres-14.21-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.22-alpine images/appscode-images-postgres-14.22-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:14.22-bookworm images/appscode-images-postgres-14.22-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.10-alpine images/appscode-images-postgres-15.10-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.10-bookworm images/appscode-images-postgres-15.10-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.12-alpine images/appscode-images-postgres-15.12-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.12-bookworm images/appscode-images-postgres-15.12-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.13-alpine images/appscode-images-postgres-15.13-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.13-bookworm images/appscode-images-postgres-15.13-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.16-alpine images/appscode-images-postgres-15.16-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.16-bookworm images/appscode-images-postgres-15.16-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.17-alpine images/appscode-images-postgres-15.17-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.17-bookworm images/appscode-images-postgres-15.17-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.5-alpine images/appscode-images-postgres-15.5-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.5-bookworm images/appscode-images-postgres-15.5-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.8-alpine images/appscode-images-postgres-15.8-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:15.8-bookworm images/appscode-images-postgres-15.8-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.1-alpine images/appscode-images-postgres-16.1-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.1-bookworm images/appscode-images-postgres-16.1-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.10-alpine images/appscode-images-postgres-16.10-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.10-bookworm images/appscode-images-postgres-16.10-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.12-alpine images/appscode-images-postgres-16.12-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.12-bookworm images/appscode-images-postgres-16.12-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.13-alpine images/appscode-images-postgres-16.13-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.13-alpine-ext images/appscode-images-postgres-16.13-alpine-ext.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.13-bookworm images/appscode-images-postgres-16.13-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.13-bookworm-ext images/appscode-images-postgres-16.13-bookworm-ext.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.4-alpine images/appscode-images-postgres-16.4-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.4-bookworm images/appscode-images-postgres-16.4-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.6-alpine images/appscode-images-postgres-16.6-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.6-bookworm images/appscode-images-postgres-16.6-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.8-alpine images/appscode-images-postgres-16.8-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.8-bookworm images/appscode-images-postgres-16.8-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.9-alpine images/appscode-images-postgres-16.9-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:16.9-bookworm images/appscode-images-postgres-16.9-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.2-alpine images/appscode-images-postgres-17.2-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.2-bookworm images/appscode-images-postgres-17.2-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.4-alpine images/appscode-images-postgres-17.4-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.4-bookworm images/appscode-images-postgres-17.4-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.5-alpine images/appscode-images-postgres-17.5-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.5-bookworm images/appscode-images-postgres-17.5-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.8-alpine images/appscode-images-postgres-17.8-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.8-bookworm images/appscode-images-postgres-17.8-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.9-alpine images/appscode-images-postgres-17.9-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.9-alpine-ext images/appscode-images-postgres-17.9-alpine-ext.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.9-bookworm images/appscode-images-postgres-17.9-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:17.9-bookworm-ext images/appscode-images-postgres-17.9-bookworm-ext.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:18.2-alpine images/appscode-images-postgres-18.2-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:18.2-bookworm images/appscode-images-postgres-18.2-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:18.3-alpine images/appscode-images-postgres-18.3-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:18.3-alpine-ext images/appscode-images-postgres-18.3-alpine-ext.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:18.3-bookworm images/appscode-images-postgres-18.3-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/postgres:18.3-bookworm-ext images/appscode-images-postgres-18.3-bookworm-ext.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kubedb-migrator-postgresql:v0.6.0 images/kubedb-kubedb-migrator-postgresql-v0.6.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/pg-coordinator:v0.50.0 images/kubedb-pg-coordinator-v0.50.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_11.22-alpine images/kubedb-postgres-archiver-v0.27.0_11.22-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_11.22-bookworm images/kubedb-postgres-archiver-v0.27.0_11.22-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_12.17-alpine images/kubedb-postgres-archiver-v0.27.0_12.17-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_12.17-bookworm images/kubedb-postgres-archiver-v0.27.0_12.17-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_13.13-alpine images/kubedb-postgres-archiver-v0.27.0_13.13-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_13.13-bookworm images/kubedb-postgres-archiver-v0.27.0_13.13-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_14.10-alpine images/kubedb-postgres-archiver-v0.27.0_14.10-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_14.10-bookworm images/kubedb-postgres-archiver-v0.27.0_14.10-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_15.5-alpine images/kubedb-postgres-archiver-v0.27.0_15.5-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_15.5-bookworm images/kubedb-postgres-archiver-v0.27.0_15.5-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_16.4-alpine images/kubedb-postgres-archiver-v0.27.0_16.4-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_16.4-bookworm images/kubedb-postgres-archiver-v0.27.0_16.4-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_17.2-alpine images/kubedb-postgres-archiver-v0.27.0_17.2-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_17.2-bookworm images/kubedb-postgres-archiver-v0.27.0_17.2-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_18.2-alpine images/kubedb-postgres-archiver-v0.27.0_18.2-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-archiver:v0.27.0_18.2-bookworm images/kubedb-postgres-archiver-v0.27.0_18.2-bookworm.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-init:0.19.0 images/kubedb-postgres-init-0.19.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres-init:0.20.0 images/kubedb-postgres-init-0.20.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.2 images/kubedb-postgres-10.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.2-v2 images/kubedb-postgres-10.2-v2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.2-v3 images/kubedb-postgres-10.2-v3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.2-v4 images/kubedb-postgres-10.2-v4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.2-v5 images/kubedb-postgres-10.2-v5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.2-v6 images/kubedb-postgres-10.2-v6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.6 images/kubedb-postgres-10.6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.6-v1 images/kubedb-postgres-10.6-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.6-v2 images/kubedb-postgres-10.6-v2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:10.6-v3 images/kubedb-postgres-10.6-v3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:11.1 images/kubedb-postgres-11.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:11.1-v1 images/kubedb-postgres-11.1-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:11.1-v2 images/kubedb-postgres-11.1-v2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:11.1-v3 images/kubedb-postgres-11.1-v3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:11.2 images/kubedb-postgres-11.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:11.2-v1 images/kubedb-postgres-11.2-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6 images/kubedb-postgres-9.6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6-v2 images/kubedb-postgres-9.6-v2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6-v3 images/kubedb-postgres-9.6-v3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6-v4 images/kubedb-postgres-9.6-v4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6-v5 images/kubedb-postgres-9.6-v5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6-v6 images/kubedb-postgres-9.6-v6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6.7 images/kubedb-postgres-9.6.7.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6.7-v2 images/kubedb-postgres-9.6.7-v2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6.7-v3 images/kubedb-postgres-9.6.7-v3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6.7-v4 images/kubedb-postgres-9.6.7-v4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6.7-v5 images/kubedb-postgres-9.6.7-v5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres:9.6.7-v6 images/kubedb-postgres-9.6.7-v6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres_exporter:v0.4.6 images/kubedb-postgres_exporter-v0.4.6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/postgres_exporter:v0.4.7 images/kubedb-postgres_exporter-v0.4.7.tar

tar -czvf images.tar.gz images
