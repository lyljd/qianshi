package image

import (
	"context"
	"errors"
	"github.com/dchest/captcha"
	"qianshi/common/result/errorx"

	"qianshi/app/captcha/cmd/api/internal/svc"
	"qianshi/app/captcha/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReloadLogic {
	return &ReloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReloadLogic) Reload(req *types.ReloadReq) error {
	if !captcha.Reload(req.Id) {
		return errorx.New(errorx.CodeParamError, errors.New("验证码id不存在"), "刷新失败")
	}

	return nil
}
