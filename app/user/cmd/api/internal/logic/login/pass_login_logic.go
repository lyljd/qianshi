package login

import (
	"context"
	"github.com/jinzhu/copier"
	__ "qianshi/app/user/cmd/rpc/pb"
	"qianshi/common/errorxs"
	"qianshi/common/result/errorx"

	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PassLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPassLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PassLoginLogic {
	return &PassLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PassLoginLogic) PassLogin(req *types.PassLoginReq, ip string) (resp *types.LoginResp, err error) {
	loginRpcResp, err := l.svcCtx.UserRpc.PassLogin(l.ctx, &__.PassLoginReq{
		Email: req.Email,
		Pass:  req.Pass,
		Ip:    ip,
	})

	if err != nil {
		if errorxs.Is(err, errorxs.ErrEmailPassWrong) {
			return nil, errorx.New(errorx.CodeParamError, err, err.Error())
		}
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	resp = new(types.LoginResp)
	if err := copier.Copy(&resp, &loginRpcResp); err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	return
}
