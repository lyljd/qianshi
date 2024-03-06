package logic

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"qianshi/app/user/model/userModel"
	"qianshi/common/errorxs"
	"qianshi/common/key"
	"time"

	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailChangeVerifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailChangeVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailChangeVerifyLogic {
	return &EmailChangeVerifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// EmailChangeVerify 该接口为更换邮箱的两用接口，根据是否只传了uid(第一次)和同时传了uid和email(第二次)决定是第一次邮箱验证还是第二次邮箱验证。
// 因为更换邮箱的流程是要先验证当前绑定的邮箱，然后再验证要绑定的邮箱，所以为了实现前端的一个窗口复用，该接口采取了两用。
// 第一次邮箱验证为【更换邮箱】验证，第二次邮箱验证为【绑定邮箱】验证。
// 当通过了第一次邮箱验证时，会向redis写入一个待绑定的key；当通过了第二次邮箱验证时，会取第一次存入key的值（也就是待更换邮箱）来更新成新邮箱。
// 当同时传了uid和email时（视为第二次邮箱验证），会先对第一次邮箱验证后在设置过新邮箱后产生的key进行判断是否存在，如果不存在则视为恶意调用。
// 所以该两用接口不存在先后调用顺序调换的安全问题。
func (l *EmailChangeVerifyLogic) EmailChangeVerify(in *__.EmailChangeVerifyReq) (*__.EmailChangeVerifyResp, error) {
	if in.Uid != 0 && in.Email == "" {
		// 开始第一次邮箱验证【更换邮箱】的流程
		u, err := userModel.QueryById(l.svcCtx.Redis, l.svcCtx.DB, in.Uid)
		if err != nil {
			return nil, err
		}

		vcodeVerifyKey := key.GetVcodeChangeEmailVerify(u.Email)
		vcode, err := l.svcCtx.Redis.Get(l.ctx, vcodeVerifyKey).Result()
		if err != nil {
			if err == redis.Nil {
				return nil, errorxs.ErrKeyNotFound
			}
			return nil, err
		}
		if vcode != in.Code {
			return nil, errorxs.ErrVcodeWrong
		}

		if err := l.svcCtx.Redis.Del(l.ctx, vcodeVerifyKey).Err(); err != nil {
			return nil, err
		}

		const ttl = time.Minute * 5
		userVerifyKey := key.GetUserChangeEmailVerify(u.Email)
		if err := l.svcCtx.Redis.SetEX(l.ctx, userVerifyKey, "", ttl).Err(); err != nil {
			return nil, err
		}

		return &__.EmailChangeVerifyResp{Ttl: int64(ttl.Seconds())}, nil
	} else if in.Uid != 0 && in.Email != "" {
		// 开始第二次邮箱验证【绑定邮箱】的流程（此次验证中uid只作为更新时删除缓存时所使用）
		vcodeVerifyKey := key.GetVcodeChangeEmailVerify(in.Email)
		vcode, err := l.svcCtx.Redis.Get(l.ctx, vcodeVerifyKey).Result()
		if err != nil {
			if err == redis.Nil {
				return nil, errorxs.ErrKeyNotFound
			}
			return nil, err
		}
		if vcode != in.Code {
			return nil, errorxs.ErrVcodeWrong
		}

		if err := l.svcCtx.Redis.Del(l.ctx, vcodeVerifyKey).Err(); err != nil {
			return nil, err
		}

		bindEmailKey := key.GetBindEmail(in.Email)
		wantChangedEmail, err := l.svcCtx.Redis.Get(l.ctx, bindEmailKey).Result()
		if err != nil {
			if err == redis.Nil {
				return nil, errorxs.ErrWrongProcessSequence
			}
			return nil, err
		}

		if err := l.svcCtx.Redis.Del(l.ctx, bindEmailKey).Err(); err != nil {
			return nil, err
		}

		if err = userModel.UpdateByEmail(l.svcCtx.Redis, l.svcCtx.DB, &userModel.User{
			Model: gorm.Model{ID: uint(in.Uid)},
			Email: wantChangedEmail,
		}, &userModel.User{
			Email: in.Email,
		}); err != nil {
			return nil, err
		}

		// 这里后端更换了邮箱后，只需要告诉其验证操作是否通过，后续要做什么事由前端自行决定
		// 这里返回的ttl是因为第一次邮箱验证需要，所以此次返回的ttl无意义
		return &__.EmailChangeVerifyResp{Ttl: 0}, nil
	} else {
		return nil, errors.New("无法识别是第几次邮箱验证")
	}
}
