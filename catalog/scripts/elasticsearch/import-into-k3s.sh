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

k3s ctr images import images/floragunncom-sg-elasticsearch-7.9.3-oss-47.1.0.tar
k3s ctr images import images/appscode-images-elastic-6.8.23.tar
k3s ctr images import images/appscode-images-elastic-7.17.15.tar
k3s ctr images import images/appscode-images-elastic-7.17.28.tar
k3s ctr images import images/appscode-images-elastic-8.17.10.tar
k3s ctr images import images/appscode-images-elastic-8.17.6.tar
k3s ctr images import images/appscode-images-elastic-8.18.2.tar
k3s ctr images import images/appscode-images-elastic-8.18.8.tar
k3s ctr images import images/appscode-images-elastic-8.19.9.tar
k3s ctr images import images/appscode-images-elastic-8.2.3.tar
k3s ctr images import images/appscode-images-elastic-8.5.3.tar
k3s ctr images import images/appscode-images-elastic-9.0.2.tar
k3s ctr images import images/appscode-images-elastic-9.0.8.tar
k3s ctr images import images/appscode-images-elastic-9.1.4.tar
k3s ctr images import images/appscode-images-elastic-9.1.9.tar
k3s ctr images import images/appscode-images-elastic-9.2.3.tar
k3s ctr images import images/appscode-images-kibana-6.8.23.tar
k3s ctr images import images/appscode-images-kibana-7.17.15.tar
k3s ctr images import images/appscode-images-kibana-7.17.28.tar
k3s ctr images import images/appscode-images-kibana-8.17.10.tar
k3s ctr images import images/appscode-images-kibana-8.17.6.tar
k3s ctr images import images/appscode-images-kibana-8.18.2.tar
k3s ctr images import images/appscode-images-kibana-8.18.8.tar
k3s ctr images import images/appscode-images-kibana-8.19.9.tar
k3s ctr images import images/appscode-images-kibana-8.2.3.tar
k3s ctr images import images/appscode-images-kibana-8.5.3.tar
k3s ctr images import images/appscode-images-kibana-9.0.2.tar
k3s ctr images import images/appscode-images-kibana-9.0.8.tar
k3s ctr images import images/appscode-images-kibana-9.1.4.tar
k3s ctr images import images/appscode-images-kibana-9.1.9.tar
k3s ctr images import images/appscode-images-kibana-9.2.3.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-1.1.0.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-1.3.13.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-1.3.20.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-2.19.2.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-2.5.0.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-3.1.0.tar
k3s ctr images import images/appscode-images-opensearch-1.1.0.tar
k3s ctr images import images/appscode-images-opensearch-1.3.13.tar
k3s ctr images import images/appscode-images-opensearch-1.3.20.tar
k3s ctr images import images/appscode-images-opensearch-2.19.2.tar
k3s ctr images import images/appscode-images-opensearch-2.5.0.tar
k3s ctr images import images/appscode-images-opensearch-3.1.0.tar
k3s ctr images import images/prometheuscommunity-elasticsearch-exporter-v1.10.0.tar
