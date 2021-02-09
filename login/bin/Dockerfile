FROM alpine

COPY ./login_linux /tmp/app/

WORKDIR /tmp/app

ENV SERVICE_TAGS=braid,login
ENV SERVICE_14101_NAME=login
ENV GOTRACEBACK=crash


EXPOSE 14101

ENTRYPOINT ./login_linux $0 $@
