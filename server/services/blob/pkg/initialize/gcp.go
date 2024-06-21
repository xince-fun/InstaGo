package initialize

import (
	"cloud.google.com/go/storage"
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"
)

func InitGCPClient(ctx context.Context) *storage.Client {
	if err := CheckGCPCredentials(); err != nil {
		klog.Fatalf("CheckGCPCredentials: %v", err)
	}
	client, err := storage.NewClient(ctx)
	if err != nil {
		klog.Fatalf("storage.NewClient: %v", err)
	}
	return client
}

func CheckGCPCredentials() error {
	creds, ok := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if !ok {
		return errors.New("error: GOOGLE_APPLICATION_CREDENTIALS not set.\nFollow https://cloud.google.com/storage/docs/reference/libraries to set it up.\n")
	}
	if _, err := os.Stat(creds); os.IsNotExist(err) {
		return errors.New("error: GOOGLE_APPLICATION_CREDENTIALS file does not exist.\n")
	}
	return nil
}
