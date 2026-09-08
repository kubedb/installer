#!/usr/bin/env python3
"""
Generate catalog/kubedb/raw-dhi/ from catalog/kubedb/raw/.

For each non-deprecated YAML in raw/, replace spec.db.image with the
Docker Hardened Image (DHI) equivalent if one exists. Files/documents
for DB types or versions without a DHI equivalent are dropped.
"""

import os
import re
import sys

RAW_DIR = os.path.join(os.path.dirname(__file__), "raw")
OUT_DIR = os.path.join(os.path.dirname(__file__), "raw-dhi")

# DHI image name per directory (db type). None means no DHI support.
DHI_IMAGE_MAP = {
    "elasticsearch": None,   # handled specially by distribution
    "mysql":         "dhi.io/mysql",
    "mongodb":       "dhi.io/mongodb",
    "redis":         "dhi.io/redis",
    "postgres":      "dhi.io/postgres",
    "memcached":     "dhi.io/memcached",
    "kafka":         "dhi.io/kafka",
    "rabbitmq":      "dhi.io/rabbitmq",
    "clickhouse":    "dhi.io/clickhouse-server",
    "zookeeper":     "dhi.io/zookeeper",
}

# Per db-type: function(version_str) -> bool, is the version supported by DHI?
# version_str is e.g. "8.0.10", "17.9", "4.2.4"
# We use major version comparison where appropriate.
def parse_version(v):
    parts = v.split(".")
    try:
        return tuple(int(x) for x in parts)
    except ValueError:
        return (0,)

def supported_mysql(v):
    t = parse_version(v)
    # 8.0.x, 8.4.x, 9.6.x
    if t[0] == 8 and t[1] in (0, 4):
        return True
    if t[0] == 9 and t[1] == 6:
        return True
    return False

def supported_mongodb(v):
    t = parse_version(v)
    # 8.0.x, 8.2.x
    return t[0] == 8 and t[1] in (0, 2)

def supported_redis(v):
    t = parse_version(v)
    # 8.x only
    return t[0] == 8

def supported_postgres(v):
    t = parse_version(v)
    # 10 through 18
    return 10 <= t[0] <= 18

def supported_memcached(v):
    t = parse_version(v)
    # 1.6.x
    return t[0] == 1 and t[1] == 6

def supported_kafka(v):
    t = parse_version(v)
    # 4.1.x, 4.2.x
    return t[0] == 4 and t[1] in (1, 2)

def supported_rabbitmq(v):
    t = parse_version(v)
    # 4.2.x
    return t[0] == 4 and t[1] == 2

def supported_clickhouse(v):
    t = parse_version(v)
    # 25.3.x, 25.8.x, 26.2.x
    if t[0] == 25 and t[1] in (3, 7, 8, 12):
        return True
    if t[0] == 26 and t[1] in (2,):
        return True
    return False

def supported_zookeeper(v):
    t = parse_version(v)
    # 3.8.x, 3.9.x
    return t[0] == 3 and t[1] in (8, 9)

# elasticsearch is handled specially by distribution
def supported_elasticsearch_elasticstack(v):
    t = parse_version(v)
    # 8.19.x, 9.1.x, 9.2.x, 9.3.x
    if t[0] == 8 and t[1] == 19:
        return True
    if t[0] == 9 and t[1] in (1, 2, 3):
        return True
    return False

def supported_opensearch(v):
    t = parse_version(v)
    # only 3.x
    return t[0] == 3

VERSION_SUPPORT = {
    "mysql":      supported_mysql,
    "mongodb":    supported_mongodb,
    "redis":      supported_redis,
    "postgres":   supported_postgres,
    "memcached":  supported_memcached,
    "kafka":      supported_kafka,
    "rabbitmq":   supported_rabbitmq,
    "clickhouse": supported_clickhouse,
    "zookeeper":  supported_zookeeper,
}


def extract_tag(image_ref):
    """Extract tag from image reference like 'registry/repo:tag' -> 'tag'"""
    if ":" not in image_ref:
        return ""
    return image_ref.rsplit(":", 1)[1]


def strip_os_suffix(tag):
    """Strip OS/variant suffix from tag: '17.9-alpine' -> '17.9', '8.0.4-bookworm' -> '8.0.4'"""
    return re.split(r"[-](?:alpine|bookworm|jammy|focal|bullseye|buster|debian|management|enterprise|official|ubi|ubi8|ubi9)", tag)[0]


def to_major_minor(version):
    """Convert 'major.minor.patch' to 'major.minor', e.g. '8.0.4' -> '8.0', '17.9' -> '17.9'"""
    parts = version.split(".")
    if len(parts) >= 3:
        return ".".join(parts[:2])
    return version


def get_dhi_image(db_type, version, distribution=None):
    """
    Returns the DHI image string for a given db_type+version, or None if not supported.
    distribution is used for elasticsearch (ElasticStack vs OpenSearch).
    """
    if db_type == "elasticsearch":
        dist = (distribution or "").lower()
        if "elasticstack" in dist or "xpack" in dist or "searchguard" in dist:
            if not supported_elasticsearch_elasticstack(version):
                return None
            return f"dhi.io/elasticsearch:{version}"
        elif "opensearch" in dist:
            if not supported_opensearch(version):
                return None
            return f"dhi.io/opensearch:{version}"
        else:
            return None

    dhi_name = DHI_IMAGE_MAP.get(db_type)
    if dhi_name is None:
        return None

    support_fn = VERSION_SUPPORT.get(db_type)
    if support_fn and not support_fn(version):
        return None

    return f"{dhi_name}:{version}"


def process_document(doc_text, db_type):
    """
    Process a single YAML document (as raw text) and return modified text or None if should be dropped.
    We use regex-based replacement to preserve comments and formatting.
    """
    # Skip deprecated documents
    if re.search(r"^\s+deprecated:\s+true\s*$", doc_text, re.MULTILINE):
        return None

    # Find spec.db.image
    db_image_match = re.search(r"(^  db:\n(?:[ \t]+\S[^\n]*\n)*?    image:\s*)(\S+)", doc_text, re.MULTILINE)
    if not db_image_match:
        # No db image field - skip
        return None

    current_image = db_image_match.group(2)
    tag = extract_tag(current_image)
    if not tag:
        return None

    version = strip_os_suffix(tag)

    # Get distribution for elasticsearch
    distribution = None
    dist_match = re.search(r"^\s+distribution:\s+(\S+)", doc_text, re.MULTILINE)
    if dist_match:
        distribution = dist_match.group(1)

    dhi_image = get_dhi_image(db_type, version, distribution)
    if dhi_image is None:
        return None

    # Use major.minor as DHI tag when the full version has a patch component
    # DHI uses major.minor tags that track the latest patch (e.g. dhi.io/redis:8.0)
    mm_version = to_major_minor(version)
    if mm_version != version:
        # Rebuild the image ref with major.minor tag
        dhi_base = dhi_image.rsplit(":", 1)[0]
        dhi_image = f"{dhi_base}:{mm_version}"

    # Replace the image
    new_doc = doc_text[:db_image_match.start(2)] + dhi_image + doc_text[db_image_match.end(2):]
    return new_doc


def process_file(src_path, db_type):
    """
    Process a raw YAML file. Returns a list of processed document strings (may be empty).
    """
    with open(src_path, "r") as f:
        content = f.read()

    # Split on document separators (keep ---)
    # Split by "^---" but preserve the separator with each doc
    raw_docs = re.split(r"(?m)^---\n", content)

    processed_docs = []
    for i, doc in enumerate(raw_docs):
        if not doc.strip():
            continue
        result = process_document(doc, db_type)
        if result is not None:
            processed_docs.append(result)

    return processed_docs


def main():
    db_types = list(DHI_IMAGE_MAP.keys()) + ["elasticsearch"]
    files_written = 0
    files_skipped = 0

    for db_type in sorted(set(db_types)):
        src_dir = os.path.join(RAW_DIR, db_type)
        if not os.path.isdir(src_dir):
            continue

        out_dir = os.path.join(OUT_DIR, db_type)

        for filename in sorted(os.listdir(src_dir)):
            if not filename.endswith(".yaml"):
                continue
            # Skip deprecated files
            if filename.startswith("deprecated-"):
                continue
            # Skip psp files
            if filename.endswith("-psp.yaml"):
                continue

            src_path = os.path.join(src_dir, filename)
            docs = process_file(src_path, db_type)

            if not docs:
                files_skipped += 1
                continue

            os.makedirs(out_dir, exist_ok=True)
            out_path = os.path.join(out_dir, filename)
            with open(out_path, "w") as f:
                for i, doc in enumerate(docs):
                    if i > 0:
                        f.write("---\n")
                    f.write(doc)
                    if not doc.endswith("\n"):
                        f.write("\n")

            files_written += 1

    print(f"Done. Files written: {files_written}, skipped (no DHI support): {files_skipped}")


if __name__ == "__main__":
    main()
