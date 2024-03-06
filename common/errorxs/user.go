package errorxs

import (
	"errors"
)

var (
	ErrEmailPassWrong        = errors.New("邮箱或密码错误")
	ErrVcodeWrong            = errors.New("验证码不匹配")
	ErrChangePassVerifyFail  = errors.New("修改密码前验证未通过")
	ErrOldPassSameAsNewPass  = errors.New("新密码与原密码相同")
	ErrCoinInsufficient      = errors.New("硬币不足")
	ErrNicknameHasExist      = errors.New("昵称已存在")
	ErrWrongProcessSequence  = errors.New("错误的流程顺序")
	ErrEmailHasBind          = errors.New("该邮箱已绑定其他账号")
	ErrChangeEmailVerifyFail = errors.New("修改邮箱前验证未通过")
)
