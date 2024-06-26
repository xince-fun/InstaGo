package po

import "github.com/xince-fun/InstaGo/server/shared/utils"

type Relation struct {
	ID         int64           `gorm:"column:id;primaryKey" json:"id"`
	FollowerID string          `gorm:"column:follower_id" json:"follower_id"`
	FolloweeID string          `gorm:"column:followee_id" json:"followee_id"`
	IsDeleted  int8            `gorm:"column:is_deleted" json:"is_deleted"`
	CreateTime utils.LocalTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime utils.LocalTime `gorm:"column:update_time" json:"update_time"`
}
