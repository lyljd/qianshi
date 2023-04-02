package handler

import (
	"net/http"
	"qianshi/common/response"
	"qianshi/service/captcha/api/internal/logic"
	"qianshi/service/captcha/api/internal/svc"
)

func getCaptchaIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetCaptchaIdLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptchaId()
		response.Response(w, resp, err)
	}
}
