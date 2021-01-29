package request

// GuestLoginReq 游客登录请求
// :port/v1/login/guest
type GuestLoginReq struct {
}

// GuestLoginRes 游客登录返回
type GuestLoginRes struct {
	Token string `json:"token"`
}

// AccountRenameReq 账号修改昵称
type AccountRenameReq struct {
	Nickname string `json:"nickname"`
}

// AccountRenameRes 账号修改昵称
type AccountRenameRes struct {
	Nickname string `json:"nickname"`
}
