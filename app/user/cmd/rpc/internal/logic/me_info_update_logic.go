package logic

import (
	"context"
	"gorm.io/gorm"
	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"
	"qianshi/app/user/model/userHomeModel"
	"qianshi/app/user/model/userModel"
	"qianshi/common/errorxs"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeInfoUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMeInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeInfoUpdateLogic {
	return &MeInfoUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MeInfoUpdateLogic) MeInfoUpdate(in *__.MeInfoUpdateReq) (*__.MeInfoUpdateResp, error) {
	queryIdResp, err := userModel.QueryById(l.svcCtx.Redis, l.svcCtx.DB, in.Id)
	if err != nil {
		return nil, err
	}

	newCoin := queryIdResp.Coin

	if in.Nickname != queryIdResp.Nickname {
		if newCoin < 5 {
			return nil, errorxs.ErrCoinInsufficient
		}

		newCoin -= 5

		_, err = userModel.QueryByNickname(l.svcCtx.DB, in.Nickname)
		if !errorxs.Is(err, errorxs.ErrRecordNotFound) {
			return nil, errorxs.ErrNicknameHasExist
		}
	}

	if err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := userModel.UpdateByIdWithNil(l.svcCtx.Redis, tx, &userModel.User{
			Model: gorm.Model{ID: uint(in.Id)},
			Email: queryIdResp.Email,
		}, map[string]any{
			"nickname":  in.Nickname,
			"signature": in.Signature,
			"coin":      newCoin,
		}); err != nil {
			return err
		}

		if err := userHomeModel.UpdateByIdWithNil(l.svcCtx.Redis, tx, &userHomeModel.UserHome{
			Model: gorm.Model{ID: uint(in.Id)},
		}, map[string]any{
			"gender":   in.Gender,
			"birthday": in.Birthday,
			"tags":     strings.Join(in.Tags, ";"),
		},
		); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &__.MeInfoUpdateResp{}, nil
}
