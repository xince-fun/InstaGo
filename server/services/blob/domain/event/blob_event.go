package event

import (
	"context"
	"fmt"
	"time"
)

type EventPublisher interface {
	Publish(context.Context, DomainEvent) error
}

type EventSubscriber interface {
	Subscribe(context.Context) (<-chan BlobUploadedEvent, func(), error)
}

type DomainEvent interface {
	fmt.Stringer
	Id() string
	OccurredOn() time.Time
}

type BlobEvent interface {
	DomainEvent
	BlobId() string
}

type BlobUploadedEvent struct {
	EventId    string
	BlobID     string
	UserID     string
	ObjectName string
	BlobType   int8
	UploadTime time.Time
}

func (e *BlobUploadedEvent) Id() string {
	return e.EventId
}

func (e *BlobUploadedEvent) OccurredOn() time.Time {
	return time.Now()
}

func (e *BlobUploadedEvent) BlobId() string {
	return e.BlobID
}

func (e *BlobUploadedEvent) String() string {
	return fmt.Sprintf("BlobUploadedEvent [EventId: %s, BlobID: %s, UserID: %s, URL: %s, UploadTime: %s]",
		e.EventId, e.BlobID, e.UserID, e.ObjectName, e.UploadTime.Format(time.RFC3339))
}
