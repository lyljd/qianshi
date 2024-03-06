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

type MeEmailVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeEmailVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeEmailVerifyLogic {
	return &MeEmailVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeEmailVerifyLogic) MeEmailVerify(req *types.MeEmailVerifyReq) (resp *types.MeEmailVerifyResp, err error) {
	ecvResp, err := l.svcCtx.UserRpc.EmailChangeVerify(l.ctx, &__.EmailChangeVerifyReq{
		Uid:   uint64(int64(ctx.GetUid(l.ctx))),
		Email: req.Email,
		Code:  req.Code,
	})

	if err != nil {
		if errorxs.Is(err, errorxs.ErrVcodeWrong) {
			return nil, errorx.New(errorx.CodeParamError, err, err.Error())
		}
		if errorxs.Is(err, errorxs.ErrKeyNotFound) {
			return nil, errorx.New(errorx.CodeParamError, err, "请先获取验证码")
		}
		if errorxs.Is(err, errorxs.ErrWrongProcessSequence) {
			return nil, errorx.New(errorx.CodeParamError, err, err.Error())
		}
		return nil, errorx.New(errorx.CodeDefault, err)
	}

	return &types.MeEmailVerifyResp{Ttl: int(ecvResp.Ttl)}, nil
}
