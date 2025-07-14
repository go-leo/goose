package gonic

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// Options interface defines methods to access all configurable options
type Options interface {
	// Returns protojson unmarshal options
	UnmarshalOptions() protojson.UnmarshalOptions

	// Returns protojson marshal options
	MarshalOptions() protojson.MarshalOptions

	// Gets the error encoder function
	ErrorEncoder() ErrorEncoder

	// Gets the response transformer
	ResponseTransformer() ResponseTransformer

	// Returns list of middlewares
	Middlewares() []gin.HandlerFunc

	// Indicates if fail-fast mode is enabled
	ShouldFailFast() bool

	// Gets validation error callback
	OnValidationErrCallback() OnValidationErrCallback
}

type options struct {
	unmarshalOptions        protojson.UnmarshalOptions
	marshalOptions          protojson.MarshalOptions
	errorEncoder            ErrorEncoder
	responseTransformer     ResponseTransformer
	middlewares             []gin.HandlerFunc
	shouldFailFast          bool
	onValidationErrCallback OnValidationErrCallback
}

// Option defines a function type for modifying options
type Option func(o *options)

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *options) UnmarshalOptions() protojson.UnmarshalOptions {
	return o.unmarshalOptions
}

func (o *options) MarshalOptions() protojson.MarshalOptions {
	return o.marshalOptions
}

func (o *options) ErrorEncoder() ErrorEncoder {
	return o.errorEncoder
}

func (o *options) ResponseTransformer() ResponseTransformer {
	return o.responseTransformer
}

func (o *options) Middlewares() []gin.HandlerFunc {
	return o.middlewares
}

func (o *options) ShouldFailFast() bool {
	return o.shouldFailFast
}

func (o *options) OnValidationErrCallback() OnValidationErrCallback {
	return o.onValidationErrCallback
}

// WithUnmarshalOptions sets protojson unmarshal options
func WithUnmarshalOptions(opts protojson.UnmarshalOptions) Option {
	return func(o *options) {
		o.unmarshalOptions = opts
	}
}

// WithMarshalOptions sets protojson marshal options
func WithMarshalOptions(opts protojson.MarshalOptions) Option {
	return func(o *options) {
		o.marshalOptions = opts
	}
}

// WithErrorEncoder configures custom error encoder
func WithErrorEncoder(encoder ErrorEncoder) Option {
	return func(o *options) {
		o.errorEncoder = encoder
	}
}

// WithResponseTransformer sets response transformer
func WithResponseTransformer(transformer ResponseTransformer) Option {
	return func(o *options) {
		o.responseTransformer = transformer
	}
}

// WithMiddlewares appends middlewares to the chain
func WithMiddlewares(middlewares ...gin.HandlerFunc) Option {
	return func(o *options) {
		o.middlewares = append(o.middlewares, middlewares...)
	}
}

// WithOnValidationErrCallback sets validation error callback
func WithOnValidationErrCallback(onValidationErrCallback OnValidationErrCallback) Option {
	return func(o *options) {
		o.onValidationErrCallback = onValidationErrCallback
	}
}

// WithFailFast enables fail-fast mode
func WithFailFast() Option {
	return func(o *options) {
		o.shouldFailFast = true
	}
}

// NewOptions creates new Options instance with defaults and applies provided options
func NewOptions(opts ...Option) Options {
	o := &options{
		unmarshalOptions:        protojson.UnmarshalOptions{},
		marshalOptions:          protojson.MarshalOptions{},
		errorEncoder:            DefaultEncodeError,
		responseTransformer:     DefaultTransformResponse,
		middlewares:             []gin.HandlerFunc{},
		shouldFailFast:          false,
		onValidationErrCallback: nil,
	}
	o = o.apply(opts...)
	return o
}
