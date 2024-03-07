package user

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"qianshi/app/user/cmd/api/internal/types"
	"qianshi/common/result"
	"qianshi/common/result/errorx"

	"qianshi/app/user/cmd/api/internal/logic/user"
	"qianshi/app/user/cmd/api/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, errorx.New(errorx.CodeParamError, err))
			return
		}

		ctx := context.WithValue(r.Context(), "uid", r.Header.Get("UID"))
		l := user.NewUserInfoLogic(ctx, svcCtx)
		if resp, err := l.UserInfo(&req); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, err)
		} else {
			result.Succ(r.Context(), w, resp)
		}
	}
}
