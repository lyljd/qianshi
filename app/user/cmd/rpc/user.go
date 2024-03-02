package main

import (
	"flag"
	"fmt"
	"qianshi/app/user/model/userHomeModel"
	"qianshi/app/user/model/userInteractionModel"
	"qianshi/app/user/model/userModel"
	"qianshi/common/xlog"

	"qianshi/app/user/cmd/rpc/internal/config"
	"qianshi/app/user/cmd/rpc/internal/server"
	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	xlog.StartCollection(c.Name, nil)
	defer xlog.StopCollection()

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		__.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	if err := ctx.DB.AutoMigrate(&userModel.User{}, &userHomeModel.UserHome{}, &userInteractionModel.UserInteraction{}); err != nil {
		panic(err)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
