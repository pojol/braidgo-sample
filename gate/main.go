package main

import (
	"braid-game/common"
	bm "braid-game/gate/middleware"
	"braid-game/gate/routes"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/labstack/echo/v4"
	"github.com/pojol/braid"
	"github.com/pojol/braid/3rd/log"
	"github.com/pojol/braid/modules/discoverconsul"
	"github.com/pojol/braid/modules/electorconsul"
	"github.com/pojol/braid/modules/grpcclient"
	"github.com/pojol/braid/modules/jaegertracing"
	"github.com/pojol/braid/modules/linkerredis"
	"github.com/pojol/braid/modules/mailboxnsq"
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
	//var kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//var nodeID = flag.String("node-id", "", "node id used for leader election")

	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	if localPort != 0 {
		addr := ":" + strconv.Itoa(localPort)

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

	b, _ := braid.New(
		NodeName,
		mailboxnsq.WithLookupAddr([]string{nsqLookupAddr}),
		mailboxnsq.WithNsqdAddr([]string{nsqdAddr}))

	b.RegistModule(
		braid.Discover(
			discoverconsul.Name,
			discoverconsul.WithConsulAddr(consulAddr)),
		braid.Client(grpcclient.Name),
		braid.Elector(
			electorconsul.Name,
			electorconsul.WithConsulAddr(consulAddr),
		),
		braid.LinkCache(linkerredis.Name,
			linkerredis.WithRedisAddr(redisAddr),
		),
		braid.Tracing(jaegertracing.Name,
			jaegertracing.WithHTTP(jaegerAddr),
			jaegertracing.WithProbabilistic(0.1),
		))

	b.Init()
	b.Run()
	defer b.Close()

	e := echo.New()
	e.Use(bm.ReqTrace())
	e.Use(bm.ReqLimit())
	routes.Regist(e)

	//go gatemid.Tick()

	go func() {
		fmt.Println(http.ListenAndServe(":6060", nil))
	}()

	if localPort == 0 {
		err = e.Start(":14001")
	} else {
		err = e.Start(":" + strconv.Itoa(localPort))
	}
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
