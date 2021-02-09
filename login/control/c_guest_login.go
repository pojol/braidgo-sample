package control

import (
	"braid-game/proto/api"
	"time"

	"github.com/google/uuid"
	"github.com/pojol/braid"
	"github.com/pojol/braid/module/mailbox"
	"github.com/pojol/braid/modules/linkerredis"
)

// GuestLogin 游客登录
func GuestLogin(res *api.GuestRegistRes) error {

	var err error
	//token = "token" + GetUniqueID()
	res.Token = uuid.New().String()

	time.AfterFunc(time.Minute, func() {
		braid.Mailbox().Pub(mailbox.Cluster, linkerredis.LinkerTopicUnlink, &mailbox.Message{
			Body: []byte(res.Token),
		})
	})

	return err
}
