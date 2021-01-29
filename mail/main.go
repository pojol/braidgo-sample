package main

import (
	"braid-game/common"
	"braid-game/mail/handle"
	"braid-game/proto/api"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/pojol/braid"
	"github.com/pojol/braid/modules/grpcserver"
	"github.com/pojol/braid/modules/jaegertracing"
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

	b, err := braid.New(NodeName)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	var rpcserver braid.Module
	if localPort == 0 {
		rpcserver = braid.Server(grpcserver.Name, grpcserver.WithListen(":14301"))
	} else {
		addr := ":" + strconv.Itoa(localPort)
		rpcserver = braid.Server(grpcserver.Name, grpcserver.WithListen(addr))

		id := strconv.Itoa(int(time.Now().UnixNano())) + addr
		err := common.Regist(common.ConsulRegistReq{
			Name:    NodeName,
			ID:      id,
			Tags:    []string{"braid", NodeName},
			Address: "127.0.0.1",
			Port:    localPort,
		}, consulAddr)
		if err != nil {
			panic(err.Error())
		}

		defer common.Deregist(id, consulAddr)
	}

	b.RegistModule(
		rpcserver,
		braid.Tracing(jaegertracing.Name,
			jaegertracing.WithHTTP(jaegerAddr),
			jaegertracing.WithProbabilistic(0.1),
		))

	api.RegisterMailServer(braid.GetServer().(*grpc.Server), &handle.MailServer{})

	b.Init()
	b.Run()
	defer b.Close()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch

}
