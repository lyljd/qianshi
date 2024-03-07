package logic

import (
	"context"
	"qianshi/app/user/model/userModel"
	"qianshi/common/errorxs"

	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSignatureUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSignatureUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSignatureUpdateLogic {
	return &UserSignatureUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSignatureUpdateLogic) UserSignatureUpdate(in *__.UserSignatureUpdateReq) (*__.UserSignatureUpdateResp, error) {
	u, err := userModel.QueryById(l.svcCtx.Redis, l.svcCtx.DB, in.Id)
	if err != nil {
		return nil, err
	}

	if in.NewSignature == u.Signature {
		return nil, errorxs.ErrOldSigSameAsNewSig
	}

	if err = userModel.UpdateByIdWithNil(l.svcCtx.Redis, l.svcCtx.DB, u, map[string]any{
		"signature": in.NewSignature,
	}); err != nil {
		return nil, err
	}

	return &__.UserSignatureUpdateResp{}, nil
}
