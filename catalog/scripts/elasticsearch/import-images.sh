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

$CMD push --allow-nondistributable-artifacts --insecure images/amazon-opendistro-for-elasticsearch-1.0.2.tar $IMAGE_REGISTRY/amazon/opendistro-for-elasticsearch:1.0.2
$CMD push --allow-nondistributable-artifacts --insecure images/amazon-opendistro-for-elasticsearch-1.1.0.tar $IMAGE_REGISTRY/amazon/opendistro-for-elasticsearch:1.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/amazon-opendistro-for-elasticsearch-1.10.1.tar $IMAGE_REGISTRY/amazon/opendistro-for-elasticsearch:1.10.1
$CMD push --allow-nondistributable-artifacts --insecure images/amazon-opendistro-for-elasticsearch-1.12.0.tar $IMAGE_REGISTRY/amazon/opendistro-for-elasticsearch:1.12.0
$CMD push --allow-nondistributable-artifacts --insecure images/amazon-opendistro-for-elasticsearch-1.13.2.tar $IMAGE_REGISTRY/amazon/opendistro-for-elasticsearch:1.13.2
$CMD push --allow-nondistributable-artifacts --insecure images/amazon-opendistro-for-elasticsearch-1.2.1.tar $IMAGE_REGISTRY/amazon/opendistro-for-elasticsearch:1.2.1
$CMD push --allow-nondistributable-artifacts --insecure images/amazon-opendistro-for-elasticsearch-1.3.0.tar $IMAGE_REGISTRY/amazon/opendistro-for-elasticsearch:1.3.0
$CMD push --allow-nondistributable-artifacts --insecure images/amazon-opendistro-for-elasticsearch-1.4.0.tar $IMAGE_REGISTRY/amazon/opendistro-for-elasticsearch:1.4.0
$CMD push --allow-nondistributable-artifacts --insecure images/amazon-opendistro-for-elasticsearch-1.6.0.tar $IMAGE_REGISTRY/amazon/opendistro-for-elasticsearch:1.6.0
$CMD push --allow-nondistributable-artifacts --insecure images/amazon-opendistro-for-elasticsearch-1.7.0.tar $IMAGE_REGISTRY/amazon/opendistro-for-elasticsearch:1.7.0
$CMD push --allow-nondistributable-artifacts --insecure images/amazon-opendistro-for-elasticsearch-1.8.0.tar $IMAGE_REGISTRY/amazon/opendistro-for-elasticsearch:1.8.0
$CMD push --allow-nondistributable-artifacts --insecure images/amazon-opendistro-for-elasticsearch-1.9.0.tar $IMAGE_REGISTRY/amazon/opendistro-for-elasticsearch:1.9.0
$CMD push --allow-nondistributable-artifacts --insecure images/floragunncom-sg-elasticsearch-6.8.1-oss-25.1.tar $IMAGE_REGISTRY/floragunncom/sg-elasticsearch:6.8.1-oss-25.1
$CMD push --allow-nondistributable-artifacts --insecure images/floragunncom-sg-elasticsearch-7.0.1-oss-35.0.0.tar $IMAGE_REGISTRY/floragunncom/sg-elasticsearch:7.0.1-oss-35.0.0
$CMD push --allow-nondistributable-artifacts --insecure images/floragunncom-sg-elasticsearch-7.1.1-oss-35.0.0.tar $IMAGE_REGISTRY/floragunncom/sg-elasticsearch:7.1.1-oss-35.0.0
$CMD push --allow-nondistributable-artifacts --insecure images/floragunncom-sg-elasticsearch-7.10.2-oss-49.0.0.tar $IMAGE_REGISTRY/floragunncom/sg-elasticsearch:7.10.2-oss-49.0.0
$CMD push --allow-nondistributable-artifacts --insecure images/floragunncom-sg-elasticsearch-7.14.2-52.3.0.tar $IMAGE_REGISTRY/floragunncom/sg-elasticsearch:7.14.2-52.3.0
$CMD push --allow-nondistributable-artifacts --insecure images/floragunncom-sg-elasticsearch-7.3.2-oss-37.0.0.tar $IMAGE_REGISTRY/floragunncom/sg-elasticsearch:7.3.2-oss-37.0.0
$CMD push --allow-nondistributable-artifacts --insecure images/floragunncom-sg-elasticsearch-7.5.2-oss-40.0.0.tar $IMAGE_REGISTRY/floragunncom/sg-elasticsearch:7.5.2-oss-40.0.0
$CMD push --allow-nondistributable-artifacts --insecure images/floragunncom-sg-elasticsearch-7.8.1-oss-43.0.0.tar $IMAGE_REGISTRY/floragunncom/sg-elasticsearch:7.8.1-oss-43.0.0
$CMD push --allow-nondistributable-artifacts --insecure images/floragunncom-sg-elasticsearch-7.9.3-oss-47.1.0.tar $IMAGE_REGISTRY/floragunncom/sg-elasticsearch:7.9.3-oss-47.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/library-busybox-1.32.0.tar $IMAGE_REGISTRY/busybox:1.32.0
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-6.8.10.tar $IMAGE_REGISTRY/elasticsearch:6.8.10
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-6.8.16.tar $IMAGE_REGISTRY/elasticsearch:6.8.16
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-7.0.1.tar $IMAGE_REGISTRY/elasticsearch:7.0.1
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-7.1.1.tar $IMAGE_REGISTRY/elasticsearch:7.1.1
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-7.12.0.tar $IMAGE_REGISTRY/elasticsearch:7.12.0
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-7.17.3.tar $IMAGE_REGISTRY/elasticsearch:7.17.3
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-7.2.1.tar $IMAGE_REGISTRY/elasticsearch:7.2.1
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-7.3.2.tar $IMAGE_REGISTRY/elasticsearch:7.3.2
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-7.4.2.tar $IMAGE_REGISTRY/elasticsearch:7.4.2
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-7.5.2.tar $IMAGE_REGISTRY/elasticsearch:7.5.2
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-7.6.2.tar $IMAGE_REGISTRY/elasticsearch:7.6.2
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-7.7.1.tar $IMAGE_REGISTRY/elasticsearch:7.7.1
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-7.8.0.tar $IMAGE_REGISTRY/elasticsearch:7.8.0
$CMD push --allow-nondistributable-artifacts --insecure images/library-elasticsearch-7.9.1.tar $IMAGE_REGISTRY/elasticsearch:7.9.1
$CMD push --allow-nondistributable-artifacts --insecure images/library-kibana-7.12.0.tar $IMAGE_REGISTRY/kibana:7.12.0
$CMD push --allow-nondistributable-artifacts --insecure images/library-kibana-7.17.3.tar $IMAGE_REGISTRY/kibana:7.17.3
$CMD push --allow-nondistributable-artifacts --insecure images/library-kibana-7.9.1.tar $IMAGE_REGISTRY/kibana:7.9.1
$CMD push --allow-nondistributable-artifacts --insecure images/opensearchproject-opensearch-dashboards-1.2.0.tar $IMAGE_REGISTRY/opensearchproject/opensearch-dashboards:1.2.0
$CMD push --allow-nondistributable-artifacts --insecure images/opensearchproject-opensearch-dashboards-1.3.2.tar $IMAGE_REGISTRY/opensearchproject/opensearch-dashboards:1.3.2
$CMD push --allow-nondistributable-artifacts --insecure images/opensearchproject-opensearch-1.2.2.tar $IMAGE_REGISTRY/opensearchproject/opensearch:1.2.2
$CMD push --allow-nondistributable-artifacts --insecure images/opensearchproject-opensearch-1.3.2.tar $IMAGE_REGISTRY/opensearchproject/opensearch:1.3.2
$CMD push --allow-nondistributable-artifacts --insecure images/prometheuscommunity-elasticsearch-exporter-v1.10.0.tar $IMAGE_REGISTRY/prometheuscommunity/elasticsearch-exporter:v1.10.0
$CMD push --allow-nondistributable-artifacts --insecure images/prometheuscommunity-elasticsearch-exporter-v1.3.0.tar $IMAGE_REGISTRY/prometheuscommunity/elasticsearch-exporter:v1.3.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-6.8.22.tar $IMAGE_REGISTRY/appscode-images/elastic:6.8.22
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-6.8.23.tar $IMAGE_REGISTRY/appscode-images/elastic:6.8.23
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.13.2.tar $IMAGE_REGISTRY/appscode-images/elastic:7.13.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.13.4.tar $IMAGE_REGISTRY/appscode-images/elastic:7.13.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.14.0.tar $IMAGE_REGISTRY/appscode-images/elastic:7.14.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.14.2.tar $IMAGE_REGISTRY/appscode-images/elastic:7.14.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.16.2.tar $IMAGE_REGISTRY/appscode-images/elastic:7.16.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.16.3.tar $IMAGE_REGISTRY/appscode-images/elastic:7.16.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.17.10.tar $IMAGE_REGISTRY/appscode-images/elastic:7.17.10
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.17.15.tar $IMAGE_REGISTRY/appscode-images/elastic:7.17.15
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.17.23.tar $IMAGE_REGISTRY/appscode-images/elastic:7.17.23
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.17.25.tar $IMAGE_REGISTRY/appscode-images/elastic:7.17.25
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.17.27.tar $IMAGE_REGISTRY/appscode-images/elastic:7.17.27
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-7.17.28.tar $IMAGE_REGISTRY/appscode-images/elastic:7.17.28
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.11.1.tar $IMAGE_REGISTRY/appscode-images/elastic:8.11.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.11.4.tar $IMAGE_REGISTRY/appscode-images/elastic:8.11.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.13.4.tar $IMAGE_REGISTRY/appscode-images/elastic:8.13.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.14.1.tar $IMAGE_REGISTRY/appscode-images/elastic:8.14.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.14.3.tar $IMAGE_REGISTRY/appscode-images/elastic:8.14.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.15.0.tar $IMAGE_REGISTRY/appscode-images/elastic:8.15.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.15.4.tar $IMAGE_REGISTRY/appscode-images/elastic:8.15.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.16.0.tar $IMAGE_REGISTRY/appscode-images/elastic:8.16.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.16.4.tar $IMAGE_REGISTRY/appscode-images/elastic:8.16.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.17.10.tar $IMAGE_REGISTRY/appscode-images/elastic:8.17.10
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.17.2.tar $IMAGE_REGISTRY/appscode-images/elastic:8.17.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.17.6.tar $IMAGE_REGISTRY/appscode-images/elastic:8.17.6
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.18.2.tar $IMAGE_REGISTRY/appscode-images/elastic:8.18.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.18.8.tar $IMAGE_REGISTRY/appscode-images/elastic:8.18.8
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.19.9.tar $IMAGE_REGISTRY/appscode-images/elastic:8.19.9
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.2.0.tar $IMAGE_REGISTRY/appscode-images/elastic:8.2.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.2.3.tar $IMAGE_REGISTRY/appscode-images/elastic:8.2.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.5.2.tar $IMAGE_REGISTRY/appscode-images/elastic:8.5.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.5.3.tar $IMAGE_REGISTRY/appscode-images/elastic:8.5.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.6.2.tar $IMAGE_REGISTRY/appscode-images/elastic:8.6.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.8.0.tar $IMAGE_REGISTRY/appscode-images/elastic:8.8.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-8.8.2.tar $IMAGE_REGISTRY/appscode-images/elastic:8.8.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-9.0.2.tar $IMAGE_REGISTRY/appscode-images/elastic:9.0.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-9.0.8.tar $IMAGE_REGISTRY/appscode-images/elastic:9.0.8
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-9.1.3.tar $IMAGE_REGISTRY/appscode-images/elastic:9.1.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-9.1.4.tar $IMAGE_REGISTRY/appscode-images/elastic:9.1.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-9.1.9.tar $IMAGE_REGISTRY/appscode-images/elastic:9.1.9
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-elastic-9.2.3.tar $IMAGE_REGISTRY/appscode-images/elastic:9.2.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-6.8.22.tar $IMAGE_REGISTRY/appscode-images/kibana:6.8.22
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-6.8.23.tar $IMAGE_REGISTRY/appscode-images/kibana:6.8.23
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.13.2.tar $IMAGE_REGISTRY/appscode-images/kibana:7.13.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.13.4.tar $IMAGE_REGISTRY/appscode-images/kibana:7.13.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.14.0.tar $IMAGE_REGISTRY/appscode-images/kibana:7.14.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.14.2.tar $IMAGE_REGISTRY/appscode-images/kibana:7.14.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.16.2.tar $IMAGE_REGISTRY/appscode-images/kibana:7.16.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.16.3.tar $IMAGE_REGISTRY/appscode-images/kibana:7.16.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.17.10.tar $IMAGE_REGISTRY/appscode-images/kibana:7.17.10
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.17.15.tar $IMAGE_REGISTRY/appscode-images/kibana:7.17.15
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.17.23.tar $IMAGE_REGISTRY/appscode-images/kibana:7.17.23
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.17.25.tar $IMAGE_REGISTRY/appscode-images/kibana:7.17.25
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.17.27.tar $IMAGE_REGISTRY/appscode-images/kibana:7.17.27
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-7.17.28.tar $IMAGE_REGISTRY/appscode-images/kibana:7.17.28
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.11.1.tar $IMAGE_REGISTRY/appscode-images/kibana:8.11.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.11.4.tar $IMAGE_REGISTRY/appscode-images/kibana:8.11.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.13.4.tar $IMAGE_REGISTRY/appscode-images/kibana:8.13.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.14.1.tar $IMAGE_REGISTRY/appscode-images/kibana:8.14.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.14.3.tar $IMAGE_REGISTRY/appscode-images/kibana:8.14.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.15.0.tar $IMAGE_REGISTRY/appscode-images/kibana:8.15.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.15.4.tar $IMAGE_REGISTRY/appscode-images/kibana:8.15.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.16.0.tar $IMAGE_REGISTRY/appscode-images/kibana:8.16.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.16.4.tar $IMAGE_REGISTRY/appscode-images/kibana:8.16.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.17.10.tar $IMAGE_REGISTRY/appscode-images/kibana:8.17.10
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.17.2.tar $IMAGE_REGISTRY/appscode-images/kibana:8.17.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.17.6.tar $IMAGE_REGISTRY/appscode-images/kibana:8.17.6
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.18.2.tar $IMAGE_REGISTRY/appscode-images/kibana:8.18.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.18.8.tar $IMAGE_REGISTRY/appscode-images/kibana:8.18.8
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.19.9.tar $IMAGE_REGISTRY/appscode-images/kibana:8.19.9
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.2.0.tar $IMAGE_REGISTRY/appscode-images/kibana:8.2.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.2.3.tar $IMAGE_REGISTRY/appscode-images/kibana:8.2.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.5.2.tar $IMAGE_REGISTRY/appscode-images/kibana:8.5.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.5.3.tar $IMAGE_REGISTRY/appscode-images/kibana:8.5.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.6.2.tar $IMAGE_REGISTRY/appscode-images/kibana:8.6.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.8.0.tar $IMAGE_REGISTRY/appscode-images/kibana:8.8.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-8.8.2.tar $IMAGE_REGISTRY/appscode-images/kibana:8.8.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-9.0.2.tar $IMAGE_REGISTRY/appscode-images/kibana:9.0.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-9.0.8.tar $IMAGE_REGISTRY/appscode-images/kibana:9.0.8
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-9.1.3.tar $IMAGE_REGISTRY/appscode-images/kibana:9.1.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-9.1.4.tar $IMAGE_REGISTRY/appscode-images/kibana:9.1.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-9.1.9.tar $IMAGE_REGISTRY/appscode-images/kibana:9.1.9
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-kibana-9.2.3.tar $IMAGE_REGISTRY/appscode-images/kibana:9.2.3
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-1.1.0.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-1.2.0.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.2.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-1.3.13.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.3.13
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-1.3.18.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.3.18
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-1.3.19.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.3.19
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-1.3.20.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.3.20
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-2.0.1.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.0.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-2.11.1.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.11.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-2.14.0.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.14.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-2.16.0.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.16.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-2.17.1.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.17.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-2.18.0.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.18.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-2.19.0.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.19.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-2.19.2.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.19.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-2.5.0.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.5.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-2.8.0.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.8.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-3.1.0.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:3.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-dashboards-3.4.0.tar $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:3.4.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-1.1.0.tar $IMAGE_REGISTRY/appscode-images/opensearch:1.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-1.2.4.tar $IMAGE_REGISTRY/appscode-images/opensearch:1.2.4
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-1.3.13.tar $IMAGE_REGISTRY/appscode-images/opensearch:1.3.13
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-1.3.18.tar $IMAGE_REGISTRY/appscode-images/opensearch:1.3.18
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-1.3.19.tar $IMAGE_REGISTRY/appscode-images/opensearch:1.3.19
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-1.3.20.tar $IMAGE_REGISTRY/appscode-images/opensearch:1.3.20
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-2.0.1.tar $IMAGE_REGISTRY/appscode-images/opensearch:2.0.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-2.11.1.tar $IMAGE_REGISTRY/appscode-images/opensearch:2.11.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-2.14.0.tar $IMAGE_REGISTRY/appscode-images/opensearch:2.14.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-2.16.0.tar $IMAGE_REGISTRY/appscode-images/opensearch:2.16.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-2.17.1.tar $IMAGE_REGISTRY/appscode-images/opensearch:2.17.1
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-2.18.0.tar $IMAGE_REGISTRY/appscode-images/opensearch:2.18.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-2.19.0.tar $IMAGE_REGISTRY/appscode-images/opensearch:2.19.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-2.19.2.tar $IMAGE_REGISTRY/appscode-images/opensearch:2.19.2
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-2.5.0.tar $IMAGE_REGISTRY/appscode-images/opensearch:2.5.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-2.8.0.tar $IMAGE_REGISTRY/appscode-images/opensearch:2.8.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-3.1.0.tar $IMAGE_REGISTRY/appscode-images/opensearch:3.1.0
$CMD push --allow-nondistributable-artifacts --insecure images/appscode-images-opensearch-3.4.0.tar $IMAGE_REGISTRY/appscode-images/opensearch:3.4.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-5.6.tar $IMAGE_REGISTRY/kubedb/elasticsearch:5.6
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-5.6-v1.tar $IMAGE_REGISTRY/kubedb/elasticsearch:5.6-v1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-5.6.16-searchguard-v2022.02.22.tar $IMAGE_REGISTRY/kubedb/elasticsearch:5.6.16-searchguard-v2022.02.22
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-5.6.4.tar $IMAGE_REGISTRY/kubedb/elasticsearch:5.6.4
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-5.6.4-v1.tar $IMAGE_REGISTRY/kubedb/elasticsearch:5.6.4-v1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-6.2.tar $IMAGE_REGISTRY/kubedb/elasticsearch:6.2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-6.2-v1.tar $IMAGE_REGISTRY/kubedb/elasticsearch:6.2-v1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-6.2.4.tar $IMAGE_REGISTRY/kubedb/elasticsearch:6.2.4
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-6.2.4-v1.tar $IMAGE_REGISTRY/kubedb/elasticsearch:6.2.4-v1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-6.3-v1.tar $IMAGE_REGISTRY/kubedb/elasticsearch:6.3-v1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-6.3.0-v1.tar $IMAGE_REGISTRY/kubedb/elasticsearch:6.3.0-v1
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-6.4.tar $IMAGE_REGISTRY/kubedb/elasticsearch:6.4
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-6.4.0.tar $IMAGE_REGISTRY/kubedb/elasticsearch:6.4.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-6.5.tar $IMAGE_REGISTRY/kubedb/elasticsearch:6.5
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-6.5.3.tar $IMAGE_REGISTRY/kubedb/elasticsearch:6.5.3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-6.8.tar $IMAGE_REGISTRY/kubedb/elasticsearch:6.8
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-7.12.0-xpack-v2021.08.23.tar $IMAGE_REGISTRY/kubedb/elasticsearch:7.12.0-xpack-v2021.08.23
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-7.13.2-xpack-v2021.08.23.tar $IMAGE_REGISTRY/kubedb/elasticsearch:7.13.2-xpack-v2021.08.23
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-7.14.0-xpack-v2021.08.23.tar $IMAGE_REGISTRY/kubedb/elasticsearch:7.14.0-xpack-v2021.08.23
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-7.16.2-xpack-v2021.12.24.tar $IMAGE_REGISTRY/kubedb/elasticsearch:7.16.2-xpack-v2021.12.24
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-7.2.tar $IMAGE_REGISTRY/kubedb/elasticsearch:7.2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-7.2.0.tar $IMAGE_REGISTRY/kubedb/elasticsearch:7.2.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-7.3.tar $IMAGE_REGISTRY/kubedb/elasticsearch:7.3
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-7.3.2.tar $IMAGE_REGISTRY/kubedb/elasticsearch:7.3.2
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-7.9.1-xpack-v2021.08.23.tar $IMAGE_REGISTRY/kubedb/elasticsearch:7.9.1-xpack-v2021.08.23
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch-8.2.0-xpack-v2022.05.24.tar $IMAGE_REGISTRY/kubedb/elasticsearch:8.2.0-xpack-v2022.05.24
$CMD push --allow-nondistributable-artifacts --insecure images/kubedb-elasticsearch_exporter-1.0.2.tar $IMAGE_REGISTRY/kubedb/elasticsearch_exporter:1.0.2
