package errorx

// code不能为负数，不然会overflow
const (
	CodeDefault     = 1
	CodeParamError  = 400
	CodeNoLogin     = 401
	CodeNoPower     = 403
	CodeNotFound    = 404
	CodeServerError = 500
)
