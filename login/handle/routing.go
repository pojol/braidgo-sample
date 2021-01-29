package handle

import (
	"braid-game/login/control"
	"braid-game/proto/api"
	"context"
)

// LoginServer 路由服务器
type LoginServer struct {
	api.LoginServer
}

// GuestRegist 游客注册
func (ls *LoginServer) GuestRegist(ctx context.Context, req *api.GuestRegistReq) (res *api.GuestRegistRes, err error) {

	res = new(api.GuestRegistRes)

	err = control.GuestLogin(res)

	return res, err
}
