package main

import (
	"braid-game/base/handle"
	"braid-game/common"
	"braid-game/proto/api"
	"fmt"
	"math/rand"
	"os"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pojol/braid-go"
	"github.com/pojol/braid-go/components"
	"github.com/pojol/braid-go/components/depends/bk8s"
	"github.com/pojol/braid-go/components/discoverk8s"
	"github.com/pojol/braid-go/components/rpcgrpc/grpcserver"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func main() {

	redis_addr := os.Getenv("REDIS_ADDR")
	ClusterTag := os.Getenv("ClusterTag")

	rand.Seed(time.Now().UnixNano())

	ServiceName := "base"

	regist := func(srv *grpc.Server) {
		api.RegisterBaseServer(srv, &handle.BaseServer{})
	}

	b, _ := braid.NewService(
		ServiceName,
		os.Getenv("POD_NAME"),
		&components.DefaultDirector{
			Opts: &components.DirectorOpts{
				ServerOpts: []grpcserver.Option{
					grpcserver.WithListen(":14201"),
					grpcserver.AppendUnaryInterceptors(grpc_prometheus.UnaryServerInterceptor),
					grpcserver.RegisterHandler(regist),
				},
				K8sCliOpts: []bk8s.Option{
					bk8s.WithConfigPath(""),
				},
				RedisCliOpts: &redis.Options{
					Addr: redis_addr,
				},
				DiscoverOpts: []discoverk8s.Option{
					discoverk8s.WithClusterTag(ClusterTag),
					discoverk8s.WithK8sClient(nil),
				},
			},
		},
	)

	b.Init()
	b.Run()
	defer b.Close()

	common.Init(ServiceName)
	common.Watch()

	fmt.Println("exit!")
}
