// wrapper_test.go
package goose

import (
	"reflect"
	"testing"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestWrapInt32Slice(t *testing.T) {
	tests := []struct {
		name string
		in   []int32
		want []*wrapperspb.Int32Value
	}{
		{"nil", nil, nil},
		{"empty", []int32{}, []*wrapperspb.Int32Value{}},
		{"values", []int32{1, -2, 3}, []*wrapperspb.Int32Value{
			wrapperspb.Int32(1), wrapperspb.Int32(-2), wrapperspb.Int32(3),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WrapInt32Slice(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapInt32Slice(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestWrapInt64Slice(t *testing.T) {
	tests := []struct {
		name string
		in   []int64
		want []*wrapperspb.Int64Value
	}{
		{"nil", nil, nil},
		{"empty", []int64{}, []*wrapperspb.Int64Value{}},
		{"values", []int64{1, -2, 3}, []*wrapperspb.Int64Value{
			wrapperspb.Int64(1), wrapperspb.Int64(-2), wrapperspb.Int64(3),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WrapInt64Slice(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapInt64Slice(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestWrapBoolSlice(t *testing.T) {
	tests := []struct {
		name string
		in   []bool
		want []*wrapperspb.BoolValue
	}{
		{"nil", nil, nil},
		{"empty", []bool{}, []*wrapperspb.BoolValue{}},
		{"values", []bool{true, false, true}, []*wrapperspb.BoolValue{
			wrapperspb.Bool(true), wrapperspb.Bool(false), wrapperspb.Bool(true),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WrapBoolSlice(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapBoolSlice(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestWrapUint32Slice(t *testing.T) {
	tests := []struct {
		name string
		in   []uint32
		want []*wrapperspb.UInt32Value
	}{
		{"nil", nil, nil},
		{"empty", []uint32{}, []*wrapperspb.UInt32Value{}},
		{"values", []uint32{1, 2, 3}, []*wrapperspb.UInt32Value{
			wrapperspb.UInt32(1), wrapperspb.UInt32(2), wrapperspb.UInt32(3),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WrapUint32Slice(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapUint32Slice(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestWrapUint64Slice(t *testing.T) {
	tests := []struct {
		name string
		in   []uint64
		want []*wrapperspb.UInt64Value
	}{
		{"nil", nil, nil},
		{"empty", []uint64{}, []*wrapperspb.UInt64Value{}},
		{"values", []uint64{1, 2, 3}, []*wrapperspb.UInt64Value{
			wrapperspb.UInt64(1), wrapperspb.UInt64(2), wrapperspb.UInt64(3),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WrapUint64Slice(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapUint64Slice(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestWrapFloat32Slice(t *testing.T) {
	tests := []struct {
		name string
		in   []float32
		want []*wrapperspb.FloatValue
	}{
		{"nil", nil, nil},
		{"empty", []float32{}, []*wrapperspb.FloatValue{}},
		{"values", []float32{1.1, -2.2, 3.3}, []*wrapperspb.FloatValue{
			wrapperspb.Float(1.1), wrapperspb.Float(-2.2), wrapperspb.Float(3.3),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WrapFloat32Slice(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapFloat32Slice(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestWrapFloat64Slice(t *testing.T) {
	tests := []struct {
		name string
		in   []float64
		want []*wrapperspb.DoubleValue
	}{
		{"nil", nil, nil},
		{"empty", []float64{}, []*wrapperspb.DoubleValue{}},
		{"values", []float64{1.1, -2.2, 3.3}, []*wrapperspb.DoubleValue{
			wrapperspb.Double(1.1), wrapperspb.Double(-2.2), wrapperspb.Double(3.3),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WrapFloat64Slice(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapFloat64Slice(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
