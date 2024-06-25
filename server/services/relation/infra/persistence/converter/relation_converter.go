package converter

import (
	"github.com/xince-fun/InstaGo/server/services/relation/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/relation/infra/persistence/po"
)

func RelationToPo(relation *entity.Relation) *po.Relation {
	return &po.Relation{
		FollowerID: relation.FollowerID,
		FolloweeID: relation.FolloweeID,
	}
}

func RelationToEntity(relation *po.Relation) *entity.Relation {
	return &entity.Relation{
		FollowerID: relation.FollowerID,
		FolloweeID: relation.FolloweeID,
	}
}
