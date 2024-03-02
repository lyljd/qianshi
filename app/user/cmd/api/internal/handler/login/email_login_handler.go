package login

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"qianshi/app/user/cmd/api/internal/logic/login"
	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"
	"qianshi/common/result"
	"qianshi/common/result/errorx"
)

func EmailLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, errorx.New(errorx.CodeParamError, err))
			return
		}

		l := login.NewEmailLoginLogic(r.Context(), svcCtx)
		if resp, err := l.EmailLogin(&req, r.Header.Get("IP")); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, err)
		} else {
			result.Succ(r.Context(), w, resp)
		}
	}
}
