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

# Syncs Grafana dashboards from the opnpulse dashboards repo into
# charts/kubedb-grafana-dashboards/dashboards. opnpulse is the source of truth;
# only transforms that are mandatory for the chart are applied:
#   1. prepend the chart-convention $shared / $alerts header lines
#   2. escape "{{...}}" string values so Helm tpl preserves Grafana legends
#      instead of evaluating them
#   3. keep the datasource template variable unconditional and first; wrap every
#      other variable in {{- if not $alerts }} ... {{- end }}
#   4. give the namespace / app variables a $shared (query) vs dedicated
#      (constant) branch so app.name / app.namespace pin them
#   5. give the top-level title a dedicated-mode variant suffixed with ns/name
# Everything else -- panels, exprs, units, id/uid -- is copied verbatim.
#
# Usage: charts/kubedb-grafana-dashboards/sync-grafana-dashboards.sh [OPNPULSE_DIR] [DB ...]
#   OPNPULSE_DIR defaults to ../../../go.open-pulse.dev/dashboards relative to repo root.
#   With no DB args, the DBs in DEFAULT_DBS are re-synced. Older DBs are not in that
#   list because their chart filenames were renamed away from the upstream basenames
#   (hazelcast_database_dashboard.json -> hazelcast-database.json), so syncing them by
#   name would add duplicates instead of updating in place. Pass such a DB explicitly
#   only after reconciling its filenames with upstream.

set -eou pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
CHART_DIR="${REPO_ROOT}/charts/kubedb-grafana-dashboards"
OPNPULSE_DIR="${1:-${REPO_ROOT}/../../../go.open-pulse.dev/dashboards}"
shift || true

if [ ! -d "$OPNPULSE_DIR" ]; then
    echo "opnpulse dashboards dir not found: $OPNPULSE_DIR" >&2
    exit 1
fi

DEFAULT_DBS=(clickhouse connectcluster elasticsearch hanadb mariadb milvus mongodb
    mssqlserver mysql neo4j oracle perconaxtradb pgbouncer postgres proxysql qdrant
    rabbitmq redis solr weaviate zookeeper)

DBS=("$@")
if [ ${#DBS[@]} -eq 0 ]; then
    DBS=("${DEFAULT_DBS[@]}")
fi

python3 - "$OPNPULSE_DIR" "$CHART_DIR" "${DBS[@]}" <<'PYEOF'
import json
import os
import re
import sys

HEADER = (
    '{{- $shared := and (eq .Values.app.name "") (eq .Values.app.namespace "") -}}\n'
    '{{- $alerts := (eq $.Values.dashboard.alerts true) -}}\n'
)

# template variables that must be pinned to the release's app in dedicated mode
CONST_VALUE = {"namespace": "$.Values.app.namespace", "app": "$.Values.app.name"}


def escape_mustaches(text):
    out = []
    for line in text.split("\n"):
        m = re.match(r'^(\s*)"([A-Za-z_][A-Za-z0-9_]*)": "(.*\{\{.*)"(,?)$', line)
        if m and "`" not in m.group(3):
            indent, key, val, comma = m.groups()
            out.append('%s"%s": {{ `"%s"` }}%s' % (indent, key, val, comma))
        else:
            out.append(line)
    return "\n".join(out)


def end_of_value(text, start):
    depth = 0
    in_str = False
    esc = False
    for i in range(start, len(text)):
        c = text[i]
        if in_str:
            if esc:
                esc = False
            elif c == "\\":
                esc = True
            elif c == '"':
                in_str = False
            continue
        if c == '"':
            in_str = True
        elif c in "{[":
            depth += 1
        elif c in "}]":
            depth -= 1
            if depth == 0:
                return i + 1
    raise ValueError("unbalanced JSON value")


def split_objects(list_body):
    items = []
    i = 0
    while i < len(list_body):
        if list_body[i] == "{":
            end = end_of_value(list_body, i)
            items.append(list_body[i:end])
            i = end
        else:
            i += 1
    return items


def templatize_var(obj_text, var_name):
    qm = re.search(r'\n(\s*)"query": ', obj_text)
    if not qm:
        raise ValueError('no "query" in %s variable' % var_name)
    indent = qm.group(1)

    # "type": "query" moves next to the query object, inside the conditional
    stripped = re.sub(r'\n%s"type": "query",' % indent, "", obj_text, count=1)
    if stripped == obj_text:
        stripped = re.sub(r',\n%s"type": "query"' % indent, "", obj_text, count=1)
        if stripped == obj_text:
            raise ValueError('no "type": "query" in %s variable' % var_name)

    qstart = stripped.index('"query": ', re.search(r'\n\s*"query": ', stripped).start())
    vstart = qstart + len('"query": ')
    vend = end_of_value(stripped, vstart)
    query_text = stripped[vstart:vend]

    block = (
        "{{- if $shared }}\n"
        '%(i)s"query": %(q)s,\n'
        '%(i)s"type": "query",\n'
        "%(i)s{{- else }}\n"
        '%(i)s"query": {{ %(v)s | quote }},\n'
        '%(i)s"type": "constant",\n'
        "%(i)s{{- end }}\n"
        "%(i)s" % {"i": indent, "q": query_text, "v": CONST_VALUE[var_name]}
    )

    tail = stripped[vend:]
    for prefix in (",", "\n", indent):
        if tail.startswith(prefix):
            tail = tail[len(prefix):]
    return stripped[:qstart] + block + tail


def transform(src, dst):
    text = open(src).read()
    if not text.endswith("\n"):
        text += "\n"

    lm = re.search(r'\n  "templating": \{\n    "list": ', text)
    if not lm:
        raise ValueError("no templating.list")
    lstart = text.index('"list": ', lm.start()) + len('"list": ')
    lend = end_of_value(text, lstart)

    datasource, rest = None, []
    for item in split_objects(text[lstart + 1:lend - 1]):
        obj = json.loads(item)
        if obj.get("type") == "datasource":
            datasource = item
        else:
            if obj.get("name") in CONST_VALUE:
                item = templatize_var(item, obj["name"])
            rest.append(item)
    if datasource is None:
        raise ValueError("no datasource variable")

    body = "      " + datasource + "\n"
    if rest:
        body += "      {{- if not $alerts }}\n      ,\n"
        body += ",\n".join("      " + r for r in rest) + "\n"
        body += "      {{- end }}\n"
    text = text[:lstart] + "[\n" + body + "    ]" + text[lend:]

    tm = re.search(r'\n  "title": "([^"]+)"(,?)\n', text)
    if not tm:
        raise ValueError("no top-level title")
    title, comma = tm.group(1), tm.group(2)
    text = text[:tm.start()] + (
        "\n  {{- if $shared }}\n"
        '  "title": "%s"%s\n'
        "  {{- else }}\n"
        '  "title": {{ printf "%s / %%s / %%s" $.Values.app.namespace $.Values.app.name | quote }}%s\n'
        "  {{- end }}\n" % (title, comma, title, comma)
    ) + text[tm.end():]

    os.makedirs(os.path.dirname(dst), exist_ok=True)
    with open(dst, "w") as f:
        f.write(HEADER + escape_mustaches(text))
    print("synced %s (%s)" % (os.path.relpath(dst), title))


up, chart, dbs = sys.argv[1], sys.argv[2], sys.argv[3:]
count = 0
for db in dbs:
    src_dir = os.path.join(up, db)
    if not os.path.isdir(src_dir):
        print("skip %s: no source dir" % db, file=sys.stderr)
        continue
    names = sorted(
        f for f in os.listdir(src_dir)
        if f.endswith(".json") and not f.endswith("-perses.json")
    )
    if not names:
        print("skip %s: no dashboards" % db, file=sys.stderr)
        continue
    for name in names:
        try:
            transform(os.path.join(src_dir, name), os.path.join(chart, "dashboards", db, name))
            count += 1
        except ValueError as e:
            print("skip %s/%s: %s" % (db, name, e), file=sys.stderr)
print("done: %d dashboards synced" % count)
PYEOF
