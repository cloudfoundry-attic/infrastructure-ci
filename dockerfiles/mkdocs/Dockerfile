FROM ubuntu:14.04
MAINTAINER https://github.com/cloudfoundry/infrastructure-ci

RUN \
      apt-get update && \
      apt-get install -y python3-pip && \
      pip3 install --upgrade pip && \
      pip3 install mkdocs-material

ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8

CMD ["/bin/sh"]
