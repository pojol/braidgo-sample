package control

import (
	"braid-game/proto/api"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pojol/braid-go"
	"github.com/pojol/braid-go/module/meta"
)

// GuestLogin 游客登录
func GuestLogin(res *api.GuestRegistRes) error {

	var err error
	//token = "token" + GetUniqueID()
	res.Token = uuid.New().String()

	time.AfterFunc(time.Minute, func() {
		fmt.Println("send unlink msg", res.Token)

		braid.Topic(meta.TopicLinkcacheUnlink).Pub(context.TODO(), &meta.Message{
			Body: []byte(res.Token),
		})
	})

	return err
}
