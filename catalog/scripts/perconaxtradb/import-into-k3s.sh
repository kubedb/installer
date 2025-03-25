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

k3s ctr images import images/appscode-images-percona-xtradb-cluster-5.7.44.tar
k3s ctr images import images/appscode-images-percona-xtradb-cluster-8.0.40.tar
k3s ctr images import images/appscode-images-percona-xtradb-cluster-8.4.3.tar
k3s ctr images import images/kubedb-percona-xtradb-coordinator-v0.26.0.tar
k3s ctr images import images/kubedb-percona-xtradb-init-0.2.3.tar
k3s ctr images import images/prom-mysqld-exporter-v0.13.0.tar
