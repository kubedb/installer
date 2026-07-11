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

k3s ctr images import images/library-mongo-3.4.17.tar
k3s ctr images import images/library-mongo-3.4.22.tar
k3s ctr images import images/library-mongo-3.6.13.tar
k3s ctr images import images/library-mongo-3.6.8.tar
k3s ctr images import images/library-mongo-4.0.11.tar
k3s ctr images import images/library-mongo-4.0.3.tar
k3s ctr images import images/library-mongo-4.0.5.tar
k3s ctr images import images/library-mongo-4.1.13.tar
k3s ctr images import images/library-mongo-4.1.4.tar
k3s ctr images import images/library-mongo-4.1.7.tar
k3s ctr images import images/library-mongo-5.0.2.tar
k3s ctr images import images/library-mongo-5.0.3.tar
k3s ctr images import images/percona-percona-server-mongodb-3.6.18.tar
k3s ctr images import images/percona-percona-server-mongodb-4.0.10.tar
k3s ctr images import images/percona-percona-server-mongodb-4.2.24.tar
k3s ctr images import images/percona-percona-server-mongodb-4.2.7-7.tar
k3s ctr images import images/percona-percona-server-mongodb-4.4.10.tar
k3s ctr images import images/percona-percona-server-mongodb-4.4.26.tar
k3s ctr images import images/percona-percona-server-mongodb-5.0.23.tar
k3s ctr images import images/percona-percona-server-mongodb-5.0.29.tar
k3s ctr images import images/percona-percona-server-mongodb-6.0.12.tar
k3s ctr images import images/percona-percona-server-mongodb-6.0.24.tar
k3s ctr images import images/percona-percona-server-mongodb-7.0.18.tar
k3s ctr images import images/percona-percona-server-mongodb-7.0.28.tar
k3s ctr images import images/percona-percona-server-mongodb-7.0.4.tar
k3s ctr images import images/percona-percona-server-mongodb-8.0.17.tar
k3s ctr images import images/percona-percona-server-mongodb-8.0.8.tar
k3s ctr images import images/appscode-images-mongo-4.2.24.tar
k3s ctr images import images/appscode-images-mongo-4.2.3.tar
k3s ctr images import images/appscode-images-mongo-4.4.26.tar
k3s ctr images import images/appscode-images-mongo-4.4.6.tar
k3s ctr images import images/appscode-images-mongo-5.0.15.tar
k3s ctr images import images/appscode-images-mongo-5.0.23.tar
k3s ctr images import images/appscode-images-mongo-5.0.26.tar
k3s ctr images import images/appscode-images-mongo-5.0.31.tar
k3s ctr images import images/appscode-images-mongo-6.0.12.tar
k3s ctr images import images/appscode-images-mongo-6.0.24.tar
k3s ctr images import images/appscode-images-mongo-6.0.5.tar
k3s ctr images import images/appscode-images-mongo-7.0.16.tar
k3s ctr images import images/appscode-images-mongo-7.0.21.tar
k3s ctr images import images/appscode-images-mongo-7.0.28.tar
k3s ctr images import images/appscode-images-mongo-7.0.5.tar
k3s ctr images import images/appscode-images-mongo-7.0.8.tar
k3s ctr images import images/appscode-images-mongo-8.0.10.tar
k3s ctr images import images/appscode-images-mongo-8.0.17.tar
k3s ctr images import images/appscode-images-mongo-8.0.3.tar
k3s ctr images import images/appscode-images-mongo-8.0.4.tar
k3s ctr images import images/kubedb-kubedb-migrator-mongodb-v0.6.0.tar
k3s ctr images import images/kubedb-mongo-3.4.tar
k3s ctr images import images/kubedb-mongo-3.4-v1.tar
k3s ctr images import images/kubedb-mongo-3.4-v2.tar
k3s ctr images import images/kubedb-mongo-3.4-v3.tar
k3s ctr images import images/kubedb-mongo-3.4-v4.tar
k3s ctr images import images/kubedb-mongo-3.4-v5.tar
k3s ctr images import images/kubedb-mongo-3.4.17.tar
k3s ctr images import images/kubedb-mongo-3.4.22.tar
k3s ctr images import images/kubedb-mongo-3.6.tar
k3s ctr images import images/kubedb-mongo-3.6-v1.tar
k3s ctr images import images/kubedb-mongo-3.6-v2.tar
k3s ctr images import images/kubedb-mongo-3.6-v3.tar
k3s ctr images import images/kubedb-mongo-3.6-v4.tar
k3s ctr images import images/kubedb-mongo-3.6-v5.tar
k3s ctr images import images/kubedb-mongo-3.6.13.tar
k3s ctr images import images/kubedb-mongo-3.6.8.tar
k3s ctr images import images/kubedb-mongo-4.0.tar
k3s ctr images import images/kubedb-mongo-4.0-v1.tar
k3s ctr images import images/kubedb-mongo-4.0-v2.tar
k3s ctr images import images/kubedb-mongo-4.0-v3.tar
k3s ctr images import images/kubedb-mongo-4.0.11.tar
k3s ctr images import images/kubedb-mongo-4.0.3.tar
k3s ctr images import images/kubedb-mongo-4.0.5.tar
k3s ctr images import images/kubedb-mongo-4.0.5-v1.tar
k3s ctr images import images/kubedb-mongo-4.0.5-v2.tar
k3s ctr images import images/kubedb-mongo-4.1.tar
k3s ctr images import images/kubedb-mongo-4.1-v1.tar
k3s ctr images import images/kubedb-mongo-4.1.13.tar
k3s ctr images import images/kubedb-mongo-4.1.4.tar
k3s ctr images import images/kubedb-mongo-4.1.7.tar
k3s ctr images import images/kubedb-mongo-4.1.7-v1.tar
k3s ctr images import images/kubedb-mongo-4.1.7-v2.tar
k3s ctr images import images/kubedb-mongo-4.2.tar
k3s ctr images import images/kubedb-mongodb-init-0.1.0.tar
k3s ctr images import images/kubedb-mongodb-init-0.3.0.tar
k3s ctr images import images/kubedb-mongodb-init-4.1-v9.tar
k3s ctr images import images/kubedb-mongodb-init-4.1.4-v9.tar
k3s ctr images import images/kubedb-mongodb-init-4.2-v9.tar
k3s ctr images import images/kubedb-mongodb-init-6.0-v11.tar
k3s ctr images import images/kubedb-mongodb_exporter-v0.20.4.tar
k3s ctr images import images/kubedb-mongodb_exporter-v0.47.2.tar
k3s ctr images import images/kubedb-wal-g-v2026.3.30_mongo.tar
