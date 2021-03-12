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
	"github.com/pojol/braid-go/modules/discoverconsul"
	"github.com/pojol/braid-go/modules/electorconsul"
	"github.com/pojol/braid-go/modules/grpcclient"
	"github.com/pojol/braid-go/modules/linkerredis"
	"github.com/pojol/braid-go/modules/mailboxnsq"
)

var (
	help bool

	consulAddr    string
	redisAddr     string
	jaegerAddr    string
	nsqLookupAddr string
	nsqdAddr      string
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
	flag.StringVar(&nsqLookupAddr, "nsqlookup", "127.0.0.1:4161", "set nsq lookup address")
	flag.StringVar(&nsqdAddr, "nsqd", "127.0.0.1:4150", "set nsqd address")
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
	routes.Linkcheckmap = make(map[string]int)

	b, _ := braid.New(
		NodeName,
		mailboxnsq.WithLookupAddr([]string{nsqLookupAddr}),
		mailboxnsq.WithNsqdAddr([]string{nsqdAddr}))

	b.RegistModule(
		braid.Discover(
			discoverconsul.Name,
			discoverconsul.WithConsulAddr(consulAddr)),
		braid.Client(grpcclient.Name),
		braid.LinkCache(linkerredis.Name,
			linkerredis.WithRedisAddr(redisAddr),
		),
		braid.Elector(
			electorconsul.Name,
			electorconsul.WithConsulAddr(consulAddr),
		),
	)

	b.Init()
	b.Run()
	defer b.Close()

	e := echo.New()
	//e.Use(bm.ReqTrace())
	//e.Use(bm.ReqLimit())
	routes.Regist(e)

	err = e.Start(":14001")
	if err != nil {
		log.Fatalf("start echo err %s", err.Error())
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)
	<-ch

	if err := e.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
}
