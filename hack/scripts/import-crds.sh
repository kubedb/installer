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

crd_dir=${1:-}

api_repo_url=https://github.com/kubedb/apimachinery.git
api_repo_tag=${KUBEDB_APIMACHINERY_TAG:-master}

if [ "$#" -ne 1 ]; then
    if [ "${api_repo_tag}" == "master" ]; then
        echo "Error: missing path_to_input_crds_directory"
        echo "Usage: import-crds.sh <path_to_input_crds_directory>"
        exit 1
    fi

    tmp_dir=$(mktemp -d -t api-XXXXXXXXXX)
    # always cleanup temp dir
    # ref: https://opensource.com/article/20/6/bash-trap
    trap \
        "{ rm -rf "${tmp_dir}"; }" \
        SIGINT SIGTERM ERR EXIT

    mkdir -p ${tmp_dir}
    pushd $tmp_dir
    git clone $api_repo_url
    repo_dir=$(ls -b1)
    cd $repo_dir
    git checkout $api_repo_tag
    popd
    crd_dir=${tmp_dir}/${repo_dir}/crds
fi

crd-importer \
    --input=${crd_dir} \
    --out=./charts/kubedb-crds/crds \
    --group=kubedb.com \
    --group=catalog.kubedb.com \
    --group=config.kubedb.com \
    --group=ops.kubedb.com \
    --group=autoscaling.kubedb.com \
    --group=dashboard.kubedb.com \
    --group=postgres.kubedb.com \
    --group=archiver.kubedb.com \
    --group=schema.kubedb.com

crd-importer \
    --input=${crd_dir} \
    --out=. --output-yaml=crds/kubedb-crds.yaml \
    --group=kubedb.com \
    --group=catalog.kubedb.com \
    --group=config.kubedb.com \
    --group=ops.kubedb.com \
    --group=autoscaling.kubedb.com \
    --group=dashboard.kubedb.com \
    --group=postgres.kubedb.com \
    --group=archiver.kubedb.com \
    --group=schema.kubedb.com

crd-importer \
    --input=${crd_dir} \
    --out=./charts/kubedb-catalog/crds \
    --group=catalog.kubedb.com

crd-importer \
    --input=${crd_dir} \
    --out=. --output-yaml=crds/kubedb-catalog-crds.yaml \
    --group=catalog.kubedb.com

crd-importer \
    --input=https://github.com/kmodules/custom-resources/raw/release-1.25/crds/metrics.appscode.com_metricsconfigurations.yaml \
    --out=./charts/kubedb-metrics/crds

crd-importer \
    --input=${crd_dir} \
    --out=./charts/kubedb-ui-server/crds \
    --group=kubedb.com

crd-importer \
    --input=https://github.com/open-viz/grafana-tools/raw/v0.0.1/crds/openviz.dev_grafanadashboards.yaml \
    --out=./charts/kubedb-grafana-dashboards/crds

crd-importer \
    --input=https://github.com/kubeops/supervisor/raw/v0.0.1/crds/supervisor.appscode.com_recommendations.yaml \
    --out=./charts/kubedb-ops-manager/crds
