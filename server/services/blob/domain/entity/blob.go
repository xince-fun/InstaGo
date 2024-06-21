package entity

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/event"
	"time"
)

type Blob struct {
	BlobID   string
	UserID   string
	URL      string
	BlobType int8
	events   []event.BlobEvent
}

func (b *Blob) NotifyUpload() error {
	id, err := uuid.NewV7()
	if err != nil {
		klog.Errorf("uuid.NewV7() failed: %v", err)
	}
	b.RaiseEvent(&event.BlobUploadedEvent{
		EventId:    id.String(),
		BlobID:     b.BlobID,
		UserID:     b.UserID,
		URL:        b.URL,
		BlobType:   b.BlobType,
		UploadTime: time.Now(),
	})
	return nil
}

func (b *Blob) RaiseEvent(event event.BlobEvent) {
	b.events = append(b.events, event)
}

func (b *Blob) ClearEvents() {
	for idx := range b.events {
		b.events[idx] = nil
	}
	b.events = nil
}

func (b *Blob) Events() []event.BlobEvent {
	return b.events
}