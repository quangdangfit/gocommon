package rabbitmq

import (
	"github.com/streadway/amqp"
)

type IRabbitMQ interface {
	NewConnection() (*amqp.Connection, error)
	CloseConnection() error
	NewChannel() (*amqp.Channel, error)
	EnsureConnection() (err error)
	CloseChannel() error
	DeclareExchange() error
	DeclareQueue() error
	BindQueue(exchange, routingKey string) error
	Setup()
	ChanelIsClosed() bool
}

type IConsumer interface {
	Consume()
}

type IPublisher interface {
	Publish(payload []byte, routingKey string, reliable bool) (err error)
}
