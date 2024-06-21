package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"time"
)

type GCPBucketManager struct {
	client *storage.Client
}

func NewGCPBucketManager(client *storage.Client) *GCPBucketManager {
	return &GCPBucketManager{
		client: client,
	}
}

func (m *GCPBucketManager) GenerateGetObjectSignedURL(ctx context.Context, bucket, object string) (string, error) {
	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(15 * time.Minute),
	}

	u, err := m.client.Bucket(bucket).SignedURL(object, opts)
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %v", bucket, err)
	}

	return u, nil
}

func (m *GCPBucketManager) GeneratePutObjectSignedURL(ctx context.Context, bucket, object string) (string, error) {
	opts := &storage.SignedURLOptions{
		Scheme: storage.SigningSchemeV4,
		Method: "PUT",
		Headers: []string{
			"Content-Type:application/octet-stream",
		},
		Expires: time.Now().Add(15 * time.Minute),
	}

	u, err := m.client.Bucket(bucket).SignedURL(object, opts)
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %v", bucket, err)
	}

	return u, nil
}
