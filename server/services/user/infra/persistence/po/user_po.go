package po

import (
	"github.com/google/uuid"
	"github.com/xince-fun/InstaGo/server/shared/utils"
)

type UserAccount struct {
	ID          int64           `gorm:"column:id;primaryKey" json:"id"`
	UserID      uuid.UUID       `gorm:"column:user_id" json:"user_id"`
	Email       string          `gorm:"column:email;default:null" json:"email"`
	PhoneNumber string          `gorm:"column:phone_number;default:null" json:"phone_number"`
	Passwd      string          `gorm:"column:passwd" json:"passwd"`
	IsDeleted   int8            `gorm:"column:is_deleted" json:"is_deleted"`
	CreateTime  utils.LocalTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime  utils.LocalTime `gorm:"column:update_time" json:"update_time"`
}

func (UserAccount) TableName() string {
	return "user_account"
}

type UserProfile struct {
	ID         int64           `gorm:"column:id;primaryKey" json:"id"`
	UserID     uuid.UUID       `gorm:"column:user_id" json:"user_id"`
	Bio        string          `gorm:"column:bio;default:null" json:"bio"`
	BirthDay   utils.LocalDate `gorm:"column:birth_day" json:"birth_day"`
	Gender     int             `gorm:"column:gender" json:"gender"`
	IsDeleted  int8            `gorm:"column:is_deleted" json:"is_deleted"`
	CreateTime utils.LocalTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime utils.LocalTime `gorm:"column:update_time" json:"update_time"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}

type UserInfo struct {
	ID         int64           `gorm:"column:id;primaryKey" json:"id"`
	UserID     uuid.UUID       `gorm:"column:user_id" json:"user_id"`
	Account    string          `gorm:"account" json:"account"`
	FullName   string          `gorm:"full_name" json:"full_name"`
	AvatarID   string          `gorm:"avatar_id;default:null" json:"avatar_id"`
	IsDeleted  int8            `gorm:"column:is_deleted" json:"is_deleted"`
	CreateTime utils.LocalTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime utils.LocalTime `gorm:"column:update_time" json:"update_time"`
}

func (UserInfo) TableName() string {
	return "user_info"
}
