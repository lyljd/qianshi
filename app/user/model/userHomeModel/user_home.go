package userHomeModel

import "gorm.io/gorm"

type UserHome struct {
	gorm.Model

	Gender   string `gorm:"default:保密"`
	Birthday string
	Tags     string
	TopImgNo int `gorm:"default:1"`
	Title    string
	Notice   string

	PostNum       int
	CollectionNum int
	FavlistNum    int
	FanNum        int
	LikeNum       int
	PlayNum       int
	ReadNum       int
}
