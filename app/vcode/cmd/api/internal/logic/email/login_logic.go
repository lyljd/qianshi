package email

import (
	"context"
	"errors"
	"fmt"
	"qianshi/common/email"
	"qianshi/common/key"
	"qianshi/common/result/errorx"
	"qianshi/common/tool"
	"time"

	"qianshi/app/vcode/cmd/api/internal/svc"
	"qianshi/app/vcode/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	vcodeExpire = time.Minute * 5
	vcodeCD     = time.Minute * 3
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	cvKey := key.GetCaptchaVerify(req.CaptchaId)
	existRes, existErr := l.svcCtx.Redis.Exists(l.ctx, cvKey).Result()
	if existErr != nil {
		return nil, errorx.New(errorx.CodeServerError, existErr)
	}
	if existRes != 1 {
		return nil, errorx.New(errorx.CodeParamError, errors.New("captcha_id不存在"), "人机验证未通过")
	}

	if err := l.svcCtx.Redis.Del(l.ctx, cvKey).Err(); err != nil {
		return nil, errorx.New(errorx.CodeParamError, err)
	}

	vlvKey := key.GetVcodeLoginVerify(req.Email)
	ttlRes, ttlErr := l.svcCtx.Redis.TTL(l.ctx, vlvKey).Result()
	if ttlErr != nil {
		return nil, errorx.New(errorx.CodeServerError, ttlErr)
	}
	gap := vcodeExpire - vcodeCD
	ttlRes -= gap
	if ttlRes > 0 {
		return &types.LoginResp{Ttl: int(ttlRes.Seconds())}, nil
	}

	vcode := tool.RandNumStr(6)
	subject := "【浅时】验证码"
	content := email.Text(fmt.Sprintf("您正在进行登录/注册操作，验证码为：%s，5分钟内有效，请勿告诉他人。", vcode))
	if err := l.svcCtx.Email.SendToOne(req.Email, subject, content); err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	if err := l.svcCtx.Redis.SetEX(l.ctx, vlvKey, vcode, vcodeExpire).Err(); err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	resp = &types.LoginResp{Ttl: -2}

	return
}
