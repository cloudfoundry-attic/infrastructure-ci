FROM relintdockerhubpushbot/cf-deployment-concourse-tasks
MAINTAINER https://github.com/cloudfoundry/infrastructure-ci

# Install gcloud
RUN echo "deb https://packages.cloud.google.com/apt cloud-sdk-trusty main" > /etc/apt/sources.list.d/google-cloud-sdk.list && \
  curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add - && \
  apt-get update && \
  apt-get -qqy install google-cloud-sdk
