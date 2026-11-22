# DBaaS open-source license restrictions

Run `go run ./catalog/kubedb/gen-dbaas-license` from the repository root to
print the restrictions in the current and legacy JSON formats. The allowlist
covers the supported open-source database editions in the catalog; it must be
reviewed when catalog editions or their licenses change. An unrestricted
database entry applies to all of its distributions.

Redis 7.2 and earlier are BSD-licensed, while Redis 7.4 through 7.8 are not
open source. Redis 8 and later are also available under AGPLv3; Valkey is
BSD-licensed. See [Redis licensing](https://redis.io/legal/licenses/).

The TimescaleDB restrictions cover the PostgreSQL versions of the catalog's
`-oss` images. The restrictions describe PostgreSQL versions and
distributions, not individual image tags, so review them if another image is
added with the same version and distribution.

The legacy v1 format supports only one restriction per database. Consequently,
its output includes only the first rule for Redis and Postgres, omitting Valkey
and TimescaleDB. Use the current format when those variants are needed.
