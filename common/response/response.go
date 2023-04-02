package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"qianshi/common/response/errorx"
)

type Body struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp any, err *errorx.Error) {
	var body Body
	if err != nil {
		body.Code = err.Code()
		body.Msg = err.Msg()
	} else {
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
