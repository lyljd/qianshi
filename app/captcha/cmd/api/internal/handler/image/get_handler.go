package image

import (
	"context"
	"net/http"
	"qianshi/common/result"
	"qianshi/common/result/errorx"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qianshi/app/captcha/cmd/api/internal/logic/image"
	"qianshi/app/captcha/cmd/api/internal/svc"
	"qianshi/app/captcha/cmd/api/internal/types"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetReq
		if err := httpx.Parse(r, &req); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, errorx.New(errorx.CodeParamError, err))
			return
		}

		ctx := context.WithValue(r.Context(), "uid", r.Header.Get("UID"))
		l := image.NewGetLogic(ctx, svcCtx)
		if resp, err := l.Get(&req); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, err)
		} else {
			result.Succ(r.Context(), w, resp)
		}
	}
}
