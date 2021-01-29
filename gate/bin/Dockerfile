FROM alpine

COPY ./gateway_linux /tmp/app/

WORKDIR /tmp/app

ENV SERVICE_TAGS=braid,gateway
#ENV SERVICE_14001_NAME=gateway
ENV SERVICE_6060_IGNORE=true
ENV GOTRACEBACK=crash

EXPOSE 14001

ENTRYPOINT ./gateway_linux $0 $@
