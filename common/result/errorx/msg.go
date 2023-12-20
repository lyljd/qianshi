package errorx

var msg = map[int]string{
	CodeDefault:     "服务器开小差了，请稍后再试",
	CodeParamError:  "请求参数异常",
	CodeNoLogin:     "请登录",
	CodeNoPower:     "无权访问",
	CodeNotFound:    "资源未找到",
	CodeServerError: "服务器繁忙，请稍后再试",
}

func getDefaultMsg() string {
	return msg[CodeDefault]
}

func getMsg(code int) string {
	if m, ok := msg[code]; ok {
		return m
	}

	return getDefaultMsg()
}
