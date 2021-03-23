package rabbitmq

import (
	"github.com/streadway/amqp"

	"github.com/quangdangfit/gocommon/logger"
)

type Publisher struct {
	*RabbitMQ
}

func NewPublisher(opts ...Option) IPublisher {
	opt := getOption(opts...)

	pub := Publisher{
		RabbitMQ: &RabbitMQ{
			config: opt.config,
		},
	}
	_, err := pub.NewConnection()
	if err != nil {
		logger.Error("Publisher create new connection failed!")
	}

	err = pub.DeclareExchange()
	if err != nil {
		logger.Error("Publisher declare exchange failed!")
	}

	return &pub
}

func (pub *Publisher) Publish(payload []byte, routingKey string, reliable bool) (err error) {
	// New channel and close after publish
	pub.EnsureConnection()
	channel, _ := pub.connection.Channel()
	defer channel.Close()

	// Reliable Publisher confirms require confirm.select support from the connection.
	var confirms chan amqp.Confirmation
	if reliable {
		if err := channel.Confirm(false); err != nil {
			logger.Errorf("Channel could not be put into confirm mode: %s", err)
			return err
		}
		confirms = channel.NotifyPublish(make(chan amqp.Confirmation, 1))
	}

	if err = channel.Publish(
		pub.config.ExchangeName, // publish to an exchange
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:     "application/json",
			ContentEncoding: "",
			Body:            payload,
			DeliveryMode:    amqp.Persistent, // 1=non-persistent, 2=persistent
			Priority:        0,               // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		logger.Error("Failed to publish message ", err)
		return err
	}

	if confirms != nil {
		defer pub.confirmOne(confirms)
	}

	return nil
}

func (pub *Publisher) confirmOne(confirms <-chan amqp.Confirmation) bool {
	confirmed := <-confirms
	if confirmed.Ack {
		logger.Info("Confirmed delivery with delivery tag: ", confirmed.DeliveryTag)
		return true
	}

	logger.Info("Failed delivery of delivery tag: ", confirmed.DeliveryTag)
	return false
}
