package jwt

import (
	"time"
)

// Define constants
const (
	DefaultTokenExpiredTime = 10 * time.Minute
	DefaultTokenSecretKey   = "secret-key"
	DefaultSigningMethod    = "HS256"
)

// Option validation option
type Option interface {
	apply(*option)
}

// option implement
type option struct {
	TokenExpiredTime time.Duration
	TokenSecretKey   string
	SigningMethod    string
}

type optionFn func(*option)

func (optFn optionFn) apply(opt *option) {
	optFn(opt)
}

// WithExpiredTime set TokenExpiredTime
func WithExpiredTime(d time.Duration) Option {
	return optionFn(func(opt *option) {
		opt.TokenExpiredTime = d
	})
}

// WithTokenSecretKey set TokenSecretKey
func WithTokenSecretKey(key string) Option {
	return optionFn(func(opt *option) {
		opt.TokenSecretKey = key
	})
}

// WithSigningMethod set SigningMethod
func WithSigningMethod(method string) Option {
	return optionFn(func(opt *option) {
		opt.SigningMethod = method
	})
}

func getDefaultOption() *option {
	return &option{
		TokenExpiredTime: DefaultTokenExpiredTime,
		TokenSecretKey:   DefaultTokenSecretKey,
		SigningMethod:    DefaultSigningMethod,
	}
}

func getOption(opts ...Option) *option {
	opt := getDefaultOption()
	for _, o := range opts {
		o.apply(opt)
	}

	return opt
}
