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

$CMD pull apache/druid:25.0.0 images/apache-druid-25.0.0.tar
$CMD pull apicurio/apicurio-registry-kafkasql:2.5.11.Final images/apicurio-apicurio-registry-kafkasql-2.5.11.Final.tar
$CMD pull apicurio/apicurio-registry-mem:2.5.11.Final images/apicurio-apicurio-registry-mem-2.5.11.Final.tar
$CMD pull clickhouse/clickhouse-keeper:24.4.1 images/clickhouse-clickhouse-keeper-24.4.1.tar
$CMD pull clickhouse/clickhouse-server:24.4.1 images/clickhouse-clickhouse-server-24.4.1.tar
$CMD pull floragunncom/sg-elasticsearch:7.9.3-oss-47.1.0 images/floragunncom-sg-elasticsearch-7.9.3-oss-47.1.0.tar
$CMD pull ghcr.io/aiven-open/karapace:3.15.0 images/aiven-open-karapace-3.15.0.tar
$CMD pull ghcr.io/appscode-images/cassandra:4.1.6 images/appscode-images-cassandra-4.1.6.tar
$CMD pull ghcr.io/appscode-images/cassandra:5.0.0 images/appscode-images-cassandra-5.0.0.tar
$CMD pull ghcr.io/appscode-images/druid:28.0.1 images/appscode-images-druid-28.0.1.tar
$CMD pull ghcr.io/appscode-images/druid:30.0.0 images/appscode-images-druid-30.0.0.tar
$CMD pull ghcr.io/appscode-images/elastic:6.8.23 images/appscode-images-elastic-6.8.23.tar
$CMD pull ghcr.io/appscode-images/elastic:7.13.4 images/appscode-images-elastic-7.13.4.tar
$CMD pull ghcr.io/appscode-images/elastic:7.14.2 images/appscode-images-elastic-7.14.2.tar
$CMD pull ghcr.io/appscode-images/elastic:7.16.3 images/appscode-images-elastic-7.16.3.tar
$CMD pull ghcr.io/appscode-images/elastic:7.17.15 images/appscode-images-elastic-7.17.15.tar
$CMD pull ghcr.io/appscode-images/elastic:7.17.23 images/appscode-images-elastic-7.17.23.tar
$CMD pull ghcr.io/appscode-images/elastic:8.11.1 images/appscode-images-elastic-8.11.1.tar
$CMD pull ghcr.io/appscode-images/elastic:8.11.4 images/appscode-images-elastic-8.11.4.tar
$CMD pull ghcr.io/appscode-images/elastic:8.13.4 images/appscode-images-elastic-8.13.4.tar
$CMD pull ghcr.io/appscode-images/elastic:8.14.1 images/appscode-images-elastic-8.14.1.tar
$CMD pull ghcr.io/appscode-images/elastic:8.14.3 images/appscode-images-elastic-8.14.3.tar
$CMD pull ghcr.io/appscode-images/elastic:8.15.0 images/appscode-images-elastic-8.15.0.tar
$CMD pull ghcr.io/appscode-images/elastic:8.2.3 images/appscode-images-elastic-8.2.3.tar
$CMD pull ghcr.io/appscode-images/elastic:8.5.3 images/appscode-images-elastic-8.5.3.tar
$CMD pull ghcr.io/appscode-images/elastic:8.6.2 images/appscode-images-elastic-8.6.2.tar
$CMD pull ghcr.io/appscode-images/elastic:8.8.2 images/appscode-images-elastic-8.8.2.tar
$CMD pull ghcr.io/appscode-images/ferretdb:1.18.0 images/appscode-images-ferretdb-1.18.0.tar
$CMD pull ghcr.io/appscode-images/ferretdb:1.23.0 images/appscode-images-ferretdb-1.23.0.tar
$CMD pull ghcr.io/appscode-images/kafka-connect-cluster:3.3.2 images/appscode-images-kafka-connect-cluster-3.3.2.tar
$CMD pull ghcr.io/appscode-images/kafka-connect-cluster:3.4.1 images/appscode-images-kafka-connect-cluster-3.4.1.tar
$CMD pull ghcr.io/appscode-images/kafka-connect-cluster:3.5.1 images/appscode-images-kafka-connect-cluster-3.5.1.tar
$CMD pull ghcr.io/appscode-images/kafka-connect-cluster:3.5.2 images/appscode-images-kafka-connect-cluster-3.5.2.tar
$CMD pull ghcr.io/appscode-images/kafka-connect-cluster:3.6.0 images/appscode-images-kafka-connect-cluster-3.6.0.tar
$CMD pull ghcr.io/appscode-images/kafka-connect-cluster:3.6.1 images/appscode-images-kafka-connect-cluster-3.6.1.tar
$CMD pull ghcr.io/appscode-images/kafka-connector-gcs:0.13.0 images/appscode-images-kafka-connector-gcs-0.13.0.tar
$CMD pull ghcr.io/appscode-images/kafka-connector-jdbc:2.6.1.final images/appscode-images-kafka-connector-jdbc-2.6.1.final.tar
$CMD pull ghcr.io/appscode-images/kafka-connector-mongodb:1.11.0 images/appscode-images-kafka-connector-mongodb-1.11.0.tar
$CMD pull ghcr.io/appscode-images/kafka-connector-mysql:2.4.2.final images/appscode-images-kafka-connector-mysql-2.4.2.final.tar
$CMD pull ghcr.io/appscode-images/kafka-connector-postgres:2.4.2.final images/appscode-images-kafka-connector-postgres-2.4.2.final.tar
$CMD pull ghcr.io/appscode-images/kafka-connector-s3:2.15.0 images/appscode-images-kafka-connector-s3-2.15.0.tar
$CMD pull ghcr.io/appscode-images/kafka-cruise-control:3.3.2 images/appscode-images-kafka-cruise-control-3.3.2.tar
$CMD pull ghcr.io/appscode-images/kafka-cruise-control:3.4.1 images/appscode-images-kafka-cruise-control-3.4.1.tar
$CMD pull ghcr.io/appscode-images/kafka-cruise-control:3.5.1 images/appscode-images-kafka-cruise-control-3.5.1.tar
$CMD pull ghcr.io/appscode-images/kafka-cruise-control:3.5.2 images/appscode-images-kafka-cruise-control-3.5.2.tar
$CMD pull ghcr.io/appscode-images/kafka-cruise-control:3.6.0 images/appscode-images-kafka-cruise-control-3.6.0.tar
$CMD pull ghcr.io/appscode-images/kafka-cruise-control:3.6.1 images/appscode-images-kafka-cruise-control-3.6.1.tar
$CMD pull ghcr.io/appscode-images/kafka-kraft:3.3.2 images/appscode-images-kafka-kraft-3.3.2.tar
$CMD pull ghcr.io/appscode-images/kafka-kraft:3.4.1 images/appscode-images-kafka-kraft-3.4.1.tar
$CMD pull ghcr.io/appscode-images/kafka-kraft:3.5.1 images/appscode-images-kafka-kraft-3.5.1.tar
$CMD pull ghcr.io/appscode-images/kafka-kraft:3.5.2 images/appscode-images-kafka-kraft-3.5.2.tar
$CMD pull ghcr.io/appscode-images/kafka-kraft:3.6.0 images/appscode-images-kafka-kraft-3.6.0.tar
$CMD pull ghcr.io/appscode-images/kafka-kraft:3.6.1 images/appscode-images-kafka-kraft-3.6.1.tar
$CMD pull ghcr.io/appscode-images/kibana:6.8.23 images/appscode-images-kibana-6.8.23.tar
$CMD pull ghcr.io/appscode-images/kibana:7.13.4 images/appscode-images-kibana-7.13.4.tar
$CMD pull ghcr.io/appscode-images/kibana:7.14.2 images/appscode-images-kibana-7.14.2.tar
$CMD pull ghcr.io/appscode-images/kibana:7.16.3 images/appscode-images-kibana-7.16.3.tar
$CMD pull ghcr.io/appscode-images/kibana:7.17.15 images/appscode-images-kibana-7.17.15.tar
$CMD pull ghcr.io/appscode-images/kibana:7.17.23 images/appscode-images-kibana-7.17.23.tar
$CMD pull ghcr.io/appscode-images/kibana:8.11.1 images/appscode-images-kibana-8.11.1.tar
$CMD pull ghcr.io/appscode-images/kibana:8.11.4 images/appscode-images-kibana-8.11.4.tar
$CMD pull ghcr.io/appscode-images/kibana:8.13.4 images/appscode-images-kibana-8.13.4.tar
$CMD pull ghcr.io/appscode-images/kibana:8.14.1 images/appscode-images-kibana-8.14.1.tar
$CMD pull ghcr.io/appscode-images/kibana:8.14.3 images/appscode-images-kibana-8.14.3.tar
$CMD pull ghcr.io/appscode-images/kibana:8.15.0 images/appscode-images-kibana-8.15.0.tar
$CMD pull ghcr.io/appscode-images/kibana:8.2.3 images/appscode-images-kibana-8.2.3.tar
$CMD pull ghcr.io/appscode-images/kibana:8.5.3 images/appscode-images-kibana-8.5.3.tar
$CMD pull ghcr.io/appscode-images/kibana:8.6.2 images/appscode-images-kibana-8.6.2.tar
$CMD pull ghcr.io/appscode-images/kibana:8.8.2 images/appscode-images-kibana-8.8.2.tar
$CMD pull ghcr.io/appscode-images/mariadb:10.10.7-jammy images/appscode-images-mariadb-10.10.7-jammy.tar
$CMD pull ghcr.io/appscode-images/mariadb:10.11.6-jammy images/appscode-images-mariadb-10.11.6-jammy.tar
$CMD pull ghcr.io/appscode-images/mariadb:10.4.32-focal images/appscode-images-mariadb-10.4.32-focal.tar
$CMD pull ghcr.io/appscode-images/mariadb:10.5.23-focal images/appscode-images-mariadb-10.5.23-focal.tar
$CMD pull ghcr.io/appscode-images/mariadb:10.6.16-focal images/appscode-images-mariadb-10.6.16-focal.tar
$CMD pull ghcr.io/appscode-images/mariadb:11.0.4-jammy images/appscode-images-mariadb-11.0.4-jammy.tar
$CMD pull ghcr.io/appscode-images/mariadb:11.1.3-jammy images/appscode-images-mariadb-11.1.3-jammy.tar
$CMD pull ghcr.io/appscode-images/mariadb:11.2.2-jammy images/appscode-images-mariadb-11.2.2-jammy.tar
$CMD pull ghcr.io/appscode-images/mariadb:11.3.2-jammy images/appscode-images-mariadb-11.3.2-jammy.tar
$CMD pull ghcr.io/appscode-images/mariadb:11.4.3-noble images/appscode-images-mariadb-11.4.3-noble.tar
$CMD pull ghcr.io/appscode-images/mariadb:11.5.2-noble images/appscode-images-mariadb-11.5.2-noble.tar
$CMD pull ghcr.io/appscode-images/memcached:1.5.22-alpine images/appscode-images-memcached-1.5.22-alpine.tar
$CMD pull ghcr.io/appscode-images/memcached:1.6.22-alpine images/appscode-images-memcached-1.6.22-alpine.tar
$CMD pull ghcr.io/appscode-images/memcached:1.6.29-alpine images/appscode-images-memcached-1.6.29-alpine.tar
$CMD pull ghcr.io/appscode-images/memcached_exporter:v0.14.3-ac images/appscode-images-memcached_exporter-v0.14.3-ac.tar
$CMD pull ghcr.io/appscode-images/mongo:4.2.24 images/appscode-images-mongo-4.2.24.tar
$CMD pull ghcr.io/appscode-images/mongo:4.4.26 images/appscode-images-mongo-4.4.26.tar
$CMD pull ghcr.io/appscode-images/mongo:5.0.23 images/appscode-images-mongo-5.0.23.tar
$CMD pull ghcr.io/appscode-images/mongo:5.0.26 images/appscode-images-mongo-5.0.26.tar
$CMD pull ghcr.io/appscode-images/mongo:6.0.12 images/appscode-images-mongo-6.0.12.tar
$CMD pull ghcr.io/appscode-images/mongo:7.0.5 images/appscode-images-mongo-7.0.5.tar
$CMD pull ghcr.io/appscode-images/mongo:7.0.8 images/appscode-images-mongo-7.0.8.tar
$CMD pull ghcr.io/appscode-images/mysql:5.7.42-debian images/appscode-images-mysql-5.7.42-debian.tar
$CMD pull ghcr.io/appscode-images/mysql:5.7.44-oracle images/appscode-images-mysql-5.7.44-oracle.tar
$CMD pull ghcr.io/appscode-images/mysql:8.0.31-oracle images/appscode-images-mysql-8.0.31-oracle.tar
$CMD pull ghcr.io/appscode-images/mysql:8.0.35-oracle images/appscode-images-mysql-8.0.35-oracle.tar
$CMD pull ghcr.io/appscode-images/mysql:8.0.36-debian images/appscode-images-mysql-8.0.36-debian.tar
$CMD pull ghcr.io/appscode-images/mysql:8.1.0-oracle images/appscode-images-mysql-8.1.0-oracle.tar
$CMD pull ghcr.io/appscode-images/mysql:8.2.0-oracle images/appscode-images-mysql-8.2.0-oracle.tar
$CMD pull ghcr.io/appscode-images/mysql:8.4.2-oracle images/appscode-images-mysql-8.4.2-oracle.tar
$CMD pull ghcr.io/appscode-images/opensearch-dashboards:1.1.0 images/appscode-images-opensearch-dashboards-1.1.0.tar
$CMD pull ghcr.io/appscode-images/opensearch-dashboards:1.2.0 images/appscode-images-opensearch-dashboards-1.2.0.tar
$CMD pull ghcr.io/appscode-images/opensearch-dashboards:1.3.13 images/appscode-images-opensearch-dashboards-1.3.13.tar
$CMD pull ghcr.io/appscode-images/opensearch-dashboards:1.3.18 images/appscode-images-opensearch-dashboards-1.3.18.tar
$CMD pull ghcr.io/appscode-images/opensearch-dashboards:2.0.1 images/appscode-images-opensearch-dashboards-2.0.1.tar
$CMD pull ghcr.io/appscode-images/opensearch-dashboards:2.11.1 images/appscode-images-opensearch-dashboards-2.11.1.tar
$CMD pull ghcr.io/appscode-images/opensearch-dashboards:2.14.0 images/appscode-images-opensearch-dashboards-2.14.0.tar
$CMD pull ghcr.io/appscode-images/opensearch-dashboards:2.16.0 images/appscode-images-opensearch-dashboards-2.16.0.tar
$CMD pull ghcr.io/appscode-images/opensearch-dashboards:2.5.0 images/appscode-images-opensearch-dashboards-2.5.0.tar
$CMD pull ghcr.io/appscode-images/opensearch-dashboards:2.8.0 images/appscode-images-opensearch-dashboards-2.8.0.tar
$CMD pull ghcr.io/appscode-images/opensearch:1.1.0 images/appscode-images-opensearch-1.1.0.tar
$CMD pull ghcr.io/appscode-images/opensearch:1.2.4 images/appscode-images-opensearch-1.2.4.tar
$CMD pull ghcr.io/appscode-images/opensearch:1.3.13 images/appscode-images-opensearch-1.3.13.tar
$CMD pull ghcr.io/appscode-images/opensearch:1.3.18 images/appscode-images-opensearch-1.3.18.tar
$CMD pull ghcr.io/appscode-images/opensearch:2.0.1 images/appscode-images-opensearch-2.0.1.tar
$CMD pull ghcr.io/appscode-images/opensearch:2.11.1 images/appscode-images-opensearch-2.11.1.tar
$CMD pull ghcr.io/appscode-images/opensearch:2.14.0 images/appscode-images-opensearch-2.14.0.tar
$CMD pull ghcr.io/appscode-images/opensearch:2.16.0 images/appscode-images-opensearch-2.16.0.tar
$CMD pull ghcr.io/appscode-images/opensearch:2.5.0 images/appscode-images-opensearch-2.5.0.tar
$CMD pull ghcr.io/appscode-images/opensearch:2.8.0 images/appscode-images-opensearch-2.8.0.tar
$CMD pull ghcr.io/appscode-images/pgpool2:4.4.5 images/appscode-images-pgpool2-4.4.5.tar
$CMD pull ghcr.io/appscode-images/pgpool2:4.4.8 images/appscode-images-pgpool2-4.4.8.tar
$CMD pull ghcr.io/appscode-images/pgpool2:4.5.0 images/appscode-images-pgpool2-4.5.0.tar
$CMD pull ghcr.io/appscode-images/pgpool2:4.5.3 images/appscode-images-pgpool2-4.5.3.tar
$CMD pull ghcr.io/appscode-images/pgpool2_exporter:v1.2.2 images/appscode-images-pgpool2_exporter-v1.2.2.tar
$CMD pull ghcr.io/appscode-images/postgres:10.23-alpine images/appscode-images-postgres-10.23-alpine.tar
$CMD pull ghcr.io/appscode-images/postgres:10.23-bullseye images/appscode-images-postgres-10.23-bullseye.tar
$CMD pull ghcr.io/appscode-images/postgres:11.22-alpine images/appscode-images-postgres-11.22-alpine.tar
$CMD pull ghcr.io/appscode-images/postgres:11.22-bookworm images/appscode-images-postgres-11.22-bookworm.tar
$CMD pull ghcr.io/appscode-images/postgres:12.17-alpine images/appscode-images-postgres-12.17-alpine.tar
$CMD pull ghcr.io/appscode-images/postgres:12.17-bookworm images/appscode-images-postgres-12.17-bookworm.tar
$CMD pull ghcr.io/appscode-images/postgres:13.13-alpine images/appscode-images-postgres-13.13-alpine.tar
$CMD pull ghcr.io/appscode-images/postgres:13.13-bookworm images/appscode-images-postgres-13.13-bookworm.tar
$CMD pull ghcr.io/appscode-images/postgres:14.10-alpine images/appscode-images-postgres-14.10-alpine.tar
$CMD pull ghcr.io/appscode-images/postgres:14.10-bookworm images/appscode-images-postgres-14.10-bookworm.tar
$CMD pull ghcr.io/appscode-images/postgres:14.13-alpine images/appscode-images-postgres-14.13-alpine.tar
$CMD pull ghcr.io/appscode-images/postgres:14.13-bookworm images/appscode-images-postgres-14.13-bookworm.tar
$CMD pull ghcr.io/appscode-images/postgres:15.5-alpine images/appscode-images-postgres-15.5-alpine.tar
$CMD pull ghcr.io/appscode-images/postgres:15.5-bookworm images/appscode-images-postgres-15.5-bookworm.tar
$CMD pull ghcr.io/appscode-images/postgres:15.8-alpine images/appscode-images-postgres-15.8-alpine.tar
$CMD pull ghcr.io/appscode-images/postgres:15.8-bookworm images/appscode-images-postgres-15.8-bookworm.tar
$CMD pull ghcr.io/appscode-images/postgres:16.1-alpine images/appscode-images-postgres-16.1-alpine.tar
$CMD pull ghcr.io/appscode-images/postgres:16.1-bookworm images/appscode-images-postgres-16.1-bookworm.tar
$CMD pull ghcr.io/appscode-images/postgres:16.4-alpine images/appscode-images-postgres-16.4-alpine.tar
$CMD pull ghcr.io/appscode-images/postgres:16.4-bookworm images/appscode-images-postgres-16.4-bookworm.tar
$CMD pull ghcr.io/appscode-images/rabbitmq:3.12.12-management-alpine images/appscode-images-rabbitmq-3.12.12-management-alpine.tar
$CMD pull ghcr.io/appscode-images/rabbitmq:3.13.2-management-alpine images/appscode-images-rabbitmq-3.13.2-management-alpine.tar
$CMD pull ghcr.io/appscode-images/redis:5.0.14-bullseye images/appscode-images-redis-5.0.14-bullseye.tar
$CMD pull ghcr.io/appscode-images/redis:6.0.20-bookworm images/appscode-images-redis-6.0.20-bookworm.tar
$CMD pull ghcr.io/appscode-images/redis:6.2.14-bookworm images/appscode-images-redis-6.2.14-bookworm.tar
$CMD pull ghcr.io/appscode-images/redis:7.0.14-bookworm images/appscode-images-redis-7.0.14-bookworm.tar
$CMD pull ghcr.io/appscode-images/redis:7.0.15-bookworm images/appscode-images-redis-7.0.15-bookworm.tar
$CMD pull ghcr.io/appscode-images/redis:7.2.3-bookworm images/appscode-images-redis-7.2.3-bookworm.tar
$CMD pull ghcr.io/appscode-images/redis:7.2.4-bookworm images/appscode-images-redis-7.2.4-bookworm.tar
$CMD pull ghcr.io/appscode-images/redis:7.4.0-bookworm images/appscode-images-redis-7.4.0-bookworm.tar
$CMD pull ghcr.io/appscode-images/singlestore-node:alma-8.1.32-e3d3cde6da images/appscode-images-singlestore-node-alma-8.1.32-e3d3cde6da.tar
$CMD pull ghcr.io/appscode-images/singlestore-node:alma-8.5.30-4f46ab16a5 images/appscode-images-singlestore-node-alma-8.5.30-4f46ab16a5.tar
$CMD pull ghcr.io/appscode-images/singlestore-node:alma-8.5.7-bf633c1a54 images/appscode-images-singlestore-node-alma-8.5.7-bf633c1a54.tar
$CMD pull ghcr.io/appscode-images/singlestore-node:alma-8.7.10-95e2357384 images/appscode-images-singlestore-node-alma-8.7.10-95e2357384.tar
$CMD pull ghcr.io/appscode-images/solr:8.11.2 images/appscode-images-solr-8.11.2.tar
$CMD pull ghcr.io/appscode-images/solr:9.4.1 images/appscode-images-solr-9.4.1.tar
$CMD pull ghcr.io/appscode-images/solr:9.6.1 images/appscode-images-solr-9.6.1.tar
$CMD pull ghcr.io/appscode-images/zookeeper:3.7.2 images/appscode-images-zookeeper-3.7.2.tar
$CMD pull ghcr.io/appscode-images/zookeeper:3.8.3 images/appscode-images-zookeeper-3.8.3.tar
$CMD pull ghcr.io/appscode-images/zookeeper:3.9.1 images/appscode-images-zookeeper-3.9.1.tar
$CMD pull ghcr.io/appscode/k8s-wait-for:v2.0 images/appscode-k8s-wait-for-v2.0.tar
$CMD pull ghcr.io/appscode/kube-rbac-proxy:v0.11.0 images/appscode-kube-rbac-proxy-v0.11.0.tar
$CMD pull ghcr.io/appscode/kubectl-nonroot:1.25 images/appscode-kubectl-nonroot-1.25.tar
$CMD pull ghcr.io/appscode/petset:v0.0.7 images/appscode-petset-v0.0.7.tar
$CMD pull ghcr.io/appscode/sidekick:v0.0.8 images/appscode-sidekick-v0.0.8.tar
$CMD pull ghcr.io/kubedb/cassandra-init:4.1.6 images/kubedb-cassandra-init-4.1.6.tar
$CMD pull ghcr.io/kubedb/cassandra-init:5.0.0 images/kubedb-cassandra-init-5.0.0.tar
$CMD pull ghcr.io/kubedb/clickhouse-init:24.4.1 images/kubedb-clickhouse-init-24.4.1.tar
$CMD pull ghcr.io/kubedb/dashboard-restic-plugin:v0.1.0 images/kubedb-dashboard-restic-plugin-v0.1.0.tar
$CMD pull ghcr.io/kubedb/dashboard-restic-plugin:v0.6.0 images/kubedb-dashboard-restic-plugin-v0.6.0.tar
$CMD pull ghcr.io/kubedb/druid-init:25.0.0 images/kubedb-druid-init-25.0.0.tar
$CMD pull ghcr.io/kubedb/druid-init:28.0.1 images/kubedb-druid-init-28.0.1.tar
$CMD pull ghcr.io/kubedb/druid-init:30.0.0 images/kubedb-druid-init-30.0.0.tar
$CMD pull ghcr.io/kubedb/elasticsearch-restic-plugin:v0.11.0 images/kubedb-elasticsearch-restic-plugin-v0.11.0.tar
$CMD pull ghcr.io/kubedb/kubedb-autoscaler:v0.33.0 images/kubedb-kubedb-autoscaler-v0.33.0.tar
$CMD pull ghcr.io/kubedb/kubedb-crd-manager:v0.3.0 images/kubedb-kubedb-crd-manager-v0.3.0.tar
$CMD pull ghcr.io/kubedb/kubedb-kibana:v0.24.0 images/kubedb-kubedb-kibana-v0.24.0.tar
$CMD pull ghcr.io/kubedb/kubedb-manifest-plugin:v0.11.0 images/kubedb-kubedb-manifest-plugin-v0.11.0.tar
$CMD pull ghcr.io/kubedb/kubedb-ops-manager:v0.35.0 images/kubedb-kubedb-ops-manager-v0.35.0.tar
$CMD pull ghcr.io/kubedb/kubedb-provisioner:v0.48.0 images/kubedb-kubedb-provisioner-v0.48.0.tar
$CMD pull ghcr.io/kubedb/kubedb-schema-manager:v0.24.0 images/kubedb-kubedb-schema-manager-v0.24.0.tar
$CMD pull ghcr.io/kubedb/kubedb-ui-server:v0.24.0 images/kubedb-kubedb-ui-server-v0.24.0.tar
$CMD pull ghcr.io/kubedb/kubedb-webhook-server:v0.24.0 images/kubedb-kubedb-webhook-server-v0.24.0.tar
$CMD pull ghcr.io/kubedb/mariadb-archiver:v0.4.0_10.10.7-jammy images/kubedb-mariadb-archiver-v0.4.0_10.10.7-jammy.tar
$CMD pull ghcr.io/kubedb/mariadb-archiver:v0.4.0_10.11.6-jammy images/kubedb-mariadb-archiver-v0.4.0_10.11.6-jammy.tar
$CMD pull ghcr.io/kubedb/mariadb-archiver:v0.4.0_10.4.32-focal images/kubedb-mariadb-archiver-v0.4.0_10.4.32-focal.tar
$CMD pull ghcr.io/kubedb/mariadb-archiver:v0.4.0_10.5.23-focal images/kubedb-mariadb-archiver-v0.4.0_10.5.23-focal.tar
$CMD pull ghcr.io/kubedb/mariadb-archiver:v0.4.0_10.6.16-focal images/kubedb-mariadb-archiver-v0.4.0_10.6.16-focal.tar
$CMD pull ghcr.io/kubedb/mariadb-archiver:v0.4.0_11.0.4-jammy images/kubedb-mariadb-archiver-v0.4.0_11.0.4-jammy.tar
$CMD pull ghcr.io/kubedb/mariadb-archiver:v0.4.0_11.1.3-jammy images/kubedb-mariadb-archiver-v0.4.0_11.1.3-jammy.tar
$CMD pull ghcr.io/kubedb/mariadb-archiver:v0.4.0_11.2.2-jammy images/kubedb-mariadb-archiver-v0.4.0_11.2.2-jammy.tar
$CMD pull ghcr.io/kubedb/mariadb-coordinator:v0.28.0 images/kubedb-mariadb-coordinator-v0.28.0.tar
$CMD pull ghcr.io/kubedb/mariadb-csi-snapshotter-plugin:v0.8.0 images/kubedb-mariadb-csi-snapshotter-plugin-v0.8.0.tar
$CMD pull ghcr.io/kubedb/mariadb-init:0.5.2 images/kubedb-mariadb-init-0.5.2.tar
$CMD pull ghcr.io/kubedb/mariadb-restic-plugin:v0.6.0 images/kubedb-mariadb-restic-plugin-v0.6.0.tar
$CMD pull ghcr.io/kubedb/mongodb-csi-snapshotter-plugin:v0.9.0 images/kubedb-mongodb-csi-snapshotter-plugin-v0.9.0.tar
$CMD pull ghcr.io/kubedb/mongodb-init:4.2-v9 images/kubedb-mongodb-init-4.2-v9.tar
$CMD pull ghcr.io/kubedb/mongodb-init:6.0-v10 images/kubedb-mongodb-init-6.0-v10.tar
$CMD pull ghcr.io/kubedb/mongodb_exporter:v0.40.0 images/kubedb-mongodb_exporter-v0.40.0.tar
$CMD pull ghcr.io/kubedb/mssql-coordinator:v0.3.0 images/kubedb-mssql-coordinator-v0.3.0.tar
$CMD pull ghcr.io/kubedb/mssql-exporter:1.1.0 images/kubedb-mssql-exporter-1.1.0.tar
$CMD pull ghcr.io/kubedb/mssql-init:2022-ubuntu-22-v3 images/kubedb-mssql-init-2022-ubuntu-22-v3.tar
$CMD pull ghcr.io/kubedb/mssqlserver-archiver:v0.0.1 images/kubedb-mssqlserver-archiver-v0.0.1.tar
$CMD pull ghcr.io/kubedb/mssqlserver-walg-plugin:v0.0.1 images/kubedb-mssqlserver-walg-plugin-v0.0.1.tar
$CMD pull ghcr.io/kubedb/mysql-archiver:v0.9.0_5.7.44 images/kubedb-mysql-archiver-v0.9.0_5.7.44.tar
$CMD pull ghcr.io/kubedb/mysql-archiver:v0.9.0_8.0.35 images/kubedb-mysql-archiver-v0.9.0_8.0.35.tar
$CMD pull ghcr.io/kubedb/mysql-archiver:v0.9.0_8.1.0 images/kubedb-mysql-archiver-v0.9.0_8.1.0.tar
$CMD pull ghcr.io/kubedb/mysql-archiver:v0.9.0_8.2.0 images/kubedb-mysql-archiver-v0.9.0_8.2.0.tar
$CMD pull ghcr.io/kubedb/mysql-coordinator:v0.26.0 images/kubedb-mysql-coordinator-v0.26.0.tar
$CMD pull ghcr.io/kubedb/mysql-csi-snapshotter-plugin:v0.9.0 images/kubedb-mysql-csi-snapshotter-plugin-v0.9.0.tar
$CMD pull ghcr.io/kubedb/mysql-init:5.7-v4 images/kubedb-mysql-init-5.7-v4.tar
$CMD pull ghcr.io/kubedb/mysql-init:8.0.31-v3 images/kubedb-mysql-init-8.0.31-v3.tar
$CMD pull ghcr.io/kubedb/mysql-init:8.4.2-v1 images/kubedb-mysql-init-8.4.2-v1.tar
$CMD pull ghcr.io/kubedb/mysql-router-init:v0.26.0 images/kubedb-mysql-router-init-v0.26.0.tar
$CMD pull ghcr.io/kubedb/mysqld-exporter:v0.13.1 images/kubedb-mysqld-exporter-v0.13.1.tar
$CMD pull ghcr.io/kubedb/percona-xtradb-coordinator:v0.21.0 images/kubedb-percona-xtradb-coordinator-v0.21.0.tar
$CMD pull ghcr.io/kubedb/percona-xtradb-init:0.2.1 images/kubedb-percona-xtradb-init-0.2.1.tar
$CMD pull ghcr.io/kubedb/pg-coordinator:v0.32.0 images/kubedb-pg-coordinator-v0.32.0.tar
$CMD pull ghcr.io/kubedb/pgbouncer:1.17.0 images/kubedb-pgbouncer-1.17.0.tar
$CMD pull ghcr.io/kubedb/pgbouncer:1.18.0 images/kubedb-pgbouncer-1.18.0.tar
$CMD pull ghcr.io/kubedb/pgbouncer:1.23.1 images/kubedb-pgbouncer-1.23.1.tar
$CMD pull ghcr.io/kubedb/pgbouncer_exporter:v0.1.1 images/kubedb-pgbouncer_exporter-v0.1.1.tar
$CMD pull ghcr.io/kubedb/postgres-archiver:v0.9.0_11.22-alpine images/kubedb-postgres-archiver-v0.9.0_11.22-alpine.tar
$CMD pull ghcr.io/kubedb/postgres-archiver:v0.9.0_11.22-bookworm images/kubedb-postgres-archiver-v0.9.0_11.22-bookworm.tar
$CMD pull ghcr.io/kubedb/postgres-archiver:v0.9.0_12.17-alpine images/kubedb-postgres-archiver-v0.9.0_12.17-alpine.tar
$CMD pull ghcr.io/kubedb/postgres-archiver:v0.9.0_12.17-bookworm images/kubedb-postgres-archiver-v0.9.0_12.17-bookworm.tar
$CMD pull ghcr.io/kubedb/postgres-archiver:v0.9.0_13.13-alpine images/kubedb-postgres-archiver-v0.9.0_13.13-alpine.tar
$CMD pull ghcr.io/kubedb/postgres-archiver:v0.9.0_13.13-bookworm images/kubedb-postgres-archiver-v0.9.0_13.13-bookworm.tar
$CMD pull ghcr.io/kubedb/postgres-archiver:v0.9.0_14.10-alpine images/kubedb-postgres-archiver-v0.9.0_14.10-alpine.tar
$CMD pull ghcr.io/kubedb/postgres-archiver:v0.9.0_14.10-bookworm images/kubedb-postgres-archiver-v0.9.0_14.10-bookworm.tar
$CMD pull ghcr.io/kubedb/postgres-archiver:v0.9.0_15.5-alpine images/kubedb-postgres-archiver-v0.9.0_15.5-alpine.tar
$CMD pull ghcr.io/kubedb/postgres-archiver:v0.9.0_15.5-bookworm images/kubedb-postgres-archiver-v0.9.0_15.5-bookworm.tar
$CMD pull ghcr.io/kubedb/postgres-archiver:v0.9.0_16.1-alpine images/kubedb-postgres-archiver-v0.9.0_16.1-alpine.tar
$CMD pull ghcr.io/kubedb/postgres-archiver:v0.9.0_16.1-bookworm images/kubedb-postgres-archiver-v0.9.0_16.1-bookworm.tar
$CMD pull ghcr.io/kubedb/postgres-csi-snapshotter-plugin:v0.9.0 images/kubedb-postgres-csi-snapshotter-plugin-v0.9.0.tar
$CMD pull ghcr.io/kubedb/postgres-init:0.15.0 images/kubedb-postgres-init-0.15.0.tar
$CMD pull ghcr.io/kubedb/postgres-restic-plugin:v0.10.0_16.1 images/kubedb-postgres-restic-plugin-v0.10.0_16.1.tar
$CMD pull ghcr.io/kubedb/provider-aws:v0.10.0 images/kubedb-provider-aws-v0.10.0.tar
$CMD pull ghcr.io/kubedb/provider-azure:v0.10.0 images/kubedb-provider-azure-v0.10.0.tar
$CMD pull ghcr.io/kubedb/provider-gcp:v0.10.0 images/kubedb-provider-gcp-v0.10.0.tar
$CMD pull ghcr.io/kubedb/proxysql-exporter:v1.1.0 images/kubedb-proxysql-exporter-v1.1.0.tar
$CMD pull ghcr.io/kubedb/proxysql:2.3.2-centos-v2 images/kubedb-proxysql-2.3.2-centos-v2.tar
$CMD pull ghcr.io/kubedb/proxysql:2.3.2-debian-v2 images/kubedb-proxysql-2.3.2-debian-v2.tar
$CMD pull ghcr.io/kubedb/proxysql:2.4.4-centos images/kubedb-proxysql-2.4.4-centos.tar
$CMD pull ghcr.io/kubedb/proxysql:2.4.4-debian images/kubedb-proxysql-2.4.4-debian.tar
$CMD pull ghcr.io/kubedb/proxysql:2.6.3-debian images/kubedb-proxysql-2.6.3-debian.tar
$CMD pull ghcr.io/kubedb/rabbitmq-init:3.12.12 images/kubedb-rabbitmq-init-3.12.12.tar
$CMD pull ghcr.io/kubedb/rabbitmq-init:3.13.2 images/kubedb-rabbitmq-init-3.13.2.tar
$CMD pull ghcr.io/kubedb/redis-coordinator:v0.27.0 images/kubedb-redis-coordinator-v0.27.0.tar
$CMD pull ghcr.io/kubedb/redis-init:0.9.0 images/kubedb-redis-init-0.9.0.tar
$CMD pull ghcr.io/kubedb/redis-restic-plugin:v0.11.0 images/kubedb-redis-restic-plugin-v0.11.0.tar
$CMD pull ghcr.io/kubedb/redis:4.0.11 images/kubedb-redis-4.0.11.tar
$CMD pull ghcr.io/kubedb/redis_exporter:1.58.0 images/kubedb-redis_exporter-1.58.0.tar
$CMD pull ghcr.io/kubedb/redis_exporter:v0.21.1 images/kubedb-redis_exporter-v0.21.1.tar
$CMD pull ghcr.io/kubedb/replication-mode-detector:v0.35.0 images/kubedb-replication-mode-detector-v0.35.0.tar
$CMD pull ghcr.io/kubedb/singlestore-coordinator:v0.3.0 images/kubedb-singlestore-coordinator-v0.3.0.tar
$CMD pull ghcr.io/kubedb/singlestore-init:8.1-v2 images/kubedb-singlestore-init-8.1-v2.tar
$CMD pull ghcr.io/kubedb/singlestore-init:8.5-v2 images/kubedb-singlestore-init-8.5-v2.tar
$CMD pull ghcr.io/kubedb/singlestore-init:8.7.10-v1 images/kubedb-singlestore-init-8.7.10-v1.tar
$CMD pull ghcr.io/kubedb/solr-init:8.11.2 images/kubedb-solr-init-8.11.2.tar
$CMD pull ghcr.io/kubedb/solr-init:9.4.1 images/kubedb-solr-init-9.4.1.tar
$CMD pull ghcr.io/kubedb/solr-init:9.6.1 images/kubedb-solr-init-9.6.1.tar
$CMD pull ghcr.io/kubedb/wal-g:v2024.5.24_mongo images/kubedb-wal-g-v2024.5.24_mongo.tar
$CMD pull ghcr.io/kubedb/zookeeper-init:3.7-v1 images/kubedb-zookeeper-init-3.7-v1.tar
$CMD pull ghcr.io/kubedb/zookeeper-restic-plugin:v0.4.0 images/kubedb-zookeeper-restic-plugin-v0.4.0.tar
$CMD pull mcr.microsoft.com/mssql/server:2022-CU12-ubuntu-22.04 images/mssql-server-2022-CU12-ubuntu-22.04.tar
$CMD pull mcr.microsoft.com/mssql/server:2022-CU14-ubuntu-22.04 images/mssql-server-2022-CU14-ubuntu-22.04.tar
$CMD pull mysql/mysql-router:8.0.31 images/mysql-mysql-router-8.0.31.tar
$CMD pull percona/percona-server-mongodb:4.2.24 images/percona-percona-server-mongodb-4.2.24.tar
$CMD pull percona/percona-server-mongodb:4.4.26 images/percona-percona-server-mongodb-4.4.26.tar
$CMD pull percona/percona-server-mongodb:5.0.23 images/percona-percona-server-mongodb-5.0.23.tar
$CMD pull percona/percona-server-mongodb:6.0.12 images/percona-percona-server-mongodb-6.0.12.tar
$CMD pull percona/percona-server-mongodb:7.0.4 images/percona-percona-server-mongodb-7.0.4.tar
$CMD pull percona/percona-xtradb-cluster:8.0.26 images/percona-percona-xtradb-cluster-8.0.26.tar
$CMD pull percona/percona-xtradb-cluster:8.0.28 images/percona-percona-xtradb-cluster-8.0.28.tar
$CMD pull percona/percona-xtradb-cluster:8.0.31 images/percona-percona-xtradb-cluster-8.0.31.tar
$CMD pull postgis/postgis:11-3.3 images/postgis-postgis-11-3.3.tar
$CMD pull postgis/postgis:12-3.4 images/postgis-postgis-12-3.4.tar
$CMD pull postgis/postgis:13-3.4 images/postgis-postgis-13-3.4.tar
$CMD pull postgis/postgis:14-3.4 images/postgis-postgis-14-3.4.tar
$CMD pull postgis/postgis:15-3.4 images/postgis-postgis-15-3.4.tar
$CMD pull postgis/postgis:16-3.4 images/postgis-postgis-16-3.4.tar
$CMD pull prom/memcached-exporter:v0.14.2 images/prom-memcached-exporter-v0.14.2.tar
$CMD pull prom/mysqld-exporter:v0.13.0 images/prom-mysqld-exporter-v0.13.0.tar
$CMD pull prometheuscommunity/elasticsearch-exporter:v1.7.0 images/prometheuscommunity-elasticsearch-exporter-v1.7.0.tar
$CMD pull prometheuscommunity/postgres-exporter:v0.15.0 images/prometheuscommunity-postgres-exporter-v0.15.0.tar
$CMD pull registry.k8s.io/git-sync/git-sync:v4.2.1 images/git-sync-git-sync-v4.2.1.tar
$CMD pull singlestore/cluster-in-a-box:alma-8.1.32-e3d3cde6da-4.0.16-1.17.6 images/singlestore-cluster-in-a-box-alma-8.1.32-e3d3cde6da-4.0.16-1.17.6.tar
$CMD pull singlestore/cluster-in-a-box:alma-8.5.22-fe61f40cd1-4.1.0-1.17.11 images/singlestore-cluster-in-a-box-alma-8.5.22-fe61f40cd1-4.1.0-1.17.11.tar
$CMD pull singlestore/cluster-in-a-box:alma-8.5.7-bf633c1a54-4.0.17-1.17.8 images/singlestore-cluster-in-a-box-alma-8.5.7-bf633c1a54-4.0.17-1.17.8.tar
$CMD pull singlestore/cluster-in-a-box:alma-8.7.10-95e2357384-4.1.0-1.17.14 images/singlestore-cluster-in-a-box-alma-8.7.10-95e2357384-4.1.0-1.17.14.tar
$CMD pull tianon/toybox:0.8.4 images/tianon-toybox-0.8.4.tar
$CMD pull timescale/timescaledb:2.14.2-pg13-oss images/timescale-timescaledb-2.14.2-pg13-oss.tar
$CMD pull timescale/timescaledb:2.14.2-pg14-oss images/timescale-timescaledb-2.14.2-pg14-oss.tar
$CMD pull timescale/timescaledb:2.14.2-pg15-oss images/timescale-timescaledb-2.14.2-pg15-oss.tar
$CMD pull timescale/timescaledb:2.14.2-pg16-oss images/timescale-timescaledb-2.14.2-pg16-oss.tar

tar -czvf images.tar.gz images