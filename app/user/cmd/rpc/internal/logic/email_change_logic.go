package logic

import (
	"context"
	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"
	"qianshi/app/user/model/userModel"
	"qianshi/common/errorxs"
	"qianshi/common/key"
	"qianshi/common/result/errorx"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailChangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailChangeLogic {
	return &EmailChangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmailChangeLogic) EmailChange(in *__.EmailChangeReq) (*__.EmailChangeResp, error) {
	u, err := userModel.QueryById(l.svcCtx.Redis, l.svcCtx.DB, in.Uid)
	if err != nil {
		return nil, err
	}

	verifyKey := key.GetUserChangeEmailVerify(u.Email)
	existRes, existErr := l.svcCtx.Redis.Exists(l.ctx, verifyKey).Result()
	if existErr != nil {
		return nil, existErr
	}
	if existRes != 1 {
		return nil, errorxs.ErrChangeEmailVerifyFail
	}

	// 判断新邮箱是否绑定其他账号
	if _, err := userModel.QueryByEmail(l.svcCtx.Redis, l.svcCtx.DB, in.NewEmail); err == nil {
		return nil, errorxs.ErrEmailHasBind
	} else if !errorxs.Is(err, errorxs.ErrKeyNotFound) && !errorxs.Is(err, errorxs.ErrRecordNotFound) {
		return nil, err
	}

	// 清除已验证key
	if err := l.svcCtx.Redis.Del(l.ctx, verifyKey).Err(); err != nil {
		return nil, errorx.New(errorx.CodeParamError, err)
	}

	// 将想更换邮箱的邮箱放入redis
	if err := l.svcCtx.Redis.SetEX(l.ctx, key.GetBindEmail(in.NewEmail), u.Email, time.Minute*5).Err(); err != nil {
		return nil, err
	}

	return &__.EmailChangeResp{}, nil
}
