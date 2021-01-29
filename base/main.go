package main

import (
	"braid-game/base/handle"
	"braid-game/common"
	"braid-game/proto"
	"braid-game/proto/api"
	"flag"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/pojol/braid"
	"github.com/pojol/braid/modules/discoverconsul"
	"github.com/pojol/braid/modules/electorconsul"
	"github.com/pojol/braid/modules/grpcclient"
	"github.com/pojol/braid/modules/grpcserver"
	"github.com/pojol/braid/modules/jaegertracing"
	"github.com/pojol/braid/modules/linkerredis"
	"github.com/pojol/braid/modules/mailboxnsq"
	"google.golang.org/grpc"
)

var (
	help bool

	consulAddr    string
	jaegerAddr    string
	nsqLookupAddr string
	nsqdAddr      string
	redisAddr     string
	localPort     int
)

func initFlag() {
	flag.BoolVar(&help, "h", false, "this help")

	flag.StringVar(&consulAddr, "consul", "http://127.0.0.1:8500", "set consul address")
	flag.StringVar(&jaegerAddr, "jaeger", "http://127.0.0.1:9411/api/v2/spans", "set jaeger address")
	flag.StringVar(&nsqLookupAddr, "nsqlookup", "127.0.0.1:4161", "set nsq lookup address")
	flag.StringVar(&nsqdAddr, "nsqd", "127.0.0.1:4150", "set nsqd address")
	flag.StringVar(&redisAddr, "redis", "redis://127.0.0.1:6379/0", "set redis address")
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
		proto.ServiceBase,
		mailboxnsq.WithLookupAddr([]string{nsqLookupAddr}),
		mailboxnsq.WithNsqdAddr([]string{nsqdAddr}))

	var rpcserver braid.Module
	if localPort == 0 {
		rpcserver = braid.Server(grpcserver.Name, grpcserver.WithListen(":14201"))
	} else {
		addr := ":" + strconv.Itoa(localPort)
		rpcserver = braid.Server(grpcserver.Name, grpcserver.WithListen(addr))

		id := strconv.Itoa(int(time.Now().UnixNano())) + addr
		err := common.Regist(common.ConsulRegistReq{
			Name:    proto.ServiceBase,
			ID:      id,
			Tags:    []string{"braid", proto.ServiceBase},
			Address: "127.0.0.1",
			Port:    localPort,
		}, consulAddr)
		if err != nil {
			panic(err.Error())
		}

		defer common.Deregist(id, consulAddr)
	}

	b.RegistModule(
		braid.Discover(
			discoverconsul.Name,
			discoverconsul.WithConsulAddr(consulAddr)),
		braid.Client(grpcclient.Name),
		rpcserver,
		braid.Elector(
			electorconsul.Name,
			electorconsul.WithConsulAddr(consulAddr),
		),
		braid.LinkCache(linkerredis.Name, linkerredis.WithRedisAddr(redisAddr)),
		braid.Tracing(jaegertracing.Name,
			jaegertracing.WithHTTP(jaegerAddr),
			jaegertracing.WithProbabilistic(0.1),
		))

	api.RegisterBaseServer(braid.GetServer().(*grpc.Server), &handle.BaseServer{})

	b.Init()
	b.Run()
	defer b.Close()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
}
