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

k3s ctr images import images/appscode-images-mariadb-10.10.7-jammy.tar
k3s ctr images import images/appscode-images-mariadb-10.11.6-jammy.tar
k3s ctr images import images/appscode-images-mariadb-10.4.32-focal.tar
k3s ctr images import images/appscode-images-mariadb-10.5.23-focal.tar
k3s ctr images import images/appscode-images-mariadb-10.6.16-focal.tar
k3s ctr images import images/appscode-images-mariadb-11.0.4-jammy.tar
k3s ctr images import images/appscode-images-mariadb-11.1.3-jammy.tar
k3s ctr images import images/appscode-images-mariadb-11.2.2-jammy.tar
k3s ctr images import images/appscode-images-mariadb-11.3.2-jammy.tar
k3s ctr images import images/appscode-images-mariadb-11.4.3-noble.tar
k3s ctr images import images/appscode-images-mariadb-11.5.2-noble.tar
k3s ctr images import images/appscode-images-mariadb-11.6.2-noble.tar
k3s ctr images import images/kubedb-mariadb-archiver-v0.11.0_10.10.7-jammy.tar
k3s ctr images import images/kubedb-mariadb-archiver-v0.11.0_10.11.6-jammy.tar
k3s ctr images import images/kubedb-mariadb-archiver-v0.11.0_10.4.32-focal.tar
k3s ctr images import images/kubedb-mariadb-archiver-v0.11.0_10.5.23-focal.tar
k3s ctr images import images/kubedb-mariadb-archiver-v0.11.0_10.6.16-focal.tar
k3s ctr images import images/kubedb-mariadb-archiver-v0.11.0_11.0.4-jammy.tar
k3s ctr images import images/kubedb-mariadb-archiver-v0.11.0_11.1.3-jammy.tar
k3s ctr images import images/kubedb-mariadb-archiver-v0.11.0_11.2.2-jammy.tar
k3s ctr images import images/kubedb-mariadb-coordinator-v0.31.0.tar
k3s ctr images import images/kubedb-mariadb-init-0.5.2.tar
