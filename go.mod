module braid-game

go 1.15

require (
	github.com/fsouza/go-dockerclient v1.7.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.2
	github.com/labstack/echo/v4 v4.1.6
	github.com/mitchellh/mapstructure v1.3.3
	github.com/pojol/braid-go v1.2.13
	github.com/pojol/httpbot v0.0.0-00010101000000-000000000000
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1
	google.golang.org/grpc v1.27.1
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
)

replace github.com/pojol/braid-go => /Users/pojol/work/go15/src/braid-go

replace github.com/pojol/httpbot => /Users/pojol/work/go15/src/httpbot
