FROM cfinfrastructure/minimal

RUN \
      apt-get update && \
      apt-get -qqy install --fix-missing \
            unzip \
      && \
      apt-get clean

# Install terraform
RUN wget https://releases.hashicorp.com/terraform/0.10.8/terraform_0.10.8_linux_amd64.zip && \
  unzip terraform_0.10.8_linux_amd64.zip && \
  rm terraform_0.10.8_linux_amd64.zip && \
  mv terraform /usr/local/bin/terraform
