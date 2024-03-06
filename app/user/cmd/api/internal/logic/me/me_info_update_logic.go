package me

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"
	__ "qianshi/app/user/cmd/rpc/pb"
	"qianshi/common/ctx"
	"qianshi/common/errorxs"
	"qianshi/common/result/errorx"
	"qianshi/common/tool"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeInfoUpdateLogic {
	return &MeInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeInfoUpdateLogic) MeInfoUpdate(req *types.MeInfoUpdateReq) error {
	req.Nickname = strings.TrimSpace(req.Nickname)
	req.Signature = strings.TrimSpace(req.Signature)
	if req.Nickname == "" {
		err := errors.New("昵称不能为空")
		return errorx.New(errorx.CodeParamError, err, err.Error())
	}
	if len(req.Nickname) > 20 {
		err := errors.New("昵称的长度最大为20")
		return errorx.New(errorx.CodeParamError, err, err.Error())
	}
	if len(req.Signature) > 50 {
		err := errors.New("签名的长度最大为50")
		return errorx.New(errorx.CodeParamError, err, err.Error())
	}
	if req.Gender != "男" && req.Gender != "女" && req.Gender != "保密" {
		err := errors.New("性别必须为[男,女,保密]中的一项")
		return errorx.New(errorx.CodeParamError, err, err.Error())
	}
	if req.Birthday != "" && !tool.CheckDateStr(req.Birthday) {
		err := errors.New("出生日期格式不符合要求或不合法")
		return errorx.New(errorx.CodeParamError, err, err.Error())
	}
	if ds := tool.CheckStrSliDup(req.Tags); len(ds) > 0 {
		err := errors.New("tags中 " + strings.Join(ds, ", "+" 重复"))
		return errorx.New(errorx.CodeParamError, err, err.Error())
	}

	rpcReq := new(__.MeInfoUpdateReq)
	if err := copier.Copy(rpcReq, req); err != nil {
		return errorx.New(errorx.CodeServerError, err)
	}
	rpcReq.Id = uint64(ctx.GetUid(l.ctx))

	if _, err := l.svcCtx.UserRpc.MeInfoUpdate(l.ctx, rpcReq); err != nil {
		if errorxs.Is(err, errorxs.ErrCoinInsufficient) {
			return errorx.New(errorx.CodeParamError, err, err.Error())
		}
		if errorxs.Is(err, errorxs.ErrNicknameHasExist) {
			return errorx.New(errorx.CodeParamError, err, err.Error())
		}
		return errorx.New(errorx.CodeDefault, err)
	}

	return nil
}
