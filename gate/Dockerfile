FROM alpine

COPY ./gateway_linux /tmp/app/

WORKDIR /tmp/app

EXPOSE 14001

ENTRYPOINT ./gateway_linux $0 $@
