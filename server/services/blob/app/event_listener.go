package app

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/event"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/repo"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/mq/amqp"
	"github.com/xince-fun/InstaGo/server/shared/consts"
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
			BlobID:     event.BlobID,
			UserID:     event.UserID,
			ObjectName: event.ObjectName,
			BlobType:   event.BlobType,
		}
		switch blob.BlobType {
		case consts.AvatarBlobType:
			if err = e.processAvatarBlob(ctx, blob); err != nil {
				klog.Errorf("process avatar blob error: %v", err)
				continue
			}
		default:
			if err = e.processOtherBlob(ctx, blob); err != nil {
				klog.Errorf("process other blob error: %v", err)
				continue
			}
		}
	}

	return nil
}

func (e *BlobEventListener) processAvatarBlob(ctx context.Context, blob *entity.Blob) error {
	blobR, err := e.blobRepo.FindBlobByUserType(ctx, blob.UserID, blob.BlobType)
	if err != nil {
		return err
	}
	if blobR != nil {
		blobR.BlobID = blob.BlobID
		blobR.ObjectName = blob.ObjectName
	} else {
		blobR = new(entity.Blob)
		err = copier.Copy(blobR, blob)
		if err != nil {
			klog.Infof("copy blob error: %v", err)
			return err
		}
	}
	return e.blobRepo.SaveBlob(ctx, blobR)
}

func (e *BlobEventListener) processOtherBlob(ctx context.Context, blob *entity.Blob) error {
	return e.blobRepo.SaveBlob(ctx, blob)
}
