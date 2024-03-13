package logic

import (
	"context"
	"fmt"
	"qianshi/common/tool"
	"time"

	"qianshi/app/authentication/cmd/rpc/internal/svc"
	"qianshi/app/authentication/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignatureCdnGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignatureCdnGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignatureCdnGetLogic {
	return &SignatureCdnGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SignatureCdnGet 没有登录的用户uid就传0
func (l *SignatureCdnGetLogic) SignatureCdnGet(in *__.SignatureCdnGetReq) (*__.SignatureCdnGetResp, error) {
	iat := time.Now().Unix()
	requestId := "0" // 阿里云cdn鉴权方法A要求参数"rand"（推荐是使用不带"-"的uuid，该参数可以用于记录所有请求签名的记录，这里用不上就待定了）

	sigKey := fmt.Sprintf("%s-%d-%s-%d-%s", in.FilePath, iat, requestId, in.Uid, l.svcCtx.Config.SigCdnGetSecret)
	md5Hash := tool.MD5(sigKey)
	url := fmt.Sprintf("%s%s?auth_key=%d-%s-%d-%s", l.svcCtx.Config.CdnUrl, in.FilePath, iat, requestId, in.Uid, md5Hash)

	return &__.SignatureCdnGetResp{Url: url}, nil
}
