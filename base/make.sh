rm base_linux

echo "build base_linux ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o base_linux main.go

# build
docker build --platform amd64 -t braidgo/sample-base . --no-cache
