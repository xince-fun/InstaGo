package repo

import (
	"context"
	"github.com/xince-fun/InstaGo/server/services/relation/domain/entity"
)

type RelationRepo interface {
	NextIdentity() (string, error)

	FindRelation(context.Context, string, string) (*entity.Relation, error)
	FindRelationNonNil(context.Context, string, string) (*entity.Relation, error)

	UpsertRelation(context.Context, *entity.Relation) error

	DeleteRelation(context.Context, *entity.Relation) error
}
