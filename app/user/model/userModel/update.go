package userModel

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"qianshi/common/key/userKey"
)

// UpdateById 修改id为obj.ID的记录，只修改obj中值不为"空值"的记录
func UpdateById(rc *redis.Client, db *gorm.DB, md *User, obj *User) error {
	if err := db.Model(md).Updates(obj).Error; err != nil {
		return err
	}

	return rc.Del(context.Background(), userKey.GetCacheUserById(md.ID), userKey.GetCacheUserByEmail(md.Email)).Err()
}
