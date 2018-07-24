FROM relintdockerhubpushbot/cf-deployment-concourse-tasks
MAINTAINER https://github.com/cloudfoundry/infrastructure-ci

ENV BOSH_CLI_VERSION 2.0.48

# Create directory for GOPATH
RUN mkdir -p /go/bin

ENV GOPATH /go

# add go and GOPATH/bin to PATH
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN \
      apt-get update && \
      apt-get -qqy install --fix-missing \
            silversearcher-ag \
            tree \
            apt-transport-https \
            openssl \
            unzip \
      && \
      apt-get clean

RUN go get -u github.com/jteeuwen/go-bindata/...

RUN wget https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-${BOSH_CLI_VERSION}-linux-amd64 && \
  mv bosh-cli-${BOSH_CLI_VERSION}-linux-amd64 /usr/local/bin/bosh && \
  chmod +x /usr/local/bin/bosh

# Install gcloud
RUN echo "deb https://packages.cloud.google.com/apt cloud-sdk-trusty main" > /etc/apt/sources.list.d/google-cloud-sdk.list && \
  curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add - && \
  apt-get update && \
  apt-get -qqy install google-cloud-sdk

# add bbl from latest commit
ADD bbl /usr/local/bin/

ADD bosh-bootloader /go/src/github.com/cloudfoundry/bosh-bootloader

RUN cp -r /go/src/github.com/cloudfoundry/bosh-bootloader /var/repos
