package repo

import (
	"context"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/event"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/mq/amqp"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/persistence"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/persistence/dal"
)

type BlobRepository interface {
	NextIdentity() (string, error)
	Save(context.Context, *entity.Blob) (err error)
	SaveBlob(context.Context, *entity.Blob) (err error)
	FindBlobByID(context.Context, string) (*entity.Blob, error)
	FindBlobByIDNonNil(context.Context, string) (*entity.Blob, error)
	FindBlobByUserType(context.Context, string, int8) (*entity.Blob, error)
	FindBlobByUserTypeNonNil(context.Context, string, int8) (*entity.Blob, error)
}

var BlobRepositorySet = wire.NewSet(
	dal.NewBlobDal,
	amqp.PublisherSet,
	amqp.NewEventPublisher,
	persistence.NewBlobRepo,
	wire.Bind(new(event.EventPublisher), new(*amqp.EventPublisher)),
	wire.Bind(new(BlobRepository), new(*persistence.BlobRepo)),
)
