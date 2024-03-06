package email

import (
	"context"
	"errors"
	__ "qianshi/app/user/cmd/rpc/pb"
	"qianshi/common/ctx"
	"qianshi/common/key"
	"qianshi/common/result/errorx"

	"qianshi/app/vcode/cmd/api/internal/svc"
	"qianshi/app/vcode/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) (resp *types.EmailResp, err error) {
	u, err := l.svcCtx.UserRpc.UserQuery(l.ctx, &__.QueryReq{Uid: uint64(ctx.GetUid(l.ctx))})
	if err != nil {
		return nil, err
	}

	// 如果在短时内验证过，则不发送验证码，直接跳过邮箱验证
	hasVerifyKey := key.GetUserChangePasswordVerify(u.Email)
	existRes, existErr := l.svcCtx.Redis.Exists(l.ctx, hasVerifyKey).Result()
	if existErr != nil {
		return nil, errorx.New(errorx.CodeServerError, existErr)
	}
	if existRes == 1 {
		ttlRes, ttlErr := l.svcCtx.Redis.TTL(l.ctx, hasVerifyKey).Result()
		if ttlErr != nil {
			return nil, errorx.New(errorx.CodeServerError, ttlErr)
		}

		return &types.EmailResp{Ttl: int(ttlRes.Seconds())}, errorx.New(errorx.CodeSuccess, errors.New("特殊的请求成功；如果在短时内验证过，则不发送验证码，直接跳过邮箱验证"))
	}

	verifyKey := key.GetVcodeChangePasswordVerify(u.Email)

	subject := "【浅时】验证码"
	contentTextTmpl := "您正在进行【修改密码】操作，验证码为：%s，5分钟内有效，请勿告诉他人。"

	return common(l.ctx, l.svcCtx, req.CaptchaId, verifyKey, u.Email, subject, contentTextTmpl, 6)
}
