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

set -eou pipefail

if [ "$#" -ne 1 ]; then
    echo "Error: missing path_to_input_crds_directory"
    echo "Usage: import-crds.sh <path_to_input_crds_directory>"
    exit 1
fi

crd-importer \
    --input=${1} \
    --out=./charts/kubedb-crds/crds

crd-importer --v=v1beta1 \
    --input=${1} \
    --out=./charts/kubedb-catalog/crds \
    --group=catalog.kubedb.com
