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

type MePowerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMePowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MePowerLogic {
	return &MePowerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MePowerLogic) MePower() (resp *types.MePowerResp, err error) {
	queryResp, err := l.svcCtx.UserRpc.UserQuery(l.ctx, &user.QueryReq{Uid: uint64(ctx.GetUid(l.ctx))})
	if err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	resp = &types.MePowerResp{Power: int(queryResp.Power)}

	return
}
