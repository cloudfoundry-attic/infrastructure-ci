FROM cfinfrastructure/golang
MAINTAINER https://github.com/cloudfoundry/infrastructure-ci

ARG GITHUB_TOKEN
ENV TERRAFORM_VERSION 0.11.0

RUN \
      apt-get update && \
      apt-get -qqy install --fix-missing \
            apt-transport-https \
            awscli \
            openssl \
            silversearcher-ag \
            unzip \
            tree \
      && \
      apt-get clean

# Install ruby-install
RUN curl https://codeload.github.com/postmodern/ruby-install/tar.gz/v0.5.0 | tar xvz -C /tmp/ && \
          cd /tmp/ruby-install-0.5.0 && \
          make install

# Install Ruby
RUN ruby-install ruby 2.2.4 -- --disable-install-rdoc

# Add ruby to PATH
ENV PATH $PATH:/home/root/.gem/ruby/2.2.4/bin:/opt/rubies/ruby-2.2.4/lib/ruby/gems/2.2.4/bin:/opt/rubies/ruby-2.2.4/bin

# Set permissions on ruby directory
RUN chmod -R 777 /opt/rubies/

# Install bundler
RUN /opt/rubies/ruby-2.2.4/bin/gem install bundler --no-rdoc --no-ri

# Install bosh_cli v1
RUN /opt/rubies/ruby-2.2.4/bin/gem install bosh_cli --no-rdoc --no-ri

# Install terraform
RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  mv terraform /usr/local/bin/terraform

# Install gcloud
RUN echo "deb https://packages.cloud.google.com/apt cloud-sdk-trusty main" > /etc/apt/sources.list.d/google-cloud-sdk.list && \
  curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add - && \
  apt-get update && \
  apt-get -qqy install google-cloud-sdk

# Install jq
RUN wget https://github.com/stedolan/jq/releases/download/jq-1.5/jq-linux64 && \
  mv jq-linux64 /usr/local/bin/jq && \
  chmod +x /usr/local/bin/jq

# Install bosh_cli v2
RUN curl -L "https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-$(curl -s https://api.github.com/repos/cloudfoundry/bosh-cli/releases | jq -r .[0].name | tr -d "v")-linux-amd64" -o "/usr/local/bin/bosh" && \
  chmod +x "/usr/local/bin/bosh"

# Install bbl
RUN curl -H "Authorization: token ${GITHUB_TOKEN}" -s https://api.github.com/repos/cloudfoundry/bosh-bootloader/releases/latest | \
  jq -r '.assets[] | .browser_download_url | select(contains("linux"))' | \
  xargs wget && \
  mv bbl-* /usr/local/bin/bbl && \
  chmod +x /usr/local/bin/bbl

# Install spiff
RUN wget https://github.com/cloudfoundry-incubator/spiff/releases/download/v1.0.7/spiff_linux_amd64 && \
  mv spiff_linux_amd64 /usr/local/bin/spiff && \
  chmod +x /usr/local/bin/spiff

# Install bosh-init
RUN wget https://s3.amazonaws.com/bosh-init-artifacts/bosh-init-0.0.98-linux-amd64 && \
  mv bosh-init-0.0.98-linux-amd64 /usr/local/bin/bosh-init && \
  chmod +x /usr/local/bin/bosh-init

RUN curl -L "https://cli.run.pivotal.io/stable?release=linux64-binary&source=github" | tar -zx && \
  chmod +x cf && \
  mv cf /usr/local/bin/cf

# Install Credhub
RUN CREDHUB_CLI_REPO="cloudfoundry-incubator/credhub-cli" && \
    CREDHUB_CLI_VERSION="$(curl -s https://api.github.com/repos/${CREDHUB_CLI_REPO}/releases | jq -r '.[0].name' | tr -d 'v')" && \
    wget -c "https://github.com/${CREDHUB_CLI_REPO}/releases/download/${CREDHUB_CLI_VERSION}/credhub-linux-${CREDHUB_CLI_VERSION}.tgz" -O - | tar -zx && \
    mv credhub /usr/local/bin/credhub && \
    chmod +x /usr/local/bin/credhub

RUN go get -u github.com/pivotal-cf/texplate

RUN go get -u github.com/jteeuwen/go-bindata/...
