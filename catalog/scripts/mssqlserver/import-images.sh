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

$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mssql-coordinator-v0.7.0.tar $IMAGE_REGISTRY/kubedb/mssql-coordinator:v0.7.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mssql-exporter-1.1.0.tar $IMAGE_REGISTRY/kubedb/mssql-exporter:1.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mssql-init-2022-ubuntu-22-v3.tar $IMAGE_REGISTRY/kubedb/mssql-init:2022-ubuntu-22-v3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-mssqlserver-archiver-v0.6.0.tar $IMAGE_REGISTRY/kubedb/mssqlserver-archiver:v0.6.0
$CMD push --allow-nondistributable-artifacts --insecure images/mssql-server-2022-CU12-ubuntu-22.04.tar $IMAGE_REGISTRY/mssql/server:2022-CU12-ubuntu-22.04
$CMD push --allow-nondistributable-artifacts --insecure images/mssql-server-2022-CU14-ubuntu-22.04.tar $IMAGE_REGISTRY/mssql/server:2022-CU14-ubuntu-22.04
$CMD push --allow-nondistributable-artifacts --insecure images/mssql-server-2022-CU16-ubuntu-22.04.tar $IMAGE_REGISTRY/mssql/server:2022-CU16-ubuntu-22.04
