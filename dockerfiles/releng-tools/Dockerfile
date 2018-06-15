FROM debian:latest
MAINTAINER https://github.com/cloudfoundry/infrastructure-ci

ENV OM_VERSION 0.37.0
ENV TERRAFORM_VERSION 0.11.7

RUN \
      apt-get update && \
      apt-get -qqy install --fix-missing \
            apt-transport-https \
            unzip \
      && \
      apt-get clean

# Install om
RUN wget https://github.com/pivotal-cf/om/releases/download/${OM_VERSION}/om-linux && \
  mv om-linux /usr/local/bin/om && \
  chmod +x /usr/local/bin/om

# Install texplate
RUN wget https://github.com/pivotal-cf/texplate/releases/download/v0.2.0/texplate_linux_amd64 && \
  mv texplate_linux_amd64 /usr/local/bin/texplate && \
  chmod +x /usr/local/bin/texplate

# Install jq
RUN wget https://github.com/stedolan/jq/releases/download/jq-1.5/jq-linux64 && \
  mv jq-linux64 /usr/local/bin/jq && \
  chmod +x /usr/local/bin/jq

# Install terraform
RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  mv terraform /usr/local/bin/terraform

