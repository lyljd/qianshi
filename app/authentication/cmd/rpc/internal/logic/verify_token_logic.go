package logic

import (
	"context"
	"encoding/json"
	"errors"
	"qianshi/common/key"
	"qianshi/common/tool"
	"strconv"
	"strings"
	"time"

	"qianshi/app/authentication/cmd/rpc/internal/svc"
	"qianshi/app/authentication/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyTokenLogic) VerifyToken(in *__.VerifyTokenReq) (*__.VerifyTokenResp, error) {
	uid, iat, err := verify(in.Token, l.svcCtx.Config.TokenSecret)
	if err != nil {
		return nil, err
	}

	expStr, err := l.svcCtx.Redis.Get(l.ctx, key.GetTokenExp(uid)).Result()
	if err != nil {
		return nil, err
	}

	exp, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return nil, err
	}

	// key不存在时exp为0，没有问题
	if iat < exp {
		return nil, errors.New("token has expired")
	}

	return &__.VerifyTokenResp{Uid: int64(uid)}, nil
}

func verify(token string, secret string) (uint, int64, error) {
	ts := strings.Split(token, ".")
	if len(ts) != 2 {
		return 0, 0, errors.New("token segment number is wrong")
	}

	ps, err := tool.Sha256(ts[0], secret, "")
	if err != nil {
		return 0, 0, err
	}

	if ps != ts[1] {
		return 0, 0, errors.New("token payload modified")
	}

	mp, err := tool.Base64Decode(ts[0])
	if err != nil {
		return 0, 0, err
	}

	var t Token
	if err := json.Unmarshal(mp, &t); err != nil {
		return 0, 0, err
	}

	if t.Exp < time.Now().Unix() {
		return 0, 0, errors.New("token has expired")
	}

	return t.Uid, t.Iat, nil
}
