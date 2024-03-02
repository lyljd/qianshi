package me

import (
	"context"
	__ "qianshi/app/user/cmd/rpc/pb"
	"qianshi/common/ctx"
	"qianshi/common/errorxs"
	"qianshi/common/result/errorx"

	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MePassVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMePassVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MePassVerifyLogic {
	return &MePassVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MePassVerifyLogic) MePassVerify(req *types.MePassVerifyReq) (err error) {
	_, err = l.svcCtx.UserRpc.PassChangeVerify(l.ctx, &__.PassChangeVerifyReq{
		Uid:  uint64(int64(ctx.GetUid(l.ctx))),
		Code: req.Code,
	})

	if err != nil {
		if errorxs.Is(err, errorxs.ErrVcodeWrong) {
			return errorx.New(errorx.CodeParamError, err, err.Error())
		}
		return errorx.New(errorx.CodeDefault, err)
	}

	return
}
