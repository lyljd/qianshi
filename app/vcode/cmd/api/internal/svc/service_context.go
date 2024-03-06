package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/zrpc"
	"qianshi/app/user/cmd/rpc/user"
	"qianshi/app/vcode/cmd/api/internal/config"
	"qianshi/common/email"
)

type ServiceContext struct {
	Config  config.Config
	Email   *email.Dialer
	Redis   *redis.Client
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Email:  email.NewExmailDialer(c.EmailSender, c.EmailUsername, c.EmailPassword),
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.RedisAddr,
			Password: c.RedisPassword,
		}),
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
