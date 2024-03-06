package me

import (
	"context"
	"net/http"
	"qianshi/common/result"

	"qianshi/app/user/cmd/api/internal/logic/me"
	"qianshi/app/user/cmd/api/internal/svc"
)

func MeAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "uid", r.Header.Get("UID"))
		l := me.NewMeAvatarLogic(ctx, svcCtx)
		if resp, err := l.MeAvatar(); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, err)
		} else {
			result.Succ(r.Context(), w, resp)
		}
	}
}
