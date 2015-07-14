FROM busybox
MAINTAINER Matt Ho

ENV ATC_TEMPLATES /opt/concourse/web/templates
ENV ATC_PUBLIC    /opt/concourse/web/public

EXPOSE 8080

ADD target /opt/concourse

CMD [ "/opt/concourse/bin/atcd" ]
