## catalog

## Generate image scripts

Uses [image-packer](https://github.com/kmodules/image-packer)

```bash
image-packer list --root-dir=charts --output-dir=catalog

image-packer generate-scripts --insecure --allow-nondistributable-artifacts \
    --output-dir=catalog \
    --src=catalog/imagelist.yaml

make add-license fmt
```
