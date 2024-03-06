package userModel

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"qianshi/common/key/userKey"
)

// UpdateById 修改id为md.ID的记录，只修改obj中值不为"空值"的记录
func UpdateById(rc *redis.Client, db *gorm.DB, md *User, obj *User) error {
	if md.ID == 0 || md.Email == "" {
		return errors.New("在userModel.UpdateById时md需要传入id和email；除了查询时需要，删除缓存时也会用到")
	}

	if err := db.Model(md).Updates(obj).Error; err != nil {
		return err
	}

	return rc.Del(context.Background(), userKey.GetCacheUserById(md.ID), userKey.GetCacheUserByEmail(md.Email)).Err()
}

// UpdateByIdWithNil 修改id为md.ID的记录，如果m中某个key的value为"空值"，该key对应的字段在数据库中依然会被修改
func UpdateByIdWithNil(rc *redis.Client, db *gorm.DB, md *User, m map[string]any) error {
	if md.ID == 0 || md.Email == "" {
		return errors.New("在userModel.UpdateByIdWithNil时md需要传入id和email；除了查询时需要，删除缓存时也会用到")
	}

	if err := db.Model(md).Updates(m).Error; err != nil {
		return err
	}

	return rc.Del(context.Background(), userKey.GetCacheUserById(md.ID), userKey.GetCacheUserByEmail(md.Email)).Err()
}

// UpdateByEmail 修改email为md.Email的记录，只修改obj中值不为"空值"的记录
func UpdateByEmail(rc *redis.Client, db *gorm.DB, md *User, obj *User) error {
	if md.ID == 0 || md.Email == "" {
		return errors.New("在userModel.UpdateByEmail时md需要传入id和email；除了查询时需要，删除缓存时也会用到")
	}

	if err := db.Model(md).Where("email = ?", md.Email).Updates(obj).Error; err != nil {
		return err
	}

	return rc.Del(context.Background(), userKey.GetCacheUserById(md.ID), userKey.GetCacheUserByEmail(md.Email)).Err()
}
