package handle

import (
	"braid-game/errcode"
	"braid-game/login/control"
	"braid-game/proto"
	"braid-game/proto/api"
	"context"
	"errors"

	"github.com/pojol/braid"
)

// LoginServer 路由服务器
type LoginServer struct {
	api.LoginServer
}

// GuestRegist 游客注册
func (ls *LoginServer) GuestRegist(ctx context.Context, req *api.GuestRegistReq) (res *api.GuestRegistRes, err error) {

	res = new(api.GuestRegistRes)

	err = control.GuestLogin(res)

	if err != nil {
		mailRes := &api.SendMailRes{}

		braid.GetClient().Invoke(ctx,
			proto.ServiceMail,
			proto.APIMailSend,
			res.Token,
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

	}

	return res, err
}
