package email

import (
	"context"
	"errors"
	"fmt"
	"qianshi/app/vcode/cmd/api/internal/svc"
	"qianshi/app/vcode/cmd/api/internal/types"
	mail "qianshi/common/email"
	"qianshi/common/key"
	"qianshi/common/result/errorx"
	"qianshi/common/tool"
	"time"
)

const (
	vcodeExpire = time.Minute * 5
	vcodeCD     = time.Minute * 3
)

func common(ctx context.Context, svcCtx *svc.ServiceContext, captchaId, verifyKey, email, subject, contentTextTmpl string, vcodeLen int) (*types.EmailResp, error) {
	cvKey := key.GetCaptchaVerify(captchaId)
	existRes, existErr := svcCtx.Redis.Exists(ctx, cvKey).Result()
	if existErr != nil {
		return nil, errorx.New(errorx.CodeServerError, existErr)
	}
	if existRes != 1 {
		return nil, errorx.New(errorx.CodeParamError, errors.New("captcha_id不存在"), "人机验证未通过")
	}

	if err := svcCtx.Redis.Del(ctx, cvKey).Err(); err != nil {
		return nil, errorx.New(errorx.CodeParamError, err)
	}

	ttlRes, ttlErr := svcCtx.Redis.TTL(ctx, verifyKey).Result()
	if ttlErr != nil {
		return nil, errorx.New(errorx.CodeServerError, ttlErr)
	}
	gap := vcodeExpire - vcodeCD
	ttlRes -= gap
	if ttlRes > 0 {
		return &types.EmailResp{Ttl: int(ttlRes.Seconds())}, nil
	}

	vcode := tool.RandStr(vcodeLen)
	content := mail.Text(fmt.Sprintf(contentTextTmpl, vcode))
	if err := svcCtx.Email.SendToOne(email, subject, content); err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	if err := svcCtx.Redis.SetEX(ctx, verifyKey, vcode, vcodeExpire).Err(); err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	return &types.EmailResp{Ttl: -2}, nil
}
