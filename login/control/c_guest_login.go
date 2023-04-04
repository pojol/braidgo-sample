package control

import (
	"braid-game/proto/api"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pojol/braid-go"
	"github.com/pojol/braid-go/module/linkcache"
	"github.com/pojol/braid-go/module/pubsub"
)

// GuestLogin 游客登录
func GuestLogin(res *api.GuestRegistRes) error {

	var err error
	//token = "token" + GetUniqueID()
	res.Token = uuid.New().String()

	time.AfterFunc(time.Minute, func() {
		fmt.Println("send unlink msg", res.Token)

		braid.Pubsub().ClusterTopic(linkcache.TopicUnlink).Pub(&pubsub.Message{
			Body: []byte(res.Token),
		})
	})

	return err
}
