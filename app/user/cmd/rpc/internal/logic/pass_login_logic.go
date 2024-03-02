package logic

import (
	"context"
	"gorm.io/gorm"
	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"
	"qianshi/app/user/model/userModel"
	"qianshi/common/errorxs"
	"qianshi/common/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type PassLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPassLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PassLoginLogic {
	return &PassLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PassLoginLogic) PassLogin(in *__.PassLoginReq) (*__.LoginResp, error) {
	pass, err := tool.Sha256(in.Pass, l.svcCtx.Config.PassSecret, "")
	if err != nil {
		return nil, err
	}

	var u userModel.User
	if l.svcCtx.DB.Where("email = ? AND password = ?", in.Email, pass).Take(&u).Error == gorm.ErrRecordNotFound {
		return nil, errorxs.ErrEmailPassWrong
	}

	return loginCommon(l.ctx, l.svcCtx, &u, in.Ip)
}
