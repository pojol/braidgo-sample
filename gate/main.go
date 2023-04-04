package main

import (
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
	"github.com/pojol/braid-go/depend"
	"github.com/pojol/braid-go/depend/blog"
	"github.com/pojol/braid-go/depend/bredis"
	"github.com/redis/go-redis/v9"
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
	b.RegisterDepend(
		depend.Logger(blog.BuildWithOption()),
		depend.Redis(bredis.BuildWithOption(&redis.Options{Addr: redisAddr})),
	)

	b.Init()
	b.Run()
	defer b.Close()

	e := echo.New()
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
