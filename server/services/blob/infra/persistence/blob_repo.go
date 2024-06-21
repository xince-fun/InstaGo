package persistence

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/event"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/persistence/dal"
	"github.com/xince-fun/InstaGo/server/shared/errno"
)

type BlobRepo struct {
	blobDal        *dal.BlobDal
	eventPublisher event.EventPublisher
}

func NewBlobRepo(blobDal *dal.BlobDal, eventPublisher event.EventPublisher) *BlobRepo {
	return &BlobRepo{
		blobDal:        blobDal,
		eventPublisher: eventPublisher,
	}
}

func (r *BlobRepo) NextIdentity() (string, error) {
	sf, err := snowflake.NewNode(1)
	if err != nil {
		return "", err
	}
	return sf.Generate().String(), nil
}

func (r *BlobRepo) Save(ctx context.Context, blob *entity.Blob) (err error) {
	for _, blobEvent := range blob.Events() {
		if err = r.eventPublisher.Publish(ctx, blobEvent); err != nil {
			return err
		}
	}
	blob.ClearEvents()
	return nil
}

func (r *BlobRepo) SaveBlob(ctx context.Context, blob *entity.Blob) (err error) {
	return r.blobDal.Upsert(ctx, blob, nil)
}

func (r *BlobRepo) FindBlobByID(ctx context.Context, blobID string) (*entity.Blob, error) {
	return r.blobDal.SelectBlobByID(ctx, blobID, nil)
}

func (r *BlobRepo) FindBlobByIDNonNil(ctx context.Context, blobID string) (*entity.Blob, error) {
	rlt, err := r.FindBlobByID(ctx, blobID)
	if err != nil {
		return nil, err
	}
	if rlt == nil {
		return nil, errno.RecordNotFound
	}
	return rlt, nil
}
