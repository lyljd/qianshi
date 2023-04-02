package logic

import (
	"context"

	"qianshi/service/captcha/api/internal/svc"
	"qianshi/service/captcha/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaPngLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaPngLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaPngLogic {
	return &GetCaptchaPngLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaPngLogic) GetCaptchaPng(req *types.GetCaptchaPngReq) error {
	// todo: add your logic here and delete this line

	return nil
}
