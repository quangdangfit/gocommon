package rabbitmq

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/streadway/amqp"

	"github.com/quangdangfit/gocommon/logger"
)

type Consumer struct {
	*RabbitMQ

	done            chan error
	consumerTag     string // Name that Consumer identifies itself to the server
	lastRecoverTime int64
	//track service current status
	currentStatus atomic.Value

	threads     int
	handleFn    func(msg amqp.Delivery) error
	errHandleFn func(err error)
}

func NewConsumer(opts ...Option) IConsumer {
	opt := getOption(opts...)

	var sub = Consumer{
		done:            make(chan error),
		lastRecoverTime: time.Now().Unix(),
		handleFn:        opt.handleFn,
		errHandleFn:     opt.errHandleFn,
		threads:         opt.threads,
		RabbitMQ: &RabbitMQ{
			config: opt.config,
		},
	}

	_, err := sub.NewConnection()
	if err != nil {
		logger.Error("Consumer create new connection failed!")
	}

	err = sub.DeclareQueue()
	if err != nil {
		logger.Error("Consumer declare queue failed!")
	}

	return &sub
}

func (c *Consumer) Reconnect() (err error) {
	err = c.CloseAll()
	if err != nil {
		logger.Error("Cannot close rabbitmq connection: ", err)
	}

	var conn *amqp.Connection
	for {
		conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", c.config.Username, c.config.Password, c.config.AMQPUrl))
		if err == nil {
			break
		}

		logger.Errorf("Failed to create new connection to AMQP: %s. Sleep %d seconds to reconnect.", err, WaitTimeReconnect)
		time.Sleep(WaitTimeReconnect * time.Second)
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	c.channel = ch
	c.connection = conn
	c.closeChan = make(chan *amqp.Error)
	conn.NotifyClose(c.closeChan)

	logger.Info("Reconnect rabbitMQ successfully!!!")
	return nil
}

func (c *Consumer) Consume() {
	c.EnsureConnection()
	c.NewChannel()
	msgC, errC := c.ConsumingMessage()
	for {
		select {
		case msg := <-msgC:
			go c.handleFn(msg)
		case err := <-errC:
			go c.errHandleFn(err)
		}
	}
}

func (c *Consumer) ConsumingMessage() (chan amqp.Delivery, chan error) {
	msgCh := make(chan amqp.Delivery)
	errCh := make(chan error)

	c.startConsuming(msgCh, errCh)
	go func() {
		for {
			closedErr := <-c.closeChan
			if closedErr != nil {
				logger.Error("rabbitMQ connection is lost, reconnecting: ", closedErr)
				err := c.Reconnect()
				if err != nil {
					logger.Error("rabbitMQ failed to reconnect: ", err)
					continue
				}
				c.startConsuming(msgCh, errCh)
			}
		}
	}()

	return msgCh, errCh
}

func (c *Consumer) startConsuming(msgCh chan amqp.Delivery, errCh chan error) {
	go func() {
		msgs, err := c.channel.Consume(
			c.config.QueueName, // queue
			"",                 // consumer
			false,              // auto-ack
			false,              // exclusive
			false,              // no-local
			false,              // no-wait
			nil,                // args
		)
		if err != nil {
			errCh <- err
		}

		forever := make(chan bool)

		go func() {
			for d := range msgs {
				msgCh <- d
			}
		}()
		<-forever
	}()
}

// CloseAll : Close connection and channel
func (c *Consumer) CloseAll() (err error) {
	if !c.connection.IsClosed() {
		err = c.connection.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
