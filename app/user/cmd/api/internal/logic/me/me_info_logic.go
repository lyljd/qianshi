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

type MeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeInfoLogic {
	return &MeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeInfoLogic) MeInfo() (resp *types.MeInfoResp, err error) {
	user, err := l.svcCtx.UserRpc.UserQuery(l.ctx, &__.QueryReq{Uid: uint64(ctx.GetUid(l.ctx))})
	if err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	userHome, err := l.svcCtx.UserRpc.UserHomeQuery(l.ctx, &__.QueryReq{Uid: uint64(ctx.GetUid(l.ctx))})
	if err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	resp = &types.MeInfoResp{
		Nickname:  user.Nickname,
		Signature: user.Signature,
		Gender:    userHome.Gender,
		Birthday:  userHome.Birthday,
		Tags:      userHome.Tags,
	}

	return
}
