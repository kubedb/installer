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

$CMD cp apache/druid:25.0.0 $IMAGE_REGISTRY/apache/druid:25.0.0
$CMD cp apicurio/apicurio-registry-kafkasql:2.5.11.Final $IMAGE_REGISTRY/apicurio/apicurio-registry-kafkasql:2.5.11.Final
$CMD cp apicurio/apicurio-registry-mem:2.5.11.Final $IMAGE_REGISTRY/apicurio/apicurio-registry-mem:2.5.11.Final
$CMD cp clickhouse/clickhouse-keeper:24.4.1 $IMAGE_REGISTRY/clickhouse/clickhouse-keeper:24.4.1
$CMD cp clickhouse/clickhouse-server:24.4.1 $IMAGE_REGISTRY/clickhouse/clickhouse-server:24.4.1
$CMD cp floragunncom/sg-elasticsearch:7.9.3-oss-47.1.0 $IMAGE_REGISTRY/floragunncom/sg-elasticsearch:7.9.3-oss-47.1.0
$CMD cp ghcr.io/aiven-open/karapace:3.15.0 $IMAGE_REGISTRY/aiven-open/karapace:3.15.0
$CMD cp ghcr.io/appscode-images/cassandra:4.1.6 $IMAGE_REGISTRY/appscode-images/cassandra:4.1.6
$CMD cp ghcr.io/appscode-images/cassandra:5.0.0 $IMAGE_REGISTRY/appscode-images/cassandra:5.0.0
$CMD cp ghcr.io/appscode-images/druid:28.0.1 $IMAGE_REGISTRY/appscode-images/druid:28.0.1
$CMD cp ghcr.io/appscode-images/druid:30.0.0 $IMAGE_REGISTRY/appscode-images/druid:30.0.0
$CMD cp ghcr.io/appscode-images/elastic:6.8.23 $IMAGE_REGISTRY/appscode-images/elastic:6.8.23
$CMD cp ghcr.io/appscode-images/elastic:7.13.4 $IMAGE_REGISTRY/appscode-images/elastic:7.13.4
$CMD cp ghcr.io/appscode-images/elastic:7.14.2 $IMAGE_REGISTRY/appscode-images/elastic:7.14.2
$CMD cp ghcr.io/appscode-images/elastic:7.16.3 $IMAGE_REGISTRY/appscode-images/elastic:7.16.3
$CMD cp ghcr.io/appscode-images/elastic:7.17.15 $IMAGE_REGISTRY/appscode-images/elastic:7.17.15
$CMD cp ghcr.io/appscode-images/elastic:7.17.23 $IMAGE_REGISTRY/appscode-images/elastic:7.17.23
$CMD cp ghcr.io/appscode-images/elastic:8.11.1 $IMAGE_REGISTRY/appscode-images/elastic:8.11.1
$CMD cp ghcr.io/appscode-images/elastic:8.11.4 $IMAGE_REGISTRY/appscode-images/elastic:8.11.4
$CMD cp ghcr.io/appscode-images/elastic:8.13.4 $IMAGE_REGISTRY/appscode-images/elastic:8.13.4
$CMD cp ghcr.io/appscode-images/elastic:8.14.1 $IMAGE_REGISTRY/appscode-images/elastic:8.14.1
$CMD cp ghcr.io/appscode-images/elastic:8.14.3 $IMAGE_REGISTRY/appscode-images/elastic:8.14.3
$CMD cp ghcr.io/appscode-images/elastic:8.15.0 $IMAGE_REGISTRY/appscode-images/elastic:8.15.0
$CMD cp ghcr.io/appscode-images/elastic:8.2.3 $IMAGE_REGISTRY/appscode-images/elastic:8.2.3
$CMD cp ghcr.io/appscode-images/elastic:8.5.3 $IMAGE_REGISTRY/appscode-images/elastic:8.5.3
$CMD cp ghcr.io/appscode-images/elastic:8.6.2 $IMAGE_REGISTRY/appscode-images/elastic:8.6.2
$CMD cp ghcr.io/appscode-images/elastic:8.8.2 $IMAGE_REGISTRY/appscode-images/elastic:8.8.2
$CMD cp ghcr.io/appscode-images/ferretdb:1.18.0 $IMAGE_REGISTRY/appscode-images/ferretdb:1.18.0
$CMD cp ghcr.io/appscode-images/ferretdb:1.23.0 $IMAGE_REGISTRY/appscode-images/ferretdb:1.23.0
$CMD cp ghcr.io/appscode-images/kafka-connect-cluster:3.3.2 $IMAGE_REGISTRY/appscode-images/kafka-connect-cluster:3.3.2
$CMD cp ghcr.io/appscode-images/kafka-connect-cluster:3.4.1 $IMAGE_REGISTRY/appscode-images/kafka-connect-cluster:3.4.1
$CMD cp ghcr.io/appscode-images/kafka-connect-cluster:3.5.1 $IMAGE_REGISTRY/appscode-images/kafka-connect-cluster:3.5.1
$CMD cp ghcr.io/appscode-images/kafka-connect-cluster:3.5.2 $IMAGE_REGISTRY/appscode-images/kafka-connect-cluster:3.5.2
$CMD cp ghcr.io/appscode-images/kafka-connect-cluster:3.6.0 $IMAGE_REGISTRY/appscode-images/kafka-connect-cluster:3.6.0
$CMD cp ghcr.io/appscode-images/kafka-connect-cluster:3.6.1 $IMAGE_REGISTRY/appscode-images/kafka-connect-cluster:3.6.1
$CMD cp ghcr.io/appscode-images/kafka-connector-gcs:0.13.0 $IMAGE_REGISTRY/appscode-images/kafka-connector-gcs:0.13.0
$CMD cp ghcr.io/appscode-images/kafka-connector-jdbc:2.6.1.final $IMAGE_REGISTRY/appscode-images/kafka-connector-jdbc:2.6.1.final
$CMD cp ghcr.io/appscode-images/kafka-connector-mongodb:1.11.0 $IMAGE_REGISTRY/appscode-images/kafka-connector-mongodb:1.11.0
$CMD cp ghcr.io/appscode-images/kafka-connector-mysql:2.4.2.final $IMAGE_REGISTRY/appscode-images/kafka-connector-mysql:2.4.2.final
$CMD cp ghcr.io/appscode-images/kafka-connector-postgres:2.4.2.final $IMAGE_REGISTRY/appscode-images/kafka-connector-postgres:2.4.2.final
$CMD cp ghcr.io/appscode-images/kafka-connector-s3:2.15.0 $IMAGE_REGISTRY/appscode-images/kafka-connector-s3:2.15.0
$CMD cp ghcr.io/appscode-images/kafka-cruise-control:3.3.2 $IMAGE_REGISTRY/appscode-images/kafka-cruise-control:3.3.2
$CMD cp ghcr.io/appscode-images/kafka-cruise-control:3.4.1 $IMAGE_REGISTRY/appscode-images/kafka-cruise-control:3.4.1
$CMD cp ghcr.io/appscode-images/kafka-cruise-control:3.5.1 $IMAGE_REGISTRY/appscode-images/kafka-cruise-control:3.5.1
$CMD cp ghcr.io/appscode-images/kafka-cruise-control:3.5.2 $IMAGE_REGISTRY/appscode-images/kafka-cruise-control:3.5.2
$CMD cp ghcr.io/appscode-images/kafka-cruise-control:3.6.0 $IMAGE_REGISTRY/appscode-images/kafka-cruise-control:3.6.0
$CMD cp ghcr.io/appscode-images/kafka-cruise-control:3.6.1 $IMAGE_REGISTRY/appscode-images/kafka-cruise-control:3.6.1
$CMD cp ghcr.io/appscode-images/kafka-kraft:3.3.2 $IMAGE_REGISTRY/appscode-images/kafka-kraft:3.3.2
$CMD cp ghcr.io/appscode-images/kafka-kraft:3.4.1 $IMAGE_REGISTRY/appscode-images/kafka-kraft:3.4.1
$CMD cp ghcr.io/appscode-images/kafka-kraft:3.5.1 $IMAGE_REGISTRY/appscode-images/kafka-kraft:3.5.1
$CMD cp ghcr.io/appscode-images/kafka-kraft:3.5.2 $IMAGE_REGISTRY/appscode-images/kafka-kraft:3.5.2
$CMD cp ghcr.io/appscode-images/kafka-kraft:3.6.0 $IMAGE_REGISTRY/appscode-images/kafka-kraft:3.6.0
$CMD cp ghcr.io/appscode-images/kafka-kraft:3.6.1 $IMAGE_REGISTRY/appscode-images/kafka-kraft:3.6.1
$CMD cp ghcr.io/appscode-images/kibana:6.8.23 $IMAGE_REGISTRY/appscode-images/kibana:6.8.23
$CMD cp ghcr.io/appscode-images/kibana:7.13.4 $IMAGE_REGISTRY/appscode-images/kibana:7.13.4
$CMD cp ghcr.io/appscode-images/kibana:7.14.2 $IMAGE_REGISTRY/appscode-images/kibana:7.14.2
$CMD cp ghcr.io/appscode-images/kibana:7.16.3 $IMAGE_REGISTRY/appscode-images/kibana:7.16.3
$CMD cp ghcr.io/appscode-images/kibana:7.17.15 $IMAGE_REGISTRY/appscode-images/kibana:7.17.15
$CMD cp ghcr.io/appscode-images/kibana:7.17.23 $IMAGE_REGISTRY/appscode-images/kibana:7.17.23
$CMD cp ghcr.io/appscode-images/kibana:8.11.1 $IMAGE_REGISTRY/appscode-images/kibana:8.11.1
$CMD cp ghcr.io/appscode-images/kibana:8.11.4 $IMAGE_REGISTRY/appscode-images/kibana:8.11.4
$CMD cp ghcr.io/appscode-images/kibana:8.13.4 $IMAGE_REGISTRY/appscode-images/kibana:8.13.4
$CMD cp ghcr.io/appscode-images/kibana:8.14.1 $IMAGE_REGISTRY/appscode-images/kibana:8.14.1
$CMD cp ghcr.io/appscode-images/kibana:8.14.3 $IMAGE_REGISTRY/appscode-images/kibana:8.14.3
$CMD cp ghcr.io/appscode-images/kibana:8.15.0 $IMAGE_REGISTRY/appscode-images/kibana:8.15.0
$CMD cp ghcr.io/appscode-images/kibana:8.2.3 $IMAGE_REGISTRY/appscode-images/kibana:8.2.3
$CMD cp ghcr.io/appscode-images/kibana:8.5.3 $IMAGE_REGISTRY/appscode-images/kibana:8.5.3
$CMD cp ghcr.io/appscode-images/kibana:8.6.2 $IMAGE_REGISTRY/appscode-images/kibana:8.6.2
$CMD cp ghcr.io/appscode-images/kibana:8.8.2 $IMAGE_REGISTRY/appscode-images/kibana:8.8.2
$CMD cp ghcr.io/appscode-images/mariadb:10.10.7-jammy $IMAGE_REGISTRY/appscode-images/mariadb:10.10.7-jammy
$CMD cp ghcr.io/appscode-images/mariadb:10.11.6-jammy $IMAGE_REGISTRY/appscode-images/mariadb:10.11.6-jammy
$CMD cp ghcr.io/appscode-images/mariadb:10.4.32-focal $IMAGE_REGISTRY/appscode-images/mariadb:10.4.32-focal
$CMD cp ghcr.io/appscode-images/mariadb:10.5.23-focal $IMAGE_REGISTRY/appscode-images/mariadb:10.5.23-focal
$CMD cp ghcr.io/appscode-images/mariadb:10.6.16-focal $IMAGE_REGISTRY/appscode-images/mariadb:10.6.16-focal
$CMD cp ghcr.io/appscode-images/mariadb:11.0.4-jammy $IMAGE_REGISTRY/appscode-images/mariadb:11.0.4-jammy
$CMD cp ghcr.io/appscode-images/mariadb:11.1.3-jammy $IMAGE_REGISTRY/appscode-images/mariadb:11.1.3-jammy
$CMD cp ghcr.io/appscode-images/mariadb:11.2.2-jammy $IMAGE_REGISTRY/appscode-images/mariadb:11.2.2-jammy
$CMD cp ghcr.io/appscode-images/mariadb:11.3.2-jammy $IMAGE_REGISTRY/appscode-images/mariadb:11.3.2-jammy
$CMD cp ghcr.io/appscode-images/mariadb:11.4.3-noble $IMAGE_REGISTRY/appscode-images/mariadb:11.4.3-noble
$CMD cp ghcr.io/appscode-images/mariadb:11.5.2-noble $IMAGE_REGISTRY/appscode-images/mariadb:11.5.2-noble
$CMD cp ghcr.io/appscode-images/memcached:1.5.22-alpine $IMAGE_REGISTRY/appscode-images/memcached:1.5.22-alpine
$CMD cp ghcr.io/appscode-images/memcached:1.6.22-alpine $IMAGE_REGISTRY/appscode-images/memcached:1.6.22-alpine
$CMD cp ghcr.io/appscode-images/memcached:1.6.29-alpine $IMAGE_REGISTRY/appscode-images/memcached:1.6.29-alpine
$CMD cp ghcr.io/appscode-images/memcached_exporter:v0.14.3-ac $IMAGE_REGISTRY/appscode-images/memcached_exporter:v0.14.3-ac
$CMD cp ghcr.io/appscode-images/mongo:4.2.24 $IMAGE_REGISTRY/appscode-images/mongo:4.2.24
$CMD cp ghcr.io/appscode-images/mongo:4.4.26 $IMAGE_REGISTRY/appscode-images/mongo:4.4.26
$CMD cp ghcr.io/appscode-images/mongo:5.0.23 $IMAGE_REGISTRY/appscode-images/mongo:5.0.23
$CMD cp ghcr.io/appscode-images/mongo:5.0.26 $IMAGE_REGISTRY/appscode-images/mongo:5.0.26
$CMD cp ghcr.io/appscode-images/mongo:6.0.12 $IMAGE_REGISTRY/appscode-images/mongo:6.0.12
$CMD cp ghcr.io/appscode-images/mongo:7.0.5 $IMAGE_REGISTRY/appscode-images/mongo:7.0.5
$CMD cp ghcr.io/appscode-images/mongo:7.0.8 $IMAGE_REGISTRY/appscode-images/mongo:7.0.8
$CMD cp ghcr.io/appscode-images/mysql:5.7.42-debian $IMAGE_REGISTRY/appscode-images/mysql:5.7.42-debian
$CMD cp ghcr.io/appscode-images/mysql:5.7.44-oracle $IMAGE_REGISTRY/appscode-images/mysql:5.7.44-oracle
$CMD cp ghcr.io/appscode-images/mysql:8.0.31-oracle $IMAGE_REGISTRY/appscode-images/mysql:8.0.31-oracle
$CMD cp ghcr.io/appscode-images/mysql:8.0.35-oracle $IMAGE_REGISTRY/appscode-images/mysql:8.0.35-oracle
$CMD cp ghcr.io/appscode-images/mysql:8.0.36-debian $IMAGE_REGISTRY/appscode-images/mysql:8.0.36-debian
$CMD cp ghcr.io/appscode-images/mysql:8.1.0-oracle $IMAGE_REGISTRY/appscode-images/mysql:8.1.0-oracle
$CMD cp ghcr.io/appscode-images/mysql:8.2.0-oracle $IMAGE_REGISTRY/appscode-images/mysql:8.2.0-oracle
$CMD cp ghcr.io/appscode-images/mysql:8.4.2-oracle $IMAGE_REGISTRY/appscode-images/mysql:8.4.2-oracle
$CMD cp ghcr.io/appscode-images/opensearch-dashboards:1.1.0 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.1.0
$CMD cp ghcr.io/appscode-images/opensearch-dashboards:1.2.0 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.2.0
$CMD cp ghcr.io/appscode-images/opensearch-dashboards:1.3.13 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.3.13
$CMD cp ghcr.io/appscode-images/opensearch-dashboards:1.3.18 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:1.3.18
$CMD cp ghcr.io/appscode-images/opensearch-dashboards:2.0.1 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.0.1
$CMD cp ghcr.io/appscode-images/opensearch-dashboards:2.11.1 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.11.1
$CMD cp ghcr.io/appscode-images/opensearch-dashboards:2.14.0 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.14.0
$CMD cp ghcr.io/appscode-images/opensearch-dashboards:2.16.0 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.16.0
$CMD cp ghcr.io/appscode-images/opensearch-dashboards:2.5.0 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.5.0
$CMD cp ghcr.io/appscode-images/opensearch-dashboards:2.8.0 $IMAGE_REGISTRY/appscode-images/opensearch-dashboards:2.8.0
$CMD cp ghcr.io/appscode-images/opensearch:1.1.0 $IMAGE_REGISTRY/appscode-images/opensearch:1.1.0
$CMD cp ghcr.io/appscode-images/opensearch:1.2.4 $IMAGE_REGISTRY/appscode-images/opensearch:1.2.4
$CMD cp ghcr.io/appscode-images/opensearch:1.3.13 $IMAGE_REGISTRY/appscode-images/opensearch:1.3.13
$CMD cp ghcr.io/appscode-images/opensearch:1.3.18 $IMAGE_REGISTRY/appscode-images/opensearch:1.3.18
$CMD cp ghcr.io/appscode-images/opensearch:2.0.1 $IMAGE_REGISTRY/appscode-images/opensearch:2.0.1
$CMD cp ghcr.io/appscode-images/opensearch:2.11.1 $IMAGE_REGISTRY/appscode-images/opensearch:2.11.1
$CMD cp ghcr.io/appscode-images/opensearch:2.14.0 $IMAGE_REGISTRY/appscode-images/opensearch:2.14.0
$CMD cp ghcr.io/appscode-images/opensearch:2.16.0 $IMAGE_REGISTRY/appscode-images/opensearch:2.16.0
$CMD cp ghcr.io/appscode-images/opensearch:2.5.0 $IMAGE_REGISTRY/appscode-images/opensearch:2.5.0
$CMD cp ghcr.io/appscode-images/opensearch:2.8.0 $IMAGE_REGISTRY/appscode-images/opensearch:2.8.0
$CMD cp ghcr.io/appscode-images/pgpool2:4.4.5 $IMAGE_REGISTRY/appscode-images/pgpool2:4.4.5
$CMD cp ghcr.io/appscode-images/pgpool2:4.4.8 $IMAGE_REGISTRY/appscode-images/pgpool2:4.4.8
$CMD cp ghcr.io/appscode-images/pgpool2:4.5.0 $IMAGE_REGISTRY/appscode-images/pgpool2:4.5.0
$CMD cp ghcr.io/appscode-images/pgpool2:4.5.3 $IMAGE_REGISTRY/appscode-images/pgpool2:4.5.3
$CMD cp ghcr.io/appscode-images/pgpool2_exporter:v1.2.2 $IMAGE_REGISTRY/appscode-images/pgpool2_exporter:v1.2.2
$CMD cp ghcr.io/appscode-images/postgres:10.23-alpine $IMAGE_REGISTRY/appscode-images/postgres:10.23-alpine
$CMD cp ghcr.io/appscode-images/postgres:10.23-bullseye $IMAGE_REGISTRY/appscode-images/postgres:10.23-bullseye
$CMD cp ghcr.io/appscode-images/postgres:11.22-alpine $IMAGE_REGISTRY/appscode-images/postgres:11.22-alpine
$CMD cp ghcr.io/appscode-images/postgres:11.22-bookworm $IMAGE_REGISTRY/appscode-images/postgres:11.22-bookworm
$CMD cp ghcr.io/appscode-images/postgres:12.17-alpine $IMAGE_REGISTRY/appscode-images/postgres:12.17-alpine
$CMD cp ghcr.io/appscode-images/postgres:12.17-bookworm $IMAGE_REGISTRY/appscode-images/postgres:12.17-bookworm
$CMD cp ghcr.io/appscode-images/postgres:13.13-alpine $IMAGE_REGISTRY/appscode-images/postgres:13.13-alpine
$CMD cp ghcr.io/appscode-images/postgres:13.13-bookworm $IMAGE_REGISTRY/appscode-images/postgres:13.13-bookworm
$CMD cp ghcr.io/appscode-images/postgres:14.10-alpine $IMAGE_REGISTRY/appscode-images/postgres:14.10-alpine
$CMD cp ghcr.io/appscode-images/postgres:14.10-bookworm $IMAGE_REGISTRY/appscode-images/postgres:14.10-bookworm
$CMD cp ghcr.io/appscode-images/postgres:14.13-alpine $IMAGE_REGISTRY/appscode-images/postgres:14.13-alpine
$CMD cp ghcr.io/appscode-images/postgres:14.13-bookworm $IMAGE_REGISTRY/appscode-images/postgres:14.13-bookworm
$CMD cp ghcr.io/appscode-images/postgres:15.5-alpine $IMAGE_REGISTRY/appscode-images/postgres:15.5-alpine
$CMD cp ghcr.io/appscode-images/postgres:15.5-bookworm $IMAGE_REGISTRY/appscode-images/postgres:15.5-bookworm
$CMD cp ghcr.io/appscode-images/postgres:15.8-alpine $IMAGE_REGISTRY/appscode-images/postgres:15.8-alpine
$CMD cp ghcr.io/appscode-images/postgres:15.8-bookworm $IMAGE_REGISTRY/appscode-images/postgres:15.8-bookworm
$CMD cp ghcr.io/appscode-images/postgres:16.1-alpine $IMAGE_REGISTRY/appscode-images/postgres:16.1-alpine
$CMD cp ghcr.io/appscode-images/postgres:16.1-bookworm $IMAGE_REGISTRY/appscode-images/postgres:16.1-bookworm
$CMD cp ghcr.io/appscode-images/postgres:16.4-alpine $IMAGE_REGISTRY/appscode-images/postgres:16.4-alpine
$CMD cp ghcr.io/appscode-images/postgres:16.4-bookworm $IMAGE_REGISTRY/appscode-images/postgres:16.4-bookworm
$CMD cp ghcr.io/appscode-images/rabbitmq:3.12.12-management-alpine $IMAGE_REGISTRY/appscode-images/rabbitmq:3.12.12-management-alpine
$CMD cp ghcr.io/appscode-images/rabbitmq:3.13.2-management-alpine $IMAGE_REGISTRY/appscode-images/rabbitmq:3.13.2-management-alpine
$CMD cp ghcr.io/appscode-images/redis:5.0.14-bullseye $IMAGE_REGISTRY/appscode-images/redis:5.0.14-bullseye
$CMD cp ghcr.io/appscode-images/redis:6.0.20-bookworm $IMAGE_REGISTRY/appscode-images/redis:6.0.20-bookworm
$CMD cp ghcr.io/appscode-images/redis:6.2.14-bookworm $IMAGE_REGISTRY/appscode-images/redis:6.2.14-bookworm
$CMD cp ghcr.io/appscode-images/redis:7.0.14-bookworm $IMAGE_REGISTRY/appscode-images/redis:7.0.14-bookworm
$CMD cp ghcr.io/appscode-images/redis:7.0.15-bookworm $IMAGE_REGISTRY/appscode-images/redis:7.0.15-bookworm
$CMD cp ghcr.io/appscode-images/redis:7.2.3-bookworm $IMAGE_REGISTRY/appscode-images/redis:7.2.3-bookworm
$CMD cp ghcr.io/appscode-images/redis:7.2.4-bookworm $IMAGE_REGISTRY/appscode-images/redis:7.2.4-bookworm
$CMD cp ghcr.io/appscode-images/redis:7.4.0-bookworm $IMAGE_REGISTRY/appscode-images/redis:7.4.0-bookworm
$CMD cp ghcr.io/appscode-images/singlestore-node:alma-8.1.32-e3d3cde6da $IMAGE_REGISTRY/appscode-images/singlestore-node:alma-8.1.32-e3d3cde6da
$CMD cp ghcr.io/appscode-images/singlestore-node:alma-8.5.30-4f46ab16a5 $IMAGE_REGISTRY/appscode-images/singlestore-node:alma-8.5.30-4f46ab16a5
$CMD cp ghcr.io/appscode-images/singlestore-node:alma-8.5.7-bf633c1a54 $IMAGE_REGISTRY/appscode-images/singlestore-node:alma-8.5.7-bf633c1a54
$CMD cp ghcr.io/appscode-images/singlestore-node:alma-8.7.10-95e2357384 $IMAGE_REGISTRY/appscode-images/singlestore-node:alma-8.7.10-95e2357384
$CMD cp ghcr.io/appscode-images/solr:8.11.2 $IMAGE_REGISTRY/appscode-images/solr:8.11.2
$CMD cp ghcr.io/appscode-images/solr:9.4.1 $IMAGE_REGISTRY/appscode-images/solr:9.4.1
$CMD cp ghcr.io/appscode-images/solr:9.6.1 $IMAGE_REGISTRY/appscode-images/solr:9.6.1
$CMD cp ghcr.io/appscode-images/zookeeper:3.7.2 $IMAGE_REGISTRY/appscode-images/zookeeper:3.7.2
$CMD cp ghcr.io/appscode-images/zookeeper:3.8.3 $IMAGE_REGISTRY/appscode-images/zookeeper:3.8.3
$CMD cp ghcr.io/appscode-images/zookeeper:3.9.1 $IMAGE_REGISTRY/appscode-images/zookeeper:3.9.1
$CMD cp ghcr.io/appscode/k8s-wait-for:v2.0 $IMAGE_REGISTRY/appscode/k8s-wait-for:v2.0
$CMD cp ghcr.io/appscode/kube-rbac-proxy:v0.11.0 $IMAGE_REGISTRY/appscode/kube-rbac-proxy:v0.11.0
$CMD cp ghcr.io/appscode/kubectl-nonroot:1.25 $IMAGE_REGISTRY/appscode/kubectl-nonroot:1.25
$CMD cp ghcr.io/appscode/petset:v0.0.7 $IMAGE_REGISTRY/appscode/petset:v0.0.7
$CMD cp ghcr.io/appscode/sidekick:v0.0.8 $IMAGE_REGISTRY/appscode/sidekick:v0.0.8
$CMD cp ghcr.io/kubedb/cassandra-init:4.1.6 $IMAGE_REGISTRY/kubedb/cassandra-init:4.1.6
$CMD cp ghcr.io/kubedb/cassandra-init:5.0.0 $IMAGE_REGISTRY/kubedb/cassandra-init:5.0.0
$CMD cp ghcr.io/kubedb/clickhouse-init:24.4.1 $IMAGE_REGISTRY/kubedb/clickhouse-init:24.4.1
$CMD cp ghcr.io/kubedb/dashboard-restic-plugin:v0.1.0 $IMAGE_REGISTRY/kubedb/dashboard-restic-plugin:v0.1.0
$CMD cp ghcr.io/kubedb/dashboard-restic-plugin:v0.6.0 $IMAGE_REGISTRY/kubedb/dashboard-restic-plugin:v0.6.0
$CMD cp ghcr.io/kubedb/druid-init:25.0.0 $IMAGE_REGISTRY/kubedb/druid-init:25.0.0
$CMD cp ghcr.io/kubedb/druid-init:28.0.1 $IMAGE_REGISTRY/kubedb/druid-init:28.0.1
$CMD cp ghcr.io/kubedb/druid-init:30.0.0 $IMAGE_REGISTRY/kubedb/druid-init:30.0.0
$CMD cp ghcr.io/kubedb/elasticsearch-restic-plugin:v0.11.0 $IMAGE_REGISTRY/kubedb/elasticsearch-restic-plugin:v0.11.0
$CMD cp ghcr.io/kubedb/kubedb-autoscaler:v0.33.0 $IMAGE_REGISTRY/kubedb/kubedb-autoscaler:v0.33.0
$CMD cp ghcr.io/kubedb/kubedb-crd-manager:v0.3.0 $IMAGE_REGISTRY/kubedb/kubedb-crd-manager:v0.3.0
$CMD cp ghcr.io/kubedb/kubedb-kibana:v0.24.0 $IMAGE_REGISTRY/kubedb/kubedb-kibana:v0.24.0
$CMD cp ghcr.io/kubedb/kubedb-manifest-plugin:v0.11.0 $IMAGE_REGISTRY/kubedb/kubedb-manifest-plugin:v0.11.0
$CMD cp ghcr.io/kubedb/kubedb-ops-manager:v0.35.0 $IMAGE_REGISTRY/kubedb/kubedb-ops-manager:v0.35.0
$CMD cp ghcr.io/kubedb/kubedb-provisioner:v0.48.0 $IMAGE_REGISTRY/kubedb/kubedb-provisioner:v0.48.0
$CMD cp ghcr.io/kubedb/kubedb-schema-manager:v0.24.0 $IMAGE_REGISTRY/kubedb/kubedb-schema-manager:v0.24.0
$CMD cp ghcr.io/kubedb/kubedb-ui-server:v0.24.0 $IMAGE_REGISTRY/kubedb/kubedb-ui-server:v0.24.0
$CMD cp ghcr.io/kubedb/kubedb-webhook-server:v0.24.0 $IMAGE_REGISTRY/kubedb/kubedb-webhook-server:v0.24.0
$CMD cp ghcr.io/kubedb/mariadb-archiver:v0.4.0_10.10.7-jammy $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.4.0_10.10.7-jammy
$CMD cp ghcr.io/kubedb/mariadb-archiver:v0.4.0_10.11.6-jammy $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.4.0_10.11.6-jammy
$CMD cp ghcr.io/kubedb/mariadb-archiver:v0.4.0_10.4.32-focal $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.4.0_10.4.32-focal
$CMD cp ghcr.io/kubedb/mariadb-archiver:v0.4.0_10.5.23-focal $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.4.0_10.5.23-focal
$CMD cp ghcr.io/kubedb/mariadb-archiver:v0.4.0_10.6.16-focal $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.4.0_10.6.16-focal
$CMD cp ghcr.io/kubedb/mariadb-archiver:v0.4.0_11.0.4-jammy $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.4.0_11.0.4-jammy
$CMD cp ghcr.io/kubedb/mariadb-archiver:v0.4.0_11.1.3-jammy $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.4.0_11.1.3-jammy
$CMD cp ghcr.io/kubedb/mariadb-archiver:v0.4.0_11.2.2-jammy $IMAGE_REGISTRY/kubedb/mariadb-archiver:v0.4.0_11.2.2-jammy
$CMD cp ghcr.io/kubedb/mariadb-coordinator:v0.28.0 $IMAGE_REGISTRY/kubedb/mariadb-coordinator:v0.28.0
$CMD cp ghcr.io/kubedb/mariadb-csi-snapshotter-plugin:v0.8.0 $IMAGE_REGISTRY/kubedb/mariadb-csi-snapshotter-plugin:v0.8.0
$CMD cp ghcr.io/kubedb/mariadb-init:0.5.2 $IMAGE_REGISTRY/kubedb/mariadb-init:0.5.2
$CMD cp ghcr.io/kubedb/mariadb-restic-plugin:v0.6.0 $IMAGE_REGISTRY/kubedb/mariadb-restic-plugin:v0.6.0
$CMD cp ghcr.io/kubedb/mongodb-csi-snapshotter-plugin:v0.9.0 $IMAGE_REGISTRY/kubedb/mongodb-csi-snapshotter-plugin:v0.9.0
$CMD cp ghcr.io/kubedb/mongodb-init:4.2-v9 $IMAGE_REGISTRY/kubedb/mongodb-init:4.2-v9
$CMD cp ghcr.io/kubedb/mongodb-init:6.0-v10 $IMAGE_REGISTRY/kubedb/mongodb-init:6.0-v10
$CMD cp ghcr.io/kubedb/mongodb_exporter:v0.40.0 $IMAGE_REGISTRY/kubedb/mongodb_exporter:v0.40.0
$CMD cp ghcr.io/kubedb/mssql-coordinator:v0.3.0 $IMAGE_REGISTRY/kubedb/mssql-coordinator:v0.3.0
$CMD cp ghcr.io/kubedb/mssql-exporter:1.1.0 $IMAGE_REGISTRY/kubedb/mssql-exporter:1.1.0
$CMD cp ghcr.io/kubedb/mssql-init:2022-ubuntu-22-v3 $IMAGE_REGISTRY/kubedb/mssql-init:2022-ubuntu-22-v3
$CMD cp ghcr.io/kubedb/mssqlserver-archiver:v0.0.1 $IMAGE_REGISTRY/kubedb/mssqlserver-archiver:v0.0.1
$CMD cp ghcr.io/kubedb/mssqlserver-walg-plugin:v0.0.1 $IMAGE_REGISTRY/kubedb/mssqlserver-walg-plugin:v0.0.1
$CMD cp ghcr.io/kubedb/mysql-archiver:v0.9.0_5.7.44 $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.9.0_5.7.44
$CMD cp ghcr.io/kubedb/mysql-archiver:v0.9.0_8.0.35 $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.9.0_8.0.35
$CMD cp ghcr.io/kubedb/mysql-archiver:v0.9.0_8.1.0 $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.9.0_8.1.0
$CMD cp ghcr.io/kubedb/mysql-archiver:v0.9.0_8.2.0 $IMAGE_REGISTRY/kubedb/mysql-archiver:v0.9.0_8.2.0
$CMD cp ghcr.io/kubedb/mysql-coordinator:v0.26.0 $IMAGE_REGISTRY/kubedb/mysql-coordinator:v0.26.0
$CMD cp ghcr.io/kubedb/mysql-csi-snapshotter-plugin:v0.9.0 $IMAGE_REGISTRY/kubedb/mysql-csi-snapshotter-plugin:v0.9.0
$CMD cp ghcr.io/kubedb/mysql-init:5.7-v4 $IMAGE_REGISTRY/kubedb/mysql-init:5.7-v4
$CMD cp ghcr.io/kubedb/mysql-init:8.0.31-v3 $IMAGE_REGISTRY/kubedb/mysql-init:8.0.31-v3
$CMD cp ghcr.io/kubedb/mysql-init:8.4.2-v1 $IMAGE_REGISTRY/kubedb/mysql-init:8.4.2-v1
$CMD cp ghcr.io/kubedb/mysql-router-init:v0.26.0 $IMAGE_REGISTRY/kubedb/mysql-router-init:v0.26.0
$CMD cp ghcr.io/kubedb/mysqld-exporter:v0.13.1 $IMAGE_REGISTRY/kubedb/mysqld-exporter:v0.13.1
$CMD cp ghcr.io/kubedb/percona-xtradb-coordinator:v0.21.0 $IMAGE_REGISTRY/kubedb/percona-xtradb-coordinator:v0.21.0
$CMD cp ghcr.io/kubedb/percona-xtradb-init:0.2.1 $IMAGE_REGISTRY/kubedb/percona-xtradb-init:0.2.1
$CMD cp ghcr.io/kubedb/pg-coordinator:v0.32.0 $IMAGE_REGISTRY/kubedb/pg-coordinator:v0.32.0
$CMD cp ghcr.io/kubedb/pgbouncer:1.17.0 $IMAGE_REGISTRY/kubedb/pgbouncer:1.17.0
$CMD cp ghcr.io/kubedb/pgbouncer:1.18.0 $IMAGE_REGISTRY/kubedb/pgbouncer:1.18.0
$CMD cp ghcr.io/kubedb/pgbouncer:1.23.1 $IMAGE_REGISTRY/kubedb/pgbouncer:1.23.1
$CMD cp ghcr.io/kubedb/pgbouncer_exporter:v0.1.1 $IMAGE_REGISTRY/kubedb/pgbouncer_exporter:v0.1.1
$CMD cp ghcr.io/kubedb/postgres-archiver:v0.9.0_11.22-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.9.0_11.22-alpine
$CMD cp ghcr.io/kubedb/postgres-archiver:v0.9.0_11.22-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.9.0_11.22-bookworm
$CMD cp ghcr.io/kubedb/postgres-archiver:v0.9.0_12.17-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.9.0_12.17-alpine
$CMD cp ghcr.io/kubedb/postgres-archiver:v0.9.0_12.17-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.9.0_12.17-bookworm
$CMD cp ghcr.io/kubedb/postgres-archiver:v0.9.0_13.13-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.9.0_13.13-alpine
$CMD cp ghcr.io/kubedb/postgres-archiver:v0.9.0_13.13-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.9.0_13.13-bookworm
$CMD cp ghcr.io/kubedb/postgres-archiver:v0.9.0_14.10-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.9.0_14.10-alpine
$CMD cp ghcr.io/kubedb/postgres-archiver:v0.9.0_14.10-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.9.0_14.10-bookworm
$CMD cp ghcr.io/kubedb/postgres-archiver:v0.9.0_15.5-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.9.0_15.5-alpine
$CMD cp ghcr.io/kubedb/postgres-archiver:v0.9.0_15.5-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.9.0_15.5-bookworm
$CMD cp ghcr.io/kubedb/postgres-archiver:v0.9.0_16.1-alpine $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.9.0_16.1-alpine
$CMD cp ghcr.io/kubedb/postgres-archiver:v0.9.0_16.1-bookworm $IMAGE_REGISTRY/kubedb/postgres-archiver:v0.9.0_16.1-bookworm
$CMD cp ghcr.io/kubedb/postgres-csi-snapshotter-plugin:v0.9.0 $IMAGE_REGISTRY/kubedb/postgres-csi-snapshotter-plugin:v0.9.0
$CMD cp ghcr.io/kubedb/postgres-init:0.15.0 $IMAGE_REGISTRY/kubedb/postgres-init:0.15.0
$CMD cp ghcr.io/kubedb/postgres-restic-plugin:v0.10.0_16.1 $IMAGE_REGISTRY/kubedb/postgres-restic-plugin:v0.10.0_16.1
$CMD cp ghcr.io/kubedb/provider-aws:v0.10.0 $IMAGE_REGISTRY/kubedb/provider-aws:v0.10.0
$CMD cp ghcr.io/kubedb/provider-azure:v0.10.0 $IMAGE_REGISTRY/kubedb/provider-azure:v0.10.0
$CMD cp ghcr.io/kubedb/provider-gcp:v0.10.0 $IMAGE_REGISTRY/kubedb/provider-gcp:v0.10.0
$CMD cp ghcr.io/kubedb/proxysql-exporter:v1.1.0 $IMAGE_REGISTRY/kubedb/proxysql-exporter:v1.1.0
$CMD cp ghcr.io/kubedb/proxysql:2.3.2-centos-v2 $IMAGE_REGISTRY/kubedb/proxysql:2.3.2-centos-v2
$CMD cp ghcr.io/kubedb/proxysql:2.3.2-debian-v2 $IMAGE_REGISTRY/kubedb/proxysql:2.3.2-debian-v2
$CMD cp ghcr.io/kubedb/proxysql:2.4.4-centos $IMAGE_REGISTRY/kubedb/proxysql:2.4.4-centos
$CMD cp ghcr.io/kubedb/proxysql:2.4.4-debian $IMAGE_REGISTRY/kubedb/proxysql:2.4.4-debian
$CMD cp ghcr.io/kubedb/proxysql:2.6.3-debian $IMAGE_REGISTRY/kubedb/proxysql:2.6.3-debian
$CMD cp ghcr.io/kubedb/rabbitmq-init:3.12.12 $IMAGE_REGISTRY/kubedb/rabbitmq-init:3.12.12
$CMD cp ghcr.io/kubedb/rabbitmq-init:3.13.2 $IMAGE_REGISTRY/kubedb/rabbitmq-init:3.13.2
$CMD cp ghcr.io/kubedb/redis-coordinator:v0.27.0 $IMAGE_REGISTRY/kubedb/redis-coordinator:v0.27.0
$CMD cp ghcr.io/kubedb/redis-init:0.9.0 $IMAGE_REGISTRY/kubedb/redis-init:0.9.0
$CMD cp ghcr.io/kubedb/redis-restic-plugin:v0.11.0 $IMAGE_REGISTRY/kubedb/redis-restic-plugin:v0.11.0
$CMD cp ghcr.io/kubedb/redis:4.0.11 $IMAGE_REGISTRY/kubedb/redis:4.0.11
$CMD cp ghcr.io/kubedb/redis_exporter:1.58.0 $IMAGE_REGISTRY/kubedb/redis_exporter:1.58.0
$CMD cp ghcr.io/kubedb/redis_exporter:v0.21.1 $IMAGE_REGISTRY/kubedb/redis_exporter:v0.21.1
$CMD cp ghcr.io/kubedb/replication-mode-detector:v0.35.0 $IMAGE_REGISTRY/kubedb/replication-mode-detector:v0.35.0
$CMD cp ghcr.io/kubedb/singlestore-coordinator:v0.3.0 $IMAGE_REGISTRY/kubedb/singlestore-coordinator:v0.3.0
$CMD cp ghcr.io/kubedb/singlestore-init:8.1-v2 $IMAGE_REGISTRY/kubedb/singlestore-init:8.1-v2
$CMD cp ghcr.io/kubedb/singlestore-init:8.5-v2 $IMAGE_REGISTRY/kubedb/singlestore-init:8.5-v2
$CMD cp ghcr.io/kubedb/singlestore-init:8.7.10-v1 $IMAGE_REGISTRY/kubedb/singlestore-init:8.7.10-v1
$CMD cp ghcr.io/kubedb/solr-init:8.11.2 $IMAGE_REGISTRY/kubedb/solr-init:8.11.2
$CMD cp ghcr.io/kubedb/solr-init:9.4.1 $IMAGE_REGISTRY/kubedb/solr-init:9.4.1
$CMD cp ghcr.io/kubedb/solr-init:9.6.1 $IMAGE_REGISTRY/kubedb/solr-init:9.6.1
$CMD cp ghcr.io/kubedb/wal-g:v2024.5.24_mongo $IMAGE_REGISTRY/kubedb/wal-g:v2024.5.24_mongo
$CMD cp ghcr.io/kubedb/zookeeper-init:3.7-v1 $IMAGE_REGISTRY/kubedb/zookeeper-init:3.7-v1
$CMD cp ghcr.io/kubedb/zookeeper-restic-plugin:v0.4.0 $IMAGE_REGISTRY/kubedb/zookeeper-restic-plugin:v0.4.0
$CMD cp mcr.microsoft.com/mssql/server:2022-CU12-ubuntu-22.04 $IMAGE_REGISTRY/mssql/server:2022-CU12-ubuntu-22.04
$CMD cp mcr.microsoft.com/mssql/server:2022-CU14-ubuntu-22.04 $IMAGE_REGISTRY/mssql/server:2022-CU14-ubuntu-22.04
$CMD cp mysql/mysql-router:8.0.31 $IMAGE_REGISTRY/mysql/mysql-router:8.0.31
$CMD cp percona/percona-server-mongodb:4.2.24 $IMAGE_REGISTRY/percona/percona-server-mongodb:4.2.24
$CMD cp percona/percona-server-mongodb:4.4.26 $IMAGE_REGISTRY/percona/percona-server-mongodb:4.4.26
$CMD cp percona/percona-server-mongodb:5.0.23 $IMAGE_REGISTRY/percona/percona-server-mongodb:5.0.23
$CMD cp percona/percona-server-mongodb:6.0.12 $IMAGE_REGISTRY/percona/percona-server-mongodb:6.0.12
$CMD cp percona/percona-server-mongodb:7.0.4 $IMAGE_REGISTRY/percona/percona-server-mongodb:7.0.4
$CMD cp percona/percona-xtradb-cluster:8.0.26 $IMAGE_REGISTRY/percona/percona-xtradb-cluster:8.0.26
$CMD cp percona/percona-xtradb-cluster:8.0.28 $IMAGE_REGISTRY/percona/percona-xtradb-cluster:8.0.28
$CMD cp percona/percona-xtradb-cluster:8.0.31 $IMAGE_REGISTRY/percona/percona-xtradb-cluster:8.0.31
$CMD cp postgis/postgis:11-3.3 $IMAGE_REGISTRY/postgis/postgis:11-3.3
$CMD cp postgis/postgis:12-3.4 $IMAGE_REGISTRY/postgis/postgis:12-3.4
$CMD cp postgis/postgis:13-3.4 $IMAGE_REGISTRY/postgis/postgis:13-3.4
$CMD cp postgis/postgis:14-3.4 $IMAGE_REGISTRY/postgis/postgis:14-3.4
$CMD cp postgis/postgis:15-3.4 $IMAGE_REGISTRY/postgis/postgis:15-3.4
$CMD cp postgis/postgis:16-3.4 $IMAGE_REGISTRY/postgis/postgis:16-3.4
$CMD cp prom/memcached-exporter:v0.14.2 $IMAGE_REGISTRY/prom/memcached-exporter:v0.14.2
$CMD cp prom/mysqld-exporter:v0.13.0 $IMAGE_REGISTRY/prom/mysqld-exporter:v0.13.0
$CMD cp prometheuscommunity/elasticsearch-exporter:v1.7.0 $IMAGE_REGISTRY/prometheuscommunity/elasticsearch-exporter:v1.7.0
$CMD cp prometheuscommunity/postgres-exporter:v0.15.0 $IMAGE_REGISTRY/prometheuscommunity/postgres-exporter:v0.15.0
$CMD cp registry.k8s.io/git-sync/git-sync:v4.2.1 $IMAGE_REGISTRY/git-sync/git-sync:v4.2.1
$CMD cp singlestore/cluster-in-a-box:alma-8.1.32-e3d3cde6da-4.0.16-1.17.6 $IMAGE_REGISTRY/singlestore/cluster-in-a-box:alma-8.1.32-e3d3cde6da-4.0.16-1.17.6
$CMD cp singlestore/cluster-in-a-box:alma-8.5.22-fe61f40cd1-4.1.0-1.17.11 $IMAGE_REGISTRY/singlestore/cluster-in-a-box:alma-8.5.22-fe61f40cd1-4.1.0-1.17.11
$CMD cp singlestore/cluster-in-a-box:alma-8.5.7-bf633c1a54-4.0.17-1.17.8 $IMAGE_REGISTRY/singlestore/cluster-in-a-box:alma-8.5.7-bf633c1a54-4.0.17-1.17.8
$CMD cp singlestore/cluster-in-a-box:alma-8.7.10-95e2357384-4.1.0-1.17.14 $IMAGE_REGISTRY/singlestore/cluster-in-a-box:alma-8.7.10-95e2357384-4.1.0-1.17.14
$CMD cp tianon/toybox:0.8.4 $IMAGE_REGISTRY/tianon/toybox:0.8.4
$CMD cp timescale/timescaledb:2.14.2-pg13-oss $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg13-oss
$CMD cp timescale/timescaledb:2.14.2-pg14-oss $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg14-oss
$CMD cp timescale/timescaledb:2.14.2-pg15-oss $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg15-oss
$CMD cp timescale/timescaledb:2.14.2-pg16-oss $IMAGE_REGISTRY/timescale/timescaledb:2.14.2-pg16-oss