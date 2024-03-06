package image

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/dchest/captcha"
	"qianshi/common/result/errorx"

	"qianshi/app/captcha/cmd/api/internal/svc"
	"qianshi/app/captcha/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.GetReq) (resp *types.GetResp, err error) {
	var buf bytes.Buffer
	// 图片设置过小会让数字重叠在一起，导致看不清
	if err := captcha.WriteImage(&buf, req.Id, 300, 90); err != nil {
		return nil, errorx.New(errorx.CodeServerError, err)
	}

	image := base64.StdEncoding.EncodeToString(buf.Bytes())
	return &types.GetResp{Image: "data:image/png;base64," + image}, nil
}
