package main

import (
	"braid-game/login/handle"
	"braid-game/proto"
	"braid-game/proto/api"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/pojol/braid"
	"github.com/pojol/braid/modules/discoverconsul"
	"github.com/pojol/braid/modules/grpcclient"
	"github.com/pojol/braid/modules/grpcserver"
	"github.com/pojol/braid/modules/mailboxnsq"
	"google.golang.org/grpc"
)

var (
	help bool

	consulAddr    string
	jaegerAddr    string
	redisAddr     string
	nsqLookupAddr string
	nsqdAddr      string
	localPort     int
)

func initFlag() {
	flag.BoolVar(&help, "h", false, "this help")

	flag.StringVar(&consulAddr, "consul", "http://127.0.0.1:8500", "set consul address")
	flag.StringVar(&nsqLookupAddr, "nsqlookup", "127.0.0.1:4161", "set nsq lookup address")
	flag.StringVar(&nsqdAddr, "nsqd", "127.0.0.1:4150", "set nsqd address")
	flag.StringVar(&redisAddr, "redis", "redis://127.0.0.1:6379/0", "set redis address")
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

	b, _ := braid.New(
		proto.ServiceLogin,
		mailboxnsq.WithLookupAddr([]string{nsqLookupAddr}),
		mailboxnsq.WithNsqdAddr([]string{nsqdAddr}))

	b.RegistModule(
		braid.Server(grpcserver.Name, grpcserver.WithListen(":14101")),
		braid.Discover(
			discoverconsul.Name,
			discoverconsul.WithConsulAddr(consulAddr),
			discoverconsul.WithBlacklist([]string{"gateway"})),
		braid.Client(grpcclient.Name),
	)

	api.RegisterLoginServer(braid.GetServer().(*grpc.Server), &handle.LoginServer{})

	b.Init()
	b.Run()
	defer b.Close()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
}
