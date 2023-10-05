package main

import (
	"braid-game/common"
	"braid-game/gate/routes"
	"braid-game/proto"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/labstack/echo/v4"
	"github.com/pojol/braid-go"
	"github.com/pojol/braid-go/components"
	"github.com/pojol/braid-go/components/depends/blog"
	"github.com/pojol/braid-go/components/discoverk8s"
	"github.com/redis/go-redis/v9"
)

const (
	// NodeName box 节点名
	NodeName = "gateway"
)

// ConsulRegistReq regist req dat
type ConsulRegistReq struct {
	ID      string   `json:"ID"`
	Name    string   `json:"Name"`
	Tags    []string `json:"Tags"`
	Address string   `json:"Address"`
	Port    int      `json:"Port"`
}

// Regist regist service 2 consul
func ServiceRegist(address string, req ConsulRegistReq) error {
	var err error
	byt, _ := json.Marshal(&req)
	reqBuf := bytes.NewBuffer(byt)
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	api := address + "/v1/agent/service/register"

	httpReq, err := http.NewRequest("PUT", api, reqBuf)
	if err != nil {
		return fmt.Errorf("failed to new request api:%v err:%v", api, err.Error())
	}
	httpReq.Header.Set("Content-type", "application/json")

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to request to server api:%v err:%v", api, err.Error())
	}
	defer httpRes.Body.Close()

	resbyt, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body api:%v err:%v", api, err.Error())
	}

	fmt.Println("register res", string(resbyt))

	return nil
}

func main() {

	redis_addr := os.Getenv("REDIS_ADDR")

	ServiceName := "gate"

	b, _ := braid.NewService(
		ServiceName,
		os.Getenv("POD_NAME"),
		&components.DefaultDirector{
			Opts: &components.DirectorOpts{
				LogOpts: []blog.Option{
					blog.WithLevel(int(blog.InfoLevel)),
					blog.WithStdout(true),
				},
				RedisCliOpts: &redis.Options{
					Addr: redis_addr,
				},
				DiscoverOpts: []discoverk8s.Option{
					discoverk8s.WithNamespace("braidgo"),
					discoverk8s.WithSelectorTag("braidgo"),
					discoverk8s.WithServicePortPairs([]discoverk8s.ServicePortPair{
						{Name: proto.ServiceLogin, Port: 14101},
					}),
				},
			},
		},
	)

	b.Init()
	b.Run()
	defer b.Close()

	e := echo.New()
	e.HideBanner = true
	routes.Regist(e)

	common.Init(ServiceName)
	common.AddWork(&common.Echo{
		E:          e,
		ExposePort: "14001",
	})

	common.Watch()
}
