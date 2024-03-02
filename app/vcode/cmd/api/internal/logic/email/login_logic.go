package email

import (
	"context"
	"qianshi/app/vcode/cmd/api/internal/svc"
	"qianshi/app/vcode/cmd/api/internal/types"
	"qianshi/common/key"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *LoginLogic) Login(req *types.EmailReq) (resp *types.EmailResp, err error) {
	verifyKey := key.GetVcodeLoginVerify(req.Email)

	subject := "【浅时】验证码"
	contentTextTmpl := "您正在进行【登录/注册】操作，验证码为：%s，5分钟内有效，请勿告诉他人。"

	return common(l.ctx, l.svcCtx, req.CaptchaId, verifyKey, req.Email, subject, contentTextTmpl, 6)
}
