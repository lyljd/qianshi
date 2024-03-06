package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	EmailSender   string `json:",env=EMAIL_SENDER"`
	EmailUsername string `json:",env=EMAIL_USERNAME"`
	EmailPassword string `json:",env=EMAIL_PASSWORD"`

	RedisAddr     string `json:",env=REDIS_ADDR"`
	RedisPassword string `json:",env=REDIS_PASSWORD,optional"`

	UserRpcConf zrpc.RpcClientConf
}
