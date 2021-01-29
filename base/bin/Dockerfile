FROM golang:alpine

COPY ./base_linux /tmp/app/

WORKDIR /tmp/app

ENV SERVICE_TAGS=braid,base
#ENV SERVICE_14201_NAME=base
ENV GOTRACEBACK=crash

EXPOSE 14201

ENTRYPOINT ./base_linux $0 $@
