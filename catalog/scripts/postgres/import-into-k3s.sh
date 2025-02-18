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
k3s ctr images import images/appscode-images-postgres-14.10-alpine.tar
k3s ctr images import images/appscode-images-postgres-14.10-bookworm.tar
k3s ctr images import images/appscode-images-postgres-14.13-alpine.tar
k3s ctr images import images/appscode-images-postgres-14.13-bookworm.tar
k3s ctr images import images/appscode-images-postgres-14.15-alpine.tar
k3s ctr images import images/appscode-images-postgres-14.15-bookworm.tar
k3s ctr images import images/appscode-images-postgres-14.17-alpine.tar
k3s ctr images import images/appscode-images-postgres-14.17-bookworm.tar
k3s ctr images import images/appscode-images-postgres-15.10-alpine.tar
k3s ctr images import images/appscode-images-postgres-15.10-bookworm.tar
k3s ctr images import images/appscode-images-postgres-15.12-alpine.tar
k3s ctr images import images/appscode-images-postgres-15.12-bookworm.tar
k3s ctr images import images/appscode-images-postgres-15.5-alpine.tar
k3s ctr images import images/appscode-images-postgres-15.5-bookworm.tar
k3s ctr images import images/appscode-images-postgres-15.8-alpine.tar
k3s ctr images import images/appscode-images-postgres-15.8-bookworm.tar
k3s ctr images import images/appscode-images-postgres-16.1-alpine.tar
k3s ctr images import images/appscode-images-postgres-16.1-bookworm.tar
k3s ctr images import images/appscode-images-postgres-16.4-alpine.tar
k3s ctr images import images/appscode-images-postgres-16.4-bookworm.tar
k3s ctr images import images/appscode-images-postgres-16.6-alpine.tar
k3s ctr images import images/appscode-images-postgres-16.6-bookworm.tar
k3s ctr images import images/appscode-images-postgres-16.8-alpine.tar
k3s ctr images import images/appscode-images-postgres-16.8-bookworm.tar
k3s ctr images import images/appscode-images-postgres-17.2-alpine.tar
k3s ctr images import images/appscode-images-postgres-17.2-bookworm.tar
k3s ctr images import images/appscode-images-postgres-17.4-alpine.tar
k3s ctr images import images/appscode-images-postgres-17.4-bookworm.tar
k3s ctr images import images/kubedb-pg-coordinator-v0.36.0.tar
k3s ctr images import images/kubedb-pg-coordinator-v0.36.0-rc.0.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0-rc.0_13.13-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0-rc.0_13.13-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0-rc.0_14.10-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0-rc.0_14.10-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0-rc.0_15.5-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0-rc.0_15.5-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0-rc.0_16.1-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0-rc.0_16.1-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0-rc.0_17.2-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0-rc.0_17.2-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_11.22-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_11.22-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_12.17-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_12.17-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_13.13-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_13.13-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_14.10-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_14.10-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_15.5-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_15.5-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_16.1-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_16.1-bookworm.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_17.2-alpine.tar
k3s ctr images import images/kubedb-postgres-archiver-v0.13.0_17.2-bookworm.tar
k3s ctr images import images/kubedb-postgres-init-0.17.1.tar
k3s ctr images import images/postgis-postgis-11-3.3.tar
k3s ctr images import images/postgis-postgis-12-3.4.tar
k3s ctr images import images/postgis-postgis-13-3.4.tar
k3s ctr images import images/postgis-postgis-14-3.4.tar
k3s ctr images import images/postgis-postgis-15-3.4.tar
k3s ctr images import images/postgis-postgis-16-3.4.tar
k3s ctr images import images/prometheuscommunity-postgres-exporter-v0.15.0.tar
k3s ctr images import images/timescale-timescaledb-2.14.2-pg13-oss.tar
k3s ctr images import images/timescale-timescaledb-2.14.2-pg14-oss.tar
k3s ctr images import images/timescale-timescaledb-2.14.2-pg15-oss.tar
k3s ctr images import images/timescale-timescaledb-2.14.2-pg16-oss.tar
