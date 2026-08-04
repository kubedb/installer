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

# Syncs Perses dashboards from the opnpulse grafana-dashboards repo into
# charts/kubedb-perses-dashboards/dashboards. opnpulse is the source of truth;
# only transforms that are mandatory for the chart are applied:
#   1. drop the "-perses" filename suffix
#   2. unwrap query_result(<expr>) -> <expr> and fix the migration labelName
#   3. delete inline "global-ds-proxy" datasource refs (panels inherit $datasource)
#   4. escape {{...}} so Helm tpl preserves Perses legends instead of evaluating them
#   5. prepend the chart-convention $shared / $alerts header lines
#
# Usage: hack/scripts/sync-perses-dashboards.sh [OPNPULSE_DIR]
#   OPNPULSE_DIR defaults to ../../opnpulse/grafana-dashboards relative to repo root.

set -eou pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
OPNPULSE_DIR="${1:-${REPO_ROOT}/../../opnpulse/grafana-dashboards}"
DEST_ROOT="${REPO_ROOT}/charts/kubedb-perses-dashboards/dashboards"

if [ ! -d "$OPNPULSE_DIR" ]; then
    echo "opnpulse dir not found: $OPNPULSE_DIR" >&2
    exit 1
fi

# DBs to sync: every DB that has perses sources in opnpulse. A DB dir with no
# *-perses.json is skipped at runtime, so listing one that lacks sources is safe.
DBS=(cassandra clickhouse connectcluster druid elasticsearch hanadb hazelcast
    ignite kafka mariadb memcached milvus mongodb mssqlserver mysql neo4j oracle
    perconaxtradb pgbouncer pgpool postgres proxysql qdrant rabbitmq redis
    singlestore solr zookeeper)

JQ_FILTER='
walk(
  if type == "object" then
    ( if (.expr? | type == "string") and (.expr | startswith("query_result(")) and (.expr | endswith(")"))
      then .expr = (.expr[13:-1]) else . end )
    | ( if .labelName? == "migration_from_grafana_not_supported" then .labelName = "app" else . end )
    | ( if (.datasource? | type == "object") and (.datasource.name? == "global-ds-proxy")
        then del(.datasource) else . end )
  else . end
)'

HEADER='{{- $shared := and (eq .Values.app.name "") (eq .Values.app.namespace "") -}}
{{- $alerts := (eq $.Values.dashboard.alerts true) -}}'

count=0
for db in "${DBS[@]}"; do
    src_dir="${OPNPULSE_DIR}/${db}"
    dst_dir="${DEST_ROOT}/${db}"
    if [ ! -d "$src_dir" ]; then
        echo "skip ${db}: no source dir" >&2
        continue
    fi
    shopt -s nullglob
    srcs=("$src_dir"/*-perses.json)
    shopt -u nullglob
    if [ ${#srcs[@]} -eq 0 ]; then
        echo "skip ${db}: no *-perses.json" >&2
        continue
    fi
    mkdir -p "$dst_dir"
    for src in "${srcs[@]}"; do
        base="$(basename "$src")"
        out="${dst_dir}/${base%-perses.json}.json"
        {
            printf '%s\n' "$HEADER"
            jq "$JQ_FILTER" "$src" |
                perl -pe 's/(\{\{[^}]*\}\})/"{{ ".chr(96).$1.chr(96)." }}"/ge'
        } >"$out"
        count=$((count + 1))
        echo "synced ${db}/${base%-perses.json}.json"
    done
done

echo "done: ${count} dashboards synced"
