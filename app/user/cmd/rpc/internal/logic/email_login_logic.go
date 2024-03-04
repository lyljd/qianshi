package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"io"
	"math/rand"
	"net/http"
	__2 "qianshi/app/authentication/cmd/rpc/pb"
	"qianshi/app/user/model/userHomeModel"
	"qianshi/app/user/model/userModel"
	"qianshi/common/errorxs"
	"qianshi/common/key"
	"qianshi/common/tool"
	"strings"

	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailLoginLogic {
	return &EmailLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmailLoginLogic) EmailLogin(in *__.EmailLoginReq) (*__.LoginResp, error) {
	loginVerifyKey := key.GetVcodeLoginVerify(in.Email)
	vcode, err := l.svcCtx.Redis.Get(l.ctx, loginVerifyKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, errorxs.ErrKeyNotFound
		}
		return nil, err
	}
	if vcode != in.Code {
		return nil, errorxs.ErrVcodeWrong
	}

	if err = l.svcCtx.Redis.Del(l.ctx, loginVerifyKey).Err(); err != nil {
		return nil, err
	}

	// 新用户自动注册账号
	u, err := userModel.QueryByEmail(l.svcCtx.Redis, l.svcCtx.DB, in.Email)
	if err != nil {
		if err != errorxs.ErrRecordNotFound {
			return nil, err
		}

		if nu, err := register(l.svcCtx.DB, in.Email); err != nil {
			return nil, err
		} else {
			u = nu
		}
	}

	resp, err := loginCommon(l.ctx, l.svcCtx, u, in.Ip)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func loginCommon(ctx context.Context, svcCtx *svc.ServiceContext, u *userModel.User, ip string) (*__.LoginResp, error) {
	// 生成token
	gtResp, err := svcCtx.AuthenticationRpc.GenerateToken(ctx, &__2.GenerateTokenReq{Uid: int64(u.ID)})
	if err != nil {
		return nil, err
	}

	// 若没有refreshToken，或refreshToken有问题（已过期），则需要重新生成refreshToken
	rft := u.RefreshToken
	if rft != "" {
		if _, err = svcCtx.AuthenticationRpc.VerifyRefreshToken(ctx, &__2.VerifyRefreshTokenReq{Token: rft}); err != nil {
			rft = ""
		}
	}

	if rft == "" {
		grtResp, err := svcCtx.AuthenticationRpc.GenerateRefreshToken(ctx, &__2.GenerateRefreshTokenReq{Uid: int64(u.ID)})
		if err != nil {
			return nil, err
		}
		rft = grtResp.Token
	}

	// 更新refreshToken，ip和ipLocation
	if err := userModel.UpdateById(svcCtx.Redis, svcCtx.DB, u, &userModel.User{
		Ip:           ip,
		IpLocation:   queryIpLocation(ip),
		RefreshToken: rft,
	}); err != nil {
		return nil, err
	}

	// TODO 查消息数和动态数
	newMessageNum := 0
	newDynamicNum := 0

	return &__.LoginResp{
		Token:         gtResp.Token,
		RefreshToken:  rft,
		Nickname:      u.Nickname,
		AvatarUrl:     u.AvatarUrl,
		NewMessageNum: int64(newMessageNum),
		NewDynamicNum: int64(newDynamicNum),
	}, nil
}

type respData struct {
	Address string `json:"address"`
}
type queryResp struct {
	Code int      `json:"code"`
	Data respData `json:"data"`
}

func queryIpLocation(ip string) string {
	resp, err := http.Get("https://searchplugin.csdn.net/api/v1/ip/get?ip=" + ip)
	if err == nil {
		if body, err := io.ReadAll(resp.Body); err == nil {
			var qr queryResp
			if err := json.Unmarshal(body, &qr); err == nil {
				if qr.Code == 200 {
					adds := strings.Split(qr.Data.Address, " ")
					if adds[0] != "中国" {
						return adds[0]
					} else if adds[1] == "台湾" || adds[1] == "香港" || adds[1] == "澳门" {
						return "中国" + adds[1]
					} else {
						return adds[1]
					}
				}
			}
		}
	}
	return "未知"
}

func register(db *gorm.DB, email string) (*userModel.User, error) {
	u := userModel.User{
		Email:     email,
		Nickname:  generateNickname(email),
		AvatarUrl: generateAvatar(),
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&u).Error; err != nil {
			return err
		}

		uh := userHomeModel.UserHome{
			Model: u.Model,
		}
		if err := tx.Create(&uh).Error; err != nil {
			return err
		}

		tx.Commit()
		return nil
	}); err != nil {
		return nil, err
	}

	return &u, nil
}

func generateNickname(email string) string {
	idx := strings.Index(email, "@")
	if idx > 15 {
		idx = 15
	}
	return email[:idx] + "_" + tool.RandNumStr(4)
}

func generateAvatar() string {
	// TODO 到时候头像会从本地转到cos上，记得改
	return fmt.Sprintf("/init-avatar/%d.jpeg", rand.Intn(7)+1)
}
