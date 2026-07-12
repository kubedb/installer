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

k3s ctr images import images/library-postgres-10.16.tar
k3s ctr images import images/library-postgres-10.16-alpine.tar
k3s ctr images import images/library-postgres-10.19-bullseye.tar
k3s ctr images import images/library-postgres-10.20-bullseye.tar
k3s ctr images import images/library-postgres-11.11.tar
k3s ctr images import images/library-postgres-11.11-alpine.tar
k3s ctr images import images/library-postgres-11.14-alpine.tar
k3s ctr images import images/library-postgres-11.14-bullseye.tar
k3s ctr images import images/library-postgres-11.15-alpine.tar
k3s ctr images import images/library-postgres-11.15-bullseye.tar
k3s ctr images import images/library-postgres-11.19-alpine.tar
k3s ctr images import images/library-postgres-11.19-bullseye.tar
k3s ctr images import images/library-postgres-11.20-alpine.tar
k3s ctr images import images/library-postgres-11.20-bullseye.tar
k3s ctr images import images/library-postgres-12.10-alpine.tar
k3s ctr images import images/library-postgres-12.10-bullseye.tar
k3s ctr images import images/library-postgres-12.13-alpine.tar
k3s ctr images import images/library-postgres-12.13-bullseye.tar
k3s ctr images import images/library-postgres-12.14-alpine.tar
k3s ctr images import images/library-postgres-12.14-bullseye.tar
k3s ctr images import images/library-postgres-12.15-alpine.tar
k3s ctr images import images/library-postgres-12.15-bullseye.tar
k3s ctr images import images/library-postgres-12.6.tar
k3s ctr images import images/library-postgres-12.6-alpine.tar
k3s ctr images import images/library-postgres-12.9-alpine.tar
k3s ctr images import images/library-postgres-12.9-bullseye.tar
k3s ctr images import images/library-postgres-13.10-alpine.tar
k3s ctr images import images/library-postgres-13.10-bullseye.tar
k3s ctr images import images/library-postgres-13.11-alpine.tar
k3s ctr images import images/library-postgres-13.11-bullseye.tar
k3s ctr images import images/library-postgres-13.2.tar
k3s ctr images import images/library-postgres-13.2-alpine.tar
k3s ctr images import images/library-postgres-13.5-alpine.tar
k3s ctr images import images/library-postgres-13.5-bullseye.tar
k3s ctr images import images/library-postgres-13.6-alpine.tar
k3s ctr images import images/library-postgres-13.6-bullseye.tar
k3s ctr images import images/library-postgres-13.9-alpine.tar
k3s ctr images import images/library-postgres-13.9-bullseye.tar
k3s ctr images import images/library-postgres-14.1-alpine.tar
k3s ctr images import images/library-postgres-14.1-bullseye.tar
k3s ctr images import images/library-postgres-14.2-alpine.tar
k3s ctr images import images/library-postgres-14.2-bullseye.tar
k3s ctr images import images/library-postgres-14.6-alpine.tar
k3s ctr images import images/library-postgres-14.6-bullseye.tar
k3s ctr images import images/library-postgres-14.7-alpine.tar
k3s ctr images import images/library-postgres-14.7-bullseye.tar
k3s ctr images import images/library-postgres-14.8-alpine.tar
k3s ctr images import images/library-postgres-14.8-bullseye.tar
k3s ctr images import images/library-postgres-15.1-alpine.tar
k3s ctr images import images/library-postgres-15.1-bullseye.tar
k3s ctr images import images/library-postgres-15.2-alpine.tar
k3s ctr images import images/library-postgres-15.2-bullseye.tar
k3s ctr images import images/library-postgres-15.3-alpine.tar
k3s ctr images import images/library-postgres-15.3-bullseye.tar
k3s ctr images import images/library-postgres-9.6.21.tar
k3s ctr images import images/library-postgres-9.6.21-alpine.tar
k3s ctr images import images/library-postgres-9.6.24-alpine.tar
k3s ctr images import images/library-postgres-9.6.24-bullseye.tar
k3s ctr images import images/postgis-postgis-11-3.3.tar
k3s ctr images import images/postgis-postgis-12-3.4.tar
k3s ctr images import images/postgis-postgis-13-3.4.tar
k3s ctr images import images/postgis-postgis-14-3.4.tar
k3s ctr images import images/postgis-postgis-15-3.4.tar
k3s ctr images import images/postgis-postgis-16-3.4.tar
k3s ctr images import images/prometheuscommunity-postgres-exporter-v0.18.1.tar
k3s ctr images import images/timescale-timescaledb-2.1.0-pg11-oss.tar
k3s ctr images import images/timescale-timescaledb-2.1.0-pg12-oss.tar
k3s ctr images import images/timescale-timescaledb-2.14.2-pg13-oss.tar
k3s ctr images import images/timescale-timescaledb-2.14.2-pg14-oss.tar
k3s ctr images import images/timescale-timescaledb-2.14.2-pg15-oss.tar
k3s ctr images import images/timescale-timescaledb-2.14.2-pg16-oss.tar
k3s ctr images import images/appscode-images-postgres-documentdb-15-0.102.0-ferretdb-2.0.0.tar
k3s ctr images import images/appscode-images-postgres-documentdb-16-0.102.0-ferretdb-2.0.0.tar
k3s ctr images import images/appscode-images-postgres-documentdb-17-0.102.0-ferretdb-2.0.0.tar
k3s ctr images import images/appscode-images-postgres-10.23-alpine.tar
k3s ctr images import images/appscode-images-postgres-10.23-bullseye.tar
k3s ctr images import images/appscode-images-postgres-11.22-alpine.tar
k3s ctr images import images/appscode-images-postgres-11.22-bookworm.tar
k3s ctr images import images/appscode-images-postgres-12.17-alpine.tar
k3s ctr images import images/appscode-images-postgres-12.17-bookworm.tar
k3s ctr images import images/appscode-images-postgres-12.22-alpine.tar
k3s ctr images import images/appscode-images-postgres-12.22-bookworm.tar
k3s ctr images import images/appscode-images-postgres-13.13-alpine.tar
k3s ctr images import images/appscode-images-postgres-13.13-bookworm.tar
k3s ctr images import images/appscode-images-postgres-13.18-alpine.tar
k3s ctr images import images/appscode-images-postgres-13.18-bookworm.tar
k3s ctr images import images/appscode-images-postgres-13.20-alpine.tar
k3s ctr images import images/appscode-images-postgres-13.20-bookworm.tar
k3s ctr images import images/appscode-images-postgres-13.21-alpine.tar
k3s ctr images import images/appscode-images-postgres-13.21-bookworm.tar
k3s ctr images import images/appscode-images-postgres-14.10-alpine.tar
k3s ctr images import images/appscode-images-postgres-14.10-bookworm.tar
k3s ctr images import images/appscode-images-postgres-14.13-alpine.tar
k3s ctr images import images/appscode-images-postgres-14.13-bookworm.tar
k3s ctr images import images/appscode-images-postgres-14.15-alpine.tar
k3s ctr images import images/appscode-images-postgres-14.15-bookworm.tar
k3s ctr images import images/appscode-images-postgres-14.17-alpine.tar
k3s ctr images import images/appscode-images-postgres-14.17-bookworm.tar
k3s ctr images import images/appscode-images-postgres-14.18-alpine.tar
k3s ctr images import images/appscode-images-postgres-14.18-bookworm.tar
k3s ctr images import images/appscode-images-postgres-14.21-alpine.tar
k3s ctr images import images/appscode-images-postgres-14.21-bookworm.tar
k3s ctr images import images/appscode-images-postgres-14.22-alpine.tar
k3s ctr images import images/appscode-images-postgres-14.22-bookworm.tar
k3s ctr images import images/appscode-images-postgres-15.10-alpine.tar
k3s ctr images import images/appscode-images-postgres-15.10-bookworm.tar
k3s ctr images import images/appscode-images-postgres-15.12-alpine.tar
k3s ctr images import images/appscode-images-postgres-15.12-bookworm.tar
k3s ctr images import images/appscode-images-postgres-15.13-alpine.tar
k3s ctr images import images/appscode-images-postgres-15.13-bookworm.tar
k3s ctr images import images/appscode-images-postgres-15.16-alpine.tar
k3s ctr images import images/appscode-images-postgres-15.16-bookworm.tar
k3s ctr images import images/appscode-images-postgres-15.17-alpine.tar
k3s ctr images import images/appscode-images-postgres-15.17-bookworm.tar
k3s ctr images import images/appscode-images-postgres-15.5-alpine.tar
k3s ctr images import images/appscode-images-postgres-15.5-bookworm.tar
k3s ctr images import images/appscode-images-postgres-15.8-alpine.tar
k3s ctr images import images/appscode-images-postgres-15.8-bookworm.tar
k3s ctr images import images/appscode-images-postgres-16.1-alpine.tar
k3s ctr images import images/appscode-images-postgres-16.1-bookworm.tar
k3s ctr images import images/appscode-images-postgres-16.10-alpine.tar
k3s ctr images import images/appscode-images-postgres-16.10-bookworm.tar
k3s ctr images import images/appscode-images-postgres-16.12-alpine.tar
k3s ctr images import images/appscode-images-postgres-16.12-bookworm.tar
k3s ctr images import images/appscode-images-postgres-16.13-alpine.tar
k3s ctr images import images/appscode-images-postgres-16.13-alpine-ext.tar
k3s ctr images import images/appscode-images-postgres-16.13-bookworm.tar
k3s ctr images import images/appscode-images-postgres-16.13-bookworm-ext.tar
k3s ctr images import images/appscode-images-postgres-16.4-alpine.tar
k3s ctr images import images/appscode-images-postgres-16.4-bookworm.tar
k3s ctr images import images/appscode-images-postgres-16.6-alpine.tar
k3s ctr images import images/appscode-images-postgres-16.6-bookworm.tar
k3s ctr images import images/appscode-images-postgres-16.8-alpine.tar
k3s ctr images import images/appscode-images-postgres-16.8-bookworm.tar
k3s ctr images import images/appscode-images-postgres-16.9-alpine.tar
k3s ctr images import images/appscode-images-postgres-16.9-bookworm.tar
k3s ctr images import images/appscode-images-postgres-17.2-alpine.tar
k3s ctr images import images/appscode-images-postgres-17.2-bookworm.tar
k3s ctr images import images/appscode-images-postgres-17.4-alpine.tar
k3s ctr images import images/appscode-images-postgres-17.4-bookworm.tar
k3s ctr images import images/appscode-images-postgres-17.5-alpine.tar
k3s ctr images import images/appscode-images-postgres-17.5-bookworm.tar
k3s ctr images import images/appscode-images-postgres-17.8-alpine.tar
k3s ctr images import images/appscode-images-postgres-17.8-bookworm.tar
k3s ctr images import images/appscode-images-postgres-17.9-alpine.tar
k3s ctr images import images/appscode-images-postgres-17.9-alpine-ext.tar
k3s ctr images import images/appscode-images-postgres-17.9-bookworm.tar
k3s ctr images import images/appscode-images-postgres-17.9-bookworm-ext.tar
k3s ctr images import images/appscode-images-postgres-18.2-alpine.tar
k3s ctr images import images/appscode-images-postgres-18.2-bookworm.tar
k3s ctr images import images/appscode-images-postgres-18.3-alpine.tar
k3s ctr images import images/appscode-images-postgres-18.3-alpine-ext.tar
k3s ctr images import images/appscode-images-postgres-18.3-bookworm.tar
k3s ctr images import images/appscode-images-postgres-18.3-bookworm-ext.tar
k3s ctr images import images/kubedb-kubedb-migrator-postgresql-v0.6.0.tar
k3s ctr images import images/kubedb-pg-coordinator-v0.50.0.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_11.22-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_11.22-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_12.17-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_12.17-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_13.13-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_13.13-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_14.10-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_14.10-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_15.5-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_15.5-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_16.4-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_16.4-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_17.2-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_17.2-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_18.2-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.27.0_18.2-bookworm.tar
k3s ctr images import images/kubedb-postgres-init-0.19.0.tar
k3s ctr images import images/kubedb-postgres-init-0.20.0.tar
k3s ctr images import images/kubedb-postgres-10.2.tar
k3s ctr images import images/kubedb-postgres-10.2-v2.tar
k3s ctr images import images/kubedb-postgres-10.2-v3.tar
k3s ctr images import images/kubedb-postgres-10.2-v4.tar
k3s ctr images import images/kubedb-postgres-10.2-v5.tar
k3s ctr images import images/kubedb-postgres-10.2-v6.tar
k3s ctr images import images/kubedb-postgres-10.6.tar
k3s ctr images import images/kubedb-postgres-10.6-v1.tar
k3s ctr images import images/kubedb-postgres-10.6-v2.tar
k3s ctr images import images/kubedb-postgres-10.6-v3.tar
k3s ctr images import images/kubedb-postgres-11.1.tar
k3s ctr images import images/kubedb-postgres-11.1-v1.tar
k3s ctr images import images/kubedb-postgres-11.1-v2.tar
k3s ctr images import images/kubedb-postgres-11.1-v3.tar
k3s ctr images import images/kubedb-postgres-11.2.tar
k3s ctr images import images/kubedb-postgres-11.2-v1.tar
k3s ctr images import images/kubedb-postgres-9.6.tar
k3s ctr images import images/kubedb-postgres-9.6-v2.tar
k3s ctr images import images/kubedb-postgres-9.6-v3.tar
k3s ctr images import images/kubedb-postgres-9.6-v4.tar
k3s ctr images import images/kubedb-postgres-9.6-v5.tar
k3s ctr images import images/kubedb-postgres-9.6-v6.tar
k3s ctr images import images/kubedb-postgres-9.6.7.tar
k3s ctr images import images/kubedb-postgres-9.6.7-v2.tar
k3s ctr images import images/kubedb-postgres-9.6.7-v3.tar
k3s ctr images import images/kubedb-postgres-9.6.7-v4.tar
k3s ctr images import images/kubedb-postgres-9.6.7-v5.tar
k3s ctr images import images/kubedb-postgres-9.6.7-v6.tar
k3s ctr images import images/kubedb-postgres_exporter-v0.4.6.tar
k3s ctr images import images/kubedb-postgres_exporter-v0.4.7.tar
