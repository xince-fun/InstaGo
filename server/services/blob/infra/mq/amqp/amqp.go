package amqp

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/xince-fun/InstaGo/server/services/blob/conf"
)

var PublisherSet = wire.NewSet(
	ProvidePExchange,
	NewPublisher,
)

type PExchange string

func ProvidePExchange() PExchange {
	return PExchange(conf.GlobalServerConf.MQConfig.Exchange)
}

type Publisher struct {
	ch       *amqp.Channel
	exchange PExchange
}

func NewPublisher(conn *amqp.Connection, exchange PExchange) (*Publisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("cannot allocate channel: %v", err)
	}

	if err = declareExchange(ch, string(exchange)); err != nil {
		return nil, fmt.Errorf("cannot declare exchange: %v", err)
	}
	return &Publisher{
		ch:       ch,
		exchange: exchange,
	}, nil
}

func (p *Publisher) Publish(ctx context.Context, msg []byte) error {
	return p.ch.PublishWithContext(ctx,
		string(p.exchange),
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
}

func declareExchange(ch *amqp.Channel, exchange string) error {
	return ch.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
}

var SubscriberSet = wire.NewSet(
	ProvideSExchange,
	NewSubscriber,
)

type SExchange string

func ProvideSExchange() SExchange {
	return SExchange(conf.GlobalServerConf.MQConfig.Exchange)
}

type Subscriber struct {
	conn     *amqp.Connection
	exchange SExchange
}

func NewSubscriber(conn *amqp.Connection, exchange SExchange) (*Subscriber, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("cannot allocate channel: %v", err)
	}
	defer ch.Close()

	if err = declareExchange(ch, string(exchange)); err != nil {
		return nil, fmt.Errorf("cannot declare exchange: %v", err)
	}

	return &Subscriber{
		conn:     conn,
		exchange: exchange,
	}, nil
}

func (s *Subscriber) SubscribeRaw(_ context.Context) (<-chan amqp.Delivery, func(), error) {
	ch, err := s.conn.Channel()
	if err != nil {
		return nil, func() {}, fmt.Errorf("cannot allocate channel: %v", err)
	}
	closeCh := func() {
		err = ch.Close()
		if err != nil {
			klog.Errorf("close channel error: %v", err)
		}
	}

	q, err := ch.QueueDeclare(
		"",
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, closeCh, fmt.Errorf("cannot declare queue: %v", err)
	}

	cleanUp := func() {
		_, err = ch.QueueDelete(q.Name, false, false, false)
		if err != nil {
			klog.Errorf("delete queue error: %v", err)
		}
		closeCh()
	}

	err = ch.QueueBind(
		q.Name,
		"",
		string(s.exchange),
		false,
		nil,
	)
	if err != nil {
		return nil, cleanUp, fmt.Errorf("cannot bind queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, cleanUp, fmt.Errorf("cannot consume: %v", err)
	}
	return msgs, cleanUp, nil
}

func (s *Subscriber) Subscribe(ctx context.Context) (<-chan []byte, func(), error) {
	msgCh, cleanUp, err := s.SubscribeRaw(ctx)
	if err != nil {
		return nil, cleanUp, err
	}

	outCh := make(chan []byte)
	go func() {
		for msg := range msgCh {
			msg := msg
			outCh <- msg.Body
		}
		close(outCh)
	}()
	return outCh, cleanUp, nil
}
