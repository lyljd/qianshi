package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"path"
	"runtime"

	"qianshi/service/vcode/api/internal/config"
	"qianshi/service/vcode/api/internal/handler"
	"qianshi/service/vcode/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	flag.Parse()

	var c config.Config
	_, filename, _, _ := runtime.Caller(0)
	configFilePath := path.Dir(filename) + "/etc/vcode.yaml"
	conf.MustLoad(configFilePath, &c)

	err := godotenv.Load("./service/vcode/api/.env")
	if err != nil {
		panic(err)
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
