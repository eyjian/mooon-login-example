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
import (
    moooncrypto "github.com/eyjian/gomooon/crypto"
    mooonutils "github.com/eyjian/gomooon/utils"
)

const (
    EmptyRequest   = 2024020201 // 空的请求
    InvalidRequest = 2024020202 // 无效请求
    UserNotExists  = 2024020203 // 用户不存在
    PasswordError  = 2024020204 // 密码错误
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
    Uid    uint32 `json:"uid"`    // 用户 ID
    Avatar string `json:"avatar"` // 头像
}

// userLoginData 用户登录数据
type loginData struct {
    password string // 用户密码
    uid      uint32 // 用户 ID
}

var loginDataTable map[string]*loginData // Key 未用户名

// 初始化登录数据
func init() {
    loginDataTable = make(map[string]*loginData)

    loginDataTable = map[string]*loginData{
        "mooon": &loginData{
            password: "123456789a",
            uid:      2024020101,
        },
        "zhangsan": &loginData{
            password: "123456789b",
            uid:      2024020102,
        },
        "wangwu": &loginData{
            password: "123456789c",
            uid:      2024020103,
        },
    }
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
        return nil, status.Error(EmptyRequest, "empty request")
    }

    // 解密请求
    err := json.Unmarshal(in.Body, &loginReq)
    if err != nil {
        logc.Errorf(l.ctx, "invalid request")
        return nil, status.Error(InvalidRequest, "empty request")
    }
    if loginReq.Username == "" || loginReq.Password == "" {
        logc.Errorf(l.ctx, "username and password are required")
        return nil, status.Error(InvalidRequest, "username and password are required")
    }

    // 检查用户是否存在
    loginData, ok := loginDataTable[loginReq.Username]
    if !ok {
        logc.Errorf(l.ctx, "user %s not exists", loginReq.Username)
        return nil, status.Error(UserNotExists, "user not exists")
    }

    // 检查密码是否正确
    if loginReq.Password != loginData.password {
        logc.Errorf(l.ctx, "user %s password error", loginReq.Username)
        return nil, status.Error(PasswordError, "password error")
    }

    // 写 cookies
    sessionCookie := mooon_login.Cookie{
        Name:   "sessionid",
        Value:  getSessionId(),
        MaxAge: 3600,
    }
    tokenCookie := mooon_login.Cookie{
        Name:     "token",
        Value:    "token example",
        HttpOnly: true,
    }
    out.HttpCookies = append(out.HttpCookies, &sessionCookie)
    out.HttpCookies = append(out.HttpCookies, &tokenCookie)

    // 写 http 头
    out.HttpHeaders = make(map[string]string)
    out.HttpHeaders["Mooon-Header"] = "example"

    // 写响应体
    loginResp.Uid = loginData.uid
    loginResp.Avatar = "https://github.com/eyjian/mooon-login-example/blob/main/avatar.png"
    out.Body, _ = json.Marshal(&loginResp)

    return &out, nil
}

func getSessionId() string {
    nonceStr := mooonutils.GetNonceStr(28)
    return moooncrypto.Md5Sum(nonceStr, false)
}
