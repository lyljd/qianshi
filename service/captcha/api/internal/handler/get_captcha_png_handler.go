package handler

import (
	"github.com/dchest/captcha"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"qianshi/service/captcha/api/internal/svc"
	"qianshi/service/captcha/api/internal/types"
)

func getCaptchaPngHandler(_ *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCaptchaPngReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		_ = captcha.WriteImage(w, req.Id, 240, 80)
	}
}
