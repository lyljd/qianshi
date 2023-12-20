package svc

import (
	"github.com/go-redis/redis/v8"
	"qianshi/app/vcode/cmd/api/internal/config"
	"qianshi/common/email"
)

type ServiceContext struct {
	Config config.Config
	Email  *email.Dialer
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Email:  email.New163Dialer(c.EmailSender, c.EmailUsername, c.EmailPassword),
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.RedisAddr,
			Password: c.RedisPassword,
		}),
	}
}
