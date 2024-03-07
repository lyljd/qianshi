package me

import (
	"context"
	"errors"
	"qianshi/app/user/cmd/rpc/user"
	"qianshi/common/ctx"
	"qianshi/common/errorxs"
	"qianshi/common/result/errorx"
	"strings"

	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeSigUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeSigUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeSigUpdateLogic {
	return &MeSigUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeSigUpdateLogic) MeSigUpdate(req *types.MeSigUpdateReq) error {
	newSignature := strings.TrimSpace(req.NewSignature)
	if len(newSignature) > 50 {
		err := errors.New("签名的最大长度为50")
		return errorx.New(errorx.CodeParamError, err, err.Error())
	}

	if _, err := l.svcCtx.UserRpc.UserSignatureUpdate(l.ctx, &user.UserSignatureUpdateReq{
		Id:           uint64(ctx.GetUid(l.ctx)),
		NewSignature: newSignature,
	}); err != nil {
		if errorxs.Is(err, errorxs.ErrOldSigSameAsNewSig) {
			return errorx.New(errorx.CodeParamError, err, err.Error())
		}
		return errorx.New(errorx.CodeServerError, err)
	}

	return nil
}
