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

CMD="./crane"

$CMD push --allow-nondistributable-artifacts --insecure images/library-busybox-1.36.tar $IMAGE_REGISTRY/busybox:1.36
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-oracle-coordinator-v0.3.0-rc.0.tar $IMAGE_REGISTRY/kubedb/oracle-coordinator:v0.3.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-oracle-ee-21.3.0.tar $IMAGE_REGISTRY/kubedb/oracle-ee:21.3.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-oracle-init-v0.1.0.tar $IMAGE_REGISTRY/kubedb/oracle-init:v0.1.0
