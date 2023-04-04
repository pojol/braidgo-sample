package main

import (
	"braid-game/login/constant"
	"braid-game/login/handle"
	"braid-game/proto/api"
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pojol/braid-go"
	"google.golang.org/grpc"
)

var (
	help bool

	consulAddr    string
	jaegerAddr    string
	redisAddr     string
	nsqLookupAddr string
	nsqdTCP       string
	nsqdHttp      string
	localPort     int
)

func initFlag() {
	flag.BoolVar(&help, "h", false, "this help")

	flag.StringVar(&consulAddr, "consul", "http://127.0.0.1:8500", "set consul address")
	flag.StringVar(&nsqLookupAddr, "nsqlookupd", "127.0.0.1:4161", "set nsq lookup address")
	flag.StringVar(&nsqdTCP, "nsqdTCP", "127.0.0.1:4150", "set nsqd address")
	flag.StringVar(&nsqdHttp, "nsqdHTTP", "127.0.0.1:4151", "set nsqd address")
	flag.StringVar(&redisAddr, "redis", "redis://127.0.0.1:6379/0", "set redis address")
	flag.StringVar(&jaegerAddr, "jaeger", "http://127.0.0.1:9411/api/v2/spans", "set jaeger address")
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

	constant.LoginRandRecord = rand.Intn(10000)

	b, _ := braid.NewService("login")
	

	api.RegisterLoginServer(braid.Server().Server().(*grpc.Server), &handle.LoginServer{})

	b.Init()
	b.Run()
	defer b.Close()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
}
