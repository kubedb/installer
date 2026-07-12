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

k3s ctr images import images/library-mysql-5.7.25.tar
k3s ctr images import images/library-mysql-5.7.29.tar
k3s ctr images import images/library-mysql-5.7.31.tar
k3s ctr images import images/library-mysql-5.7.33.tar
k3s ctr images import images/library-mysql-5.7.35.tar
k3s ctr images import images/library-mysql-5.7.36.tar
k3s ctr images import images/library-mysql-8.0.14.tar
k3s ctr images import images/library-mysql-8.0.17.tar
k3s ctr images import images/library-mysql-8.0.20.tar
k3s ctr images import images/library-mysql-8.0.21.tar
k3s ctr images import images/library-mysql-8.0.23.tar
k3s ctr images import images/library-mysql-8.0.27.tar
k3s ctr images import images/library-mysql-8.0.3.tar
k3s ctr images import images/mysql-mysql-router-8.0.27.tar
k3s ctr images import images/mysql-mysql-server-8.0.27.tar
k3s ctr images import images/prom-mysqld-exporter-v0.13.0.tar
k3s ctr images import images/appscode-images-mysql-router-8.0.45.tar
k3s ctr images import images/appscode-images-mysql-router-8.4.8.tar
k3s ctr images import images/appscode-images-mysql-router-9.0.1.tar
k3s ctr images import images/appscode-images-mysql-router-9.1.0.tar
k3s ctr images import images/appscode-images-mysql-router-9.4.0.tar
k3s ctr images import images/appscode-images-mysql-router-9.6.0.tar
k3s ctr images import images/appscode-images-mysql-5.7.41-oracle.tar
k3s ctr images import images/appscode-images-mysql-5.7.42-debian.tar
k3s ctr images import images/appscode-images-mysql-5.7.44-oracle.tar
k3s ctr images import images/appscode-images-mysql-8.0.29-oracle.tar
k3s ctr images import images/appscode-images-mysql-8.0.31-oracle.tar
k3s ctr images import images/appscode-images-mysql-8.0.32-oracle.tar
k3s ctr images import images/appscode-images-mysql-8.0.35-oracle.tar
k3s ctr images import images/appscode-images-mysql-8.0.36-debian.tar
k3s ctr images import images/appscode-images-mysql-8.1.0-oracle.tar
k3s ctr images import images/appscode-images-mysql-8.2.0-oracle.tar
k3s ctr images import images/appscode-images-mysql-8.4.2-oracle.tar
k3s ctr images import images/appscode-images-mysql-8.4.3-oracle.tar
k3s ctr images import images/appscode-images-mysql-8.4.8-oracle.tar
k3s ctr images import images/appscode-images-mysql-9.0.1-oracle.tar
k3s ctr images import images/appscode-images-mysql-9.1.0-oracle.tar
k3s ctr images import images/appscode-images-mysql-9.4.0-oracle.tar
k3s ctr images import images/appscode-images-mysql-9.6.0-oracle.tar
k3s ctr images import images/kubedb-kubedb-migrator-mysql-v0.6.0.tar
k3s ctr images import images/kubedb-mysql-archiver-v0.27.0_5.7.44.tar
k3s ctr images import images/kubedb-mysql-archiver-v0.27.0_8.0.35.tar
k3s ctr images import images/kubedb-mysql-archiver-v0.27.0_8.1.0.tar
k3s ctr images import images/kubedb-mysql-archiver-v0.27.0_8.2.0.tar
k3s ctr images import images/kubedb-mysql-archiver-v0.27.0_8.4.3.tar
k3s ctr images import images/kubedb-mysql-archiver-v0.27.0_9.1.0.tar
k3s ctr images import images/kubedb-mysql-coordinator-v0.44.0.tar
k3s ctr images import images/kubedb-mysql-init-0.1.0.tar
k3s ctr images import images/kubedb-mysql-init-5.7.tar
k3s ctr images import images/kubedb-mysql-init-5.7-v3.tar
k3s ctr images import images/kubedb-mysql-init-5.7-v4.tar
k3s ctr images import images/kubedb-mysql-init-5.7-v9.tar
k3s ctr images import images/kubedb-mysql-init-8.0.17.tar
k3s ctr images import images/kubedb-mysql-init-8.0.21.tar
k3s ctr images import images/kubedb-mysql-init-8.0.26-v3.tar
k3s ctr images import images/kubedb-mysql-init-8.0.3.tar
k3s ctr images import images/kubedb-mysql-init-8.0.3-v2.tar
k3s ctr images import images/kubedb-mysql-router-init-v0.44.0.tar
k3s ctr images import images/kubedb-mysql-5.tar
k3s ctr images import images/kubedb-mysql-5-v1.tar
k3s ctr images import images/kubedb-mysql-5.7.tar
k3s ctr images import images/kubedb-mysql-5.7-v1.tar
k3s ctr images import images/kubedb-mysql-5.7-v2.tar
k3s ctr images import images/kubedb-mysql-5.7.25.tar
k3s ctr images import images/kubedb-mysql-5.7.25-v1.tar
k3s ctr images import images/kubedb-mysql-5.7.25-v2.tar
k3s ctr images import images/kubedb-mysql-5.7.29.tar
k3s ctr images import images/kubedb-mysql-5.7.31.tar
k3s ctr images import images/kubedb-mysql-5.7.31-v1.tar
k3s ctr images import images/kubedb-mysql-5.7.33.tar
k3s ctr images import images/kubedb-mysql-5.7.35.tar
k3s ctr images import images/kubedb-mysql-8.tar
k3s ctr images import images/kubedb-mysql-8-v1.tar
k3s ctr images import images/kubedb-mysql-8.0.tar
k3s ctr images import images/kubedb-mysql-8.0-v1.tar
k3s ctr images import images/kubedb-mysql-8.0-v2.tar
k3s ctr images import images/kubedb-mysql-8.0.14.tar
k3s ctr images import images/kubedb-mysql-8.0.14-v1.tar
k3s ctr images import images/kubedb-mysql-8.0.14-v2.tar
k3s ctr images import images/kubedb-mysql-8.0.20.tar
k3s ctr images import images/kubedb-mysql-8.0.20-v1.tar
k3s ctr images import images/kubedb-mysql-8.0.21.tar
k3s ctr images import images/kubedb-mysql-8.0.21-v1.tar
k3s ctr images import images/kubedb-mysql-8.0.23.tar
k3s ctr images import images/kubedb-mysql-8.0.26.tar
k3s ctr images import images/kubedb-mysql-8.0.3.tar
k3s ctr images import images/kubedb-mysql-8.0.3-v1.tar
k3s ctr images import images/kubedb-mysql-8.0.3-v2.tar
k3s ctr images import images/kubedb-mysqld-exporter-v0.18.0.tar
k3s ctr images import images/kubedb-replication-mode-detector-v0.53.0.tar
