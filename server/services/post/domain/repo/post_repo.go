package repo

import (
	"context"
	"github.com/xince-fun/InstaGo/server/services/post/domain/entity"
)

type PostRepository interface {
	NextIdentity() (string, error)

	UpsertPost(context.Context, *entity.Post) error
}
