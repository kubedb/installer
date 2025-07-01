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

$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/singlestore-node:alma-8.1.32-e3d3cde6da images/appscode-images-singlestore-node-alma-8.1.32-e3d3cde6da.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/singlestore-node:alma-8.5.30-4f46ab16a5 images/appscode-images-singlestore-node-alma-8.5.30-4f46ab16a5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/singlestore-node:alma-8.5.7-bf633c1a54 images/appscode-images-singlestore-node-alma-8.5.7-bf633c1a54.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/singlestore-node:alma-8.7.10-95e2357384 images/appscode-images-singlestore-node-alma-8.7.10-95e2357384.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/singlestore-node:alma-8.7.21-f0b8de04d5 images/appscode-images-singlestore-node-alma-8.7.21-f0b8de04d5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/singlestore-node:alma-8.9.3-bfa36a984a images/appscode-images-singlestore-node-alma-8.9.3-bfa36a984a.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/singlestore-coordinator:v0.11.0 images/kubedb-singlestore-coordinator-v0.11.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/singlestore-init:8.1-v2 images/kubedb-singlestore-init-8.1-v2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/singlestore-init:8.5-v2 images/kubedb-singlestore-init-8.5-v2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/singlestore-init:8.7.10-v1 images/kubedb-singlestore-init-8.7.10-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure singlestore/cluster-in-a-box:alma-8.1.32-e3d3cde6da-4.0.16-1.17.6 images/singlestore-cluster-in-a-box-alma-8.1.32-e3d3cde6da-4.0.16-1.17.6.tar
$CMD pull --allow-nondistributable-artifacts --insecure singlestore/cluster-in-a-box:alma-8.5.22-fe61f40cd1-4.1.0-1.17.11 images/singlestore-cluster-in-a-box-alma-8.5.22-fe61f40cd1-4.1.0-1.17.11.tar
$CMD pull --allow-nondistributable-artifacts --insecure singlestore/cluster-in-a-box:alma-8.5.7-bf633c1a54-4.0.17-1.17.8 images/singlestore-cluster-in-a-box-alma-8.5.7-bf633c1a54-4.0.17-1.17.8.tar
$CMD pull --allow-nondistributable-artifacts --insecure singlestore/cluster-in-a-box:alma-8.7.10-95e2357384-4.1.0-1.17.14 images/singlestore-cluster-in-a-box-alma-8.7.10-95e2357384-4.1.0-1.17.14.tar

tar -czvf images.tar.gz images
