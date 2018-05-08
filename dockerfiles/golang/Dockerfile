FROM cfinfrastructure/minimal
MAINTAINER https://github.com/cloudfoundry/infrastructure-ci

ENV GO_VERSION 1.10

# Install go
RUN wget https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz && \
  tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
  rm -rf go${GO_VERSION}.linux-amd64.tar.gz
RUN chown -R testuser:testuser /usr/local/go/pkg

# Create directory for GOPATH
RUN mkdir -p /go/bin

# set GOPATH
ENV GOPATH /go

# add go and GOPATH/bin to PATH
ENV PATH $PATH:$GOPATH/bin:/usr/local/go/bin

# install test dependencies
RUN go get github.com/onsi/ginkgo/...
RUN go get github.com/onsi/gomega/...

RUN chown -R testuser:testuser /go
