package me

import (
	"context"
	"net/http"
	"qianshi/app/user/cmd/api/internal/logic/me"
	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/common/result"
)

func MeCoinHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "uid", r.Header.Get("UID"))
		l := me.NewMeCoinLogic(ctx, svcCtx)
		if resp, err := l.MeCoin(); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, err)
		} else {
			result.Succ(r.Context(), w, resp)
		}
	}
}
