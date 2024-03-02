package email

import (
	"context"
	"qianshi/common/key"

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

func (l *ChangePasswordLogic) ChangePassword(req *types.EmailReq) (resp *types.EmailResp, err error) {
	verifyKey := key.GetVcodeChangePasswordVerify(req.Email)

	subject := "【浅时】验证码"
	contentTextTmpl := "您正在进行【修改密码】操作，验证码为：%s，5分钟内有效，请勿告诉他人。"

	return common(l.ctx, l.svcCtx, req.CaptchaId, verifyKey, req.Email, subject, contentTextTmpl, 6)
}
