FROM cfinfrastructure/golang
MAINTAINER https://github.com/cloudfoundry/infrastructure-ci

ENV TERRAFORM_VERSION 0.11.0

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

# Install terraform
RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  mv terraform /usr/local/bin/terraform
