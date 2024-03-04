package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"qianshi/app/gateway/internal/config"
	"qianshi/app/gateway/internal/server"
	"qianshi/app/gateway/internal/svc"
	"qianshi/common/xlog"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	xlog.StartCollection(c.Name, nil)
	defer xlog.StopCollection()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	logx.DisableStat()
	server.Start(svc.NewServiceContext(c))
}
