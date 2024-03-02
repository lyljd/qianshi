package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/zrpc"
	"qianshi/app/authentication/cmd/rpc/internal/config"
	"qianshi/app/user/cmd/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	Redis   *redis.Client
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.RedisAddr,
			Password: c.RedisPassword,
		}),
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
