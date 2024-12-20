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

$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/singlestore-node:alma-8.1.32-e3d3cde6da $IMAGE_REGISTRY/appscode-images/singlestore-node:alma-8.1.32-e3d3cde6da
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/singlestore-node:alma-8.5.30-4f46ab16a5 $IMAGE_REGISTRY/appscode-images/singlestore-node:alma-8.5.30-4f46ab16a5
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/singlestore-node:alma-8.5.7-bf633c1a54 $IMAGE_REGISTRY/appscode-images/singlestore-node:alma-8.5.7-bf633c1a54
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/singlestore-node:alma-8.7.10-95e2357384 $IMAGE_REGISTRY/appscode-images/singlestore-node:alma-8.7.10-95e2357384
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/singlestore-node:alma-8.7.21-f0b8de04d5 $IMAGE_REGISTRY/appscode-images/singlestore-node:alma-8.7.21-f0b8de04d5
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/singlestore-node:alma-8.9.3-bfa36a984a $IMAGE_REGISTRY/appscode-images/singlestore-node:alma-8.9.3-bfa36a984a
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/singlestore-coordinator:v0.5.0 $IMAGE_REGISTRY/kubedb/singlestore-coordinator:v0.5.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/singlestore-init:8.1-v2 $IMAGE_REGISTRY/kubedb/singlestore-init:8.1-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/singlestore-init:8.5-v2 $IMAGE_REGISTRY/kubedb/singlestore-init:8.5-v2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/singlestore-init:8.7.10-v1 $IMAGE_REGISTRY/kubedb/singlestore-init:8.7.10-v1
$CMD cp --allow-nondistributable-artifacts --insecure singlestore/cluster-in-a-box:alma-8.1.32-e3d3cde6da-4.0.16-1.17.6 $IMAGE_REGISTRY/singlestore/cluster-in-a-box:alma-8.1.32-e3d3cde6da-4.0.16-1.17.6
$CMD cp --allow-nondistributable-artifacts --insecure singlestore/cluster-in-a-box:alma-8.5.22-fe61f40cd1-4.1.0-1.17.11 $IMAGE_REGISTRY/singlestore/cluster-in-a-box:alma-8.5.22-fe61f40cd1-4.1.0-1.17.11
$CMD cp --allow-nondistributable-artifacts --insecure singlestore/cluster-in-a-box:alma-8.5.7-bf633c1a54-4.0.17-1.17.8 $IMAGE_REGISTRY/singlestore/cluster-in-a-box:alma-8.5.7-bf633c1a54-4.0.17-1.17.8
$CMD cp --allow-nondistributable-artifacts --insecure singlestore/cluster-in-a-box:alma-8.7.10-95e2357384-4.1.0-1.17.14 $IMAGE_REGISTRY/singlestore/cluster-in-a-box:alma-8.7.10-95e2357384-4.1.0-1.17.14
