#!/usr/bin/env bash

source cf-deployment-concourse-tasks/shared-functions

main() {
  bosh -v

  setup_bosh_env_vars

  bosh delete-config -n --name=dns --type=runtime
}

main $@
