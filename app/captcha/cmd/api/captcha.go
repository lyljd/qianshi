package main

import (
	"flag"
	"fmt"
	"qianshi/common/xlog"

	"qianshi/app/captcha/cmd/api/internal/config"
	"qianshi/app/captcha/cmd/api/internal/handler"
	"qianshi/app/captcha/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	"github.com/joho/godotenv"
)

var configFile = flag.String("f", "etc/captcha.yaml", "the config file")

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	xlog.StartCollection(c.Name, nil)
	defer xlog.StopCollection()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
