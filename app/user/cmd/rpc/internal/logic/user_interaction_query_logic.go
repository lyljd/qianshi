package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
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
	var u userInteractionModel.UserInteraction
	if l.svcCtx.DB.Take(&u, in.Uid).Error == gorm.ErrRecordNotFound {
		return nil, errors.New("记录不存在")
	}

	var resp __.UserInteractionQueryResp
	if err := copier.Copy(&resp, &u); err != nil {
		return nil, err
	}
	resp.Id = uint64(u.ID)
	resp.CreatedAt = u.CreatedAt.Unix()
	resp.UpdatedAt = u.UpdatedAt.Unix()
	if !u.DeletedAt.Valid {
		resp.DeletedAt = 0
	} else {
		resp.DeletedAt = u.DeletedAt.Time.Unix()
	}

	return &resp, nil
}
