package amqp

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/event"
)

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
			var event event.BlobUploadedEvent
			if err := sonic.Unmarshal(msg, &event); err != nil {
				klog.Errorf("unmarshal event error: %v", err)
				continue
			}
			outCh <- event
		}
	}()

	return outCh, cleanUp, err
}
