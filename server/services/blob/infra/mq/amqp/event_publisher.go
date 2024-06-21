package amqp

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/xince-fun/InstaGo/server/services/blob/domain/event"
)

type EventPublisher struct {
	publisher *Publisher
}

func NewEventPublisher(publisher *Publisher) *EventPublisher {
	return &EventPublisher{
		publisher: publisher,
	}
}

func (ep *EventPublisher) Publish(ctx context.Context, event event.DomainEvent) error {
	body, err := sonic.Marshal(event)
	if err != nil {
		return err
	}
	return ep.publisher.Publish(ctx, body)
}
