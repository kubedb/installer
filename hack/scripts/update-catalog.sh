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

# image-packer embeds the resourceeditors hub, but the reloader used by
# kmodules.xyz/resource-metadata will shadow the embedded copy with
# /tmp/hub/resourceeditors if that directory exists. A stale/empty leftover
# there makes image-packer load zero editors and emit a broken catalog, so
# remove it before running.
rm -rf /tmp/hub/resourceeditors

image-packer list --root-dir=charts --output-dir=catalog

image-packer generate-scripts --insecure --allow-nondistributable-artifacts \
    --output-dir=catalog \
    --src=catalog/imagelist.yaml

declare -a components=(
    aerospike
    cassandra
    clickhouse
    documentdb
    druid
    elasticsearch
    hazelcast
    kafka
    kafkaconnector
    mariadb
    memcached
    mongodb
    mssqlserver
    mysql
    neo4j
    oracle
    operator
    perconaxtradb
    pgbouncer
    pgpool
    postgres
    proxysql
    rabbitmq
    redis
    schemaregistry
    singlestore
    solr
    zookeeper
)

for component in "${components[@]}"; do
    image-packer generate-scripts --insecure --allow-nondistributable-artifacts \
        --output-dir=catalog/scripts/$component \
        --src=catalog/scripts/$component/imagelist.yaml
done

make add-license fmt
