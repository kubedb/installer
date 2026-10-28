#!/usr/bin/env python3

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

"""Check that the exporter metrics used by the NATS dashboards exist in a real scrape.

Usage: hack/scripts/check-nats-dashboard-metrics.py <metrics.txt> <dashboard.json>...

<metrics.txt> is the output of GET /metrics on a prometheus-nats-exporter (port 7777 of the
<name>-stats Service) run with -connz -routez -subz -varz -jsz=all -use_internal_server_id, against a
NATS that has at least one stream and one consumer. The dashboards are the chart templates under
charts/kubedb-{grafana,perses}-dashboards/dashboards/nats (Helm header lines are stripped).
Only gnatsd_*, jetstream_*, go_*, process_* and promhttp_* names are checked; kube-state, cadvisor
and kubedb_com_nats_* metrics come from other sources.
"""

import json
import re
import sys

EXPORTER = re.compile(r"^(gnatsd_|jetstream_|go_|process_|promhttp_)")
EXTERNAL = re.compile(r"^(kube_|kubelet_|container_|node_namespace_pod_container:|kubedb_com_nats_)")
FUNCS = set(
    "sum avg min max count rate irate increase by on group_left group_right without topk ceil "
    "clamp_min histogram_quantile or and unless vector time label_values up".split()
)


def exprs(o):
    if isinstance(o, dict):
        for k, v in o.items():
            if k in ("expr", "query", "definition") and isinstance(v, str):
                yield v
            elif k == "matchers":
                yield from v
            else:
                yield from exprs(v)
    elif isinstance(o, list):
        for v in o:
            yield from exprs(v)


def load(path):
    t = open(path).read()
    t = re.sub(r"^\{\{-[^\n]*\}\}\n", "", t, flags=re.M)
    t = re.sub(r'\{\{ `("(?:[^`])*")` \}\}', lambda m: m.group(1), t)
    t = re.sub(r"\n\s*\{\{-?\s*(if|else|end)[^\n]*\}\}", "", t)

    def sub(m):
        start = t.rfind("\n", 0, m.start()) + 1
        return "X" if t[start:m.start()].count('"') % 2 else '"X"'

    return json.loads(re.sub(r"\{\{[^\n]*?\}\}", sub, t))


have = {l.split("{")[0].split(" ")[0] for l in open(sys.argv[1]) if l.strip() and not l.startswith("#")}
have |= {n + s for n in list(have) for s in ("_count", "_sum", "_bucket")}
missing_total = 0
for path in sys.argv[2:]:
    names = set()
    for e in exprs(load(path)):
        e = re.sub(r'"(?:[^"\\]|\\.)*"', '""', e)
        for m in re.finditer(r"[A-Za-z_:][A-Za-z0-9_:]*", e):
            n = m.group(0)
            if n not in FUNCS and not n.startswith("__") and (EXPORTER.match(n) or EXTERNAL.match(n)):
                names.add(n)
    missing = sorted(n for n in names if EXPORTER.match(n) and n not in have)
    print("%s: %d metric names, %d missing from exporter" % (path, len(names), len(missing)))
    for n in missing:
        print("  MISSING", n)
    missing_total += len(missing)
sys.exit(1 if missing_total else 0)
