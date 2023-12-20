package svc

import (
	"github.com/go-redis/redis/v8"
	"qianshi/app/captcha/cmd/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.RedisAddr,
			Password: c.RedisPassword,
		}),
	}
}
