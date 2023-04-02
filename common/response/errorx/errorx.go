package errorx

const DefaultCode = -1

type Error struct {
	code int
	msg  string
}

func New(code int, msg string) *Error {
	return &Error{code: code, msg: msg}
}

func NewDefault(msg string) *Error {
	return New(DefaultCode, msg)
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}
