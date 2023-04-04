package main

import (
	"braid-game/base/handle"
	"braid-game/proto/api"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pojol/braid-go"
	"github.com/pojol/braid-go/depend"
	"github.com/pojol/braid-go/depend/bconsul"
	"github.com/pojol/braid-go/depend/blog"
	"github.com/pojol/braid-go/depend/bredis"
	"github.com/pojol/braid-go/mock"
	"github.com/pojol/braid-go/module/elector"
	"github.com/pojol/braid-go/module/modules"
	"github.com/pojol/braid-go/module/pubsub"
	"github.com/pojol/braid-go/module/rpc/client"
	"github.com/pojol/braid-go/module/rpc/server"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	mock.Init()

	b, _ := braid.NewService("base")
	b.RegisterDepend(
		depend.Logger(blog.BuildWithOption()),
		depend.Redis(bredis.BuildWithOption(&redis.Options{Addr: mock.RedisAddr})),
		depend.Consul(
			bconsul.BuildWithOption(bconsul.WithAddress([]string{mock.ConsulAddr})),
		),
	)

	b.RegisterModule(
		modules.Pubsub(
			pubsub.WithLookupAddr([]string{mock.NSQLookupdAddr}),
			pubsub.WithNsqdAddr([]string{mock.NsqdAddr}, []string{mock.NsqdHttpAddr}),
		),
		modules.Client(
			client.AppendInterceptors(grpc_prometheus.UnaryClientInterceptor),
		),
		modules.Server(
			server.WithListen(":14201"),
			server.AppendInterceptors(grpc_prometheus.UnaryServerInterceptor),
		),
		modules.Discover(),
		modules.Elector(
			elector.WithLockTick(3*time.Second)),
	)

	api.RegisterBaseServer(braid.Server().Server().(*grpc.Server), &handle.BaseServer{})

	b.Init()
	b.Run()
	defer b.Close()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
}
