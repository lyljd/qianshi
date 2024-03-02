package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"qianshi/app/user/model/userModel"
	"qianshi/common/errorxs"
	"qianshi/common/key"
	"time"

	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PassChangeVerifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPassChangeVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PassChangeVerifyLogic {
	return &PassChangeVerifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PassChangeVerifyLogic) PassChangeVerify(in *__.PassChangeVerifyReq) (*__.PassChangeVerifyResp, error) {
	var u userModel.User
	if l.svcCtx.DB.Select("email").Take(&u, in.Uid).Error == gorm.ErrRecordNotFound {
		return nil, errors.New("记录不存在")
	}

	vcodeVerifyKey := key.GetVcodeChangePasswordVerify(u.Email)
	vcode, err := l.svcCtx.Redis.Get(l.ctx, vcodeVerifyKey).Result()
	if err != nil {
		return nil, err
	}
	if vcode != in.Code {
		return nil, errorxs.ErrVcodeWrong
	}

	if err := l.svcCtx.Redis.Del(l.ctx, vcodeVerifyKey).Err(); err != nil {
		return nil, err
	}

	userVerifyKey := key.GetUserChangePasswordVerify(u.Email)
	if err := l.svcCtx.Redis.SetEX(l.ctx, userVerifyKey, "", time.Minute*5).Err(); err != nil {
		return nil, err
	}

	return &__.PassChangeVerifyResp{}, nil
}
