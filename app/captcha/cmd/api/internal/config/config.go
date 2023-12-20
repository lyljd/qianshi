package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	RedisAddr     string `json:",env=REDIS_ADDR"`
	RedisPassword string `json:",env=REDIS_PASSWORD,optional"`
}
