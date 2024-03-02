package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"
	"qianshi/app/user/model/userHomeModel"
	"strings"
)

type UserHomeQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserHomeQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomeQueryLogic {
	return &UserHomeQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserHomeQueryLogic) UserHomeQuery(in *__.QueryReq) (*__.UserHomeQueryResp, error) {
	var u userHomeModel.UserHome
	if l.svcCtx.DB.Take(&u, in.Uid).Error == gorm.ErrRecordNotFound {
		return nil, errors.New("记录不存在")
	}

	var resp __.UserHomeQueryResp
	if err := copier.Copy(&resp, &u); err != nil {
		return nil, err
	}
	if u.Tags != "" {
		resp.Tags = strings.Split(u.Tags, ";")
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
