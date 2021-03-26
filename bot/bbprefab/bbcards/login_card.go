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

	"github.com/pojol/httpbot/prefab"
)

// GuestLoginCard 游客登录
type GuestLoginCard struct {
	URL   string
	delay time.Duration
	md    *bbprefab.BotData
	Base  *prefab.Card
}

// NewGuestLoginCard 生成账号创建预制
func NewGuestLoginCard(md *bbprefab.BotData) *GuestLoginCard {
	return &GuestLoginCard{
		Base:  prefab.NewCardWithConfig(),
		URL:   "http://123.207.198.57:14001/v1/login/guest",
		delay: time.Millisecond,
		md:    md,
	}
}

// GetURL 获取服务器地址
func (card *GuestLoginCard) GetURL() string { return card.URL }

func (card *GuestLoginCard) GetName() string { return "GuestLoginCard" }

// GetClient 获取 http.client
func (card *GuestLoginCard) GetClient() *http.Client { return nil }

// GetHeader get card header
func (card *GuestLoginCard) GetHeader() map[string]string { return card.Base.Header }

// GetMethod get method
func (card *GuestLoginCard) GetMethod() string { return card.Base.Method }

// SetDelay 设置卡片之间调用的延迟
func (card *GuestLoginCard) SetDelay(delay time.Duration) { card.delay = delay }

// GetDelay 获取卡片之间调用的延迟
func (card *GuestLoginCard) GetDelay() time.Duration { return card.delay }

// Enter 序列化传入消息体
func (card *GuestLoginCard) Enter() []byte {

	b := []byte{}

	return b
}

// Leave 反序列化返回消息
func (card *GuestLoginCard) Leave(res *http.Response) error {

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

	return nil
}
