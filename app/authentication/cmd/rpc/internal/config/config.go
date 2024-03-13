package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	TokenSecret         string `json:",env=TOKEN_SECRET"`
	TokenMinutes        int    `json:",env=TOKEN_MINUTES"`
	RefreshTokenSecret  string `json:",env=REFRESH_TOKEN_SECRET"`
	RefreshTokenMinutes int    `json:",env=REFRESH_TOKEN_MINUTES"`
	CdnUrl              string `json:",env=CDN_URL"`
	SigCdnGetSecret     string `json:",env=SIG_CDN_SECRET"`

	RedisAddr     string `json:",env=REDIS_ADDR"`
	RedisPassword string `json:",env=REDIS_PASSWORD,optional"`

	UserRpcConf zrpc.RpcClientConf
}
