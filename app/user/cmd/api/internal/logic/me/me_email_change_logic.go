package me

import (
	"context"
	"errors"
	__ "qianshi/app/user/cmd/rpc/pb"
	"qianshi/common/ctx"
	"qianshi/common/errorxs"
	"qianshi/common/result/errorx"
	"qianshi/common/tool"

	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeEmailChangeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeEmailChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeEmailChangeLogic {
	return &MeEmailChangeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeEmailChangeLogic) MeEmailChange(req *types.MeEmailChangeReq) error {
	if !tool.CheckEmail(req.Email) {
		err := errors.New("传入的邮箱不满足邮箱格式")
		return errorx.New(errorx.CodeParamError, err, err.Error())
	}

	_, err := l.svcCtx.UserRpc.EmailChange(l.ctx, &__.EmailChangeReq{
		Uid:      uint64(ctx.GetUid(l.ctx)),
		NewEmail: req.Email,
	})
	if err != nil {
		if errorxs.Is(err, errorxs.ErrChangeEmailVerifyFail) {
			return errorx.New(errorx.CodeDefault, err, err.Error())
		}
		if errorxs.Is(err, errorxs.ErrEmailHasBind) {
			return errorx.New(errorx.CodeDefault, err, err.Error())
		}
		return errorx.New(errorx.CodeServerError, err)
	}

	return nil
}
