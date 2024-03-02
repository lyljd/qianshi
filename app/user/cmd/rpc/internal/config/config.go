package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	MysqlDsn      string `json:",env=MYSQL_DSN"`
	RedisAddr     string `json:",env=REDIS_ADDR"`
	RedisPassword string `json:",env=REDIS_PASSWORD,optional"`

	PassSecret string `json:",env=PASS_SECRET"`

	AuthenticationRpcConf zrpc.RpcClientConf
}
