// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	login "qianshi/app/user/cmd/api/internal/handler/login"
	me "qianshi/app/user/cmd/api/internal/handler/me"
	"qianshi/app/user/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/email",
				Handler: login.EmailLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/pass",
				Handler: login.PassLoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/user/login"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/exp",
				Handler: me.MeExpHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: me.MeInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/security",
				Handler: me.MeSecurityHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/coin",
				Handler: me.MeCoinHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/pass/verify",
				Handler: me.MePassVerifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/pass/change",
				Handler: me.MePassChangeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/user/me"),
	)
}