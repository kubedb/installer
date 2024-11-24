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

$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/memcached:1.5.22-alpine images/appscode-images-memcached-1.5.22-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/memcached:1.6.22-alpine images/appscode-images-memcached-1.6.22-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/memcached:1.6.29-alpine images/appscode-images-memcached-1.6.29-alpine.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/memcached_exporter:v0.14.3-ac images/appscode-images-memcached_exporter-v0.14.3-ac.tar
$CMD pull --allow-nondistributable-artifacts --insecure prom/memcached-exporter:v0.14.2 images/prom-memcached-exporter-v0.14.2.tar

tar -czvf images.tar.gz images
