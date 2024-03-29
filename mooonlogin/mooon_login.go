// Code generated by goctl. DO NOT EDIT.
// Source: mooon_login.proto

package mooonlogin

import (
	"context"

	"mooon-login-example/pb/mooon_login"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Cookie    = mooon_login.Cookie
	LoginReq  = mooon_login.LoginReq
	LoginResp = mooon_login.LoginResp

	MooonLogin interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	}

	defaultMooonLogin struct {
		cli zrpc.Client
	}
)

func NewMooonLogin(cli zrpc.Client) MooonLogin {
	return &defaultMooonLogin{
		cli: cli,
	}
}

func (m *defaultMooonLogin) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := mooon_login.NewMooonLoginClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
