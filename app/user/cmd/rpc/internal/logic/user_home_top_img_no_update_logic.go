package logic

import (
	"context"
	"gorm.io/gorm"
	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"
	"qianshi/app/user/model/userHomeModel"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomeTopImgNoUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserHomeTopImgNoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomeTopImgNoUpdateLogic {
	return &UserHomeTopImgNoUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserHomeTopImgNoUpdateLogic) UserHomeTopImgNoUpdate(in *__.UserHomeTopImgNoUpdateReq) (*__.UserHomeTopImgNoUpdateResp, error) {
	if err := userHomeModel.UpdateById(l.svcCtx.Redis, l.svcCtx.DB, &userHomeModel.UserHome{Model: gorm.Model{
		ID: uint(in.Id)},
	}, &userHomeModel.UserHome{
		TopImgNo: int(in.NewTopImgNo),
	}); err != nil {
		return nil, err
	}

	return &__.UserHomeTopImgNoUpdateResp{}, nil
}
