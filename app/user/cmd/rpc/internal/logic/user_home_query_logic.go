package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
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
	uh, err := userHomeModel.QueryById(l.svcCtx.Redis, l.svcCtx.DB, in.Uid)
	if err != nil {
		return nil, err
	}

	var resp __.UserHomeQueryResp
	if err := copier.Copy(&resp, &uh); err != nil {
		return nil, err
	}
	if uh.Tags != "" {
		resp.Tags = strings.Split(uh.Tags, ";")
	}
	resp.Id = uint64(uh.ID)
	resp.CreatedAt = uh.CreatedAt.Unix()
	resp.UpdatedAt = uh.UpdatedAt.Unix()
	if !uh.DeletedAt.Valid {
		resp.DeletedAt = 0
	} else {
		resp.DeletedAt = uh.DeletedAt.Time.Unix()
	}

	return &resp, nil
}
