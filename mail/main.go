package main

import (
	"braid-game/mail/constant"
	"braid-game/mail/handle"
	"braid-game/proto/api"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pojol/braid-go"
	"github.com/pojol/braid-go/components"
	"github.com/pojol/braid-go/components/depends/bk8s"
	"github.com/pojol/braid-go/components/discoverk8s"
	"github.com/pojol/braid-go/components/rpcgrpc/grpcserver"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

var (
	ServiceName = "mail"
)

func main() {

	constant.MailRandRecord = rand.Intn(10000)

	redis_addr := os.Getenv("REDIS_ADDR")

	fmt.Println("new service")
	b, _ := braid.NewService(
		"mail",
		os.Getenv("POD_NAME"),
		&components.DefaultDirector{
			Opts: &components.DirectorOpts{
				ServerOpts: []grpcserver.Option{
					grpcserver.WithListen(":14301"),
					grpcserver.AppendUnaryInterceptors(grpc_prometheus.UnaryServerInterceptor),
					grpcserver.RegisterHandler(func(srv *grpc.Server) {
						api.RegisterMailServer(srv, &handle.MailServer{})
					}),
				},
				K8sCliOpts: []bk8s.Option{
					bk8s.WithConfigPath(""),
				},
				RedisCliOpts: &redis.Options{
					Addr: redis_addr,
				},
				DiscoverOpts: []discoverk8s.Option{
					discoverk8s.WithNamespace("braidgo"),
					discoverk8s.WithSelectorTag("braidgo"),
				},
			},
		},
	)

	fmt.Println("init service")
	b.Init()
	fmt.Println("run service")
	b.Run()
	defer b.Close()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch

}
