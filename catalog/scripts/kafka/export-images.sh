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

$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-connect-cluster:3.5.2 images/appscode-images-kafka-connect-cluster-3.5.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-connect-cluster:3.6.1 images/appscode-images-kafka-connect-cluster-3.6.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-connect-cluster:3.7.2 images/appscode-images-kafka-connect-cluster-3.7.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-connect-cluster:3.8.1 images/appscode-images-kafka-connect-cluster-3.8.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-connect-cluster:3.9.0 images/appscode-images-kafka-connect-cluster-3.9.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-connect-cluster:4.0.0 images/appscode-images-kafka-connect-cluster-4.0.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-cruise-control:3.5.2 images/appscode-images-kafka-cruise-control-3.5.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-cruise-control:3.6.1 images/appscode-images-kafka-cruise-control-3.6.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-cruise-control:3.7.2 images/appscode-images-kafka-cruise-control-3.7.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-cruise-control:3.8.1 images/appscode-images-kafka-cruise-control-3.8.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-cruise-control:3.9.0 images/appscode-images-kafka-cruise-control-3.9.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-cruise-control:4.0.0 images/appscode-images-kafka-cruise-control-4.0.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-kraft:3.5.2 images/appscode-images-kafka-kraft-3.5.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-kraft:3.6.1 images/appscode-images-kafka-kraft-3.6.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-kraft:3.7.2 images/appscode-images-kafka-kraft-3.7.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-kraft:3.8.1 images/appscode-images-kafka-kraft-3.8.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka-kraft:3.9.0 images/appscode-images-kafka-kraft-3.9.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kafka:4.0.0 images/appscode-images-kafka-4.0.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/kafka-init:4.0-v1 images/kubedb-kafka-init-4.0-v1.tar

tar -czvf images.tar.gz images
