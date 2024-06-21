package po

import (
	"github.com/xince-fun/InstaGo/server/shared/utils"
)

type Blob struct {
	ID         int64           `gorm:"column:id;primaryKey" json:"id"`
	UserID     string          `gorm:"column:user_id" json:"user_id"`
	BlobID     string          `gorm:"column:blob_id" json:"blobID"`
	URL        string          `gorm:"column:url" json:"url"`
	BlobType   int8            `gorm:"column:blob_type" json:"blob_type"`
	IsDeleted  int8            `gorm:"column:is_deleted" json:"is_deleted"`
	CreateTime utils.LocalTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime utils.LocalTime `gorm:"column:update_time" json:"update_time"`
}

func (Blob) TableName() string {
	return "blob"
}
