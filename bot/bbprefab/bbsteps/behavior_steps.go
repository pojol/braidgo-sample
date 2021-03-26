package bbsteps

import (
	"braid-game/bot/bbprefab"
	"braid-game/bot/bbprefab/bbcards"

	"github.com/pojol/httpbot/timeline"
)

// NewRenameStep rename
func NewRenameStep(md *bbprefab.BotData) *timeline.Step {

	step := timeline.NewStep()

	step.AddCard(bbcards.NewRenameCard(md))

	return step
}
