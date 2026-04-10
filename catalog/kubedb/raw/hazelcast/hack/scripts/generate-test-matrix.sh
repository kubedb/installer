#!/usr/bin/env bash

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

show_help() {
    echo "/ok-to-test test=test_branch_or_default_master installer=installer_branch_or_default_master k8s=(*|comma separated versions) db=*** versions=comma_separated_versions profiles=comma_separated_profiles ssl"
}

k8sVersions=(v1.34.0)

elasticsearchVersions=(xpack-8.5.2 opensearch-2.5.0)
kafkaVersions=(3.4.0)
mariadbVersions=(10.10.2)
memcachedVersions=(1.5.22)
mongodbVersions=(4.4.6)
mysqlVersions=(8.0.32 5.7.41)
perconaXtraDBVersions=(8.0.28)
pgbouncerVersions=(1.18.0)
postgresVersions=(15.1)
proxysqlVersions=(2.4.4-debian)
redisVersions=(7.0.9 6.2.11)

declare -A CATALOG
# store array as a comma separated string as map value
CATALOG['elasticsearch']=$(echo ${elasticsearchVersions[@]})
CATALOG['kafka']=$(echo ${kafkaVersions[@]})
CATALOG['mariadb']=$(echo ${mariadbVersions[@]})
CATALOG['memcached']=$(echo ${memcachedVersions[@]})
CATALOG['mongodb']=$(echo ${mongodbVersions[@]})
CATALOG['mysql']=$(echo ${mysqlVersions[@]})
CATALOG['percona-xtradb']=$(echo ${perconaXtraDBVersions[@]})
CATALOG['pgbouncer']=$(echo ${pgbouncerVersions[@]})
CATALOG['postgres']=$(echo ${postgresVersions[@]})
CATALOG['proxysql']=$(echo ${proxysqlVersions[@]})
CATALOG['redis']=$(echo ${redisVersions[@]})

declare -a k8s=()
test='master'
installer='master'
# detect db from git repo name, if name is not a key in CATALOG, set it to blank
db=${GITHUB_REPOSITORY#"${GITHUB_REPOSITORY_OWNER}/"}
if [ ${CATALOG[$db]+_} ]; then
    echo "Running test for $db"
else
    db=
fi
declare -a versions=()
target=
profiles='all'
ssl=('false')

oldIFS=$IFS
IFS=' '
read -ra COMMENT <<<"$@"
IFS=$oldIFS

for ((i = 0; i < ${#COMMENT[@]}; i++)); do
    entry="${COMMENT[$i]}"

    case "$entry" in
        '/ok-to-test') ;;

        test*)
            test=$(echo $entry | sed -e 's/^[^=]*=//g')
            ;;

        installer*)
            installer=$(echo $entry | sed -e 's/^[^=]*=//g')
            ;;

        k8s*)
            v=$(echo $entry | sed -e 's/^[^=]*=//g')
            oldIFS=$IFS
            IFS=','
            read -ra k8s <<<"$v"
            IFS=$oldIFS
            ;;

        db*)
            db=$(echo $entry | sed -e 's/^[^=]*=//g')
            ;;

        versions*)
            v=$(echo $entry | sed -e 's/^[^=]*=//g')
            oldIFS=$IFS
            IFS=','
            read -ra versions <<<"$v"
            IFS=$oldIFS
            ;;

        target*)
            target=$(echo $entry | sed -e 's/^[^=]*=//g')
            ;;

        profiles*)
            profiles=$(echo $entry | sed -e 's/^[^=]*=//g')
            ;;

        ssl*)
            v=$(echo $entry | sed -e 's/^[^=]*=//g')
            oldIFS=$IFS
            IFS=','
            read -ra ssl <<<"$v"
            IFS=$oldIFS
            ;;

        *)
            show_help
            exit 1
            ;;
    esac
done

if [ -z "$db" ]; then
    echo "missing db=*** parameter"
    exit 1
fi

if [ ${#k8s[@]} -eq 0 ] || [ ${k8s[0]} == "*" ]; then
    # assign array to a variable
    k8s=("${k8sVersions[@]}")
fi

# https://wiki.nix-pro.com/view/BASH_associative_arrays#Check_if_key_exists
if [ ${CATALOG[$db]+_} ]; then
    if [ ${#versions[@]} -eq 0 ] || [ ${versions[0]} == "*" ]; then
        # convert string back to an array
        oldIFS=$IFS
        IFS=' '
        read -ra versions <<<"${CATALOG[$db]}"
        IFS=$oldIFS
    fi
else
    echo "Unknonwn database: $s"
    exit 1
fi

echo "test = $test"
echo "installer = $installer"
echo "k8s = ${k8s[@]}"
echo "db = $db"
echo "versions = ${versions[@]}"
echo "target = $target"
echo "profiles = ${profiles}"
echo "ssl = ${ssl[@]}"

matrix=()
for k in ${k8s[@]}; do
    for v in ${versions[@]}; do
        for s in ${ssl[@]}; do
            matrix+=($(jq -n -c --arg k "$k" --arg d "$db" --arg v "$v" --arg t "$target" --arg p "$profiles" --arg s "$s" '{"k8s":$k,"db":$d,"version":$v,"target":$t,"profiles":$p,"ssl":$s}'))
        done
    done
done

# https://stackoverflow.com/a/63046305/244009
function join() {
    local IFS="$1"
    shift
    echo "$*"
}
matrix=$(echo '{"include":['$(join , ${matrix[@]})']}')
echo matrix=$matrix >>$GITHUB_OUTPUT
echo test_ref=$test >>$GITHUB_OUTPUT
echo installer_ref=$installer >>$GITHUB_OUTPUT
