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

crd_dir=${1:-}/apimachinery/crds
update_kubedb_crds=true

api_repo_url=https://github.com/kubedb/apimachinery.git
api_repo_tag=${KUBEDB_APIMACHINERY_TAG:-mssql}

if [ "$#" -ne 1 ]; then
    if [ "${api_repo_tag}" == "master" ]; then
        echo "Skipping updating kubedb/apimachinery crds"
        echo "To update use: import-crds.sh <path_to_input_crds_directory>"
        update_kubedb_crds=false
    else
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
fi

if [ "$update_kubedb_crds" = true ] && [ -d ${crd_dir} ]; then
    crd-importer \
        --input=${crd_dir} \
        --out=./charts/kubedb-crds/crds \
        --group=kubedb.com \
        --group=catalog.kubedb.com \
        --group=config.kubedb.com \
        --group=ops.kubedb.com \
        --group=autoscaling.kubedb.com \
        --group=elasticsearch.kubedb.com \
        --group=kafka.kubedb.com \
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
        --group=elasticsearch.kubedb.com \
        --group=kafka.kubedb.com \
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
fi

KEDACORE_HTTP_ADD_ON_TAG=${KEDACORE_HTTP_ADD_ON_TAG:-v0.8.0}
KMODULES_CUSTOM_RESOURCES_TAG=${KMODULES_CUSTOM_RESOURCES_TAG:-v0.30.0}
KMODULES_RESOURCE_METADATA_TAG=${KMODULES_RESOURCE_METADATA_TAG:-v0.18.9}
KUBEOPS_SUPERVISOR_TAG=${KUBEOPS_SUPERVISOR_TAG:-v0.0.5}
KUBERNETES_SIGS_GATEWAY_API_TAG=${KUBERNETES_SIGS_GATEWAY_API_TAG:-v1.0.0}
KUBESTASH_APIMACHINERY_TAG=${KUBESTASH_APIMACHINERY_TAG:-v0.10.0}
OPEN_VIZ_APIMACHINERY_TAG=${OPEN_VIZ_APIMACHINERY_TAG:-v0.0.7}

crd-importer \
    --input=https://github.com/kubestash/apimachinery/raw/${KUBESTASH_APIMACHINERY_TAG}/crds/addons.kubestash.com_addons.yaml \
    --input=https://github.com/kubestash/apimachinery/raw/${KUBESTASH_APIMACHINERY_TAG}/crds/addons.kubestash.com_functions.yaml \
    --out=./charts/kubedb-kubestash-catalog/crds

crd-importer \
    --input=https://github.com/kmodules/resource-metadata/raw/${KMODULES_RESOURCE_METADATA_TAG}/crds/node.k8s.appscode.com_nodetopologies.yaml \
    --out=./charts/kubedb-autoscaler/crds

crd-importer \
    --input=https://github.com/kmodules/custom-resources/raw/${KMODULES_CUSTOM_RESOURCES_TAG}/crds/metrics.appscode.com_metricsconfigurations.yaml \
    --out=./charts/kubedb-metrics/crds

crd-importer \
    --input=https://github.com/open-viz/apimachinery/raw/${OPEN_VIZ_APIMACHINERY_TAG}/crds/openviz.dev_grafanadashboards.yaml \
    --out=./charts/kubedb-grafana-dashboards/crds

crd-importer \
    --input=https://github.com/kubeops/supervisor/raw/${KUBEOPS_SUPERVISOR_TAG}/crds/supervisor.appscode.com_recommendations.yaml \
    --out=./charts/kubedb-ops-manager/crds

# dashboard charts
crd-importer \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_gateways.yaml \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_httproutes.yaml \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_referencegrants.yaml \
    --out=./charts/dbgate/crds
crd-importer \
    --input=https://github.com/kedacore/http-add-on/raw/${KEDACORE_HTTP_ADD_ON_TAG}/config/crd/bases/http.keda.sh_httpscaledobjects.yaml \
    --labels 'app.kubernetes.io/managed-by=Helm' \
    --annotations 'meta.helm.sh/release-name=keda-add-ons-http,meta.helm.sh/release-namespace=keda' \
    --out=./charts/dbgate/crds

crd-importer \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_gateways.yaml \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_httproutes.yaml \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_referencegrants.yaml \
    --out=./charts/kafka-ui/crds
crd-importer \
    --input=https://github.com/kedacore/http-add-on/raw/${KEDACORE_HTTP_ADD_ON_TAG}/config/crd/bases/http.keda.sh_httpscaledobjects.yaml \
    --labels 'app.kubernetes.io/managed-by=Helm' \
    --annotations 'meta.helm.sh/release-name=keda-add-ons-http,meta.helm.sh/release-namespace=keda' \
    --out=./charts/kafka-ui/crds

crd-importer \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_gateways.yaml \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_httproutes.yaml \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_referencegrants.yaml \
    --out=./charts/mongo-ui/crds
crd-importer \
    --input=https://github.com/kedacore/http-add-on/raw/${KEDACORE_HTTP_ADD_ON_TAG}/config/crd/bases/http.keda.sh_httpscaledobjects.yaml \
    --labels 'app.kubernetes.io/managed-by=Helm' \
    --annotations 'meta.helm.sh/release-name=keda-add-ons-http,meta.helm.sh/release-namespace=keda' \
    --out=./charts/mongo-ui/crds

crd-importer \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_gateways.yaml \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_httproutes.yaml \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_referencegrants.yaml \
    --out=./charts/pgadmin/crds
crd-importer \
    --input=https://github.com/kedacore/http-add-on/raw/${KEDACORE_HTTP_ADD_ON_TAG}/config/crd/bases/http.keda.sh_httpscaledobjects.yaml \
    --labels 'app.kubernetes.io/managed-by=Helm' \
    --annotations 'meta.helm.sh/release-name=keda-add-ons-http,meta.helm.sh/release-namespace=keda' \
    --out=./charts/pgadmin/crds

crd-importer \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_gateways.yaml \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_httproutes.yaml \
    --input=https://github.com/kubernetes-sigs/gateway-api/raw/${KUBERNETES_SIGS_GATEWAY_API_TAG}/config/crd/standard/gateway.networking.k8s.io_referencegrants.yaml \
    --out=./charts/phpmyadmin/crds
crd-importer \
    --input=https://github.com/kedacore/http-add-on/raw/${KEDACORE_HTTP_ADD_ON_TAG}/config/crd/bases/http.keda.sh_httpscaledobjects.yaml \
    --labels 'app.kubernetes.io/managed-by=Helm' \
    --annotations 'meta.helm.sh/release-name=keda-add-ons-http,meta.helm.sh/release-namespace=keda' \
    --out=./charts/phpmyadmin/crds

{
    crd_dir=${1:-}/provider-aws/package/crds
    update_aws_crds=true

    repo_url=https://github.com/kubedb/provider-aws.git
    repo_tag=${KUBEDB_PROVIDER_AWS_TAG:-main}

    if [ "$#" -ne 1 ]; then
        if [ "${repo_tag}" == "main" ]; then
            echo "Skipping updating kubedb/provider-aws crds"
            echo "To update use: import-crds.sh <path_to_input_crds_directory>"
            update_aws_crds=false
        else
            tmp_dir=$(mktemp -d -t api-XXXXXXXXXX)
            # always cleanup temp dir
            # ref: https://opensource.com/article/20/6/bash-trap
            trap \
                "{ rm -rf "${tmp_dir}"; }" \
                SIGINT SIGTERM ERR EXIT

            mkdir -p ${tmp_dir}
            pushd $tmp_dir
            git clone $repo_url
            repo_dir=$(ls -b1)
            cd $repo_dir
            git checkout $repo_tag
            popd
            crd_dir=${tmp_dir}/${repo_dir}/package/crds
        fi
    fi

    if [ "$update_aws_crds" = true ] && [ -d ${crd_dir} ]; then
        crd-importer \
            --input=${crd_dir} \
            --out=./charts/kubedb-provider-aws/crds
    fi
}
{
    crd_dir=${1:-}/provider-azure/package/crds
    update_azure_crds=true

    repo_url=https://github.com/kubedb/provider-azure.git
    repo_tag=${KUBEDB_PROVIDER_AZURE_TAG:-main}

    if [ "$#" -ne 1 ]; then
        if [ "${repo_tag}" == "main" ]; then
            echo "Skipping updating kubedb/provider-azure crds"
            echo "To update use: import-crds.sh <path_to_input_crds_directory>"
            update_azure_crds=false
        else
            tmp_dir=$(mktemp -d -t api-XXXXXXXXXX)
            # always cleanup temp dir
            # ref: https://opensource.com/article/20/6/bash-trap
            trap \
                "{ rm -rf "${tmp_dir}"; }" \
                SIGINT SIGTERM ERR EXIT

            mkdir -p ${tmp_dir}
            pushd $tmp_dir
            git clone $repo_url
            repo_dir=$(ls -b1)
            cd $repo_dir
            git checkout $repo_tag
            popd
            crd_dir=${tmp_dir}/${repo_dir}/package/crds
        fi
    fi

    if [ "$update_azure_crds" = true ] && [ -d ${crd_dir} ]; then
        crd-importer \
            --input=${crd_dir} \
            --out=./charts/kubedb-provider-azure/crds
    fi
}
{
    crd_dir=${1:-}/provider-gcp/package/crds
    update_gcp_crds=true

    repo_url=https://github.com/kubedb/provider-gcp.git
    repo_tag=${KUBEDB_PROVIDER_GCP_TAG:-main}

    if [ "$#" -ne 1 ]; then
        if [ "${repo_tag}" == "main" ]; then
            echo "Skipping updating kubedb/provider-gcp crds"
            echo "To update use: import-crds.sh <path_to_input_crds_directory>"
            update_gcp_crds=false
        else
            tmp_dir=$(mktemp -d -t api-XXXXXXXXXX)
            # always cleanup temp dir
            # ref: https://opensource.com/article/20/6/bash-trap
            trap \
                "{ rm -rf "${tmp_dir}"; }" \
                SIGINT SIGTERM ERR EXIT

            mkdir -p ${tmp_dir}
            pushd $tmp_dir
            git clone $repo_url
            repo_dir=$(ls -b1)
            cd $repo_dir
            git checkout $repo_tag
            popd
            crd_dir=${tmp_dir}/${repo_dir}/package/crds
        fi
    fi

    if [ "$update_gcp_crds" = true ] && [ -d ${crd_dir} ]; then
        crd-importer \
            --input=${crd_dir} \
            --out=./charts/kubedb-provider-gcp/crds
    fi
}
