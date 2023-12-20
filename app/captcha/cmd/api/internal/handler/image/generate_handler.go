package image

import (
	"net/http"
	"qianshi/common/result"

	"qianshi/app/captcha/cmd/api/internal/logic/image"
	"qianshi/app/captcha/cmd/api/internal/svc"
)

func GenerateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := image.NewGenerateLogic(r.Context(), svcCtx)
		if resp, err := l.Generate(); err != nil {
			result.Fail(r.Context(), w, svcCtx.Config.Mode, err)
		} else {
			result.Succ(r.Context(), w, resp)
		}
	}
}
