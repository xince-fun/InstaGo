package app

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/blob/conf"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/repo"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/object/minio"
	"github.com/xince-fun/InstaGo/server/shared/errno"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/blob"
	"github.com/xince-fun/InstaGo/server/shared/utils"
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
	blobID, err := s.blobRepo.NextIdentity()
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.BlobSrvError)
		return resp, nil
	}

	objectName := fmt.Sprintf("%s/%d/%s", req.UserId, req.BlobType, blobID)
	bucketName := conf.GlobalServerConf.BucketConfig.AvatarBucket

	url, err := s.bucketManager.GeneratePutObjectSignedURL(ctx, bucketName, objectName, time.Duration(req.Timeout)*time.Second)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.BlobSrvError)
		return resp, nil
	}

	resp = &blob.GeneratePutPreSignedUrlResponse{
		BaseResp: utils.BuildBaseResp(nil),
		Url:      url,
		Id:       blobID,
	}
	return resp, nil
}

func (s *BlobApplicationService) GenerateGetPreSignedUrl(ctx context.Context, req *blob.GenerateGetPreSignedUrlRequest) (resp *blob.GenerateGetPreSignedUrlResponse, err error) {

	blobR, err := s.blobRepo.FindBlobByIDNonNil(ctx, req.BlobId)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.BlobSrvError)
		return resp, nil
	}

	bucketName := conf.GlobalServerConf.BucketConfig.AvatarBucket
	url, err := s.bucketManager.GenerateGetObjectSignedURL(ctx, bucketName, blobR.URL, time.Duration(req.Timeout)*time.Second)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.BlobSrvError)
		return resp, nil
	}

	resp = &blob.GenerateGetPreSignedUrlResponse{
		BaseResp: utils.BuildBaseResp(nil),
		Url:      url,
	}
	return resp, nil
}

func (s *BlobApplicationService) NotifyBlobUpload(ctx context.Context, req *blob.NotifyBlobUploadRequest) (resp *blob.NotifyBlobUploadResponse, err error) {
	resp = new(blob.NotifyBlobUploadResponse)

	b := &entity.Blob{}
	if err = b.NotifyUpload(); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.BlobSrvError)
		return resp, nil
	}

	if err = s.blobRepo.Save(ctx, b); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.BlobSrvError)
		return resp, nil
	}

	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}
