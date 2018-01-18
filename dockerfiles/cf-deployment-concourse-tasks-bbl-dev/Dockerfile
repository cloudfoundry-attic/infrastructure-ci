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
      && \
      apt-get clean

# go install latest bosh-bootloader develop
ADD bosh-bootloader /go/src/github.com/cloudfoundry/bosh-bootloader
RUN cd /go/src/github.com/cloudfoundry/bosh-bootloader > /dev/null && \
  git rev-parse HEAD && go install github.com/cloudfoundry/bosh-bootloader/bbl

RUN cp -r /go/src/github.com/cloudfoundry/bosh-bootloader /var/repos

RUN wget https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-${BOSH_CLI_VERSION}-linux-amd64 && \
  mv bosh-cli-${BOSH_CLI_VERSION}-linux-amd64 /usr/local/bin/bosh && \
  chmod +x /usr/local/bin/bosh
