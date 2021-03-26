package bstrategy

import (
	"braid-game/bot/bbprefab"
	"braid-game/bot/bbprefab/bbsteps"
	"net/http"

	bot "github.com/pojol/httpbot"
)

// FactoryDefault 默认的机器人工厂
func FactoryDefault(addr string, client *http.Client) *bot.Bot {
	md := &bbprefab.BotData{}

	b := bot.New(md, client)

	b.Timeline.AddStep(bbsteps.NewGuestLoginStep(md))
	for i := 0; i < 10; i++ {
		b.Timeline.AddStep(bbsteps.NewRenameStep(md))
	}

	return b
}
