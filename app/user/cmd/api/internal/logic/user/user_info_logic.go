package user

import (
	"context"
	"github.com/jinzhu/copier"
	"qianshi/app/user/cmd/rpc/user"
	"qianshi/common/errorxs"
	"qianshi/common/result/errorx"
	"time"

	"qianshi/app/user/cmd/api/internal/svc"
	"qianshi/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	uid := req.Id

	u, err := l.svcCtx.UserRpc.UserQuery(l.ctx, &user.QueryReq{Uid: uid})
	if err != nil {
		if errorxs.Is(err, errorxs.ErrKeyNotFound) || errorxs.Is(err, errorxs.ErrRecordNotFound) {
			return nil, errorx.New(errorx.CodeNotFound, err, "用户不存在")
		}
		return nil, errorx.New(errorx.CodeServerError, err)
	}
	uh, err := l.svcCtx.UserRpc.UserHomeQuery(l.ctx, &user.QueryReq{Uid: uid})
	if err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	resp = new(types.UserInfoResp)
	if err := copier.Copy(resp, u); err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}
	if err := copier.Copy(resp, uh); err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	resp.Uid, resp.IsVip = int(uid), time.Unix(u.VipExpire, 0).After(time.Now())

	// TODO 查是否关注
	resp.IsFocu = false

	// TODO 查是否拉黑
	resp.IsBlock = false

	// TODO 查投稿数、合集数、收藏夹数
	resp.PostNum, resp.CollectionNum, resp.FavlistNum = 0, 0, 0

	return
}
