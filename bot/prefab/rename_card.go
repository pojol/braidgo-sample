package prefab

import (
	"braid-game/proto/request"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/pojol/httpbot/card"
)

// RenameCard 改名
type RenameCard struct {
	Base  *card.Card
	URL   string
	delay time.Duration
	md    *BotData
}

// NewRenameCard 修改昵称
func NewRenameCard(md *BotData) *RenameCard {
	return &RenameCard{
		Base:  card.NewCardWithConfig(),
		URL:   "http://127.0.0.1:14001/v1/base/rename",
		delay: time.Millisecond,
		md:    md,
	}
}

// GetURL 获取服务器地址
func (card *RenameCard) GetURL() string { return card.URL }

func (card *RenameCard) GetName() string { return "RenameCard" }

// GetClient 获取 http.client
func (card *RenameCard) GetClient() *http.Client { return nil }

// GetHeader get card header
func (card *RenameCard) GetHeader() map[string]string { return card.Base.Header }

// GetMethod get method
func (card *RenameCard) GetMethod() string { return card.Base.Method }

// SetDelay 设置卡片之间调用的延迟
func (card *RenameCard) SetDelay(delay time.Duration) { card.delay = delay }

// GetDelay 获取卡片之间调用的延迟
func (card *RenameCard) GetDelay() time.Duration { return card.delay }

// Enter 序列化传入消息体
func (card *RenameCard) Enter() []byte {

	card.Base.Inject(card)

	card.Base.Header["token"] = card.md.AccToken

	req := request.AccountRenameReq{
		Nickname: "newname",
	}

	b, err := json.Marshal(&req)
	if err != nil {
		fmt.Println(card.GetURL(), "proto.Marshal err", err)
	}

	return b
}

// Leave 反序列化返回消息
func (card *RenameCard) Leave(res *http.Response) error {

	var err error

	errcode, _ := strconv.Atoi(res.Header["Errcode"][0])
	if errcode != 0 {
		err = fmt.Errorf("request err %v", errcode)
		goto EXT
	}

	err = card.Base.Assert()

EXT:
	return err
}
