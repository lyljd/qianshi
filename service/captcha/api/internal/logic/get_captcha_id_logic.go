package logic

import (
	"context"
	"github.com/dchest/captcha"
	"github.com/zeromicro/go-zero/core/logx"
	"qianshi/common/response/errorx"
	"qianshi/service/captcha/api/internal/svc"
	"qianshi/service/captcha/api/internal/types"
)

type GetCaptchaIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaIdLogic {
	return &GetCaptchaIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaIdLogic) GetCaptchaId() (resp *types.GetCaptchaIdResp, err *errorx.Error) {
	id := captcha.New()
	return &types.GetCaptchaIdResp{
		Id:  id,
		Src: l.svcCtx.ServerAddr + "/api/captcha/" + id,
	}, nil
}
