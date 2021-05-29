package arrange

import (
	"braid-game/bot/prefab"
	"net/http"

	bot "github.com/pojol/httpbot"
)

// NewStrategyDefault 默认的机器人工厂
func NewStrategyDefault(fmd interface{}, client *http.Client) *bot.Bot {
	md := &prefab.BotData{}

	b := bot.New(md,
		client,
		bot.WithName("default bot"))

	b.Timeline.AddStep(NewGuestLoginStep(md))
	for i := 0; i < 10; i++ {
		b.Timeline.AddStep(NewRenameStep(md))
	}

	return b
}
