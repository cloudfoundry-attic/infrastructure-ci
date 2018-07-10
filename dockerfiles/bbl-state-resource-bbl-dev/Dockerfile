FROM cfinfrastructure/cf-deployment-concourse-tasks-bbl-dev as bbl-dev

RUN which bbl

RUN bbl version

FROM cfinfrastructure/bbl-state-resource
MAINTAINER https://github.com/cloudfoundry/infrastructure-ci

RUN which bbl

RUN bbl version

COPY --from=bbl-dev /go/bin/bbl /usr/local/bin/bbl

RUN bbl version
