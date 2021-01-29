package bbprefab

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// BotData bot mapping data
type BotData struct {
	AccToken string `json:"acctoken"`
}

// Refresh refresh metadata
func (md *BotData) Refresh(meta interface{}) {
	err := mapstructure.Decode(meta, md)
	if err != nil {
		fmt.Println("refresh metadata err", err)
	}
}
