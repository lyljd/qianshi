package errorxs

import (
	"errors"
)

var (
	ErrEmailPassWrong       = errors.New("邮箱或密码错误")
	ErrVcodeWrong           = errors.New("验证码不匹配")
	ErrChangePassVerifyFail = errors.New("修改密码前验证未通过")
	ErrOldPassSameAsNewPass = errors.New("新密码与原密码相同")
)
