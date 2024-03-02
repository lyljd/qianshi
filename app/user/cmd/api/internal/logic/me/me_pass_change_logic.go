package me

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	__ "qianshi/app/user/cmd/rpc/pb"
	"qianshi/common/ctx"
	"qianshi/common/errorxs"
	"qianshi/common/result/errorx"

	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MePassChangeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMePassChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MePassChangeLogic {
	return &MePassChangeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MePassChangeLogic) MePassChange(req *types.MePassChangeReq) (resp *types.MePassChangeResp, err error) {
	if len(req.Pass) < 6 || len(req.Pass) > 20 {
		err := errors.New("密码长度需在6~20位内")
		return nil, errorx.New(errorx.CodeParamError, err, err.Error())
	}

	pcResp, err := l.svcCtx.UserRpc.PassChange(l.ctx, &__.PassChangeReq{
		Uid:  uint64(ctx.GetUid(l.ctx)),
		Pass: req.Pass,
	})
	if err != nil {
		if errorxs.Is(err, errorxs.ErrChangePassVerifyFail) {
			return nil, errorx.New(errorx.CodeDefault, err, err.Error())
		}
		if errorxs.Is(err, errorxs.ErrOldPassSameAsNewPass) {
			return nil, errorx.New(errorx.CodeDefault, err, err.Error())
		}
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	resp = new(types.MePassChangeResp)
	if err := copier.Copy(resp, pcResp); err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	return
}
