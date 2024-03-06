package login

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"
	__ "qianshi/app/user/cmd/rpc/pb"
	"qianshi/common/errorxs"
	"qianshi/common/result/errorx"
)

type EmailLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailLoginLogic {
	return &EmailLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailLoginLogic) EmailLogin(req *types.EmailLoginReq, ip string) (resp *types.LoginResp, err error) {
	loginRpcResp, err := l.svcCtx.UserRpc.EmailLogin(l.ctx, &__.EmailLoginReq{
		Email: req.Email,
		Code:  req.Code,
		Ip:    ip,
	})

	if err != nil {
		if errorxs.Is(err, errorxs.ErrVcodeWrong) {
			return nil, errorx.New(errorx.CodeParamError, err, err.Error())
		}
		if errorxs.Is(err, errorxs.ErrKeyNotFound) {
			return nil, errorx.New(errorx.CodeParamError, err, "请先获取验证码")
		}
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	resp = new(types.LoginResp)
	if err := copier.Copy(&resp, &loginRpcResp); err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	return
}
