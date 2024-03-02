package logic

import (
	"context"
	"errors"
	__2 "qianshi/app/user/cmd/rpc/pb"

	"qianshi/app/authentication/cmd/rpc/internal/svc"
	"qianshi/app/authentication/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyRefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyRefreshTokenLogic {
	return &VerifyRefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyRefreshTokenLogic) VerifyRefreshToken(in *__.VerifyRefreshTokenReq) (*__.VerifyRefreshTokenResp, error) {
	uid, _, err := verify(in.Token, l.svcCtx.Config.RefreshTokenSecret)
	if err != nil {
		return nil, err
	}

	query, err := l.svcCtx.UserRpc.UserQuery(l.ctx, &__2.QueryReq{Uid: uint64(uid)})
	if err != nil {
		return nil, err
	}
	if query.RefreshToken != in.Token {
		return nil, errors.New("传入的refreshToken与数据库中的refreshToken不匹配")
	}

	return &__.VerifyRefreshTokenResp{Uid: int64(uid)}, nil
}
