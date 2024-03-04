package userInteractionModel

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"math/rand"
	"qianshi/common/errorxs"
	"qianshi/common/generic"
	"qianshi/common/key/userKey"
	"qianshi/common/xlog"
	"time"
)

func QueryById[T generic.ID](rc *redis.Client, db *gorm.DB, id T) (*UserInteraction, error) {
	// 查缓存
	var obj UserInteraction
	cacheKey := userKey.GetCacheUserInteractionById(id)

	getRes, err := rc.Get(context.Background(), cacheKey).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err != redis.Nil {
		if getRes == "" {
			return nil, errorxs.ErrRecordNotFound
		}

		if err = json.Unmarshal([]byte(getRes), &obj); err != nil {
			return nil, err
		}

		return &obj, nil
	}

	// 查数据库
	if err = db.Take(&obj, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 写空缓存（1分钟）
			if err := rc.Set(context.Background(), cacheKey, "", time.Minute).Err(); err != nil {
				xlog.Warn("空缓存写入失败，数据库有缓存穿透风险！" + err.Error())
			}
			return nil, errorxs.ErrRecordNotFound
		}
		return nil, err
	}

	// 写缓存（24~72小时）
	jm, _ := json.Marshal(obj)

	if err := rc.Set(context.Background(), cacheKey, jm, time.Hour*time.Duration(rand.Intn(49)+24)).Err(); err != nil {
		xlog.Warn("缓存写入失败，数据库有缓存穿透风险！" + err.Error())
		return &obj, nil
	}

	return &obj, nil
}
