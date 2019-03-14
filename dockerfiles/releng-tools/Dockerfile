FROM python:3.6.5-jessie
MAINTAINER https://github.com/cloudfoundry/infrastructure-ci

ENV TERRAFORM_VERSION 0.11.7
ENV BOSH_CLI_VERSION 5.1.2

RUN \
      apt-get update && \
      apt-get -qqy install --fix-missing \
            apt-transport-https \
            unzip \
            wget \
            curl \
            lsb-release \
            gnupg \
      && \
      apt-get clean

# Install jq
RUN wget https://github.com/stedolan/jq/releases/download/jq-1.5/jq-linux64 && \
  mv jq-linux64 /usr/local/bin/jq && \
  chmod +x /usr/local/bin/jq

# Install om
RUN wget "https://github.com/pivotal-cf/om/releases/download/$(curl -s https://api.github.com/repos/pivotal-cf/om/releases | jq -r '.[0].name' | tr -d "v")/om-linux" && \
  mv om-linux /usr/local/bin/om && \
  chmod +x /usr/local/bin/om

# Install texplate
RUN wget https://github.com/pivotal-cf/texplate/releases/download/v0.2.0/texplate_linux_amd64 && \
  mv texplate_linux_amd64 /usr/local/bin/texplate && \
  chmod +x /usr/local/bin/texplate

# Install terraform
RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  mv terraform /usr/local/bin/terraform

# Install gcloud
RUN export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)" && \
  echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
  curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
  apt-get update -y && apt-get install google-cloud-sdk -y

# Install bosh_cli v2
RUN wget https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-${BOSH_CLI_VERSION}-linux-amd64 && \
  mv bosh-cli-${BOSH_CLI_VERSION}-linux-amd64 /usr/local/bin/bosh && \
  chmod +x /usr/local/bin/bosh

# Install yq
RUN pip install yq

#Install credhub
RUN wget https://github.com/cloudfoundry-incubator/credhub-cli/releases/download/2.2.1/credhub-linux-2.2.1.tgz && \
  tar -xzvf credhub-linux-2.2.1.tgz -C /usr/local/bin && \
  chmod +x /usr/local/bin/credhub
