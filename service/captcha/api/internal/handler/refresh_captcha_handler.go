package handler

import (
	"net/http"
	"qianshi/common/response"
	"qianshi/service/captcha/api/internal/logic"
	"qianshi/service/captcha/api/internal/svc"
	"qianshi/service/captcha/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func refreshCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RefreshCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRefreshCaptchaLogic(r.Context(), svcCtx)
		err := l.RefreshCaptcha(&req)
		response.Response(w, nil, err)
	}
}
