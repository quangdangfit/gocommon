package rabbitmq

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"

	"github.com/quangdangfit/gocommon/logger"
)

// Rabbitmq constant
const (
	WaitTimeReconnect = 5
)

type RabbitMQ struct {
	config          *Config
	connection      *amqp.Connection
	channel         *amqp.Channel
	errorChan       chan *amqp.Error
	failedChan      chan *amqp.Delivery
	isClosed        bool
	channelIsClosed bool
	closeChan       chan *amqp.Error
}

func New(conf *Config) (IRabbitMQ, error) {
	r := &RabbitMQ{
		config: conf,
	}

	_, err := r.NewConnection()
	if err != nil {
		return nil, err
	}

	_, err = r.NewChannel()
	if err != nil {
		return nil, err
	}

	r.Setup()

	return r, nil
}

func (mq *RabbitMQ) NewConnection() (*amqp.Connection, error) {
	conn, err := amqp.Dial(mq.config.AMQPUrl)
	for err != nil {
		fmt.Printf("Failed to create new connection to AMQP: %s. Sleep %d seconds to reconnect", err, WaitTimeReconnect)
		time.Sleep(WaitTimeReconnect * time.Second)
		conn, err = amqp.Dial(mq.config.AMQPUrl)
	}
	mq.connection = conn

	return conn, nil
}

func (mq *RabbitMQ) CloseConnection() error {
	if mq.isClosed {
		return nil
	}
	mq.CloseChannel()

	if mq.connection != nil {
		if err := mq.connection.Close(); err != nil {
			return err
		}
		mq.connection = nil
	}

	mq.isClosed = true
	return nil
}

func (mq *RabbitMQ) NewChannel() (*amqp.Channel, error) {
	mq.EnsureConnection()

	if mq.connection == nil || mq.connection.IsClosed() {
		logger.Error("Connection is not open, cannot create new channel")
		return nil, fmt.Errorf("Connection is not open")
	}

	channel, err := mq.connection.Channel()
	if err != nil {
		logger.Error("Failed to new channel: ", err)
		return nil, err
	}
	mq.channel = channel
	mq.channelIsClosed = false
	logger.Info("New channel successfully")
	return channel, nil
}

func (mq *RabbitMQ) EnsureConnection() (err error) {
	if mq.connection == nil || mq.connection.IsClosed() {
		_, err = mq.NewConnection()
		if err != nil {
			return err
		}
	}
	return nil
}

func (mq *RabbitMQ) CloseChannel() error {
	if mq.isClosed || mq.channelIsClosed {
		return nil
	}
	logger.Info("Close channel")
	if mq.channel != nil {
		_ = mq.channel.Close()
		mq.channel = nil
		mq.channelIsClosed = true
	}

	return nil
}

func (mq *RabbitMQ) DeclareExchange() error {
	mq.NewChannel()
	defer mq.CloseChannel()

	if mq.ChanelIsClosed() {
		logger.Error("Channel is not open, cannot declare exchange")
	}

	if err := mq.channel.ExchangeDeclare(
		mq.config.ExchangeName, // name
		mq.config.ExchangeType, // type
		true,                   // durable
		false,                  // auto-deleted
		false,                  // internal
		false,                  // noWait
		nil,                    // arguments
	); err != nil {
		logger.Error("Failed to declare exchange: ", err)
		return err
	}

	logger.Info("Declared exchange: ", mq.config.ExchangeName)
	return nil
}

func (mq *RabbitMQ) DeclareQueue() error {
	mq.NewChannel()
	defer mq.CloseChannel()

	if mq.ChanelIsClosed() {
		logger.Error("Channel is not open, cannot declare exchange")
	}

	if _, err := mq.channel.QueueDeclare(
		mq.config.QueueName,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		logger.Error("Failed to declare queue: ", err)
		return err
	}

	logger.Info("Declared queue: ", mq.config.QueueName)
	return nil
}

func (mq *RabbitMQ) BindQueue(exchange, routingKey string) error {
	if err := mq.channel.QueueBind(
		mq.config.QueueName, // name
		routingKey,          // key
		exchange,            // exchange
		false,               //noWait
		nil,                 // args
	); err != nil {
		return err
	}
	return nil
}

func (mq *RabbitMQ) Setup() {
	mq.DeclareExchange()
	mq.DeclareQueue()
}

func (mq *RabbitMQ) ChanelIsClosed() bool {
	if mq.channel == nil || mq.channelIsClosed {
		return true
	}
	return false
}
