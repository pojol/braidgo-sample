package handle

import (
	"braid-game/mail/control"
	"braid-game/mail/model"
	"braid-game/proto/api"
	"context"
)

// MailServer 邮件服务
type MailServer struct {
	api.MailServer
}

// Send 发送邮件请求（内部
func (ms *MailServer) Send(ctx context.Context, req *api.SendMailReq) (*api.SendMailRes, error) {
	res := new(api.SendMailRes)

	errcode := control.SendNormal(req.Accountid, model.MailDat{
		Title: req.Body.Title,
		Txt:   req.Body.Txt,
	})
	res.Errcode = int32(errcode)

	return res, nil
}
