package rabbitmq

import (
	"github.com/streadway/amqp"
)

const (
	ConsumerThreads = 10
)

// Option rabbitmq consumerOption
type Option interface {
	apply(*consumerOption)
}

// consumerOption implement
type consumerOption struct {
	handleFn    func(msg amqp.Delivery) error
	errHandleFn func(err error)
	threads     int
	config      *Config
}

type optionFn func(*consumerOption)

func (optFn optionFn) apply(opt *consumerOption) {
	optFn(opt)
}

// WithHandleFn set handleFn
func WithHandleFn(handleFn func(msg amqp.Delivery) error) Option {
	return optionFn(func(opt *consumerOption) {
		opt.handleFn = handleFn
	})
}

// WithErrorHandleFn set errHandleFn
func WithErrorHandleFn(errHandleFn func(err error)) Option {
	return optionFn(func(opt *consumerOption) {
		opt.errHandleFn = errHandleFn
	})
}

// WithThreads set threads
func WithThreads(threads int) Option {
	return optionFn(func(opt *consumerOption) {
		opt.threads = threads
	})
}

// WithConfig set config
func WithConfig(config *Config) Option {
	return optionFn(func(opt *consumerOption) {
		opt.config = config
	})
}

func getDefaultOption() *consumerOption {
	return &consumerOption{
		handleFn:    nil,
		errHandleFn: nil,
		threads:     ConsumerThreads,
		config:      &Config{},
	}
}

func getOption(opts ...Option) *consumerOption {
	opt := getDefaultOption()
	for _, o := range opts {
		o.apply(opt)
	}

	return opt
}
