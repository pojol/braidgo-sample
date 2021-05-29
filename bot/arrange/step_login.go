package arrange

import (
	"braid-game/bot/prefab"

	"github.com/pojol/httpbot/timeline"
)

// NewGuestLoginStep guest
func NewGuestLoginStep(md *prefab.BotData) *timeline.Step {

	step := timeline.NewStep("login_step")

	step.AddCard(prefab.NewGuestLoginCard(md))

	return step
}
