FROM golang:1.11.4-stretch

RUN apt-get update && \
    apt-get -y install supervisor
RUN echo '[supervisord]\n\
nodaemon=true\n\
\n\
[program:frontend]\n\
command=/go/bin/frontend\n\
autorestart=true\n\
priority=100\n\
stdout_logfile=/dev/stdout\n\
stdout_logfile_maxbytes=0\n\
\n\
[program:backend]\n\
command=/go/bin/backend\n\
autorestart=true\n\
priority=200\n\
stdout_logfile=/dev/stdout\n\
stdout_logfile_maxbytes=0\n ' \
> /etc/supervisor/conf.d/megu.conf

RUN go get -v \
    github.com/go-sql-driver/mysql \
    github.com/gorilla/context \
    github.com/gorilla/mux \
    github.com/gorilla/sessions \
    github.com/kelseyhightower/envconfig \
    github.com/satori/go.uuid \
    golang.org/x/crypto/bcrypt \
    gopkg.in/gomail.v2 \
    gopkg.in/resty.v1

ENV APPNAME=bitbucket.org/mendelgusmao/me_gu
ENV APPDIR=/go/src/$APPNAME

ADD . $APPDIR

RUN go install \
    $APPNAME/frontend \
    $APPNAME/backend && \
    rm $APPDIR/frontend/templates/*.go && \
    mv $APPDIR/frontend/templates /srv && \
    rm -rf $APPDIR

EXPOSE 8000 8001

CMD ["/usr/bin/supervisord"]
