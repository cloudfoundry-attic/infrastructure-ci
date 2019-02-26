#!/bin/bash -exu

ROOT=${PWD}

function download_terraform() {
  local platform
  local zip_name

  platform=${1}
  zip_name="terraform_${TF_VERSION}_${platform}.zip"

  ZIP_URL="${URL_BASE}/${TF_VERSION}/${zip_name}"
  wget ${ZIP_URL} -q
  unzip -o ${zip_name}

  rm -rf ${zip_name}
}

pushd ${ROOT}/terraform-binaries > /dev/null
  for platform in "darwin_amd64" "linux_386" "linux_amd64"; do
    download_terraform ${platform}
    mv terraform "terraform_${platform}"
  done
  for platform in "windows_amd64"; do
    download_terraform ${platform}
    mv terraform.exe "terraform_${platform}.exe"
  done
popd > /dev/null
