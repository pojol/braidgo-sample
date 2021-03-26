package bbsteps

import (
	"braid-game/bot/bbprefab"
	"braid-game/bot/bbprefab/bbcards"

	"github.com/pojol/httpbot/timeline"
)

// NewGuestLoginStep guest
func NewGuestLoginStep(md *bbprefab.BotData) *timeline.Step {

	step := timeline.NewStep()

	step.AddCard(bbcards.NewGuestLoginCard(md))

	return step
}
