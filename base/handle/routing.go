package handle

import (
	"braid-game/errcode"
	"braid-game/proto"
	"braid-game/proto/api"
	"context"
	"errors"

	"github.com/pojol/braid-go"
)

// BaseServer 基础业务服务节点
type BaseServer struct {
	api.BaseServer
}

// AccRename rename
func (bs *BaseServer) AccRename(ctx context.Context, req *api.AccRenameReq) (res *api.AccRenameRes, err error) {

	res = new(api.AccRenameRes)
	res.Nickname = req.Nickname

	mailRes := &api.SendMailRes{}

	braid.Send(ctx,
		proto.ServiceMail,
		proto.APIMailSend,
		req.Token,
		&api.SendMailReq{
			Accountid: "acc_xx",
			Body: &api.MailBody{
				Title: "welcome!",
			},
		},
		mailRes,
	)

	if mailRes.Errcode != int32(errcode.Succ) {
		return res, errors.New("send mail fail")
	}

	return res, err
}
