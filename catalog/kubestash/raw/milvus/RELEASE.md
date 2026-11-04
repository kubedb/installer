# Milvus object-storage TLS rollout

This catalog change must be promoted only after the corresponding worker images are published:

- `ghcr.io/kubedb/milvus-archiver:v0.2.1` (shared storage resolver, CA renewal and verified listener TLS/mTLS)
- `ghcr.io/kubedb/milvus-backup-plugin:v0.1.1` (logical trust bundle and physical restore namespace handling)
- `ghcr.io/kubedb/milvus-csi-snapshotter-plugin:v0.2.1` (matching shared resolver)

These tags are proposed release coordinates, not a claim that the images are already published. Do not merge or deploy this catalog until all three manifests are available and the updated provisioner/webhook APIs are deployed. The test cluster uses VM-built `tls-backup-test` images loaded directly into k3d; that tag must never be promoted as a release.

Keep Milvus listener TLS (`spec.tls`), live object-storage trust (`spec.objectStorage.tls`), and backup destination trust (`BackupStorage` credentials `CA_CERT_DATA`) independent. Existing custom pod-local CA paths must migrate to the typed object-storage CA selector before backup workers can use them.
