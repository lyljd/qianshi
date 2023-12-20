package image

import (
	"context"
	"errors"
	"github.com/dchest/captcha"
	"qianshi/common/key"
	"qianshi/common/result/errorx"
	"time"

	"qianshi/app/captcha/cmd/api/internal/svc"
	"qianshi/app/captcha/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyLogic) Verify(req *types.VerifyReq) error {
	if !captcha.VerifyString(req.Id, req.Code) {
		return errorx.New(errorx.CodeParamError, errors.New("验证码id不存在或未正确输入验证码"), "验证未通过")
	}

	if err := l.svcCtx.Redis.SetNX(l.ctx, key.GetCaptchaVerify(req.Id), "", time.Second*10).Err(); err != nil {
		return errorx.New(errorx.CodeServerError, err)
	}

	return nil
}
