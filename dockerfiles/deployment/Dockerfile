FROM cfinfrastructure/golang
MAINTAINER https://github.com/cloudfoundry/infrastructure-ci

ENV TERRAFORM_VERSION 0.11.0
ENV BOSH_CLI_VERSION 2.0.44

RUN \
      apt-get update && \
      apt-get -qqy install --fix-missing \
            apt-transport-https \
            awscli \
            openssl \
            unzip \
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

# Install bosh_cli v2
RUN wget https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-${BOSH_CLI_VERSION}-linux-amd64 && \
  mv bosh-cli-${BOSH_CLI_VERSION}-linux-amd64 /usr/local/bin/bosh && \
  chmod +x /usr/local/bin/bosh

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

# Install bbl
RUN curl -s https://api.github.com/repos/cloudfoundry/bosh-bootloader/releases/latest | \
  sed -n '/browser_download_url.*linux/p' | \
  cut -d\" -f4 | xargs wget && \
  mv bbl-* /usr/local/bin/bbl && \
  chmod +x /usr/local/bin/bbl

# Install jq
RUN wget https://github.com/stedolan/jq/releases/download/jq-1.5/jq-linux64 && \
  mv jq-linux64 /usr/local/bin/jq && \
  chmod +x /usr/local/bin/jq

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
