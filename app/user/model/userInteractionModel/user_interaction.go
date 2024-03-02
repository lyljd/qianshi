package userInteractionModel

import "gorm.io/gorm"

type UserInteraction struct {
	gorm.Model

	PlayNum             int
	PlayNumIncr         int
	VideoCommentNum     int
	VideoCommentNumIncr int
	DanmuNum            int
	DanmuNumIncr        int
	VideoLikeNum        int
	VideoLikeNumIncr    int
	CoinNum             int
	CoinNumIncr         int
	VideoStarNum        int
	VideoStarNumIncr    int
	VideoShareNum       int
	VideoShareNumIncr   int

	ReadNum            int
	ReadNumIncr        int
	ReadCommentNum     int
	ReadCommentNumIncr int
	ReadLikeNum        int
	ReadLikeNumIncr    int
	ReadStarNum        int
	ReadStarNumIncr    int
	ReadShareNum       int
	ReadShareNumIncr   int
}
