package common

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	consul "github.com/hashicorp/consul/api"
	"github.com/pojol/braid-go/components/depends/bconsul"
	"github.com/pojol/braid-go/module"
	"github.com/pojol/braid-go/module/meta"
)

type Regist struct {
	ID     string
	Name   string
	Tags   []string
	Addr   string
	Port   int
	Client *bconsul.Client
}

func (r *Regist) Execute(topic module.ITopic, wg *sync.WaitGroup) {

	ch, _ := topic.Sub(context.TODO(), "Regist")
	ch.Arrived(func(m *meta.Message) error {
		fmt.Println("work regist signal arrived")
		err := r.Client.ServiceDeregister(r.ID)
		if err != nil {
			fmt.Println("deregister err", err.Error())
		}

		wg.Done()
		fmt.Println("work regist done!", atomic.AddInt32(&wgcnt, -1))
		return nil
	})

	err := r.Client.ServiceRegister(consul.AgentServiceRegistration{
		ID:      r.ID,
		Name:    r.Name,
		Tags:    r.Tags,
		Address: r.Addr,
		Port:    r.Port,
	})
	if err != nil {
		log.Fatalf("regist consul err %s", err.Error())
	}
}
