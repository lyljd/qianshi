package result

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"qianshi/common/result/errorx"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Err  string `json:"err,omitempty"`
}

func Succ(ctx context.Context, w http.ResponseWriter, data any) {
	resp := &Response{
		Data: data,
	}

	httpx.OkJsonCtx(ctx, w, resp)
}

// Fail data只传第一个，多的不传
func Fail(ctx context.Context, w http.ResponseWriter, mode string, err error, data ...any) {
	errX := errorx.Convert(err)
	resp := &Response{
		Code: errX.Code(),
		Msg:  errX.Msg(),
	}
	if mode != "pro" {
		resp.Err = errX.Err()
	}
	if data != nil && data[0] != nil {
		resp.Data = data[0]
	}

	httpx.OkJsonCtx(ctx, w, resp)
}
