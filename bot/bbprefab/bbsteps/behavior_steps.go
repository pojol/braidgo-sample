package bbsteps

import (
	"braid-game/bot/bbprefab"
	"braid-game/bot/bbprefab/bbcards"

	"github.com/pojol/gobot/prefab"
)

// NewRenameStep rename
func NewRenameStep(md *bbprefab.BotData) *prefab.Step {

	step := prefab.NewStep()

	step.AddCard(bbcards.NewRenameCard(md))

	return step
}
