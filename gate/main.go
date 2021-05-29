package main

import (
	bm "braid-game/gate/middleware"
	"braid-game/gate/routes"
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "net/http/pprof"

	"github.com/labstack/echo/v4"
	"github.com/pojol/braid-go"
	"github.com/pojol/braid-go/modules/discoverconsul"
	"github.com/pojol/braid-go/modules/electorconsul"
	"github.com/pojol/braid-go/modules/jaegertracing"
	"github.com/pojol/braid-go/modules/linkerredis"
	"github.com/pojol/braid-go/modules/pubsubnsq"
)

var (
	help bool

	consulAddr    string
	redisAddr     string
	jaegerAddr    string
	nsqLookupAddr string
	nsqdTCP       string
	nsqdHttp      string
	localPort     int
)

const (
	// NodeName box 节点名
	NodeName = "gateway"
)

func initFlag() {
	flag.BoolVar(&help, "h", false, "this help")

	flag.StringVar(&consulAddr, "consul", "http://127.0.0.1:8500", "set consul address")
	flag.StringVar(&redisAddr, "redis", "redis://127.0.0.1:6379/0", "set redis address")
	flag.StringVar(&jaegerAddr, "jaeger", "http://127.0.0.1:9411/api/v2/spans", "set jaeger address")
	flag.StringVar(&nsqLookupAddr, "nsqlookupd", "127.0.0.1:4161", "set nsq lookup address")
	flag.StringVar(&nsqdTCP, "nsqdTCP", "127.0.0.1:4150", "set nsqd address")
	flag.StringVar(&nsqdHttp, "nsqdHTTP", "127.0.0.1:4151", "set nsqd address")
	flag.IntVar(&localPort, "localPort", 0, "run locally")
}

func main() {

	initFlag()
	var err error

	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	b, _ := braid.NewService("gate")
	b.Register(
		braid.Module(braid.LoggerZap),
		braid.Module(braid.PubsubNsq,
			pubsubnsq.WithLookupAddr([]string{nsqLookupAddr}),
			pubsubnsq.WithNsqdAddr([]string{nsqdTCP}, []string{nsqdHttp}),
		),
		braid.Module(braid.DiscoverConsul,
			discoverconsul.WithConsulAddr(consulAddr),
			discoverconsul.WithBlacklist([]string{"gateway"}),
		),
		braid.Module(braid.LinkcacheRedis,
			linkerredis.WithRedisAddr(redisAddr),
			linkerredis.WithMode(linkerredis.LinkerRedisModeLocal),
		),
		braid.Module(braid.ElectorConsul, electorconsul.WithConsulAddr(consulAddr)),
		braid.Module(braid.TracerJaeger,
			jaegertracing.WithHTTP(jaegerAddr),
			jaegertracing.WithProbabilistic(1),
		),
		braid.Module(braid.BalancerSWRR),
		braid.Module(braid.ClientGRPC),
	)

	b.Init()
	b.Run()
	defer b.Close()

	e := echo.New()
	e.Use(bm.ReqTrace())
	//e.Use(bm.ReqLimit())
	routes.Regist(e)

	err = e.Start(":14001")
	if err != nil {
		log.Fatalf("start echo err %s", err.Error())
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM)
	<-ch

	if err := e.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
}
