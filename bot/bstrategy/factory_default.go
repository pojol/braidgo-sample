package bstrategy

import (
	"braid-game/bot/bbprefab"
	"braid-game/bot/bbprefab/bbsteps"
	"net/http"
	"time"

	bot "github.com/pojol/gobot"
)

// FactoryDefault 默认的机器人工厂
func FactoryDefault(addr string, client *http.Client) *bot.Bot {
	md := &bbprefab.BotData{}

	bot := bot.New(bot.BotConfig{
		Addr:   addr,
		Report: false,
	}, client, md)

	bot.Timeline.AddStep(bbsteps.NewGuestLoginStep(md))
	for i := 0; i < 10; i++ {
		bot.Timeline.AddDelayStep(bbsteps.NewRenameStep(md), time.Millisecond*100)
	}

	return bot
}
