package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"qianshi/app/user/cmd/api/internal/config"
	"qianshi/app/user/cmd/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
