package app

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/event"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/repo"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/mq/amqp"
)

var BlobEventListenerSet = wire.NewSet(
	amqp.SubscriberSet,
	amqp.NewEventSubscriber,
	NewBlobEventListener,
	wire.Bind(new(event.EventSubscriber), new(*amqp.EventSubscriber)),
)

type BlobEventListener struct {
	blobRepo        repo.BlobRepository
	eventSubscriber event.EventSubscriber
}

func NewBlobEventListener(blobRepo repo.BlobRepository, eventSubscriber event.EventSubscriber) *BlobEventListener {
	return &BlobEventListener{
		blobRepo:        blobRepo,
		eventSubscriber: eventSubscriber,
	}
}

func (e *BlobEventListener) Start(ctx context.Context) error {
	eventCh, cleanUp, err := e.eventSubscriber.Subscribe(ctx)
	if err != nil {
		klog.Errorf("subscribe event error: %v", err)
		return err
	}
	defer cleanUp()

	for event := range eventCh {
		blob := &entity.Blob{
			BlobID:   event.BlobID,
			UserID:   event.UserID,
			URL:      event.URL,
			BlobType: event.BlobType,
		}
		if err = e.blobRepo.Save(ctx, blob); err != nil {
			klog.Errorf("save blob error: %v", err)
			continue
		}
	}

	return nil
}
