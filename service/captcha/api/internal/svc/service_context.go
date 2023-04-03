package svc

import (
	"github.com/go-redis/redis/v8"
	"os"
	"qianshi/service/captcha/api/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	RedisCli   *redis.Client
	ServerAddr string
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisCli: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Username: os.Getenv("REDIS_USERNAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
		}),
		ServerAddr: os.Getenv("SERVER_ADDR"),
	}
}
