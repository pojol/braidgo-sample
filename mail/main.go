package main

import (
	"braid-game/mail/handle"
	"braid-game/proto/api"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/pojol/braid"
	"github.com/pojol/braid/modules/grpcserver"
	"google.golang.org/grpc"
)

var (
	help bool

	consulAddr string
	jaegerAddr string
	localPort  int

	// NodeName 节点名
	NodeName = "mail"
)

func initFlag() {
	flag.BoolVar(&help, "h", false, "this help")

	flag.StringVar(&consulAddr, "consul", "http://127.0.0.1:8500", "set consul address")
	flag.StringVar(&jaegerAddr, "jaeger", "http://127.0.0.1:9411/api/v2/spans", "set jaeger address")
	flag.IntVar(&localPort, "localPort", 0, "run locally")
}

func main() {
	initFlag()

	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	b, _ := braid.New(NodeName)

	b.RegistModule(braid.Server(grpcserver.Name, grpcserver.WithListen(":14301")))

	api.RegisterMailServer(braid.GetServer().(*grpc.Server), &handle.MailServer{})

	b.Init()
	b.Run()
	defer b.Close()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch

}
