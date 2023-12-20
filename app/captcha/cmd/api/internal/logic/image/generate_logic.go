package image

import (
	"context"
	"github.com/dchest/captcha"

	"qianshi/app/captcha/cmd/api/internal/svc"
	"qianshi/app/captcha/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateLogic {
	return &GenerateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateLogic) Generate() (resp *types.GenerateResp, err error) {
	id := captcha.NewLen(4)
	resp = &types.GenerateResp{Id: id}
	return
}
