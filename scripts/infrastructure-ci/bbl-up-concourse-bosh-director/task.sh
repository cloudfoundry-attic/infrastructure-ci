#!/bin/bash -exu

function main() {
  local root_dir
  root_dir=${1}

  trap 'commit_bbl_state "${root_dir}"' EXIT

  local bbl_state_dir
  if [[ "${BBL_IAAS}" == "gcp" ]]; then
    bbl_state_dir="infra-ci"
  else
    bbl_state_dir="vsphere-concourse"
  fi

  pushd "${root_dir}/infrastructure-ci-bbl-states/${bbl_state_dir}" > /dev/null

    if [[ "${BBL_IAAS}" == "gcp" ]]; then
      bbl plan --lb-type concourse
    else
      bbl plan
    fi

    bbl up

  popd > /dev/null

}

function commit_bbl_state() {
  local root_dir
  root_dir=${1}

  local bbl_commit

  pushd "${root_dir}/bosh-bootloader" > /dev/null
    bbl_commit=$(git rev-parse HEAD)
  popd > /dev/null

  cp -r "${root_dir}/infrastructure-ci-bbl-states/." "${root_dir}/updated-bbl-states"

  pushd "${root_dir}/updated-bbl-states" > /dev/null

    git config --global user.email "ifra@pivotal.io"
    git config --global user.name "Infra CI Bot"

    git checkout master

    git add .
    git commit -m "update ${BBL_IAAS} concourse bbl director to bbl commit ${bbl_commit}"
  popd > /dev/null

}

main ${PWD}
