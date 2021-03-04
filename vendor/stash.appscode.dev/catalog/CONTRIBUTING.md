## Test Catalog Installation Locally

In order to test whether the installation script works or not for a new catalog or catalog version without publishing the catalog chart, you can follow the following steps:

**Deploy a local chart repository:**

Deploy a local chart repository as shown below:

```console
# create a directory where we will store the charts
$ mkdir local-repo

# run chart server
$ docker run --rm -it \
  -p 8080:8080 \
  -v $HOME/local-repo:/charts \
  -e STORAGE=local \
  -e STORAGE_LOCAL_ROOTDIR=/charts \
  chartmuseum/chartmuseum
```

**Publish catalog chart to the local repository:**

Publish the catalog chart to the local repository. An example of publishing `stash-postgres` chart for [stashed/postgres](https://github.com/stashed/postgres) repository is shown below.

```console
$ helm package charts/stash-postgres
$ mv ./stash-postgres-11.2.tgz $HOME/local-repo/
$ helm repo index $HOME/local-repo/
```

**Set `APPSCODE_CHART_REGISTRY_URL` env to point your local repository:**

```console
export APPSCODE_CHART_REGISTRY_URL=http://localhost:8080
```

Now, you can use the installation scripts to install catalogs from your local repository.
