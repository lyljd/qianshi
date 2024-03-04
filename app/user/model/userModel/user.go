package userModel

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string
	Password     string
	Nickname     string
	AvatarUrl    string
	Signature    string
	Power        int
	VipExpire    int
	Ip           string
	IpLocation   string
	Exp          int
	Level        int `gorm:"default:1"`
	Coin         int
	FollowNum    int
	RefreshToken string
}
