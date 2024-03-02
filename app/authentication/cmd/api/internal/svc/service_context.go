package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"qianshi/app/authentication/cmd/api/internal/config"
	"qianshi/app/authentication/cmd/rpc/authentication"
)

type ServiceContext struct {
	Config            config.Config
	AuthenticationRpc authentication.Authentication
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		AuthenticationRpc: authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpcConf)),
	}
}
