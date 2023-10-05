FROM alpine

COPY ./mail_linux /home/app/

WORKDIR /home/app

EXPOSE 14301

ENTRYPOINT ./mail_linux $0 $@
