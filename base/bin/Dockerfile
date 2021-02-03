FROM golang:alpine

COPY ./base_linux /tmp/app/

WORKDIR /tmp/app

ENV SERVICE_TAGS=braid,base
ENV SERVICE_14201_NAME=base
ENV GOTRACEBACK=crash

# health
ENV SERVICE_80_CHECK_HTTP=/health/endpoint/path
ENV SERVICE_80_CHECK_INTERVAL=15s
ENV SERVICE_80_CHECK_TIMEOUT=1s

EXPOSE 14201

ENTRYPOINT ./base_linux $0 $@
