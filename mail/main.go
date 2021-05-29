package main

import (
	"braid-game/mail/constant"
	"braid-game/mail/handle"
	"braid-game/proto/api"
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pojol/braid-go"
	"github.com/pojol/braid-go/modules/grpcserver"
	"github.com/pojol/braid-go/modules/pubsubnsq"
	"google.golang.org/grpc"
)

var (
	help bool

	consulAddr    string
	jaegerAddr    string
	localPort     int
	nsqLookupAddr string
	nsqdTCP       string
	nsqdHttp      string

	// NodeName 节点名
	NodeName = "mail"
)

func initFlag() {
	flag.BoolVar(&help, "h", false, "this help")

	flag.StringVar(&consulAddr, "consul", "http://127.0.0.1:8500", "set consul address")
	flag.StringVar(&jaegerAddr, "jaeger", "http://127.0.0.1:9411/api/v2/spans", "set jaeger address")
	flag.StringVar(&nsqLookupAddr, "nsqlookupd", "127.0.0.1:4161", "set nsq lookup address")
	flag.StringVar(&nsqdTCP, "nsqdTCP", "127.0.0.1:4150", "set nsqd address")
	flag.StringVar(&nsqdHttp, "nsqdHTTP", "127.0.0.1:4151", "set nsqd address")
	flag.IntVar(&localPort, "localPort", 0, "run locally")
}

func main() {
	initFlag()
	rand.Seed(time.Now().UnixNano())

	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	constant.MailRandRecord = rand.Intn(10000)

	b, _ := braid.NewService("mail")
	b.Register(
		braid.Module(braid.LoggerZap),
		braid.Module(braid.PubsubNsq,
			pubsubnsq.WithLookupAddr([]string{nsqLookupAddr}),
			pubsubnsq.WithNsqdAddr([]string{nsqdTCP}, []string{nsqdHttp}),
		),
		braid.Module(braid.ServerGRPC, grpcserver.WithListen(":14301")),
	)

	api.RegisterMailServer(braid.Server().Server().(*grpc.Server), &handle.MailServer{})

	b.Init()
	b.Run()
	defer b.Close()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch

}
