package logic

import (
	"context"

	"mooon-login-example/internal/svc"
	"mooon-login-example/pb/mooon_login"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *mooon_login.LoginReq) (*mooon_login.LoginResp, error) {
	// todo: add your logic here and delete this line
	logx.Infof("Body: %s\n", in.Body)
	var loginResp mooon_login.LoginResp

	_ = in
	httpCookie := mooon_login.Cookie{
		Name:  "sid",
		Value: "example",
	}
	loginResp.HttpCookies = append(loginResp.HttpCookies, &httpCookie)

    loginResp.HttpHeaders = make(map[string]string)
	loginResp.HttpHeaders["Mooon-Header"] = "example"
	loginResp.Body = []byte("{\"\"mooon\":\"example\"}")

	return &loginResp, nil
}