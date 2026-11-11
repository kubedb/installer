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

# Syncs the NATS CRDs from a kubedb/apimachinery checkout into this repo, the
# way crd-importer does (descriptions stripped, app.kubernetes.io/name label).
# Safe to rerun; usage: hack/sync-nats-crds.sh [path-to-apimachinery]

set -eou pipefail

api_dir=${1:-$(pwd)/../../apimachinery.worktrees/nats}
repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)

python3 -I - "${api_dir}/crds" "${repo_root}" <<'PY'
import os, re, sys
import yaml

src, root = sys.argv[1], sys.argv[2]

kubedb = [
    "autoscaling.kubedb.com_natsautoscalers.yaml",
    "catalog.kubedb.com_natsversions.yaml",
    "gitops.kubedb.com_natses.yaml",
    "kubedb.com_natses.yaml",
    "ops.kubedb.com_natsopsrequests.yaml",
]
catalog = "catalog.kubedb.com_natsversions.yaml"
group_order = ["archiver", "autoscaling", "catalog", "config", "elasticsearch",
               "gitops", "kafka", "kubedb", "ops", "postgres", "schema"]


class Dumper(yaml.SafeDumper):
    def choose_scalar_style(self):
        s = super().choose_scalar_style()
        return '"' if s == "'" else s


def strip(o):
    if isinstance(o, dict):
        return {k: strip(v) for k, v in o.items()
                if not (k == "description" and isinstance(v, str))}
    if isinstance(o, list):
        return [strip(x) for x in o]
    return o


def render(name):
    d = strip(yaml.safe_load(open(os.path.join(src, name))))
    d.setdefault("metadata", {}).setdefault("labels", {})["app.kubernetes.io/name"] = "kubedb"
    for k in ("annotations", "creationTimestamp"):
        d["metadata"].pop(k, None)
    d.pop("status", None)
    return yaml.dump(d, Dumper=Dumper, default_flow_style=False, sort_keys=True)


def write(path, text):
    os.makedirs(os.path.dirname(path), exist_ok=True)
    with open(path, "w") as f:
        f.write(text)


def doc_key(doc):
    m = re.search(r"^  name: (\S+)$", doc, re.M)
    plural, _, group = m.group(1).partition(".")
    return group, plural


def gidx(d):
    g = doc_key(d)[0]
    return group_order.index("kubedb" if g == "kubedb.com" else g.split(".")[0])


def merge_bundle(path, names):
    text = open(path).read()
    docs = [d.strip("\n") for d in re.split(r"^---$", text, flags=re.M) if d.strip()]
    docs = [d for d in docs if not doc_key(d)[1].startswith("nats")]
    for name in names:
        new = render(name).strip("\n")
        group, plural = doc_key(new)
        gi = gidx(new)
        pos = None
        last = None
        for i, d in enumerate(docs):
            if gidx(d) == gi:
                last = i
                if pos is None and doc_key(d)[1] > plural:
                    pos = i
        if pos is None:
            pos = (last + 1) if last is not None else len(docs)
        docs.insert(pos, new)
    write(path, "\n---\n".join(d + "\n" for d in docs))


for n in kubedb:
    write(os.path.join(root, "charts/kubedb-crds/crds", n), render(n))
write(os.path.join(root, "charts/kubedb-catalog/crds", catalog), render(catalog))
write(os.path.join(root, "charts/kubedb-certified-crds/crds", catalog), render(catalog))
merge_bundle(os.path.join(root, "crds/kubedb-crds.yaml"), kubedb)
merge_bundle(os.path.join(root, "crds/kubedb-catalog-crds.yaml"), [catalog])
PY
