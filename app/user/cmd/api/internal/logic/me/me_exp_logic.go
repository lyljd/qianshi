package me

import (
	"context"
	"github.com/jinzhu/copier"
	__ "qianshi/app/user/cmd/rpc/pb"
	"qianshi/common/ctx"
	"qianshi/common/result/errorx"

	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeExpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeExpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeExpLogic {
	return &MeExpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeExpLogic) MeExp() (resp *types.MeExpResp, err error) {
	query, err := l.svcCtx.UserRpc.UserQuery(l.ctx, &__.QueryReq{Uid: uint64(ctx.GetUid(l.ctx))})
	if err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	resp = new(types.MeExpResp)
	if err := copier.Copy(&resp, &query); err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	return
}
