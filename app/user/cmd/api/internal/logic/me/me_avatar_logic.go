package me

import (
	"context"
	"qianshi/app/user/cmd/rpc/user"
	"qianshi/common/ctx"
	"qianshi/common/result/errorx"

	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeAvatarLogic {
	return &MeAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeAvatarLogic) MeAvatar() (resp *types.MeAvatarResp, err error) {
	queryResp, err := l.svcCtx.UserRpc.UserQuery(l.ctx, &user.QueryReq{Uid: uint64(ctx.GetUid(l.ctx))})
	if err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	resp = &types.MeAvatarResp{AvatarUrl: queryResp.AvatarUrl}

	return
}
