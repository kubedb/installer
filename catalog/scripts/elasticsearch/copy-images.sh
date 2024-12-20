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
mv /tmp/crane .

CMD="./crane"

$CMD cp --allow-nondistributable-artifacts --insecure floragunncom/sg-elasticsearch:7.9.3-oss-47.1.0 $IMAGE_REGISTRY/floragunncom/sg-elasticsearch:7.9.3-oss-47.1.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:6.8.23 $IMAGE_REGISTRY/appscode-images/elastic:6.8.23
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.13.4 $IMAGE_REGISTRY/appscode-images/elastic:7.13.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.14.2 $IMAGE_REGISTRY/appscode-images/elastic:7.14.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.16.3 $IMAGE_REGISTRY/appscode-images/elastic:7.16.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.17.15 $IMAGE_REGISTRY/appscode-images/elastic:7.17.15
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.17.23 $IMAGE_REGISTRY/appscode-images/elastic:7.17.23
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:7.17.25 $IMAGE_REGISTRY/appscode-images/elastic:7.17.25
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.11.1 $IMAGE_REGISTRY/appscode-images/elastic:8.11.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.11.4 $IMAGE_REGISTRY/appscode-images/elastic:8.11.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.13.4 $IMAGE_REGISTRY/appscode-images/elastic:8.13.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.14.1 $IMAGE_REGISTRY/appscode-images/elastic:8.14.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.14.3 $IMAGE_REGISTRY/appscode-images/elastic:8.14.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.15.0 $IMAGE_REGISTRY/appscode-images/elastic:8.15.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.15.4 $IMAGE_REGISTRY/appscode-images/elastic:8.15.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.16.0 $IMAGE_REGISTRY/appscode-images/elastic:8.16.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.2.3 $IMAGE_REGISTRY/appscode-images/elastic:8.2.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.5.3 $IMAGE_REGISTRY/appscode-images/elastic:8.5.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.6.2 $IMAGE_REGISTRY/appscode-images/elastic:8.6.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/elastic:8.8.2 $IMAGE_REGISTRY/appscode-images/elastic:8.8.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:6.8.23 $IMAGE_REGISTRY/appscode-images/kibana:6.8.23
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.13.4 $IMAGE_REGISTRY/appscode-images/kibana:7.13.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.14.2 $IMAGE_REGISTRY/appscode-images/kibana:7.14.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.16.3 $IMAGE_REGISTRY/appscode-images/kibana:7.16.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.17.15 $IMAGE_REGISTRY/appscode-images/kibana:7.17.15
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.17.23 $IMAGE_REGISTRY/appscode-images/kibana:7.17.23
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:7.17.25 $IMAGE_REGISTRY/appscode-images/kibana:7.17.25
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.11.1 $IMAGE_REGISTRY/appscode-images/kibana:8.11.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.11.4 $IMAGE_REGISTRY/appscode-images/kibana:8.11.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.13.4 $IMAGE_REGISTRY/appscode-images/kibana:8.13.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.14.1 $IMAGE_REGISTRY/appscode-images/kibana:8.14.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.14.3 $IMAGE_REGISTRY/appscode-images/kibana:8.14.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.15.0 $IMAGE_REGISTRY/appscode-images/kibana:8.15.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.15.4 $IMAGE_REGISTRY/appscode-images/kibana:8.15.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.16.0 $IMAGE_REGISTRY/appscode-images/kibana:8.16.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.2.3 $IMAGE_REGISTRY/appscode-images/kibana:8.2.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.5.3 $IMAGE_REGISTRY/appscode-images/kibana:8.5.3
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.6.2 $IMAGE_REGISTRY/appscode-images/kibana:8.6.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/kibana:8.8.2 $IMAGE_REGISTRY/appscode-images/kibana:8.8.2
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:1.1.0 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.1.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:1.2.0 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.2.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:1.3.13 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.3.13
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:1.3.18 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.3.18
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:1.3.19 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.3.19
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.0.1 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.0.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.11.1 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.11.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.14.0 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.14.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.16.0 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.16.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.17.1 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.17.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.18.0 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.18.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.5.0 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.5.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch-dashboards:2.8.0 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.8.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:1.1.0 $IMAGE_REGISTRY/appscode-images/opensearch:1.1.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:1.2.4 $IMAGE_REGISTRY/appscode-images/opensearch:1.2.4
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:1.3.13 $IMAGE_REGISTRY/appscode-images/opensearch:1.3.13
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:1.3.18 $IMAGE_REGISTRY/appscode-images/opensearch:1.3.18
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:1.3.19 $IMAGE_REGISTRY/appscode-images/opensearch:1.3.19
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.0.1 $IMAGE_REGISTRY/appscode-images/opensearch:2.0.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.11.1 $IMAGE_REGISTRY/appscode-images/opensearch:2.11.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.14.0 $IMAGE_REGISTRY/appscode-images/opensearch:2.14.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.16.0 $IMAGE_REGISTRY/appscode-images/opensearch:2.16.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.17.1 $IMAGE_REGISTRY/appscode-images/opensearch:2.17.1
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.18.0 $IMAGE_REGISTRY/appscode-images/opensearch:2.18.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.5.0 $IMAGE_REGISTRY/appscode-images/opensearch:2.5.0
$CMD cp --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/opensearch:2.8.0 $IMAGE_REGISTRY/appscode-images/opensearch:2.8.0
$CMD cp --allow-nondistributable-artifacts --insecure prometheuscommunity/elasticsearch-exporter:v1.7.0 $IMAGE_REGISTRY/prometheuscommunity/elasticsearch-exporter:v1.7.0
