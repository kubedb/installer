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
k3s ctr images import images/appscode-images-elastic-7.13.4.tar
k3s ctr images import images/appscode-images-elastic-7.14.2.tar
k3s ctr images import images/appscode-images-elastic-7.16.3.tar
k3s ctr images import images/appscode-images-elastic-7.17.15.tar
k3s ctr images import images/appscode-images-elastic-7.17.23.tar
k3s ctr images import images/appscode-images-elastic-8.11.1.tar
k3s ctr images import images/appscode-images-elastic-8.11.4.tar
k3s ctr images import images/appscode-images-elastic-8.13.4.tar
k3s ctr images import images/appscode-images-elastic-8.14.1.tar
k3s ctr images import images/appscode-images-elastic-8.14.3.tar
k3s ctr images import images/appscode-images-elastic-8.15.0.tar
k3s ctr images import images/appscode-images-elastic-8.2.3.tar
k3s ctr images import images/appscode-images-elastic-8.5.3.tar
k3s ctr images import images/appscode-images-elastic-8.6.2.tar
k3s ctr images import images/appscode-images-elastic-8.8.2.tar
k3s ctr images import images/appscode-images-kibana-6.8.23.tar
k3s ctr images import images/appscode-images-kibana-7.13.4.tar
k3s ctr images import images/appscode-images-kibana-7.14.2.tar
k3s ctr images import images/appscode-images-kibana-7.16.3.tar
k3s ctr images import images/appscode-images-kibana-7.17.15.tar
k3s ctr images import images/appscode-images-kibana-7.17.23.tar
k3s ctr images import images/appscode-images-kibana-8.11.1.tar
k3s ctr images import images/appscode-images-kibana-8.11.4.tar
k3s ctr images import images/appscode-images-kibana-8.13.4.tar
k3s ctr images import images/appscode-images-kibana-8.14.1.tar
k3s ctr images import images/appscode-images-kibana-8.14.3.tar
k3s ctr images import images/appscode-images-kibana-8.15.0.tar
k3s ctr images import images/appscode-images-kibana-8.2.3.tar
k3s ctr images import images/appscode-images-kibana-8.5.3.tar
k3s ctr images import images/appscode-images-kibana-8.6.2.tar
k3s ctr images import images/appscode-images-kibana-8.8.2.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-1.1.0.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-1.2.0.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-1.3.13.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-1.3.18.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-2.0.1.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-2.11.1.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-2.14.0.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-2.16.0.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-2.5.0.tar
k3s ctr images import images/appscode-images-opensearch-dashboards-2.8.0.tar
k3s ctr images import images/appscode-images-opensearch-1.1.0.tar
k3s ctr images import images/appscode-images-opensearch-1.2.4.tar
k3s ctr images import images/appscode-images-opensearch-1.3.13.tar
k3s ctr images import images/appscode-images-opensearch-1.3.18.tar
k3s ctr images import images/appscode-images-opensearch-2.0.1.tar
k3s ctr images import images/appscode-images-opensearch-2.11.1.tar
k3s ctr images import images/appscode-images-opensearch-2.14.0.tar
k3s ctr images import images/appscode-images-opensearch-2.16.0.tar
k3s ctr images import images/appscode-images-opensearch-2.5.0.tar
k3s ctr images import images/appscode-images-opensearch-2.8.0.tar
k3s ctr images import images/prometheuscommunity-elasticsearch-exporter-v1.7.0.tar