package me

import (
	"context"
	"net/http"
	"qianshi/common/result"
	"qianshi/common/result/errorx"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qianshi/app/user/cmd/api/internal/logic/me"
	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"
)

func MeEmailVerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MeEmailVerifyReq
		if err := httpx.Parse(r, &req); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, errorx.New(errorx.CodeParamError, err))
			return
		}

		ctx := context.WithValue(r.Context(), "uid", r.Header.Get("UID"))
		l := me.NewMeEmailVerifyLogic(ctx, svcCtx)
		if resp, err := l.MeEmailVerify(&req); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, err)
		} else {
			result.Succ(r.Context(), w, resp)
		}
	}
}
