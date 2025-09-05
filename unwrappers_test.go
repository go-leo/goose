// unwrappers_test.go
package goose

import (
	"reflect"
	"testing"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUnwrapBoolSlice(t *testing.T) {
	tests := []struct {
		name string
		s    []*wrapperspb.BoolValue
		want []bool
	}{
		{
			name: "nil slice",
			s:    nil,
			want: nil,
		},
		{
			name: "empty slice",
			s:    []*wrapperspb.BoolValue{},
			want: []bool{},
		},
		{
			name: "single value",
			s:    []*wrapperspb.BoolValue{wrapperspb.Bool(true)},
			want: []bool{true},
		},
		{
			name: "multiple values",
			s:    []*wrapperspb.BoolValue{wrapperspb.Bool(true), wrapperspb.Bool(false), wrapperspb.Bool(true)},
			want: []bool{true, false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnwrapBoolSlice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapBoolSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnwrapInt32Slice(t *testing.T) {
	tests := []struct {
		name string
		s    []*wrapperspb.Int32Value
		want []int32
	}{
		{
			name: "nil slice",
			s:    nil,
			want: nil,
		},
		{
			name: "empty slice",
			s:    []*wrapperspb.Int32Value{},
			want: []int32{},
		},
		{
			name: "single value",
			s:    []*wrapperspb.Int32Value{wrapperspb.Int32(42)},
			want: []int32{42},
		},
		{
			name: "multiple values",
			s:    []*wrapperspb.Int32Value{wrapperspb.Int32(1), wrapperspb.Int32(2), wrapperspb.Int32(3)},
			want: []int32{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnwrapInt32Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapInt32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnwrapInt64Slice(t *testing.T) {
	tests := []struct {
		name string
		s    []*wrapperspb.Int64Value
		want []int64
	}{
		{
			name: "nil slice",
			s:    nil,
			want: nil,
		},
		{
			name: "empty slice",
			s:    []*wrapperspb.Int64Value{},
			want: []int64{},
		},
		{
			name: "single value",
			s:    []*wrapperspb.Int64Value{wrapperspb.Int64(42)},
			want: []int64{42},
		},
		{
			name: "multiple values",
			s:    []*wrapperspb.Int64Value{wrapperspb.Int64(1), wrapperspb.Int64(2), wrapperspb.Int64(3)},
			want: []int64{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnwrapInt64Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapInt64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnwrapFloat32Slice(t *testing.T) {
	tests := []struct {
		name string
		s    []*wrapperspb.FloatValue
		want []float32
	}{
		{
			name: "nil slice",
			s:    nil,
			want: nil,
		},
		{
			name: "empty slice",
			s:    []*wrapperspb.FloatValue{},
			want: []float32{},
		},
		{
			name: "single value",
			s:    []*wrapperspb.FloatValue{wrapperspb.Float(3.14)},
			want: []float32{3.14},
		},
		{
			name: "multiple values",
			s:    []*wrapperspb.FloatValue{wrapperspb.Float(1.1), wrapperspb.Float(2.2), wrapperspb.Float(3.3)},
			want: []float32{1.1, 2.2, 3.3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnwrapFloat32Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapFloat32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnwrapFloat64Slice(t *testing.T) {
	tests := []struct {
		name string
		s    []*wrapperspb.DoubleValue
		want []float64
	}{
		{
			name: "nil slice",
			s:    nil,
			want: nil,
		},
		{
			name: "empty slice",
			s:    []*wrapperspb.DoubleValue{},
			want: []float64{},
		},
		{
			name: "single value",
			s:    []*wrapperspb.DoubleValue{wrapperspb.Double(3.14159)},
			want: []float64{3.14159},
		},
		{
			name: "multiple values",
			s:    []*wrapperspb.DoubleValue{wrapperspb.Double(1.1), wrapperspb.Double(2.2), wrapperspb.Double(3.3)},
			want: []float64{1.1, 2.2, 3.3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnwrapFloat64Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapFloat64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnwrapUint32Slice(t *testing.T) {
	tests := []struct {
		name string
		s    []*wrapperspb.UInt32Value
		want []uint32
	}{
		{
			name: "nil slice",
			s:    nil,
			want: nil,
		},
		{
			name: "empty slice",
			s:    []*wrapperspb.UInt32Value{},
			want: []uint32{},
		},
		{
			name: "single value",
			s:    []*wrapperspb.UInt32Value{wrapperspb.UInt32(42)},
			want: []uint32{42},
		},
		{
			name: "multiple values",
			s:    []*wrapperspb.UInt32Value{wrapperspb.UInt32(1), wrapperspb.UInt32(2), wrapperspb.UInt32(3)},
			want: []uint32{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnwrapUint32Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapUint32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnwrapUint64Slice(t *testing.T) {
	tests := []struct {
		name string
		s    []*wrapperspb.UInt64Value
		want []uint64
	}{
		{
			name: "nil slice",
			s:    nil,
			want: nil,
		},
		{
			name: "empty slice",
			s:    []*wrapperspb.UInt64Value{},
			want: []uint64{},
		},
		{
			name: "single value",
			s:    []*wrapperspb.UInt64Value{wrapperspb.UInt64(42)},
			want: []uint64{42},
		},
		{
			name: "multiple values",
			s:    []*wrapperspb.UInt64Value{wrapperspb.UInt64(1), wrapperspb.UInt64(2), wrapperspb.UInt64(3)},
			want: []uint64{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnwrapUint64Slice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapUint64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnwrapStringSlice(t *testing.T) {
	tests := []struct {
		name string
		s    []*wrapperspb.StringValue
		want []string
	}{
		{
			name: "nil slice",
			s:    nil,
			want: nil,
		},
		{
			name: "empty slice",
			s:    []*wrapperspb.StringValue{},
			want: []string{},
		},
		{
			name: "single value",
			s:    []*wrapperspb.StringValue{wrapperspb.String("hello")},
			want: []string{"hello"},
		},
		{
			name: "multiple values",
			s:    []*wrapperspb.StringValue{wrapperspb.String("a"), wrapperspb.String("b"), wrapperspb.String("c")},
			want: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnwrapStringSlice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnwrapBytesSlice(t *testing.T) {
	tests := []struct {
		name string
		s    []*wrapperspb.BytesValue
		want [][]byte
	}{
		{
			name: "nil slice",
			s:    nil,
			want: nil,
		},
		{
			name: "empty slice",
			s:    []*wrapperspb.BytesValue{},
			want: [][]byte{},
		},
		{
			name: "single value",
			s:    []*wrapperspb.BytesValue{wrapperspb.Bytes([]byte("hello"))},
			want: [][]byte{[]byte("hello")},
		},
		{
			name: "multiple values",
			s:    []*wrapperspb.BytesValue{wrapperspb.Bytes([]byte("a")), wrapperspb.Bytes([]byte("b")), wrapperspb.Bytes([]byte("c"))},
			want: [][]byte{[]byte("a"), []byte("b"), []byte("c")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnwrapBytesSlice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapBytesSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
