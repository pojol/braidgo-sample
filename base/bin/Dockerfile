FROM alpine

COPY ./base_linux /home/app/

WORKDIR /home/app

ENV SERVICE_TAGS=braid,base
ENV SERVICE_14201_NAME=base

ENV GOTRACEBACK=crash

EXPOSE 14201

ENTRYPOINT ./base_linux $0 $@