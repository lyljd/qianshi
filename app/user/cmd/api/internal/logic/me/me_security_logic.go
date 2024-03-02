package me

import (
	"context"
	__ "qianshi/app/user/cmd/rpc/pb"
	"qianshi/common/ctx"
	"qianshi/common/result/errorx"

	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeSecurityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeSecurityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeSecurityLogic {
	return &MeSecurityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeSecurityLogic) MeSecurity() (resp *types.MeSecurityResp, err error) {
	query, err := l.svcCtx.UserRpc.UserQuery(l.ctx, &__.QueryReq{Uid: uint64(ctx.GetUid(l.ctx))})
	if err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	resp = &types.MeSecurityResp{
		IsSetPassword: query.Password != "",
		Email:         query.Email,
	}

	return
}
