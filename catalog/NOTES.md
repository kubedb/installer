# Catalog Notes

- `imagelist.yaml` — generated, aggregated image list.
- `README.md` — generated CVE report.
- `VersionMatrix.md` — generated DB <-> operator version matrix.
- `kubedb/raw/**` — Raw DBVersion YAMLs.
- `kubestash/raw/**` — Raw addon/function YAMLs.
- `scripts/{db}/imagelist.yaml` — per-component image list, regenerate via `./hack/scripts/update-catalog.sh`.

Image mirroring scripts:

| Scenario | Script(s) |
|---|---|
| Direct mirror to your registry | `copy-images.sh` |
| Air-gapped via tarball | `export-images.sh` then `import-images.sh` |
| Air-gapped k3s | `export-images.sh` then `import-into-k3s.sh` |
