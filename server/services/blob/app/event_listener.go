package app

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/event"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/repo"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/cache"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/mq/amqp"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/sal"
	"github.com/xince-fun/InstaGo/server/shared/consts"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/user"
)

var BlobEventListenerSet = wire.NewSet(
	amqp.SubscriberSet,
	sal.UserManagerSet,
	amqp.NewEventSubscriber,
	NewBlobEventListener,
	wire.Bind(new(event.EventSubscriber), new(*amqp.EventSubscriber)),
	wire.Bind(new(UserManager), new(*sal.UserManager)),
)

type BlobEventListener struct {
	blobRepo        repo.BlobRepository
	eventSubscriber event.EventSubscriber
	cacheManager    CacheManager
	userManager     UserManager
}

type UserManager interface {
	UpdateAvatarInfo(context.Context, *user.UpdateAvatarInfoRequest) (*user.UpdateAvatarInfoResponse, error)
}

func NewBlobEventListener(blobRepo repo.BlobRepository, eventSubscriber event.EventSubscriber, cacheManager CacheManager,
	userManager UserManager) *BlobEventListener {
	return &BlobEventListener{
		blobRepo:        blobRepo,
		eventSubscriber: eventSubscriber,
		cacheManager:    cacheManager,
		userManager:     userManager,
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
		go func() {
			if err = e.cacheManager.Set(ctx, blob.BlobID, &cache.BlobItem{BlobID: blob.BlobID, ObjectName: blob.ObjectName}); err != nil {
				klog.Errorf("set cache error: %v", err)
			}
		}()
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
	_, err = e.userManager.UpdateAvatarInfo(ctx, &user.UpdateAvatarInfoRequest{
		UserId:   blobR.UserID,
		AvatarId: blobR.BlobID,
	})
	if err != nil {
		return err
	}
	return e.blobRepo.SaveBlob(ctx, blobR)
}

func (e *BlobEventListener) processOtherBlob(ctx context.Context, blob *entity.Blob) error {
	return e.blobRepo.SaveBlob(ctx, blob)
}
