FROM cfinfrastructure/cf-deployment-concourse-bbl-dev
MAINTAINER https://github.com/cloudfoundry/infrastructure-ci

RUN \
      apt-get update && \
      apt-get -qqy install --fix-missing \
            apt-transport-https \
            openssl \
            unzip \
      && \
      apt-get clean

# Install gcloud
RUN echo "deb https://packages.cloud.google.com/apt cloud-sdk-trusty main" > /etc/apt/sources.list.d/google-cloud-sdk.list && \
  curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add - && \
  apt-get update && \
  apt-get -qqy install google-cloud-sdk
