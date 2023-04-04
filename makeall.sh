#! /bin/sh

cd build

rm gateway_linux
echo "build gateway_linux ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway_linux ../gate/main.go

docker build -f ../gate/bin/Dockerfile -t braid-sample/gateway .

rm login_linux
echo "build login_linux ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o login_linux ../login/main.go

docker build -f ../login/bin/Dockerfile -t braid-sample/login .


rm mail_linux
echo "build mail_linux ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mail_linux ../mail/main.go

docker build -f ../mail/bin/Dockerfile -t braid-sample/mail .