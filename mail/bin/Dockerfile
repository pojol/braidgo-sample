FROM alpine

COPY ./mail_linux /home/app/

WORKDIR /home/app

ENV SERVICE_TAGS=braid,mail
ENV SERVICE_14301_NAME=mail
ENV GOTRACEBACK=crash


EXPOSE 14301

ENTRYPOINT ./mail_linux $0 $@
