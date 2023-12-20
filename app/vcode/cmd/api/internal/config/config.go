package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	EmailSender   string `json:",env=Email_Sender"`
	EmailUsername string `json:",env=Email_Username"`
	EmailPassword string `json:",env=Email_Password"`

	RedisAddr     string `json:",env=REDIS_ADDR"`
	RedisPassword string `json:",env=REDIS_PASSWORD,optional"`
}
