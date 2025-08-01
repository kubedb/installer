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

$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:10.10.7-jammy images/appscode-images-mariadb-10.10.7-jammy.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:10.11.6-jammy images/appscode-images-mariadb-10.11.6-jammy.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:10.4.32-focal images/appscode-images-mariadb-10.4.32-focal.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:10.5.23-focal images/appscode-images-mariadb-10.5.23-focal.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:10.6.16-focal images/appscode-images-mariadb-10.6.16-focal.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:11.0.4-jammy images/appscode-images-mariadb-11.0.4-jammy.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:11.1.3-jammy images/appscode-images-mariadb-11.1.3-jammy.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:11.2.2-jammy images/appscode-images-mariadb-11.2.2-jammy.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:11.3.2-jammy images/appscode-images-mariadb-11.3.2-jammy.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:11.4.3-noble images/appscode-images-mariadb-11.4.3-noble.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:11.5.2-noble images/appscode-images-mariadb-11.5.2-noble.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/mariadb:11.6.2-noble images/appscode-images-mariadb-11.6.2-noble.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.17.0-rc.0_10.10.7-jammy images/kubedb-mariadb-archiver-v0.17.0-rc.0_10.10.7-jammy.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.17.0-rc.0_10.11.6-jammy images/kubedb-mariadb-archiver-v0.17.0-rc.0_10.11.6-jammy.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.17.0-rc.0_10.4.32-focal images/kubedb-mariadb-archiver-v0.17.0-rc.0_10.4.32-focal.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.17.0-rc.0_10.5.23-focal images/kubedb-mariadb-archiver-v0.17.0-rc.0_10.5.23-focal.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.17.0-rc.0_10.6.16-focal images/kubedb-mariadb-archiver-v0.17.0-rc.0_10.6.16-focal.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.17.0-rc.0_11.0.4-jammy images/kubedb-mariadb-archiver-v0.17.0-rc.0_11.0.4-jammy.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.17.0-rc.0_11.1.3-jammy images/kubedb-mariadb-archiver-v0.17.0-rc.0_11.1.3-jammy.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-archiver:v0.17.0-rc.0_11.2.2-jammy images/kubedb-mariadb-archiver-v0.17.0-rc.0_11.2.2-jammy.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-coordinator:v0.37.0-rc.0 images/kubedb-mariadb-coordinator-v0.37.0-rc.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/mariadb-init:0.7.0 images/kubedb-mariadb-init-0.7.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure mariadb/maxscale:24.02.4 images/mariadb-maxscale-24.02.4.tar

tar -czvf images.tar.gz images
