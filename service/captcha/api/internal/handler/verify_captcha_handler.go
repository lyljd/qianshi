package handler

import (
	"net/http"
	"qianshi/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qianshi/service/captcha/api/internal/logic"
	"qianshi/service/captcha/api/internal/svc"
	"qianshi/service/captcha/api/internal/types"
)

func verifyCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewVerifyCaptchaLogic(r.Context(), svcCtx)
		err := l.VerifyCaptcha(&req)
		response.Response(w, nil, err)
	}
}
