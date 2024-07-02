package amqp

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7/pkg/notification"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/event"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type S3Event struct {
	EventName string               `json:"EventName"`
	Key       string               `json:"Key"`
	Records   []notification.Event `json:"Records"`
}

type EventSubscriber struct {
	subscriber *Subscriber
}

func NewEventSubscriber(subscriber *Subscriber) *EventSubscriber {
	return &EventSubscriber{
		subscriber: subscriber,
	}
}

func (es *EventSubscriber) Subscribe(ctx context.Context) (<-chan event.BlobUploadedEvent, func(), error) {
	msgCh, cleanUp, err := es.subscriber.Subscribe(ctx)

	outCh := make(chan event.BlobUploadedEvent)
	go func() {
		defer close(outCh)
		for msg := range msgCh {
			var event S3Event
			if err := sonic.Unmarshal(msg, &event); err != nil {
				klog.Errorf("unmarshal event error: %v", err)
				continue
			}
			bEvent := parseS3EventKey(&event)
			if bEvent.BlobID == "" {
				continue
			}
			outCh <- bEvent
		}
	}()

	return outCh, cleanUp, err
}

func parseS3EventKey(s3e *S3Event) event.BlobUploadedEvent {
	// key: bucket/userid/blobtype/blobid
	suf := filepath.Ext(s3e.Key)
	key := strings.TrimSuffix(s3e.Key, suf)
	sl := strings.Split(key, "/")
	if len(sl) != 4 {
		return event.BlobUploadedEvent{}
	}
	objectName := strings.TrimPrefix(key, sl[0]+"/")
	blobType, err := strconv.ParseUint(sl[2], 10, 8)
	if err != nil {
		return event.BlobUploadedEvent{}
	}
	if len(s3e.Records) == 0 {
		return event.BlobUploadedEvent{}
	}
	et, err := time.Parse(time.RFC3339, s3e.Records[0].EventTime)
	if err != nil {
		return event.BlobUploadedEvent{}
	}
	id, err := uuid.NewV7()
	if err != nil {
		return event.BlobUploadedEvent{}
	}
	return event.BlobUploadedEvent{
		EventId:    id.String(),
		BlobID:     sl[3],
		UserID:     sl[1],
		BlobType:   int8(blobType),
		ObjectName: objectName,
		UploadTime: et,
	}
}
