package minio

import (
	"context"
	"github.com/google/wire"
	"time"

	"github.com/minio/minio-go/v7"
)

var BucketManagerSet = wire.NewSet(
	NewMinioBucketManager,
)

type MinioBucketManager struct {
	client *minio.Client
}

func NewMinioBucketManager(client *minio.Client) *MinioBucketManager {
	return &MinioBucketManager{
		client: client,
	}
}

func (m *MinioBucketManager) GenerateGetObjectSignedURL(ctx context.Context, bucket, object string, timeOut time.Duration) (string, error) {
	url, err := m.client.PresignedGetObject(ctx, bucket, object, timeOut, nil)
	if err != nil {
		return "", err
	}
	return url.String(), err
}

func (m *MinioBucketManager) GeneratePutObjectSignedURL(ctx context.Context, bucket, object string, timeOut time.Duration) (string, error) {
	url, err := m.client.PresignedPutObject(ctx, bucket, object, timeOut)
	if err != nil {
		return "", err
	}
	return url.String(), err
}
