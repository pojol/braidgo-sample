package main

import (
	"braid-game/base/handle"
	"braid-game/common"
	"braid-game/proto"
	"braid-game/proto/api"
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pojol/braid-go"
	"github.com/pojol/braid-go/modules/discoverconsul"
	"github.com/pojol/braid-go/modules/grpcclient"
	"github.com/pojol/braid-go/modules/grpcserver"
	"github.com/pojol/braid-go/modules/jaegertracing"
	"github.com/pojol/braid-go/modules/linkerredis"
	"github.com/pojol/braid-go/modules/mailboxnsq"
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
	rand.Seed(time.Now().UnixNano())

	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	b, _ := braid.New(
		proto.ServiceBase,
		mailboxnsq.WithLookupAddr([]string{nsqLookupAddr}),
		mailboxnsq.WithNsqdAddr([]string{nsqdAddr}))

	b.RegistModule(
		braid.Discover(
			discoverconsul.Name,
			discoverconsul.WithConsulAddr(consulAddr)),
		braid.LinkCache(linkerredis.Name,
			linkerredis.WithRedisAddr(redisAddr),
			linkerredis.WithMode(linkerredis.LinkerRedisModeLocal),
		),
		braid.Client(grpcclient.Name),
		braid.Server(grpcserver.Name, grpcserver.WithListen(":14201")),
		braid.Tracing(
			jaegertracing.Name,
			jaegertracing.WithHTTP(jaegerAddr),
			jaegertracing.WithProbabilistic(1),
			jaegertracing.WithSpanFactory(
				jaegertracing.SpanFactory{
					Name:    "tracer_span_echo",
					Factory: jaegertracing.CreateEchoTraceSpan(),
				},
				jaegertracing.SpanFactory{
					Name:    "tracer_span_redis",
					Factory: jaegertracing.CreateRedisSpanFactory(),
				},
				jaegertracing.SpanFactory{
					Name:    "tracer_span_methon",
					Factory: common.CreateMethonSpanFactory(),
				},
			),
		),
	)

	api.RegisterBaseServer(braid.GetServer().(*grpc.Server), &handle.BaseServer{})

	b.Init()
	b.Run()
	defer b.Close()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
}
