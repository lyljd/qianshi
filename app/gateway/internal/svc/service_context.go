package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"qianshi/app/authentication/cmd/rpc/authentication"
	"qianshi/app/gateway/internal/config"
	"qianshi/app/user/cmd/rpc/user"
)

type ServiceContext struct {
	Config            config.Config
	AuthenticationRpc authentication.Authentication
	UserRpc           user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		AuthenticationRpc: authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpcConf)),
		UserRpc:           user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
