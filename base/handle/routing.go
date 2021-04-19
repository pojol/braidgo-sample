package handle

import (
	"braid-game/errcode"
	"braid-game/proto"
	"braid-game/proto/api"
	"context"
	"errors"
	"fmt"

	"github.com/pojol/braid-go"
	"github.com/pojol/braid-go/module/tracer"
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

	tracing := braid.Tracer()
	var span tracer.ISpan

	if tracing != nil {
		span, err = tracing.GetSpan("tracer_span_methon")
		if err != nil {
			fmt.Println("get span err", err.Error())
		}
		if span != nil {
			span.Begin(ctx)
			span.SetTag("methon", "AccRename")
			span.SetTag("nickname", req.Nickname)
			defer span.End(ctx)
		}
	}

	braid.GetClient().Invoke(ctx,
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

	span.SetTag("errcode", mailRes.Errcode)
	if mailRes.Errcode != int32(errcode.Succ) {
		return res, errors.New("send mail fail")
	}

	return res, err
}
