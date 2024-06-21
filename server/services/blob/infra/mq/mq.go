package mq

import "context"

type Publisher interface {
	Publish(context.Context, []byte) error
}

type Subscriber interface {
	Subscribe(ctx context.Context) (chan []byte, func(), error)
}

type MessageQueue interface {
	Publisher
	Subscriber
}
