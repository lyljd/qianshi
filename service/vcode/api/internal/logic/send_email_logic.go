package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"qianshi/common/response/errorx"
	"qianshi/service/vcode/api/internal/svc"
	"qianshi/service/vcode/api/internal/svc/randx"
	"qianshi/service/vcode/api/internal/types"
	"regexp"
	"strconv"
	"time"
)

type SendEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailLogic) SendEmail(req *types.SendEmailReq) (resp *types.SendEmailResp, err *errorx.Error) {
	captchaVerifyKey := "captcha:verify:" + req.Cid
	find := l.svcCtx.RedisCli.Exists(context.Background(), captchaVerifyKey).Val()
	if find == 0 {
		return nil, errorx.NewDefault("验证码不存在或已过期")
	}
	l.svcCtx.RedisCli.Del(context.Background(), captchaVerifyKey)

	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	ok := re.MatchString(req.Email)
	if !ok {
		return nil, errorx.NewDefault("请输入正确的邮箱")
	}

	vcodeKey := "vcode:" + req.Email
	restTtl := int(l.svcCtx.RedisCli.TTL(context.Background(), vcodeKey).Val().Seconds())
	restCd := restTtl - (l.svcCtx.VcodeTTL - l.svcCtx.VcodeCD)
	if restCd > 0 {
		return &types.SendEmailResp{CD: restCd}, errorx.NewDefault("请在" + strconv.Itoa(restCd) + "秒后再试")
	}
	vcode := randx.New()
	l.svcCtx.RedisCli.SetEX(context.Background(), vcodeKey, vcode, time.Second*time.Duration(l.svcCtx.VcodeTTL))

	content := fmt.Sprintf(l.svcCtx.EmailContent, vcode)
	er := l.svcCtx.EmailCli.SendOne(req.Email, l.svcCtx.EmailTitle, content)
	if er != nil {
		return nil, errorx.NewDefault(er.Error())
	}
	return nil, nil
}
