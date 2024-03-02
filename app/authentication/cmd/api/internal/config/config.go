package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	RefreshTokenSecret  string `json:",env=REFRESH_TOKEN_SECRET"`
	RefreshTokenMinutes int    `json:",env=REFRESH_TOKEN_MINUTES"`

	AuthenticationRpcConf zrpc.RpcClientConf
}
