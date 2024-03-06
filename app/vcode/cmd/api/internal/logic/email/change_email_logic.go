package email

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	__ "qianshi/app/user/cmd/rpc/pb"
	"qianshi/app/vcode/cmd/api/internal/svc"
	"qianshi/app/vcode/cmd/api/internal/types"
	"qianshi/common/ctx"
	"qianshi/common/errorxs"
	"qianshi/common/key"
	"qianshi/common/result/errorx"
)

type ChangeEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeEmailLogic {
	return &ChangeEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeEmailLogic) ChangeEmail(req *types.ChangeEmailReq) (resp *types.EmailResp, err error) {
	// 请求时如果上传的email为空则为【更换邮箱】，不为空则为【绑定邮箱】
	email, option := req.Email, "绑定邮箱"

	if email == "" { // 更换邮箱
		u, err := l.svcCtx.UserRpc.UserQuery(l.ctx, &__.QueryReq{Uid: uint64(ctx.GetUid(l.ctx))})
		if err != nil {
			return nil, err
		}

		email, option = u.Email, "更换邮箱"

		// 如果在短时内验证过，则不发送验证码，直接跳过邮箱验证
		hasVerifyKey := key.GetUserChangeEmailVerify(email)
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
	} else { // 绑定邮箱
		// 判断要绑定的邮箱是否已经被绑定了
		if _, err := l.svcCtx.UserRpc.UserQuery(l.ctx, &__.QueryReq{Email: email}); err == nil {
			err := errors.New("要绑定的邮箱已经被绑定了")
			return nil, errorx.New(errorx.CodeParamError, err, err.Error())
		} else if !errorxs.Is(err, errorxs.ErrKeyNotFound) && !errorxs.Is(err, errorxs.ErrRecordNotFound) {
			return nil, errorx.New(errorx.CodeServerError, errors.New("查询要绑定的邮箱是否已经被绑定了失败！err: "+err.Error()))
		}
	}

	verifyKey := key.GetVcodeChangeEmailVerify(email)

	subject := "【浅时】验证码"
	contentTextTmpl := "您正在进行【" + option + "】操作，验证码为：%s，5分钟内有效，请勿告诉他人。"

	return common(l.ctx, l.svcCtx, req.CaptchaId, verifyKey, email, subject, contentTextTmpl, 6)
}
