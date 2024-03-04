package userKey

import (
	"qianshi/common/generic"
	"strconv"
)

func GetCacheUserById[T generic.ID](id T) string {
	return "cache:user:" + strconv.Itoa(int(id))
}

func GetCacheUserByEmail(email string) string {
	return "cache:user:" + email
}

func GetCacheUserHomeById[T generic.ID](id T) string {
	return "cache:user_home:" + strconv.Itoa(int(id))
}

func GetCacheUserInteractionById[T generic.ID](id T) string {
	return "cache:user_interaction:" + strconv.Itoa(int(id))
}
