package logic

import (
	"context"
	"github.com/dchest/captcha"
	"qianshi/common/response/errorx"
	"qianshi/service/captcha/api/internal/svc"
	"qianshi/service/captcha/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshCaptchaLogic {
	return &RefreshCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshCaptchaLogic) RefreshCaptcha(req *types.RefreshCaptchaReq) *errorx.Error {
	ok := captcha.Reload(req.Id)
	if ok {
		return nil
	}
	return errorx.NewDefault("刷新失败")
}
