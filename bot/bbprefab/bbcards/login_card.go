package bbcards

import (
	"braid-game/bot/bbprefab"
	"braid-game/proto/request"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// GuestLoginCard 游客登录
type GuestLoginCard struct {
	URL    string
	delay  time.Duration
	md     *bbprefab.BotData
	header map[string]string
	method string
}

// NewGuestLoginCard 生成账号创建预制
func NewGuestLoginCard(md *bbprefab.BotData) *GuestLoginCard {
	return &GuestLoginCard{
		URL:   "/v1/login/guest",
		delay: time.Millisecond,
		md:    md,
		header: map[string]string{
			"Content-type": "application/json",
		},
		method: "POST",
	}
}

// GetURL 获取服务器地址
func (card *GuestLoginCard) GetURL() string { return card.URL }

// GetHeader get card header
func (card *GuestLoginCard) GetHeader() map[string]string { return nil }

// GetMethod get method
func (card *GuestLoginCard) GetMethod() string { return card.method }

// SetDelay 设置卡片之间调用的延迟
func (card *GuestLoginCard) SetDelay(delay time.Duration) { card.delay = delay }

// GetDelay 获取卡片之间调用的延迟
func (card *GuestLoginCard) GetDelay() time.Duration { return card.delay }

// Marshal 序列化传入消息体
func (card *GuestLoginCard) Marshal() []byte {

	b := []byte{}

	return b
}

// Unmarshal 反序列化返回消息
func (card *GuestLoginCard) Unmarshal(res *http.Response) {

	errcode, _ := strconv.Atoi(res.Header["Errcode"][0])
	if errcode != 0 {
		fmt.Println(res.Request.URL, card.GetURL(), "request err", errcode)
	}

	cres := request.GuestLoginRes{}
	b, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(b, &cres)
	if err != nil {
		fmt.Println(res.Request.URL, card.GetURL(), "json.Unmarshal", errcode, "token", cres.Token)
	}

	card.md.AccToken = cres.Token
}
