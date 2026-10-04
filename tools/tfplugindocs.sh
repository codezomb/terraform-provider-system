#!/bin/sh

# Runs a tfplugindocs command, generate by default, e.g. tools/tfplugindocs.sh validate
# The provider schema is exported with OpenTofu and passed to tfplugindocs, which otherwise downloads and runs another cli.

set -eu

work="$(mktemp -d)"
trap 'rm -rf "${work}"' EXIT

go build -o "${work}/terraform-provider-system" .

cat > "${work}/tofu.rc" <<RC
provider_installation {
  dev_overrides {
    "neuspaces/system" = "${work}"
  }

  direct {}
}
RC

cat > "${work}/main.tf" <<TF
terraform {
  required_providers {
    system = {
      source = "neuspaces/system"
    }
  }
}
TF

# tfplugindocs finds the provider by its short name
TF_CLI_CONFIG_FILE="${work}/tofu.rc" tofu -chdir="${work}" providers schema -json \
  | sed 's|"registry.opentofu.org/neuspaces/system"|"system"|' > "${work}/schema.json"

go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs "${1:-generate}" --provider-name system --providers-schema "${work}/schema.json"
