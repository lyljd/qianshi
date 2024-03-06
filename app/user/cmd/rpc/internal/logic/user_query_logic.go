package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"
	"qianshi/app/user/model/userModel"
)

type UserQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserQueryLogic {
	return &UserQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserQueryLogic) UserQuery(in *__.QueryReq) (*__.UserQueryResp, error) {
	// 传入了uid则根据id来查询，没传就看是否传了email，如果传入了email则根据email来查询，没传就报错
	var u *userModel.User
	var err error

	if in.Uid != 0 {
		u, err = userModel.QueryById(l.svcCtx.Redis, l.svcCtx.DB, in.Uid)
		if err != nil {
			return nil, err
		}
	} else if in.Email != "" {
		u, err = userModel.QueryByEmail(l.svcCtx.Redis, l.svcCtx.DB, in.Email)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("user查询关键字为空")
	}

	var resp __.UserQueryResp
	if err = copier.Copy(&resp, &u); err != nil {
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
