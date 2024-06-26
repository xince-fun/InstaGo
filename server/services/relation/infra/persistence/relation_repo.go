package persistence

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/xince-fun/InstaGo/server/services/relation/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/relation/infra/persistence/converter"
	"github.com/xince-fun/InstaGo/server/services/relation/infra/persistence/dal"
	"github.com/xince-fun/InstaGo/server/shared/errno"
)

func NewRelationRepo(relationDal *dal.RelationDal) *RelationRepo {
	return &RelationRepo{
		relationDal: relationDal,
	}
}

type RelationRepo struct {
	relationDal *dal.RelationDal
}

func (r *RelationRepo) NextIdentity() (string, error) {
	sf, err := snowflake.NewNode(1)
	if err != nil {
		return "", err
	}
	return sf.Generate().String(), nil
}

func (r *RelationRepo) FindRelation(ctx context.Context, followerID, followeeID string) (*entity.Relation, error) {
	relationPo, err := r.relationDal.SelectRelation(ctx, followerID, followeeID, nil)
	return converter.RelationToEntity(relationPo), err
}

func (r *RelationRepo) FindRelationNonNil(ctx context.Context, followerID, followeeID string) (*entity.Relation, error) {
	rlt, err := r.FindRelation(ctx, followerID, followeeID)
	if err != nil {
		return nil, err
	}
	if rlt == nil {
		return nil, errno.RecordNotFound
	}
	return rlt, nil
}

func (r *RelationRepo) UpsertRelation(ctx context.Context, relation *entity.Relation) (err error) {
	tx := r.relationDal.BeginTransaction(ctx)
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	if count, err := r.relationDal.CountRelation(ctx, relation.FollowerID, relation.FolloweeID, tx); err != nil {
		return errno.RelationDBError
	} else if count > 0 {
		return errno.RelationExistError
	}

	if err = r.relationDal.UpsertRelation(ctx, relation, tx); err != nil {
		return errno.RelationDBError
	}

	return nil
}

func (r *RelationRepo) DeleteRelation(ctx context.Context, relation *entity.Relation) (err error) {
	tx := r.relationDal.BeginTransaction(ctx)
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	relationPo, err := r.relationDal.SelectRelation(ctx, relation.FollowerID, relation.FolloweeID, tx)
	if err != nil {
		return errno.RelationDBError
	} else if relationPo == nil {
		return errno.RecordNotFound
	}

	if err = r.relationDal.DeleteRelation(ctx, relationPo, tx); err != nil {
		return errno.RelationDBError
	}

	return nil
}

func (r *RelationRepo) CountFollowee(ctx context.Context, followerID string) (int64, error) {
	return r.relationDal.CountFollowee(ctx, followerID, nil)
}

func (r *RelationRepo) CountFollower(ctx context.Context, followeeID string) (int64, error) {
	return r.relationDal.CountFollower(ctx, followeeID, nil)
}

func (r *RelationRepo) GetFolloweeList(ctx context.Context, followerId string, offset int, size int) ([]*entity.Relation, error) {
	relationPos, err := r.relationDal.SelectFolloweeList(ctx, followerId, offset, size, nil)
	return converter.RelationToEntityList(relationPos), err
}

func (r *RelationRepo) GetFollowerList(ctx context.Context, followeeId string, offset int, size int) ([]*entity.Relation, error) {
	relationPos, err := r.relationDal.SelectFollowerList(ctx, followeeId, offset, size, nil)
	return converter.RelationToEntityList(relationPos), err
}
