package logic

import (
	"context"
	__2 "qianshi/app/authentication/cmd/rpc/pb"
	"qianshi/app/user/model/userModel"
	"qianshi/common/errorxs"
	"qianshi/common/key"
	"qianshi/common/result/errorx"
	"qianshi/common/tool"
	"time"

	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PassChangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPassChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PassChangeLogic {
	return &PassChangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PassChangeLogic) PassChange(in *__.PassChangeReq) (*__.PassChangeResp, error) {
	u, err := userModel.QueryById(l.svcCtx.Redis, l.svcCtx.DB, in.Uid)
	if err != nil {
		return nil, err
	}

	verifyKey := key.GetUserChangePasswordVerify(u.Email)
	existRes, existErr := l.svcCtx.Redis.Exists(l.ctx, verifyKey).Result()
	if existErr != nil {
		return nil, existErr
	}
	if existRes != 1 {
		return nil, errorxs.ErrChangePassVerifyFail
	}

	// 加密密码
	pass, err := tool.Sha256(in.Pass, l.svcCtx.Config.PassSecret, "")
	if err != nil {
		return nil, err
	}

	// 如果新密码和旧密码相同，则不清除已验证key（还有机会重新传密码）
	if u.Password == pass {
		return nil, errorxs.ErrOldPassSameAsNewPass
	}

	// 清除已验证key
	if err := l.svcCtx.Redis.Del(l.ctx, verifyKey).Err(); err != nil {
		return nil, errorx.New(errorx.CodeParamError, err)
	}

	// 生成token
	gtResp, err := l.svcCtx.AuthenticationRpc.GenerateToken(l.ctx, &__2.GenerateTokenReq{Uid: int64(u.ID)})
	if err != nil {
		return nil, err
	}

	// 生成refreshToken
	grtResp, err := l.svcCtx.AuthenticationRpc.GenerateRefreshToken(l.ctx, &__2.GenerateRefreshTokenReq{Uid: int64(u.ID)})
	if err != nil {
		return nil, err
	}

	// 更新密码和refreshToken
	if err := userModel.UpdateById(l.svcCtx.Redis, l.svcCtx.DB, u, &userModel.User{
		Password:     pass,
		RefreshToken: grtResp.Token,
	}); err != nil {
		return nil, err
	}

	// 让之前颁发的所有token均失效
	teKey := key.GetTokenExp(uint(in.Uid))
	if err := l.svcCtx.Redis.SetEX(l.ctx, teKey, time.Now().Unix(), time.Minute*60).Err(); err != nil {
		return nil, err
	}

	return &__.PassChangeResp{
		Token:        gtResp.Token,
		RefreshToken: grtResp.Token,
	}, nil
}
