package repo

import (
	"context"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/relation/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/relation/infra/persistence"
	"github.com/xince-fun/InstaGo/server/services/relation/infra/persistence/dal"
)

type RelationRepo interface {
	NextIdentity() (string, error)

	FindRelation(context.Context, string, string) (*entity.Relation, error)
	FindRelationNonNil(context.Context, string, string) (*entity.Relation, error)

	GetFolloweeList(context.Context, string, int, int) ([]*entity.Relation, error)
	GetFollowerList(context.Context, string, int, int) ([]*entity.Relation, error)

	UpsertRelation(context.Context, *entity.Relation) error
	DeleteRelation(context.Context, *entity.Relation) error

	CountFollowee(context.Context, string) (int64, error)
	CountFollower(context.Context, string) (int64, error)

	IsFollow(context.Context, string, string) (bool, error)
}

var RelationRepositorySet = wire.NewSet(
	dal.NewRelationDal,
	persistence.NewRelationRepo,
	wire.Bind(new(RelationRepo), new(*persistence.RelationRepo)),
)
