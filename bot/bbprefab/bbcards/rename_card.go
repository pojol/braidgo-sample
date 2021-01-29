package bbcards

import (
	"braid-game/bot/bbprefab"
	"braid-game/proto/request"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// RenameCard 改名
type RenameCard struct {
	URL    string
	delay  time.Duration
	md     *bbprefab.BotData
	header map[string]string
	method string
}

// NewRenameCard 修改昵称
func NewRenameCard(md *bbprefab.BotData) *RenameCard {
	return &RenameCard{
		URL:   "/v1/base/rename",
		delay: time.Millisecond,
		md:    md,
		header: map[string]string{
			"Content-type": "application/json",
		},
		method: "POST",
	}
}

// GetURL 获取服务器地址
func (card *RenameCard) GetURL() string { return card.URL }

// GetHeader get card header
func (card *RenameCard) GetHeader() map[string]string {
	return card.header
}

// GetMethod get method
func (card *RenameCard) GetMethod() string { return card.method }

// SetDelay 设置卡片之间调用的延迟
func (card *RenameCard) SetDelay(delay time.Duration) { card.delay = delay }

// GetDelay 获取卡片之间调用的延迟
func (card *RenameCard) GetDelay() time.Duration { return card.delay }

// Marshal 序列化传入消息体
func (card *RenameCard) Marshal() []byte {

	card.header["token"] = card.md.AccToken

	req := request.AccountRenameReq{
		Nickname: "newname",
	}

	b, err := json.Marshal(&req)
	if err != nil {
		fmt.Println(card.GetURL(), "proto.Marshal err", err)
	}

	return b
}

// Unmarshal 反序列化返回消息
func (card *RenameCard) Unmarshal(res *http.Response) {

	errcode, _ := strconv.Atoi(res.Header["Errcode"][0])
	if errcode != 0 {
		fmt.Println(res.Request.URL, card.GetURL(), "request err", errcode)
	}
}
