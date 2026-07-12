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

$CMD pull --allow-nondistributable-artifacts --insecure docker.io/amazon/opendistro-for-elasticsearch:1.0.2 images/amazon-opendistro-for-elasticsearch-1.0.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/amazon/opendistro-for-elasticsearch:1.1.0 images/amazon-opendistro-for-elasticsearch-1.1.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/amazon/opendistro-for-elasticsearch:1.10.1 images/amazon-opendistro-for-elasticsearch-1.10.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/amazon/opendistro-for-elasticsearch:1.12.0 images/amazon-opendistro-for-elasticsearch-1.12.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/amazon/opendistro-for-elasticsearch:1.13.2 images/amazon-opendistro-for-elasticsearch-1.13.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/amazon/opendistro-for-elasticsearch:1.2.1 images/amazon-opendistro-for-elasticsearch-1.2.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/amazon/opendistro-for-elasticsearch:1.3.0 images/amazon-opendistro-for-elasticsearch-1.3.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/amazon/opendistro-for-elasticsearch:1.4.0 images/amazon-opendistro-for-elasticsearch-1.4.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/amazon/opendistro-for-elasticsearch:1.6.0 images/amazon-opendistro-for-elasticsearch-1.6.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/amazon/opendistro-for-elasticsearch:1.7.0 images/amazon-opendistro-for-elasticsearch-1.7.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/amazon/opendistro-for-elasticsearch:1.8.0 images/amazon-opendistro-for-elasticsearch-1.8.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/amazon/opendistro-for-elasticsearch:1.9.0 images/amazon-opendistro-for-elasticsearch-1.9.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/floragunncom/sg-elasticsearch:6.8.1-oss-25.1 images/floragunncom-sg-elasticsearch-6.8.1-oss-25.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/floragunncom/sg-elasticsearch:7.0.1-oss-35.0.0 images/floragunncom-sg-elasticsearch-7.0.1-oss-35.0.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/floragunncom/sg-elasticsearch:7.1.1-oss-35.0.0 images/floragunncom-sg-elasticsearch-7.1.1-oss-35.0.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/floragunncom/sg-elasticsearch:7.10.2-oss-49.0.0 images/floragunncom-sg-elasticsearch-7.10.2-oss-49.0.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/floragunncom/sg-elasticsearch:7.14.2-52.3.0 images/floragunncom-sg-elasticsearch-7.14.2-52.3.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/floragunncom/sg-elasticsearch:7.3.2-oss-37.0.0 images/floragunncom-sg-elasticsearch-7.3.2-oss-37.0.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/floragunncom/sg-elasticsearch:7.5.2-oss-40.0.0 images/floragunncom-sg-elasticsearch-7.5.2-oss-40.0.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/floragunncom/sg-elasticsearch:7.8.1-oss-43.0.0 images/floragunncom-sg-elasticsearch-7.8.1-oss-43.0.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/floragunncom/sg-elasticsearch:7.9.3-oss-47.1.0 images/floragunncom-sg-elasticsearch-7.9.3-oss-47.1.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/busybox:1.32.0 images/library-busybox-1.32.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:6.8.10 images/library-elasticsearch-6.8.10.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:6.8.16 images/library-elasticsearch-6.8.16.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:7.0.1 images/library-elasticsearch-7.0.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:7.1.1 images/library-elasticsearch-7.1.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:7.12.0 images/library-elasticsearch-7.12.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:7.17.3 images/library-elasticsearch-7.17.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:7.2.1 images/library-elasticsearch-7.2.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:7.3.2 images/library-elasticsearch-7.3.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:7.4.2 images/library-elasticsearch-7.4.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:7.5.2 images/library-elasticsearch-7.5.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:7.6.2 images/library-elasticsearch-7.6.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:7.7.1 images/library-elasticsearch-7.7.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:7.8.0 images/library-elasticsearch-7.8.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/elasticsearch:7.9.1 images/library-elasticsearch-7.9.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/kibana:7.12.0 images/library-kibana-7.12.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/kibana:7.17.3 images/library-kibana-7.17.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/library/kibana:7.9.1 images/library-kibana-7.9.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/opensearchproject/opensearch-dashboards:1.2.0 images/opensearchproject-opensearch-dashboards-1.2.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/opensearchproject/opensearch-dashboards:1.3.2 images/opensearchproject-opensearch-dashboards-1.3.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/opensearchproject/opensearch:1.2.2 images/opensearchproject-opensearch-1.2.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/opensearchproject/opensearch:1.3.2 images/opensearchproject-opensearch-1.3.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/prometheuscommunity/elasticsearch-exporter:v1.10.0 images/prometheuscommunity-elasticsearch-exporter-v1.10.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure docker.io/prometheuscommunity/elasticsearch-exporter:v1.3.0 images/prometheuscommunity-elasticsearch-exporter-v1.3.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:6.8.22 images/appscode-images-elastic-6.8.22.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:6.8.23 images/appscode-images-elastic-6.8.23.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.13.2 images/appscode-images-elastic-7.13.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.13.4 images/appscode-images-elastic-7.13.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.14.0 images/appscode-images-elastic-7.14.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.14.2 images/appscode-images-elastic-7.14.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.16.2 images/appscode-images-elastic-7.16.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.16.3 images/appscode-images-elastic-7.16.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.17.10 images/appscode-images-elastic-7.17.10.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.17.15 images/appscode-images-elastic-7.17.15.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.17.23 images/appscode-images-elastic-7.17.23.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.17.25 images/appscode-images-elastic-7.17.25.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.17.27 images/appscode-images-elastic-7.17.27.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.17.28 images/appscode-images-elastic-7.17.28.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.11.1 images/appscode-images-elastic-8.11.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.11.4 images/appscode-images-elastic-8.11.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.13.4 images/appscode-images-elastic-8.13.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.14.1 images/appscode-images-elastic-8.14.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.14.3 images/appscode-images-elastic-8.14.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.15.0 images/appscode-images-elastic-8.15.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.15.4 images/appscode-images-elastic-8.15.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.16.0 images/appscode-images-elastic-8.16.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.16.4 images/appscode-images-elastic-8.16.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.17.10 images/appscode-images-elastic-8.17.10.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.17.2 images/appscode-images-elastic-8.17.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.17.6 images/appscode-images-elastic-8.17.6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.18.2 images/appscode-images-elastic-8.18.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.18.8 images/appscode-images-elastic-8.18.8.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.19.9 images/appscode-images-elastic-8.19.9.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.2.0 images/appscode-images-elastic-8.2.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.2.3 images/appscode-images-elastic-8.2.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.5.2 images/appscode-images-elastic-8.5.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.5.3 images/appscode-images-elastic-8.5.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.6.2 images/appscode-images-elastic-8.6.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.8.0 images/appscode-images-elastic-8.8.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.8.2 images/appscode-images-elastic-8.8.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:9.0.2 images/appscode-images-elastic-9.0.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:9.0.8 images/appscode-images-elastic-9.0.8.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:9.1.3 images/appscode-images-elastic-9.1.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:9.1.4 images/appscode-images-elastic-9.1.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:9.1.9 images/appscode-images-elastic-9.1.9.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:9.2.3 images/appscode-images-elastic-9.2.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:6.8.22 images/appscode-images-kibana-6.8.22.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:6.8.23 images/appscode-images-kibana-6.8.23.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.13.2 images/appscode-images-kibana-7.13.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.13.4 images/appscode-images-kibana-7.13.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.14.0 images/appscode-images-kibana-7.14.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.14.2 images/appscode-images-kibana-7.14.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.16.2 images/appscode-images-kibana-7.16.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.16.3 images/appscode-images-kibana-7.16.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.17.10 images/appscode-images-kibana-7.17.10.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.17.15 images/appscode-images-kibana-7.17.15.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.17.23 images/appscode-images-kibana-7.17.23.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.17.25 images/appscode-images-kibana-7.17.25.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.17.27 images/appscode-images-kibana-7.17.27.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.17.28 images/appscode-images-kibana-7.17.28.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.11.1 images/appscode-images-kibana-8.11.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.11.4 images/appscode-images-kibana-8.11.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.13.4 images/appscode-images-kibana-8.13.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.14.1 images/appscode-images-kibana-8.14.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.14.3 images/appscode-images-kibana-8.14.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.15.0 images/appscode-images-kibana-8.15.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.15.4 images/appscode-images-kibana-8.15.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.16.0 images/appscode-images-kibana-8.16.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.16.4 images/appscode-images-kibana-8.16.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.17.10 images/appscode-images-kibana-8.17.10.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.17.2 images/appscode-images-kibana-8.17.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.17.6 images/appscode-images-kibana-8.17.6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.18.2 images/appscode-images-kibana-8.18.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.18.8 images/appscode-images-kibana-8.18.8.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.19.9 images/appscode-images-kibana-8.19.9.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.2.0 images/appscode-images-kibana-8.2.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.2.3 images/appscode-images-kibana-8.2.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.5.2 images/appscode-images-kibana-8.5.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.5.3 images/appscode-images-kibana-8.5.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.6.2 images/appscode-images-kibana-8.6.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.8.0 images/appscode-images-kibana-8.8.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.8.2 images/appscode-images-kibana-8.8.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:9.0.2 images/appscode-images-kibana-9.0.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:9.0.8 images/appscode-images-kibana-9.0.8.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:9.1.3 images/appscode-images-kibana-9.1.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:9.1.4 images/appscode-images-kibana-9.1.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:9.1.9 images/appscode-images-kibana-9.1.9.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:9.2.3 images/appscode-images-kibana-9.2.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:1.1.0 images/appscode-images-opensearch-dashboards-1.1.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:1.2.0 images/appscode-images-opensearch-dashboards-1.2.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:1.3.13 images/appscode-images-opensearch-dashboards-1.3.13.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:1.3.18 images/appscode-images-opensearch-dashboards-1.3.18.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:1.3.19 images/appscode-images-opensearch-dashboards-1.3.19.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:1.3.20 images/appscode-images-opensearch-dashboards-1.3.20.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.0.1 images/appscode-images-opensearch-dashboards-2.0.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.11.1 images/appscode-images-opensearch-dashboards-2.11.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.14.0 images/appscode-images-opensearch-dashboards-2.14.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.16.0 images/appscode-images-opensearch-dashboards-2.16.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.17.1 images/appscode-images-opensearch-dashboards-2.17.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.18.0 images/appscode-images-opensearch-dashboards-2.18.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.19.0 images/appscode-images-opensearch-dashboards-2.19.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.19.2 images/appscode-images-opensearch-dashboards-2.19.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.5.0 images/appscode-images-opensearch-dashboards-2.5.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.8.0 images/appscode-images-opensearch-dashboards-2.8.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:3.1.0 images/appscode-images-opensearch-dashboards-3.1.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:3.4.0 images/appscode-images-opensearch-dashboards-3.4.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:1.1.0 images/appscode-images-opensearch-1.1.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:1.2.4 images/appscode-images-opensearch-1.2.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:1.3.13 images/appscode-images-opensearch-1.3.13.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:1.3.18 images/appscode-images-opensearch-1.3.18.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:1.3.19 images/appscode-images-opensearch-1.3.19.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:1.3.20 images/appscode-images-opensearch-1.3.20.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.0.1 images/appscode-images-opensearch-2.0.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.11.1 images/appscode-images-opensearch-2.11.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.14.0 images/appscode-images-opensearch-2.14.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.16.0 images/appscode-images-opensearch-2.16.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.17.1 images/appscode-images-opensearch-2.17.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.18.0 images/appscode-images-opensearch-2.18.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.19.0 images/appscode-images-opensearch-2.19.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.19.2 images/appscode-images-opensearch-2.19.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.5.0 images/appscode-images-opensearch-2.5.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.8.0 images/appscode-images-opensearch-2.8.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:3.1.0 images/appscode-images-opensearch-3.1.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:3.4.0 images/appscode-images-opensearch-3.4.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:5.6 images/kubedb-elasticsearch-5.6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:5.6-v1 images/kubedb-elasticsearch-5.6-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:5.6.16-searchguard-v2022.02.22 images/kubedb-elasticsearch-5.6.16-searchguard-v2022.02.22.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:5.6.4 images/kubedb-elasticsearch-5.6.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:5.6.4-v1 images/kubedb-elasticsearch-5.6.4-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:6.2 images/kubedb-elasticsearch-6.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:6.2-v1 images/kubedb-elasticsearch-6.2-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:6.2.4 images/kubedb-elasticsearch-6.2.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:6.2.4-v1 images/kubedb-elasticsearch-6.2.4-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:6.3-v1 images/kubedb-elasticsearch-6.3-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:6.3.0-v1 images/kubedb-elasticsearch-6.3.0-v1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:6.4 images/kubedb-elasticsearch-6.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:6.4.0 images/kubedb-elasticsearch-6.4.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:6.5 images/kubedb-elasticsearch-6.5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:6.5.3 images/kubedb-elasticsearch-6.5.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:6.8 images/kubedb-elasticsearch-6.8.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:7.12.0-xpack-v2021.08.23 images/kubedb-elasticsearch-7.12.0-xpack-v2021.08.23.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:7.13.2-xpack-v2021.08.23 images/kubedb-elasticsearch-7.13.2-xpack-v2021.08.23.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:7.14.0-xpack-v2021.08.23 images/kubedb-elasticsearch-7.14.0-xpack-v2021.08.23.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:7.16.2-xpack-v2021.12.24 images/kubedb-elasticsearch-7.16.2-xpack-v2021.12.24.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:7.2 images/kubedb-elasticsearch-7.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:7.2.0 images/kubedb-elasticsearch-7.2.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:7.3 images/kubedb-elasticsearch-7.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:7.3.2 images/kubedb-elasticsearch-7.3.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:7.9.1-xpack-v2021.08.23 images/kubedb-elasticsearch-7.9.1-xpack-v2021.08.23.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch:8.2.0-xpack-v2022.05.24 images/kubedb-elasticsearch-8.2.0-xpack-v2022.05.24.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubedb/elasticsearch_exporter:1.0.2 images/kubedb-elasticsearch_exporter-1.0.2.tar

tar -czvf images.tar.gz images
