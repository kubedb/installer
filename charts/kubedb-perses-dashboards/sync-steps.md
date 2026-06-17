# Plan: Sync Perses dashboards from opnpulse into `charts/kubedb-perses-dashboards`

## Goal
Re-derive **all** `charts/kubedb-perses-dashboards/dashboards/<db>/*.json` from the updated
`/Users/arnobkumarsaha/go/src/opnpulse/grafana-dashboards/<db>/*-perses.json`.
opnpulse is the **source of truth**: preserve its content; transform only what is a broken
Grafana-migration artifact or what breaks Helm `tpl` rendering. Done via a **re-runnable script**.

## Source-of-truth decisions (from review)
1. Scope: **re-sync ALL** DBs that have perses files in opnpulse (overwrite existing chart files + add missing).
2. Conversion: **scripted** (re-runnable for future syncs).
3. Field content: **keep opnpulse's newer content** (format/unit, dataLink drill-downs, value-column names,
   real `seriesNameFormat` legends). Only strip/fix what is mandatory (below).

## Conversion recipe (opnpulse `*-perses.json` -> chart `dashboards/<db>/<name>.json`)
Mandatory transforms only:
1. **Filename**: drop `-perses` suffix (`postgres_summary_dashboard-perses.json` -> `postgres_summary_dashboard.json`).
2. **Header**: prepend the two chart-convention lines (matches existing files):
   ```
   {{- $shared := and (eq .Values.app.name "") (eq .Values.app.namespace "") -}}
   {{- $alerts := (eq $.Values.dashboard.alerts true) -}}
   ```
3. **Fix migration artifacts** (broken Grafana->Perses leftovers, not real content):
   - Unwrap `query_result(<expr>)` -> `<expr>` in variable `expr`.
   - `"labelName": "migration_from_grafana_not_supported"` -> `"labelName": "app"`.
4. **Strip inline datasource refs** so panels inherit the dashboard `$datasource` variable
   (the kubedb deployment has no `global-ds-proxy` datasource -> would break every panel):
   remove every `"datasource": { "kind": "PrometheusDatasource", "name": "global-ds-proxy" }` object.
5. **Templating safety for `{{ }}`** in JSON string values (e.g. `"seriesNameFormat": "{{pod}}"`):
   Helm `tpl` (applied in `templates/shared/dashboard.yaml`) would try to evaluate `{{pod}}` and fail/empty it.
   - **Default: escape** so the Perses legend is preserved: `{{pod}}` -> `{{ "{{pod}}" }}` (renders back to `{{pod}}`).
   - This is strictly better than the old behavior (existing files **blanked** these to `""`, losing legends).
     Because we re-sync ALL files, every DB becomes consistent under the escape approach.
   - **OPEN QUESTION for approval**: escape (preserve legends, my recommendation) vs blank (match old shipped files)?
6. Keep `id`/`uid` as-is — the template already does `omit ... "id" "uid"`, so they are harmless.
   (`colorMode`, `format`, `unit`, `dataLink`, `metricLabel`, value-column names are opnpulse content -> **kept**.)

## DB / file scope
Re-synced from opnpulse (perses present): cassandra(3), druid(3), elasticsearch(3), hazelcast(3),
kafka(3), mariadb(5), memcached(3), mongodb(3), mysql(4), pgbouncer(3), pgpool(3), postgres(3),
proxysql(3), rabbitmq(3), redis(5), singlestore(3), solr(3), zookeeper(3) = 18 DBs.

Net-new files this brings in (currently missing in chart): redis `redis_pod`, `redis_shards`, `sentinel_pod`;
rabbitmq `rabbitmq-pod`, `rabbitmq-summary`; proxysql `proxysql-summary`; singlestore `singlestore_summary`;
zookeeper `zookeeper_pod`.

Left untouched (no perses source in opnpulse, cannot re-sync): ignite, mssqlserver, perconaxtradb (chart-only).
Out of scope (no opnpulse perses at all): clickhouse, hanadb, milvus, neo4j, oracle, qdrant, connectcluster, db2, weaviate.

No edits needed to `values.yaml` featureGates or `data/resources.yaml` — every synced DB is already listed there.

## Implementation steps
1. Write `hack/scripts/sync-perses-dashboards.sh` (license header, `set -euo pipefail`):
   - Arg: `OPNPULSE_DIR` (default `../opnpulse/grafana-dashboards`).
   - For each DB dir in scope, for each `*-perses.json`:
     - `jq` pass: `del(.. | .datasource? // empty | select(.name=="global-ds-proxy"))`-style removal of the
       `global-ds-proxy` datasource objects; `query_result(...)` unwrap + labelName fix (string ops on the
       compacted JSON or jq `gsub`).
     - escape `{{...}}` -> `{{ "{{...}}" }}` (sed/perl on string values).
     - prepend header lines; write to `dashboards/<db>/<name>.json`.
   - Idempotent: re-running with unchanged opnpulse yields no diff.
2. Run the script.
3. Validate: `helm template charts/kubedb-perses-dashboards` renders without error (all PersesDashboard
   objects valid YAML/JSON); spot-check redis/rabbitmq/postgres output.
4. Optional: `make ct CT_COMMAND=lint TEST_CHARTS=kubedb-perses-dashboards`.
5. `git diff --stat` review; confirm only `dashboards/**` + the new script changed.

## Risks / notes
- Escape-vs-blank for `{{ }}` (step 5) is the one behavioral choice — needs your nod.
- jq drops insignificant whitespace; output will be reformatted vs hand-edited files (large but mechanical diff).
  If you want minimal diffs on already-synced DBs, we can match opnpulse's exact JSON formatting instead.
- Branch: `arnob/perses-sync`. plan.md stays untracked and is deleted before commit.
