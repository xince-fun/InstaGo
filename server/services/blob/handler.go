package main

import (
	"context"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/blob/app"
	blob "github.com/xince-fun/InstaGo/server/shared/kitex_gen/blob"
)

var BlobServiceImplSet = wire.NewSet(
	app.BlobApplicationServiceSet,
	app.BlobEventListenerSet,
	NewBlobServiceImpl,
)

func NewBlobServiceImpl(blobAppService *app.BlobApplicationService, listener *app.BlobEventListener) *BlobServiceImpl {
	return &BlobServiceImpl{
		app:      blobAppService,
		listener: listener,
	}
}

// BlobServiceImpl implements the last service interface defined in the IDL.
type BlobServiceImpl struct {
	app      *app.BlobApplicationService
	listener *app.BlobEventListener
}

// GeneratePutPreSignedUrl implements the BlobServiceImpl interface.
func (s *BlobServiceImpl) GeneratePutPreSignedUrl(ctx context.Context, req *blob.GeneratePutPreSignedUrlRequest) (resp *blob.GeneratePutPreSignedUrlResponse, err error) {
	return s.app.GeneratePutPreSignedUrl(ctx, req)
}

// GenerateGetPreSignedUrl implements the BlobServiceImpl interface.
func (s *BlobServiceImpl) GenerateGetPreSignedUrl(ctx context.Context, req *blob.GenerateGetPreSignedUrlRequest) (resp *blob.GenerateGetPreSignedUrlResponse, err error) {
	return s.app.GenerateGetPreSignedUrl(ctx, req)
}

// NotifyBlobUpload implements the BlobServiceImpl interface.
func (s *BlobServiceImpl) NotifyBlobUpload(ctx context.Context, req *blob.NotifyBlobUploadRequest) (resp *blob.NotifyBlobUploadResponse, err error) {
	return s.app.NotifyBlobUpload(ctx, req)
}
