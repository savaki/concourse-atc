FROM busybox
MAINTAINER Matt Ho

ENV ATC_TEMPLATES /opt/concourse/web/templates
ENV ATC_PUBLIC    /opt/concourse/web/public

EXPOSE 8080

ADD target/bin /opt/concourse/bin
ADD target/web /opt/concourse/web

CMD [ "/opt/concourse/bin/atcd" ]
