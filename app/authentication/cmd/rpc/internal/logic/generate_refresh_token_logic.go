package logic

import (
	"context"

	"qianshi/app/authentication/cmd/rpc/internal/svc"
	"qianshi/app/authentication/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateRefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateRefreshTokenLogic {
	return &GenerateRefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateRefreshTokenLogic) GenerateRefreshToken(in *__.GenerateRefreshTokenReq) (*__.GenerateRefreshTokenResp, error) {
	token, err := generate(uint(in.Uid), l.svcCtx.Config.RefreshTokenMinutes, l.svcCtx.Config.RefreshTokenSecret)
	if err != nil {
		return nil, err
	}

	return &__.GenerateRefreshTokenResp{Token: token}, nil
}
