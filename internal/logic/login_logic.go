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
    Username string `json:"username,required"` // 用户名
    Password string `json:"password,required"` // 密码
}

// LoginResp 登录响应
type LoginResp struct {
    Uid uint32 `json:"uid"`// 用户 ID
    Avatar string `json:"avatar"`// 头像
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
    var loginResp LoginResp
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
    if loginReq.Username == "" || loginReq.Password == "" {
        logc.Errorf(l.ctx, "username and password are required")
        return nil, status.Error(InvalidRequest, "username and password are required")
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
    loginResp.uid = 20240202
    loginResp.Avatar = "https://github.com/eyjian/mooon-login-example/blob/main/avatar.png"
    out.Body, _ = json.Marshal(&loginResp)

    return &out, nil
}