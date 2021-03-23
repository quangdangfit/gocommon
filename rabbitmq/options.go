package rabbitmq

import (
	"github.com/streadway/amqp"
)

const (
	ConsumerThreads = 10
)

// Option rabbitmq option
type Option interface {
	apply(*option)
}

// option implement
type option struct {
	handleFn    func(msg amqp.Delivery) error
	errHandleFn func(err error)
	threads     int
	config      *Config
}

type optionFn func(*option)

func (optFn optionFn) apply(opt *option) {
	optFn(opt)
}

// WithHandleFn set handleFn
func WithHandleFn(handleFn func(msg amqp.Delivery) error) Option {
	return optionFn(func(opt *option) {
		opt.handleFn = handleFn
	})
}

// WithErrorHandleFn set errHandleFn
func WithErrorHandleFn(errHandleFn func(err error)) Option {
	return optionFn(func(opt *option) {
		opt.errHandleFn = errHandleFn
	})
}

// WithThreads set threads
func WithThreads(threads int) Option {
	return optionFn(func(opt *option) {
		opt.threads = threads
	})
}

// WithConfig set config
func WithConfig(config *Config) Option {
	return optionFn(func(opt *option) {
		opt.config = config
	})
}

func getDefaultOption() *option {
	return &option{
		handleFn:    nil,
		errHandleFn: nil,
		threads:     ConsumerThreads,
		config:      &Config{},
	}
}

func getOption(opts ...Option) *option {
	opt := getDefaultOption()
	for _, o := range opts {
		o.apply(opt)
	}

	return opt
}
