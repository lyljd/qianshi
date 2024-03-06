package email

import (
	"context"
	"net/http"
	"qianshi/common/result"
	"qianshi/common/result/errorx"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qianshi/app/vcode/cmd/api/internal/logic/email"
	"qianshi/app/vcode/cmd/api/internal/svc"
	"qianshi/app/vcode/cmd/api/internal/types"
)

func ChangeEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangeEmailReq
		if err := httpx.Parse(r, &req); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, errorx.New(errorx.CodeParamError, err))
			return
		}

		ctx := context.WithValue(r.Context(), "uid", r.Header.Get("UID"))
		l := email.NewChangeEmailLogic(ctx, svcCtx)
		if resp, err := l.ChangeEmail(&req); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, err, resp)
		} else {
			result.Succ(r.Context(), w, resp)
		}
	}
}
