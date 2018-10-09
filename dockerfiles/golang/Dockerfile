FROM cfinfrastructure/minimal
MAINTAINER https://github.com/cloudfoundry/infrastructure-ci

# Install go
ADD go /usr/local/go
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
