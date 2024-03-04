package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
	__ "qianshi/app/authentication/cmd/rpc/pb"
	"qianshi/app/gateway/internal/router"
	"qianshi/app/gateway/internal/svc"
	__2 "qianshi/app/user/cmd/rpc/pb"
	"strconv"
	"strings"
)

func Start(svcCtx *svc.ServiceContext) {
	r := gin.Default()

	r.Any("/api/v1/*path", func(c *gin.Context) {
		paths := strings.Split(c.Param("path")[1:], "/")

		find, to, needLogin, power := search(router.Routers, paths)
		if !find {
			c.String(404, "404 Not found")
			return
		}

		targetURL, err := url.Parse(to)
		if err != nil {
			c.String(500, "500 Internal Server Error")
			return
		}

		if needLogin || power > 0 {
			token := c.Request.Header.Get("Token")
			if token == "" {
				c.String(401, "401 Unauthorized")
				return
			}

			verifyResp, err := svcCtx.AuthenticationRpc.VerifyToken(context.Background(), &__.VerifyTokenReq{Token: token})
			if err != nil {
				c.String(401, "401 Unauthorized")
				return
			}

			queryResp, err := svcCtx.UserRpc.UserQuery(context.Background(), &__2.QueryReq{Uid: uint64(verifyResp.Uid)})
			if err != nil {
				c.String(500, "500 Internal Server Error")
				return
			}

			if int(queryResp.Power) < power {
				c.String(403, "403 Forbidden")
				return
			}

			c.Request.Header.Set("UID", strconv.FormatInt(verifyResp.Uid, 10))
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		c.Request.Host = targetURL.Host
		c.Request.URL.Host = targetURL.Host
		c.Request.URL.Scheme = targetURL.Scheme

		c.Request.Header.Set("IP", getIP(c.Request))

		proxy.ServeHTTP(c.Writer, c.Request)
	})

	_ = r.Run(fmt.Sprintf("%s:%d", svcCtx.Config.Host, svcCtx.Config.Port))
}
