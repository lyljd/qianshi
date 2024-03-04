package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Name string
	Host string `json:",default=0.0.0.0"`
	Port int
	Mode string `json:",default=pro,options=dev|test|rt|pre|pro"`

	AuthenticationRpcConf zrpc.RpcClientConf
	UserRpcConf           zrpc.RpcClientConf
}
