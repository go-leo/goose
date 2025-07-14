package gonic

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func dummyErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
}

func dummyResponseTransformer(ctx context.Context, resp proto.Message) proto.Message {
	return nil
}

func TestOptions_Defaults(t *testing.T) {
	opts := NewOptions()

	// 检查默认值
	if !reflect.DeepEqual(opts.UnmarshalOptions(), protojson.UnmarshalOptions{}) {
		t.Errorf("default UnmarshalOptions not empty")
	}
	if !reflect.DeepEqual(opts.MarshalOptions(), protojson.MarshalOptions{}) {
		t.Errorf("default MarshalOptions not empty")
	}
	if opts.ErrorEncoder() == nil {
		t.Errorf("default ErrorEncoder is nil")
	}
	if opts.ResponseTransformer() == nil {
		t.Errorf("default ResponseTransformer is nil")
	}
}

func TestOptions_WithOptions(t *testing.T) {
	unmarshalOpt := protojson.UnmarshalOptions{AllowPartial: true}
	marshalOpt := protojson.MarshalOptions{EmitUnpopulated: true}
	var customErrorEncoder ErrorEncoder = dummyErrorEncoder
	var customResponseTransformer ResponseTransformer = dummyResponseTransformer

	opts := NewOptions(
		WithUnmarshalOptions(unmarshalOpt),
		WithMarshalOptions(marshalOpt),
		WithErrorEncoder(customErrorEncoder),
		WithResponseTransformer(customResponseTransformer),
	)

	if !reflect.DeepEqual(opts.UnmarshalOptions(), unmarshalOpt) {
		t.Errorf("UnmarshalOptions not set correctly")
	}
	if !reflect.DeepEqual(opts.MarshalOptions(), marshalOpt) {
		t.Errorf("MarshalOptions not set correctly")
	}
	// if opts.ErrorEncoder() != customErrorEncoder {
	// 	t.Errorf("ErrorEncoder not set correctly")
	// }
	// if opts.ResponseTransformer() != customResponseTransformer {
	// 	t.Errorf("ResponseTransformer not set correctly")
	// }
}

func TestOptions_Apply(t *testing.T) {
	o := &options{}
	opt1 := WithUnmarshalOptions(protojson.UnmarshalOptions{DiscardUnknown: true})
	opt2 := WithMarshalOptions(protojson.MarshalOptions{UseProtoNames: true})
	o.apply(opt1, opt2)

	if !o.unmarshalOptions.DiscardUnknown {
		t.Errorf("Apply did not set UnmarshalOptions.DiscardUnknown")
	}
	if !o.marshalOptions.UseProtoNames {
		t.Errorf("Apply did not set MarshalOptions.UseProtoNames")
	}
}
