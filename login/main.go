package main

import (
	"braid-game/login/constant"
	"braid-game/login/handle"
	"braid-game/proto/api"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/redis/go-redis/v9"

	"github.com/pojol/braid-go"
	"github.com/pojol/braid-go/components"
	"github.com/pojol/braid-go/components/depends/bk8s"
	"github.com/pojol/braid-go/components/discoverk8s"
	"github.com/pojol/braid-go/components/rpcgrpc/grpcserver"
	"google.golang.org/grpc"
)

func main() {

	constant.LoginRandRecord = rand.Intn(10000)

	redis_addr := os.Getenv("REDIS_ADDR")
	fmt.Println("redis addr", redis_addr)

	ServiceName := "login"

	b, _ := braid.NewService(
		ServiceName,
		os.Getenv("POD_NAME"),
		&components.DefaultDirector{
			Opts: &components.DirectorOpts{
				ServerOpts: []grpcserver.Option{
					grpcserver.WithListen(":14101"),
					grpcserver.AppendUnaryInterceptors(grpc_prometheus.UnaryServerInterceptor),
					grpcserver.RegisterHandler(func(srv *grpc.Server) {
						api.RegisterLoginServer(srv, &handle.LoginServer{})
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

	b.Init()
	b.Run()
	defer b.Close()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch

	fmt.Println("exit!")
}
