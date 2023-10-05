package routes

import (
	"braid-game/proto"
	"braid-game/proto/api"
	"braid-game/proto/request"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pojol/braid-go"

	"github.com/labstack/echo/v4"
)

func loginGuest(ctx echo.Context) error {
	req := &api.GuestRegistReq{}
	res := &api.GuestRegistRes{}
	byt := []byte{}

	fmt.Println("login guest!")
	err := braid.Send(ctx.Request().Context(),
		proto.ServiceLogin,
		proto.APILoginGuest,
		"",
		req,
		res,
	)
	if err != nil {
		goto EXT
	}

	byt, err = json.Marshal(request.GuestLoginRes{
		Token: res.Token,
	})

EXT:
	if err != nil {
		ctx.Response().Header().Set("Errcode", "-1")
		ctx.Response().Header().Set("Errmsg", err.Error())
	} else {
		ctx.Response().Header().Set("Errcode", "0")
	}

	ctx.Blob(http.StatusOK, "text/plain; charset=utf-8", byt)
	return nil
}

func baseAccRename(ctx echo.Context) error {
	var err error
	byt := []byte{}
	var body []byte
	var errcode string
	jreq := &request.AccountRenameReq{}
	req := &api.AccRenameReq{}
	res := &api.AccRenameRes{}
	token := ctx.Request().Header.Get("token")

	if token == "" {
		errcode = "-2" // tmp
		err = errors.New("token is not available")
		goto EXT
	}

	body, err = ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		errcode = "-3"
		goto EXT
	}

	err = json.Unmarshal(body, jreq)
	if err != nil {
		errcode = "-4"
		goto EXT
	}

	req.Nickname = jreq.Nickname
	req.Token = token

	err = braid.Send(ctx.Request().Context(),
		proto.ServiceBase,
		proto.APIBaseAccRename,
		token,
		req,
		res)
	if err != nil {
		goto EXT
	}

	byt, err = json.Marshal(request.AccountRenameRes{
		Nickname: res.Nickname,
	})

EXT:
	status := http.StatusOK
	if err != nil {
		ctx.Response().Header().Set("Errcode", errcode)
		ctx.Response().Header().Set("Errmsg", err.Error())
		status = http.StatusBadRequest
	} else {
		ctx.Response().Header().Set("Errcode", "0")
	}

	ctx.Blob(status, "text/plain; charset=utf-8", byt)
	return nil
}

// Regist regist
func Regist(e *echo.Echo) {
	e.POST("/v1/login/guest", loginGuest)
	e.POST("/v1/base/rename", baseAccRename)
}
