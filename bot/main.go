package main

import (
	"braid-game/bot/bstrategy"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pojol/httpbot/factory"
)

var (
	help bool

	// target server addr
	target string

	// robot number
	num int

	// increase
	increase bool

	strategyParm string

	lifetime int
)

func initFlag() {
	flag.BoolVar(&help, "h", false, "this help")

	flag.StringVar(&target, "target", "http://localhost:14001", "set target server address")
	flag.IntVar(&num, "num", 1, "robot number")
	flag.BoolVar(&increase, "increase", false, "incremental robot in every second")
	flag.IntVar(&lifetime, "lifetime", 60, "life time by second")
	flag.StringVar(&strategyParm, "strategy", "default", "robot strategy")

}

/*
        +---->login1+--+--->mail1
        |              |
        +---->login2+--+--->mail2
gate1+->+
        |
        +---->base1
        |
        +---->base2
*/

func main() {

	initFlag()

	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	ports := []string{"http://127.0.0.1:14001" /*, "14003"*/}
	fmt.Println("targets", ports)
	fmt.Println("num", num)
	fmt.Println("increase", increase)
	fmt.Println("lifetime", lifetime)
	fmt.Println("strategy", strategyParm)

	rand.Seed(time.Now().UnixNano())

	mode := ""
	if increase {
		mode = factory.FactoryModeIncrease
	} else {
		mode = factory.FactoryModeStatic
	}

	rand.Seed(time.Now().UnixNano())

	client := &http.Client{}
	f, err := factory.Create(
		factory.WithCreateNum(num),
		factory.WithRunMode(mode),
		factory.WithClient(client),
		factory.WithLifeTime(time.Duration(lifetime)*time.Second),
	)
	if err != nil {
		panic(err)
	}

	f.Append(bstrategy.Default, bstrategy.FactoryDefault)
	f.Run()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	<-ch
	f.Close()
	time.Sleep(time.Second)
}
