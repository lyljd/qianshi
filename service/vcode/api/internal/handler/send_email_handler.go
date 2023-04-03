package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"qianshi/common/response"
	"qianshi/service/vcode/api/internal/logic"
	"qianshi/service/vcode/api/internal/svc"
	"qianshi/service/vcode/api/internal/types"
)

func sendEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendEmailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendEmailLogic(r.Context(), svcCtx)
		resp, err := l.SendEmail(&req)
		if resp == nil {
			response.Response(w, nil, err)
			return
		}
		response.Response(w, resp, err)
	}
}
