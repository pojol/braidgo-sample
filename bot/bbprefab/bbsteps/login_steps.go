package bbsteps

import (
	"braid-game/bot/bbprefab"
	"braid-game/bot/bbprefab/bbcards"

	"github.com/pojol/httpbot/prefab"
)

// NewGuestLoginStep guest
func NewGuestLoginStep(md *bbprefab.BotData) *prefab.Step {

	step := prefab.NewStep()

	step.AddCard(bbcards.NewGuestLoginCard(md))

	return step
}
