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

$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/percona-xtradb-coordinator:v0.22.0 images/kubedb-percona-xtradb-coordinator-v0.22.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/percona-xtradb-init:0.2.1 images/kubedb-percona-xtradb-init-0.2.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure percona/percona-xtradb-cluster:8.0.26 images/percona-percona-xtradb-cluster-8.0.26.tar
$CMD pull --allow-nondistributable-artifacts --insecure percona/percona-xtradb-cluster:8.0.28 images/percona-percona-xtradb-cluster-8.0.28.tar
$CMD pull --allow-nondistributable-artifacts --insecure percona/percona-xtradb-cluster:8.0.31 images/percona-percona-xtradb-cluster-8.0.31.tar
$CMD pull --allow-nondistributable-artifacts --insecure prom/mysqld-exporter:v0.13.0 images/prom-mysqld-exporter-v0.13.0.tar

tar -czvf images.tar.gz images