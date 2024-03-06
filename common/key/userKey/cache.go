package userKey

import (
	"qianshi/common/generic"
	"strconv"
)

func GetCacheUserById[T generic.ID](id T) string {
	return "cache:user:id:" + strconv.Itoa(int(id))
}

func GetCacheUserByEmail(email string) string {
	return "cache:user:email:" + email
}

func GetCacheUserHomeById[T generic.ID](id T) string {
	return "cache:user_home:id:" + strconv.Itoa(int(id))
}

func GetCacheUserInteractionById[T generic.ID](id T) string {
	return "cache:user_interaction:id:" + strconv.Itoa(int(id))
}
