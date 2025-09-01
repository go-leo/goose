package client

import (
	"net/http"

	"github.com/go-leo/goose"
	"google.golang.org/protobuf/encoding/protojson"
)

// Options interface defines methods to access all configurable options
type Options interface {
	// Returns http client
	Client() *http.Client

	// Returns protojson unmarshal options
	UnmarshalOptions() protojson.UnmarshalOptions

	// Returns protojson marshal options
	MarshalOptions() protojson.MarshalOptions

	// Returns error encoder
	ErrorDecoder() goose.ErrorDecoder

	// Returns error factory
	ErrorFactory() goose.ErrorFactory

	// Returns list of middlewares
	Middlewares() []Middleware

	// Indicates if fail-fast mode is enabled
	ShouldFailFast() bool

	// Gets validation error callback
	OnValidationErrCallback() goose.OnValidationErrCallback
}

type options struct {
	client                  *http.Client
	unmarshalOptions        protojson.UnmarshalOptions
	marshalOptions          protojson.MarshalOptions
	errorDecoder            goose.ErrorDecoder
	errorFactory            goose.ErrorFactory
	middlewares             []Middleware
	shouldFailFast          bool
	onValidationErrCallback goose.OnValidationErrCallback
}

// Option defines a function type for modifying options
type Option func(o *options)

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *options) Client() *http.Client {
	return o.client
}

func (o *options) UnmarshalOptions() protojson.UnmarshalOptions {
	return o.unmarshalOptions
}

func (o *options) MarshalOptions() protojson.MarshalOptions {
	return o.marshalOptions
}

func (o *options) ErrorDecoder() goose.ErrorDecoder {
	return o.errorDecoder
}

func (o *options) ErrorFactory() goose.ErrorFactory {
	return o.errorFactory
}

func (o *options) Middlewares() []Middleware {
	return o.middlewares
}

func (o *options) ShouldFailFast() bool {
	return o.shouldFailFast
}

func (o *options) OnValidationErrCallback() goose.OnValidationErrCallback {
	return o.onValidationErrCallback
}

// Client set *http.Client
func Client(client *http.Client) Option {
	return func(o *options) {
		o.client = client
	}
}

// UnmarshalOptions sets protojson unmarshal options
func UnmarshalOptions(opts protojson.UnmarshalOptions) Option {
	return func(o *options) {
		o.unmarshalOptions = opts
	}
}

// MarshalOptions sets protojson marshal options
func MarshalOptions(opts protojson.MarshalOptions) Option {
	return func(o *options) {
		o.marshalOptions = opts
	}
}

// ErrorEncoder configures custom error decoder
func ErrorEncoder(decoder goose.ErrorDecoder) Option {
	return func(o *options) {
		o.errorDecoder = decoder
	}
}

// ErrorFactory set error factory
func ErrorFactory(factory goose.ErrorFactory) Option {
	return func(o *options) {
		o.errorFactory = factory
	}
}

// Middlewares appends middlewares to the chain
func Middlewares(middlewares ...Middleware) Option {
	return func(o *options) {
		o.middlewares = append(o.middlewares, middlewares...)
	}
}

// FailFast enables fail-fast mode
func FailFast() Option {
	return func(o *options) {
		o.shouldFailFast = true
	}
}

// WithOnValidationErrCallback sets validation error callback
func OnValidationErrCallback(OnValidationErrCallback goose.OnValidationErrCallback) Option {
	return func(o *options) {
		o.onValidationErrCallback = OnValidationErrCallback
	}
}

// NewOptions creates new Options instance with defaults and applies provided options
func NewOptions(opts ...Option) Options {
	o := &options{
		client:                  &http.Client{},
		unmarshalOptions:        protojson.UnmarshalOptions{},
		marshalOptions:          protojson.MarshalOptions{},
		errorDecoder:            goose.DefaultDecodeError,
		errorFactory:            goose.DefaultErrorFactory,
		middlewares:             []Middleware{},
		shouldFailFast:          false,
		onValidationErrCallback: nil,
	}
	o = o.apply(opts...)
	return o
}
