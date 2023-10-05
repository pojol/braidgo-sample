package common

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"

	braid "github.com/pojol/braid-go"
	"github.com/pojol/braid-go/module"
	"github.com/pojol/braid-go/module/meta"
)

type WorkGroup struct {
	wg    sync.WaitGroup
	topic module.ITopic

	lst []IWork
}

type IWork interface {
	Execute(topic module.ITopic, wg *sync.WaitGroup)
}

var wg *WorkGroup
var wgcnt = int32(0)

func Init(serviceName string) {
	wg = &WorkGroup{
		wg:    sync.WaitGroup{},
		topic: braid.Topic(serviceName + "_exit_signal"),
	}
}

func AddWork(w IWork) error {
	wg.lst = append(wg.lst, w)
	return nil
}

func Watch() {

	for _, v := range wg.lst {
		wg.wg.Add(1)
		go v.Execute(wg.topic, &wg.wg)

		fmt.Println("add work", atomic.AddInt32(&wgcnt, 1))
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			c := <-ch
			switch c {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT: // 收到退出信号
				fmt.Println("recv signal", c)
				wg.topic.Pub(context.TODO(), &meta.Message{Body: []byte("nil")})
				return
			case syscall.SIGHUP: // 已经关闭
			default:
				return
			}
		}
	}()

	// wait
	fmt.Println("<<<<<<<<<<<<<<<<")
	wg.wg.Wait()
	fmt.Println(">>>>>>>>>>>>>>>>")
}
