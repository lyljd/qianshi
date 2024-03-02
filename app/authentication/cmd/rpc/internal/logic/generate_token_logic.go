package logic

import (
	"context"
	"encoding/json"
	"qianshi/common/tool"
	"time"

	"qianshi/app/authentication/cmd/rpc/internal/svc"
	"qianshi/app/authentication/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type Token struct {
	Iat int64 `json:"iat"`
	Exp int64 `json:"exp"`
	Uid uint  `json:"uid"`
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *__.GenerateTokenReq) (*__.GenerateTokenResp, error) {
	token, err := generate(uint(in.Uid), l.svcCtx.Config.TokenMinutes, l.svcCtx.Config.TokenSecret)
	if err != nil {
		return nil, err
	}

	return &__.GenerateTokenResp{Token: token}, nil
}

func generate(uid uint, minutes int, secret string) (string, error) {
	t := &Token{
		Iat: time.Now().Unix(),
		Exp: time.Now().Add(time.Minute * time.Duration(minutes)).Unix(),
		Uid: uid,
	}

	jt, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	payload := tool.Base64Encode(jt)
	signature, err := tool.Sha256(payload, secret, "")
	if err != nil {
		return "", err
	}

	return payload + "." + signature, nil
}
