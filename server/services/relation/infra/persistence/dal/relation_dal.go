package dal

import (
	"context"
	"errors"
	"github.com/xince-fun/InstaGo/server/services/relation/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/relation/infra/persistence/converter"
	"github.com/xince-fun/InstaGo/server/services/relation/infra/persistence/po"
	"github.com/xince-fun/InstaGo/server/shared/utils"
	"gorm.io/gorm"
	"time"
)

func NewRelationDal(db *gorm.DB) *RelationDal {
	return &RelationDal{
		db: db,
	}
}

type RelationDal struct {
	db *gorm.DB
}

func (d *RelationDal) UpsertRelation(ctx context.Context, relation *entity.Relation, tx *gorm.DB) error {
	relationPo := converter.RelationToPo(relation)

	if tx == nil {
		tx = d.db
	}
	now := utils.LocalTime(time.Now())
	relationPo.UpdateTime = now
	relationPo.CreateTime = now

	res := tx.WithContext(ctx).Save(relationPo)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (d *RelationDal) SelectRelation(ctx context.Context, followerID, followeeID string, tx *gorm.DB) (*po.Relation, error) {
	relationPo := po.Relation{}

	if tx == nil {
		tx = d.db
	}

	res := tx.WithContext(ctx).Where(map[string]interface{}{"follower_id": followerID, "followee_id": followeeID, "is_deleted": 0}).First(&relationPo)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &relationPo, nil
}

func (d *RelationDal) CountRelation(ctx context.Context, followerID, followeeID string, tx *gorm.DB) (int64, error) {
	if tx == nil {
		tx = d.db
	}

	var count int64
	res := tx.WithContext(ctx).Model(&po.Relation{}).Where(map[string]interface{}{"follower_id": followerID, "followee_id": followeeID, "is_deleted": 0}).Count(&count)
	if res.Error != nil {
		return 0, res.Error
	}
	return count, nil
}

func (d *RelationDal) DeleteRelation(ctx context.Context, relation *po.Relation, tx *gorm.DB) error {
	if tx == nil {
		tx = d.db
	}

	relation.UpdateTime = utils.LocalTime(time.Now())
	relation.IsDeleted = 1

	return tx.WithContext(ctx).Model(relation).Updates(map[string]interface{}{"is_deleted": relation.IsDeleted, "update_time": relation.UpdateTime}).Error
}

func (d *RelationDal) BeginTransaction(ctx context.Context) *gorm.DB {
	return d.db.WithContext(ctx).Begin()
}
