package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"qianshi/app/authentication/cmd/rpc/authentication"
	"qianshi/app/user/cmd/rpc/internal/config"
)

type ServiceContext struct {
	Config            config.Config
	DB                *gorm.DB
	Redis             *redis.Client
	AuthenticationRpc authentication.Authentication
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.MysqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.RedisAddr,
			Password: c.RedisPassword,
		}),
		AuthenticationRpc: authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpcConf)),
	}
}
