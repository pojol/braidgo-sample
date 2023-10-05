FROM alpine

COPY ./login_linux /tmp/app/

WORKDIR /tmp/app

EXPOSE 14101

ENTRYPOINT ./login_linux $0 $@
