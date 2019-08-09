# Pipelines for the Cloud Foundry Infrastructure team
- terraforming-(azure, gcp, aws, vsphere)
- bosh-bootloader
- bbl-latest
- socks5-proxy
- leftovers
- az-automation (cli for creating an azure service principal)

## Reconfigure

Use the `./reconfigure PIPELINE_NAME` to `fly set-pipeline` with the
necessary environment variables.

## Design notes

- [bosh-bootloader versioning](docs/bosh-bootloader-versioning.md)
