package arrange

import (
	"braid-game/bot/prefab"

	"github.com/pojol/httpbot/timeline"
)

// NewRenameStep rename
func NewRenameStep(md *prefab.BotData) *timeline.Step {

	step := timeline.NewStep("rename_step")

	step.AddCard(prefab.NewRenameCard(md))

	return step
}
