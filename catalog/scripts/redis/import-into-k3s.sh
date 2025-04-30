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

k3s ctr images import images/appscode-images-redis-5.0.14-bullseye.tar
k3s ctr images import images/appscode-images-redis-6.0.20-bookworm.tar
k3s ctr images import images/appscode-images-redis-6.2.14-bookworm.tar
k3s ctr images import images/appscode-images-redis-6.2.16-bookworm.tar
k3s ctr images import images/appscode-images-redis-7.0.14-bookworm.tar
k3s ctr images import images/appscode-images-redis-7.0.15-bookworm.tar
k3s ctr images import images/appscode-images-redis-7.2.3-bookworm.tar
k3s ctr images import images/appscode-images-redis-7.2.4-bookworm.tar
k3s ctr images import images/appscode-images-redis-7.2.6-bookworm.tar
k3s ctr images import images/appscode-images-redis-7.4.0-bookworm.tar
k3s ctr images import images/appscode-images-redis-7.4.1-bookworm.tar
k3s ctr images import images/appscode-images-valkey-7.2.5.tar
k3s ctr images import images/appscode-images-valkey-7.2.9.tar
k3s ctr images import images/appscode-images-valkey-8.0.3.tar
k3s ctr images import images/appscode-images-valkey-8.1.1.tar
k3s ctr images import images/kubedb-redis-coordinator-v0.33.0.tar
k3s ctr images import images/kubedb-redis-init-0.10.0.tar
k3s ctr images import images/kubedb-redis-4.0.11.tar
k3s ctr images import images/kubedb-redis_exporter-1.66.0.tar
