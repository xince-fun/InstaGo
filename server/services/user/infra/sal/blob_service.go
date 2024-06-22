package sal

import (
	"context"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/blob"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/blob/blobservice"
)

var BlobManagerSet = wire.NewSet(
	NewBlobManager,
)

type BlobManager struct {
	client blobservice.Client
}

func NewBlobManager(client blobservice.Client) *BlobManager {
	return &BlobManager{
		client: client,
	}
}

func (b *BlobManager) UploadBlob(ctx context.Context, req *blob.GeneratePutPreSignedUrlRequest) (*blob.GeneratePutPreSignedUrlResponse, error) {
	return b.client.GeneratePutPreSignedUrl(ctx, req)
}

func (b *BlobManager) GetBlob(ctx context.Context, req *blob.GenerateGetPreSignedUrlRequest) (*blob.GenerateGetPreSignedUrlResponse, error) {
	return b.client.GenerateGetPreSignedUrl(ctx, req)
}

func (b *BlobManager) NotifyBlobUpload(ctx context.Context, req *blob.NotifyBlobUploadRequest) (*blob.NotifyBlobUploadResponse, error) {
	return b.client.NotifyBlobUpload(ctx, req)
}
