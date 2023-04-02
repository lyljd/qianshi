package logic

import (
	"context"
	"github.com/dchest/captcha"
	"qianshi/common/response/errorx"

	"qianshi/service/captcha/api/internal/svc"
	"qianshi/service/captcha/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCaptchaLogic {
	return &VerifyCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyCaptchaLogic) VerifyCaptcha(req *types.VerifyCaptchaReq) *errorx.Error {
	ok := captcha.VerifyString(req.Id, req.Digits)
	if ok {
		return nil
	}
	return errorx.NewDefault("验证码错误")
}
