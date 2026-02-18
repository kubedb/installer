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

$CMD push --allow-nondistributable-artifacts --insecure images/floragunncom-sg-elasticsearch-7.9.3-oss-47.1.0.tar $IMAGE_REGISTRY/floragunncom/sg-elasticsearch:7.9.3-oss-47.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-6.8.23.tar $IMAGE_REGISTRY/appscode-images/elastic:6.8.23
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.17.15.tar $IMAGE_REGISTRY/appscode-images/elastic:7.17.15
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.17.28.tar $IMAGE_REGISTRY/appscode-images/elastic:7.17.28
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.17.10.tar $IMAGE_REGISTRY/appscode-images/elastic:8.17.10
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.17.6.tar $IMAGE_REGISTRY/appscode-images/elastic:8.17.6
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.18.2.tar $IMAGE_REGISTRY/appscode-images/elastic:8.18.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.18.8.tar $IMAGE_REGISTRY/appscode-images/elastic:8.18.8
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.19.9.tar $IMAGE_REGISTRY/appscode-images/elastic:8.19.9
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.2.3.tar $IMAGE_REGISTRY/appscode-images/elastic:8.2.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.5.3.tar $IMAGE_REGISTRY/appscode-images/elastic:8.5.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-9.0.2.tar $IMAGE_REGISTRY/appscode-images/elastic:9.0.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-9.0.8.tar $IMAGE_REGISTRY/appscode-images/elastic:9.0.8
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-9.1.4.tar $IMAGE_REGISTRY/appscode-images/elastic:9.1.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-9.1.9.tar $IMAGE_REGISTRY/appscode-images/elastic:9.1.9
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-9.2.3.tar $IMAGE_REGISTRY/appscode-images/elastic:9.2.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-6.8.23.tar $IMAGE_REGISTRY/appscode-images/kibana:6.8.23
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.17.15.tar $IMAGE_REGISTRY/appscode-images/kibana:7.17.15
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.17.28.tar $IMAGE_REGISTRY/appscode-images/kibana:7.17.28
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.17.10.tar $IMAGE_REGISTRY/appscode-images/kibana:8.17.10
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.17.6.tar $IMAGE_REGISTRY/appscode-images/kibana:8.17.6
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.18.2.tar $IMAGE_REGISTRY/appscode-images/kibana:8.18.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.18.8.tar $IMAGE_REGISTRY/appscode-images/kibana:8.18.8
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.19.9.tar $IMAGE_REGISTRY/appscode-images/kibana:8.19.9
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.2.3.tar $IMAGE_REGISTRY/appscode-images/kibana:8.2.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.5.3.tar $IMAGE_REGISTRY/appscode-images/kibana:8.5.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-9.0.2.tar $IMAGE_REGISTRY/appscode-images/kibana:9.0.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-9.0.8.tar $IMAGE_REGISTRY/appscode-images/kibana:9.0.8
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-9.1.4.tar $IMAGE_REGISTRY/appscode-images/kibana:9.1.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-9.1.9.tar $IMAGE_REGISTRY/appscode-images/kibana:9.1.9
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-9.2.3.tar $IMAGE_REGISTRY/appscode-images/kibana:9.2.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-1.3.13.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.3.13
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-1.3.20.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.3.20
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-2.19.2.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.19.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-2.5.0.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.5.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-3.1.0.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:3.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-3.4.0.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:3.4.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-1.3.13.tar $IMAGE_REGISTRY/appscode-images/opensearch:1.3.13
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-1.3.20.tar $IMAGE_REGISTRY/appscode-images/opensearch:1.3.20
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-2.19.2.tar $IMAGE_REGISTRY/appscode-images/opensearch:2.19.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-2.5.0.tar $IMAGE_REGISTRY/appscode-images/opensearch:2.5.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-3.1.0.tar $IMAGE_REGISTRY/appscode-images/opensearch:3.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-3.4.0.tar $IMAGE_REGISTRY/appscode-images/opensearch:3.4.0
$CMD push --allow-nondistributable-artifacts --insecure images/prometheuscommunity-elasticsearch-exporter-v1.10.0.tar $IMAGE_REGISTRY/prometheuscommunity/elasticsearch-exporter:v1.10.0
