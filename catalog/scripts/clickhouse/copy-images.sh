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

$CMD cp --allow-nondistributable-artifacts --insecure clickhouse/clickhouse-keeper:24.4.1 $IMAGE_REGISTRY/clickhouse/clickhouse-keeper:24.4.1
$CMD cp --allow-nondistributable-artifacts --insecure clickhouse/clickhouse-keeper:25.7.1 $IMAGE_REGISTRY/clickhouse/clickhouse-keeper:25.7.1
$CMD cp --allow-nondistributable-artifacts --insecure clickhouse/clickhouse-server:24.4.1 $IMAGE_REGISTRY/clickhouse/clickhouse-server:24.4.1
$CMD cp --allow-nondistributable-artifacts --insecure clickhouse/clickhouse-server:25.7.1 $IMAGE_REGISTRY/clickhouse/clickhouse-server:25.7.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/clickhouse-init:24.4.1-v3 $IMAGE_REGISTRY/kubedb/clickhouse-init:24.4.1-v3
