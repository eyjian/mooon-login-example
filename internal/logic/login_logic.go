package logic

import (
	"context"
    "encoding/json"
    "github.com/zeromicro/go-zero/core/logc"
    "google.golang.org/grpc/status"

    "mooon-login-example/internal/svc"
	"mooon-login-example/pb/mooon_login"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
    EmptyRequest   = 2024020201 // 空的请求
    InvalidRequest = 2024020202 // 无效请求
)

type LoginLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

// LoginReq 登录请求
type LoginReq struct {
    Username string `json:"username"` // 用户名
    Password string `json:"password"` // 密码
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
    var loginReq LoginReq
    var out mooon_login.LoginResp

    // 判断请求是否为空
    if len(in.Body) == 0 {
        logc.Error(l.ctx, "empty request")
        return nil, status.Error(EmptyRequest,"empty request")
    }

    // 解密请求
    err := json.Unmarshal(in.Body, &loginReq)
    if err != nil {
        logc.Errorf(l.ctx, "invalid request")
        return nil, status.Error(InvalidRequest,"empty request")
    }

    // 写 cookies
    httpCookie := mooon_login.Cookie{
        Name:  "sid",
        Value: "example",
    }
    out.HttpCookies = append(out.HttpCookies, &httpCookie)

    // 写 http 头
    out.HttpHeaders = make(map[string]string)
    out.HttpHeaders["Mooon-Header"] = "example"

    // 写响应体
    out.Body = []byte("{\"mooon\":\"example\"}")

    return &out, nil
}