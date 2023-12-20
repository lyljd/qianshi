package errorx

import (
	"encoding/json"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type Error struct {
	code int
	errorDesc
}

type errorDesc struct {
	Err string `json:"err"`
	Msg string `json:"msg"`
}

func New(code int, err error, msgs ...string) error {
	ed := errorDesc{
		Err: err.Error(),
		Msg: strings.Join(msgs, "；"),
	}
	if ed.Msg == "" {
		ed.Msg = getMsg(code)
	}
	desc, _ := json.Marshal(ed)

	return status.Error(codes.Code(code), string(desc))
}

func Convert(err error) *Error {
	if err == nil {
		return &Error{
			code: CodeDefault,
			errorDesc: errorDesc{
				Err: "err为空",
				Msg: getDefaultMsg(),
			},
		}
	}

	s, _ := status.FromError(err)
	var desc errorDesc
	_ = json.Unmarshal([]byte(s.Message()), &desc)

	return &Error{
		code: int(s.Code()),
		errorDesc: errorDesc{
			Err: desc.Err,
			Msg: desc.Msg,
		},
	}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Err() string {
	return e.errorDesc.Err
}

func (e *Error) Msg() string {
	return e.errorDesc.Msg
}
