FROM relintdockerhubpushbot/cf-deployment-concourse-tasks
MAINTAINER https://github.com/cloudfoundry/infrastructure-ci

# Create directory for GOPATH
RUN mkdir -p /go/bin

# set GOPATH
ENV GOPATH /go

# add go and GOPATH/bin to PATH
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# go install latest bosh-bootloader develop
RUN git clone https://github.com/cloudfoundry/bosh-bootloader /go/src/github.com/cloudfoundry/bosh-bootloader && \
  cd /go/src/github.com/cloudfoundry/bosh-bootloader > /dev/null && \
  git rev-parse HEAD && go install github.com/cloudfoundry/bosh-bootloader/bbl

# Install bosh_cli v2.0.44
RUN wget https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-2.0.44-linux-amd64 && \
  mv bosh-cli-2.0.44-linux-amd64 /usr/local/bin/bosh && \
  chmod +x /usr/local/bin/bosh
