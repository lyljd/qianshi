package me

import (
	"context"
	"net/http"
	"qianshi/app/user/cmd/api/internal/logic/me"
	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/common/result"
)

func MeExpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "uid", r.Header.Get("UID"))
		l := me.NewMeExpLogic(ctx, svcCtx)
		if resp, err := l.MeExp(); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, err)
		} else {
			result.Succ(r.Context(), w, resp)
		}
	}
}
