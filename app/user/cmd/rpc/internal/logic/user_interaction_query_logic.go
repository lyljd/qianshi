package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"qianshi/app/user/model/userInteractionModel"

	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInteractionQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInteractionQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInteractionQueryLogic {
	return &UserInteractionQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInteractionQueryLogic) UserInteractionQuery(in *__.QueryReq) (*__.UserInteractionQueryResp, error) {
	ui, err := userInteractionModel.QueryById(l.svcCtx.Redis, l.svcCtx.DB, in.Uid)
	if err != nil {
		return nil, err
	}

	var resp __.UserInteractionQueryResp
	if err := copier.Copy(&resp, &ui); err != nil {
		return nil, err
	}
	resp.Id = uint64(ui.ID)
	resp.CreatedAt = ui.CreatedAt.Unix()
	resp.UpdatedAt = ui.UpdatedAt.Unix()
	if !ui.DeletedAt.Valid {
		resp.DeletedAt = 0
	} else {
		resp.DeletedAt = ui.DeletedAt.Time.Unix()
	}

	return &resp, nil
}
