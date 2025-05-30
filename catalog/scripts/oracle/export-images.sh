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

$CMD pull --allow-nondistributable-artifacts --insecure busybox:1.36 images/library-busybox-1.36.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/oracle-coordinator:v0.1.0 images/kubedb-oracle-coordinator-v0.1.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/oracle-ee:21.3.0 images/kubedb-oracle-ee-21.3.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/oracle-init:v0.1.0 images/kubedb-oracle-init-v0.1.0.tar

tar -czvf images.tar.gz images
