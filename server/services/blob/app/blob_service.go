package app

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/blob/conf"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/repo"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/object/minio"
	"github.com/xince-fun/InstaGo/server/shared/errno"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/blob"
	"time"
)

var BlobApplicationServiceSet = wire.NewSet(
	repo.BlobRepositorySet,
	minio.BucketManagerSet,
	NewBlobApplicationService,
	wire.Bind(new(BucketManager), new(*minio.MinioBucketManager)),
)

type BlobApplicationService struct {
	blobRepo      repo.BlobRepository
	bucketManager BucketManager
}

type BucketManager interface {
	GenerateGetObjectSignedURL(context.Context, string, string, time.Duration) (string, error)
	GeneratePutObjectSignedURL(context.Context, string, string, time.Duration) (string, error)
}

func NewBlobApplicationService(blobRepo repo.BlobRepository, bucketManager BucketManager) *BlobApplicationService {
	return &BlobApplicationService{
		blobRepo:      blobRepo,
		bucketManager: bucketManager,
	}
}

func (s *BlobApplicationService) GeneratePutPreSignedUrl(ctx context.Context, req *blob.GeneratePutPreSignedUrlRequest) (resp *blob.GeneratePutPreSignedUrlResponse, err error) {
	resp = new(blob.GeneratePutPreSignedUrlResponse)

	blobID, err := s.blobRepo.NextIdentity()
	if err != nil {
		return nil, errno.BlobSrvError
	}

	objectName := fmt.Sprintf("%s/%d/%s", req.UserId, req.BlobType, blobID)
	bucketName := conf.GlobalServerConf.BucketConfig.AvatarBucket

	url, err := s.bucketManager.GeneratePutObjectSignedURL(ctx, bucketName, objectName, time.Duration(req.Timeout)*time.Second)
	if err != nil {
		return nil, errno.BlobSrvError
	}

	resp = &blob.GeneratePutPreSignedUrlResponse{
		Url:        url,
		Id:         blobID,
		ObjectName: objectName,
	}
	return resp, nil
}

func (s *BlobApplicationService) GenerateGetPreSignedUrl(ctx context.Context, req *blob.GenerateGetPreSignedUrlRequest) (resp *blob.GenerateGetPreSignedUrlResponse, err error) {
	resp = new(blob.GenerateGetPreSignedUrlResponse)

	blobR, err := s.blobRepo.FindBlobByIDNonNil(ctx, req.BlobId)
	if err != nil {
		return resp, errno.RecordNotFound
	}
	bucketName := conf.GlobalServerConf.BucketConfig.AvatarBucket
	url, err := s.bucketManager.GenerateGetObjectSignedURL(ctx, bucketName, blobR.ObjectName, time.Duration(req.Timeout)*time.Second)
	if err != nil {
		klog.Infof("error: %v", err)
		return resp, errno.BlobSrvError
	}

	resp = &blob.GenerateGetPreSignedUrlResponse{
		Url: url,
	}
	return resp, nil
}

func (s *BlobApplicationService) NotifyBlobUpload(ctx context.Context, req *blob.NotifyBlobUploadRequest) (resp *blob.NotifyBlobUploadResponse, err error) {
	resp = new(blob.NotifyBlobUploadResponse)

	b := &entity.Blob{
		BlobID:     req.BlobId,
		UserID:     req.UserId,
		ObjectName: req.ObjectName,
		BlobType:   req.BlobType,
	}
	if err = b.NotifyUpload(); err != nil {
		return nil, errno.BlobSrvError
	}

	if err = s.blobRepo.Save(ctx, b); err != nil {
		return resp, errno.BlobSrvError
	}

	return resp, nil
}
