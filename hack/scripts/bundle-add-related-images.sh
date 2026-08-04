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

# Adds spec.relatedImages (derived from the image/containerImage references
# already present in the CSV) and spec.replaces to a ClusterServiceVersion
# manifest. Edits the file in place via awk/sed line insertion rather than a
# YAML round-trip, so unrelated formatting is left untouched.
#
# Usage: bundle-add-related-images.sh <csv-file> [prev-version]
# prev-version is the previous release tag (e.g. v2026.6.19); when omitted,
# spec.replaces is left unset.

set -eou pipefail

CSV_FILE="$1"
PREV_VERSION="${2:-}"

# Normalize to a "v"-prefixed tag: pass through v2026.6.19 as-is, add the
# prefix to bare 2026.6.19.
case "$PREV_VERSION" in
    v*) ;;
    ?*) PREV_VERSION="v$PREV_VERSION" ;;
esac

images=$(grep -E '^ +(image|containerImage): ' "$CSV_FILE" | awk '{print $NF}' | sort -u)

blockfile=$(mktemp)
trap 'rm -f "$blockfile"' EXIT

{
    echo "  relatedImages:"
    while IFS= read -r img; do
        name=$(printf '%s' "$img" | sed -E 's#.*/##; s/@.*//; s/:.*//')
        echo "  - image: $img"
        echo "    name: $name"
    done <<<"$images"
    if [ -n "$PREV_VERSION" ]; then
        echo "  replaces: kubedb-installer.$PREV_VERSION"
    fi
} >"$blockfile"

awk -v blockfile="$blockfile" '
  /^  relatedImages:$/ { skip=1; next }
  skip && /^  [A-Za-z]/ { skip=0 }
  skip { next }
  /^  replaces:/ { next }
  /^  version:/ {
    while ((getline line < blockfile) > 0) print line
    close(blockfile)
  }
  { print }
' "$CSV_FILE" >"$CSV_FILE.tmp" && mv "$CSV_FILE.tmp" "$CSV_FILE"
