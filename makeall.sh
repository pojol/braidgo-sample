#! /bin/sh

cd build

rm gateway_linux
echo "build gateway_linux ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway_linux /Users/pojol/github/braidgo-sample/gate/main.go

docker build -f /Users/pojol/github/braidgo-sample/gate/Dockerfile -t braid-sample/gateway .

rm login_linux
echo "build login_linux ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o login_linux /Users/pojol/github/braidgo-sample/login/main.go

docker build -f /Users/pojol/github/braidgo-sample/login/Dockerfile -t braid-sample/login .


rm mail_linux
echo "build mail_linux ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mail_linux /Users/pojol/github/braidgo-sample/mail/main.go

docker build -f /Users/pojol/github/braidgo-sample/mail/Dockerfile -t braid-sample/mail .
