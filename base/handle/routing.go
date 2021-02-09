package handle

import (
	"braid-game/proto/api"
	"context"
)

// BaseServer 基础业务服务节点
type BaseServer struct {
	api.BaseServer
}

// AccRename rename
func (bs *BaseServer) AccRename(ctx context.Context, req *api.AccRenameReq) (res *api.AccRenameRes, err error) {

	res = new(api.AccRenameRes)
	res.Nickname = req.Nickname

	return res, err
}
