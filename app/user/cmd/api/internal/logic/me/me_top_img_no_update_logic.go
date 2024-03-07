package me

import (
	"context"
	"errors"
	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"
	__ "qianshi/app/user/cmd/rpc/pb"
	"qianshi/common/ctx"
	"qianshi/common/result/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeTopImgNoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeTopImgNoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeTopImgNoUpdateLogic {
	return &MeTopImgNoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeTopImgNoUpdateLogic) MeTopImgNoUpdate(req *types.MeTopImgNoUpdateReq) error {
	if req.TopImgNo < 1 || req.TopImgNo > 4 {
		err := errors.New("头图编号范围必须在1~4之中")
		return errorx.New(errorx.CodeParamError, err, err.Error())
	}

	if _, err := l.svcCtx.UserRpc.UserHomeTopImgNoUpdate(l.ctx, &__.UserHomeTopImgNoUpdateReq{
		Id:          uint64(ctx.GetUid(l.ctx)),
		NewTopImgNo: int64(req.TopImgNo),
	}); err != nil {
		return err
	}

	return nil
}
