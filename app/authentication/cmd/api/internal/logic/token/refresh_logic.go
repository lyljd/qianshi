package token

import (
	"context"
	__ "qianshi/app/authentication/cmd/rpc/pb"
	"qianshi/common/result/errorx"

	"qianshi/app/authentication/cmd/api/internal/svc"
	"qianshi/app/authentication/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh(req *types.TokenRefreshReq) (resp *types.TokenRefreshResp, err error) {
	verifyResp, err := l.svcCtx.AuthenticationRpc.VerifyRefreshToken(
		l.ctx,
		&__.VerifyRefreshTokenReq{Token: req.RefreshToken},
	)
	if err != nil {
		return nil, errorx.New(errorx.CodeDefault, err)
	}

	generateResp, err := l.svcCtx.AuthenticationRpc.GenerateToken(
		l.ctx,
		&__.GenerateTokenReq{Uid: verifyResp.Uid},
	)
	if err != nil {
		return nil, errorx.New(errorx.CodeDefault, err)
	}

	resp = &types.TokenRefreshResp{Token: generateResp.Token}
	return
}
