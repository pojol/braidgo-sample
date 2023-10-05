package common

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pojol/braid-go/module"
	"github.com/pojol/braid-go/module/meta"
)

type Echo struct {
	E          *echo.Echo
	ExposePort string
}

func (we *Echo) Execute(topic module.ITopic, wg *sync.WaitGroup) {

	ch, err := topic.Sub(context.TODO(), "Echo")
	ch.Arrived(func(m *meta.Message) error {
		fmt.Println("work echo signal arrived")

		defer func() {
			wg.Done()
			fmt.Println("work echo done!", atomic.AddInt32(&wgcnt, -1))
		}()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if err := we.E.Shutdown(ctx); err != nil {
			panic(err)
		}

		return nil
	})

	err = we.E.Start(":" + we.ExposePort)
	if err != nil {
		fmt.Println("echo start err", err.Error())
	}
}
