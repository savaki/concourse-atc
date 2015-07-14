FROM ubuntu:14.04
MAINTAINER Matt Ho

ENV ATC_TEMPLATES /opt/concourse/web/templates
ENV ATC_PUBLIC    /opt/concourse/web/public
ENV ATC_ATC       /opt/concourse/bin/atc

EXPOSE 8080

ADD target/bin /opt/concourse/bin
ADD target/web /opt/concourse/web

CMD [ "/opt/concourse/bin/atcd" ]
