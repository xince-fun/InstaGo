package dal

import (
	"context"
	"errors"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/persistence/converter"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/persistence/po"
	"github.com/xince-fun/InstaGo/server/shared/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type BlobDal struct {
	db *gorm.DB
}

func NewBlobDal(db *gorm.DB) *BlobDal {
	return &BlobDal{db: db}
}

func (d *BlobDal) Upsert(ctx context.Context, blob *entity.Blob, tx *gorm.DB) (err error) {
	blobPo := converter.BlobToPo(blob)

	now := utils.LocalTime(time.Now())
	blobPo.UpdateTime = now
	blobPo.CreateTime = now

	if tx == nil {
		tx = d.db
	}

	err = tx.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "blob_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"blob_type", "blob_url", "is_deleted", "update_time"}),
	}).Create(blobPo).Error

	return err
}

func (d *BlobDal) SelectBlobByID(ctx context.Context, blobID string, tx *gorm.DB) (*entity.Blob, error) {
	blobPo := po.Blob{}

	if tx == nil {
		tx = d.db
	}

	res := tx.WithContext(ctx).Where(map[string]interface{}{"blob_id": blobID, "is_deleted": 0}).First(&blobPo)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	blob := converter.BlobToEntity(&blobPo)
	return blob, nil
}
