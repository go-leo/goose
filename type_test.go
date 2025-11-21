// wrapper_test.go
package goose

import (
	"math"
	"net/url"
	"reflect"
	"testing"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestParseBool(t *testing.T) {
	tests := []struct {
		in    string
		want  bool
		isErr bool
	}{
		{"true", true, false},
		{"false", false, false},
		{"1", true, false},
		{"0", false, false},
		{"t", true, false},
		{"f", false, false},
		{"invalid", false, true},
	}
	for _, tt := range tests {
		got, err := ParseBool[bool](tt.in)
		if (err != nil) != tt.isErr {
			t.Errorf("ParseBool(%q) error = %v, wantErr %v", tt.in, err, tt.isErr)
		}
		if err == nil && got != tt.want {
			t.Errorf("ParseBool(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestParseInt(t *testing.T) {
	type testCase[T any] struct {
		in      string
		base    int
		bitSize int
		want    T
		isErr   bool
	}
	tests := []testCase[int64]{
		{"123", 10, 64, 123, false},
		{"-42", 10, 64, -42, false},
		{"7b", 16, 64, 123, false},
		{"invalid", 10, 64, 0, true},
	}
	for _, tt := range tests {
		got, err := ParseInt[int64](tt.in, tt.base, tt.bitSize)
		if (err != nil) != tt.isErr {
			t.Errorf("ParseInt(%q) error = %v, wantErr %v", tt.in, err, tt.isErr)
		}
		if err == nil && got != tt.want {
			t.Errorf("ParseInt(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestParseUint(t *testing.T) {
	type testCase[T any] struct {
		in      string
		base    int
		bitSize int
		want    T
		isErr   bool
	}
	tests := []testCase[uint64]{
		{"123", 10, 64, 123, false},
		{"7b", 16, 64, 123, false},
		{"-1", 10, 64, 0, true},
		{"invalid", 10, 64, 0, true},
	}
	for _, tt := range tests {
		got, err := ParseUint[uint64](tt.in, tt.base, tt.bitSize)
		if (err != nil) != tt.isErr {
			t.Errorf("ParseUint(%q) error = %v, wantErr %v", tt.in, err, tt.isErr)
		}
		if err == nil && got != tt.want {
			t.Errorf("ParseUint(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestParseFloat(t *testing.T) {
	type testCase[T any] struct {
		in      string
		bitSize int
		want    T
		isErr   bool
	}
	tests := []testCase[float64]{
		{"3.14", 64, 3.14, false},
		{"-2.5", 64, -2.5, false},
		{"invalid", 64, 0, true},
	}
	for _, tt := range tests {
		got, err := ParseFloat[float64](tt.in, tt.bitSize)
		if (err != nil) != tt.isErr {
			t.Errorf("ParseFloat(%q) error = %v, wantErr %v", tt.in, err, tt.isErr)
		}
		if err == nil && got != tt.want {
			t.Errorf("ParseFloat(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestParseBoolSlice(t *testing.T) {
	tests := []struct {
		in    []string
		want  []bool
		isErr bool
	}{
		{[]string{"true", "false", "1"}, []bool{true, false, true}, false},
		{[]string{"t", "f"}, []bool{true, false}, false},
		{[]string{"true", "invalid"}, nil, true},
		{nil, nil, false},
	}
	for _, tt := range tests {
		got, err := ParseBoolSlice[bool](tt.in)
		if (err != nil) != tt.isErr {
			t.Errorf("ParseBoolSlice(%v) error = %v, wantErr %v", tt.in, err, tt.isErr)
		}
		if err == nil && !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ParseBoolSlice(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestParseIntSlice(t *testing.T) {
	tests := []struct {
		in      []string
		base    int
		bitSize int
		want    []int64
		isErr   bool
	}{
		{[]string{"1", "2", "3"}, 10, 64, []int64{1, 2, 3}, false},
		{[]string{"7b", "2a"}, 16, 64, []int64{123, 42}, false},
		{[]string{"1", "invalid"}, 10, 64, nil, true},
		{nil, 10, 64, nil, false},
	}
	for _, tt := range tests {
		got, err := ParseIntSlice[int64](tt.in, tt.base, tt.bitSize)
		if (err != nil) != tt.isErr {
			t.Errorf("ParseIntSlice(%v) error = %v, wantErr %v", tt.in, err, tt.isErr)
		}
		if err == nil && !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ParseIntSlice(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestParseUintSlice(t *testing.T) {
	tests := []struct {
		in      []string
		base    int
		bitSize int
		want    []uint64
		isErr   bool
	}{
		{[]string{"1", "2", "3"}, 10, 64, []uint64{1, 2, 3}, false},
		{[]string{"7b", "2a"}, 16, 64, []uint64{123, 42}, false},
		{[]string{"1", "-1"}, 10, 64, nil, true},
		{nil, 10, 64, nil, false},
	}
	for _, tt := range tests {
		got, err := ParseUintSlice[uint64](tt.in, tt.base, tt.bitSize)
		if (err != nil) != tt.isErr {
			t.Errorf("ParseUintSlice(%v) error = %v, wantErr %v", tt.in, err, tt.isErr)
		}
		if err == nil && !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ParseUintSlice(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestParseFloatSlice(t *testing.T) {
	tests := []struct {
		in      []string
		bitSize int
		want    []float64
		isErr   bool
	}{
		{[]string{"1.1", "2.2", "3.3"}, 64, []float64{1.1, 2.2, 3.3}, false},
		{[]string{"1.1", "invalid"}, 64, nil, true},
		{nil, 64, nil, false},
	}
	for _, tt := range tests {
		got, err := ParseFloatSlice[float64](tt.in, tt.bitSize)
		if (err != nil) != tt.isErr {
			t.Errorf("ParseFloatSlice(%v) error = %v, wantErr %v", tt.in, err, tt.isErr)
		}
		if err == nil && !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ParseFloatSlice(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestParseBytesSlice(t *testing.T) {
	tests := []struct {
		in   []string
		want [][]byte
	}{
		{[]string{"abc", "123"}, [][]byte{[]byte("abc"), []byte("123")}},
		{[]string{}, [][]byte{}},
		{nil, nil},
	}
	for _, tt := range tests {
		got := ParseBytesSlice(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ParseBytesSlice(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

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

// TestFormatBool is a test function for FormatBool.
func TestFormatBool(t *testing.T) {
	tests := []struct {
		name string
		b    bool
		want string
	}{
		{"true", true, "true"},
		{"false", false, "false"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatBool(tt.b); got != tt.want {
				t.Errorf("FormatBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 下面是单元测试代码
func TestFormatBoolSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []bool // 使用原生bool类型，泛型参数会在测试中实例化
		expected []string
	}{
		{
			name:     "nil slice",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty slice",
			input:    []bool{},
			expected: []string{},
		},
		{
			name:     "slice with true",
			input:    []bool{true},
			expected: []string{"true"},
		},
		{
			name:     "slice with false",
			input:    []bool{false},
			expected: []string{"false"},
		},
		{
			name:     "slice with true and false",
			input:    []bool{true, false},
			expected: []string{"true", "false"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FormatBoolSlice(tt.input) // 使用测试数据调用待测函数
			if len(actual) != len(tt.expected) {
				t.Errorf("FormatBoolSlice(%v) expected length %d, actual length %d", tt.input, len(tt.expected), len(actual))
			}
			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("FormatBoolSlice(%v) expected %v, actual %v at index %d", tt.input, tt.expected, actual, i)
				}
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	form := url.Values{}
	form.Set("a", "true")
	v, err := GetBool[bool](form, "a")
	if err != nil || v != true {
		t.Errorf("GetBool(a) = %v, %v; want true, nil", v, err)
	}
	v, err = GetBool[bool](form, "notfound")
	if err != nil || v != false {
		t.Errorf("GetBool(notfound) = %v, %v; want false, nil", v, err)
	}
}

func TestGetBoolPtr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "true")
	ptr, err := GetBoolPtr[bool](form, "a")
	if err != nil || ptr == nil || *ptr != true {
		t.Errorf("GetBoolPtr(a) = %v, %v; want ptr to true, nil", ptr, err)
	}
}

func TestGetBoolSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"true", "false"}
	got, err := GetBoolSlice[bool](form, "a")
	want := []bool{true, false}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetBoolSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
	got, err = GetBoolSlice[bool](form, "notfound")
	if err != nil || got != nil {
		t.Errorf("GetBoolSlice(notfound) = %v, %v; want nil, nil", got, err)
	}
}

func TestGetInt(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	form.Set("b", "-42")
	form.Set("invalid", "abc")

	v, err := GetInt[int64](form, "a")
	if err != nil || v != 123 {
		t.Errorf("GetInt[int64](a) = %v, %v; want 123, nil", v, err)
	}
	v, err = GetInt[int64](form, "b")
	if err != nil || v != -42 {
		t.Errorf("GetInt[int64](b) = %v, %v; want -42, nil", v, err)
	}
	v, err = GetInt[int64](form, "invalid")
	if err == nil {
		t.Errorf("GetInt[int64](invalid) = %v, %v; want error", v, err)
	}
	v, err = GetInt[int64](form, "notfound")
	if err != nil || v != 0 {
		t.Errorf("GetInt[int64](notfound) = %v, %v; want 0, nil", v, err)
	}
}

func TestGetIntPtr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	ptr, err := GetIntPtr[int64](form, "a")
	if err != nil || ptr == nil || *ptr != 123 {
		t.Errorf("GetIntPtr[int64](a) = %v, %v; want ptr to 123, nil", ptr, err)
	}
	ptr, err = GetIntPtr[int64](form, "notfound")
	if err != nil || ptr == nil || *ptr != 0 {
		t.Errorf("GetIntPtr[int64](notfound) = %v, %v; want ptr to 0, nil", ptr, err)
	}
}

func TestGetIntSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1", "2", "3"}
	form["b"] = []string{"x", "2"}
	got, err := GetIntSlice[int64](form, "a")
	want := []int64{1, 2, 3}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetIntSlice[int64](a) = %v, %v; want %v, nil", got, err, want)
	}
	_, err = GetIntSlice[int64](form, "b")
	if err == nil {
		t.Errorf("GetIntSlice[int64](b) should return error")
	}
	got, err = GetIntSlice[int64](form, "notfound")
	if err != nil || got != nil {
		t.Errorf("GetIntSlice[int64](notfound) = %v, %v; want nil, nil", got, err)
	}
}

func TestGetUint(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	form.Set("invalid", "-1")
	v, err := GetUint[uint64](form, "a")
	if err != nil || v != 123 {
		t.Errorf("GetUint[uint64](a) = %v, %v; want 123, nil", v, err)
	}
	_, err = GetUint[uint64](form, "invalid")
	if err == nil {
		t.Errorf("GetUint[uint64](invalid) should return error")
	}
	v, err = GetUint[uint64](form, "notfound")
	if err != nil || v != 0 {
		t.Errorf("GetUint[uint64](notfound) = %v, %v; want 0, nil", v, err)
	}
}

func TestGetUintPtr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	ptr, err := GetUintPtr[uint64](form, "a")
	if err != nil || ptr == nil || *ptr != 123 {
		t.Errorf("GetUintPtr[uint64](a) = %v, %v; want ptr to 123, nil", ptr, err)
	}
	ptr, err = GetUintPtr[uint64](form, "notfound")
	if err != nil || ptr == nil || *ptr != 0 {
		t.Errorf("GetUintPtr[uint64](notfound) = %v, %v; want ptr to 0, nil", ptr, err)
	}
}

func TestGetUintSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1", "2", "3"}
	form["b"] = []string{"-1", "2"}
	got, err := GetUintSlice[uint64](form, "a")
	want := []uint64{1, 2, 3}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetUintSlice[uint64](a) = %v, %v; want %v, nil", got, err, want)
	}
	_, err = GetUintSlice[uint64](form, "b")
	if err == nil {
		t.Errorf("GetUintSlice[uint64](b) should return error")
	}
	got, err = GetUintSlice[uint64](form, "notfound")
	if err != nil || got != nil {
		t.Errorf("GetUintSlice[uint64](notfound) = %v, %v; want nil, nil", got, err)
	}
}

func TestGetFloat(t *testing.T) {
	form := url.Values{}
	form.Set("a", "3.14")
	form.Set("invalid", "abc")
	v, err := GetFloat[float64](form, "a")
	if err != nil || math.Abs(v-3.14) > 1e-9 {
		t.Errorf("GetFloat[float64](a) = %v, %v; want 3.14, nil", v, err)
	}
	_, err = GetFloat[float64](form, "invalid")
	if err == nil {
		t.Errorf("GetFloat[float64](invalid) should return error")
	}
	v, err = GetFloat[float64](form, "notfound")
	if err != nil || v != 0 {
		t.Errorf("GetFloat[float64](notfound) = %v, %v; want 0, nil", v, err)
	}
}

func TestGetFloatPtr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "3.14")
	ptr, err := GetFloatPtr[float64](form, "a")
	if err != nil || ptr == nil || math.Abs(*ptr-3.14) > 1e-9 {
		t.Errorf("GetFloatPtr[float64](a) = %v, %v; want ptr to 3.14, nil", ptr, err)
	}
	ptr, err = GetFloatPtr[float64](form, "notfound")
	if err != nil || ptr == nil || *ptr != 0 {
		t.Errorf("GetFloatPtr[float64](notfound) = %v, %v; want ptr to 0, nil", ptr, err)
	}
}

func TestGetFloatSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1.1", "2.2", "3.3"}
	form["b"] = []string{"x", "2.2"}
	got, err := GetFloatSlice[float64](form, "a")
	want := []float64{1.1, 2.2, 3.3}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetFloatSlice[float64](a) = %v, %v; want %v, nil", got, err, want)
	}
	_, err = GetFloatSlice[float64](form, "b")
	if err == nil {
		t.Errorf("GetFloatSlice[float64](b) should return error")
	}
}

func TestGetBoolValue(t *testing.T) {
	form := url.Values{}
	form.Set("a", "true")
	got, err := GetBoolValue(form, "a")
	want := wrapperspb.Bool(true)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetBoolValue(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetBoolValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"true", "false"}
	got, err := GetBoolValueSlice(form, "a")
	want := []*wrapperspb.BoolValue{wrapperspb.Bool(true), wrapperspb.Bool(false)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetBoolValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetInt32Value(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	got, err := GetInt32Value(form, "a")
	want := wrapperspb.Int32(123)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetInt32Value(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetInt32ValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1", "2"}
	got, err := GetInt32ValueSlice(form, "a")
	want := []*wrapperspb.Int32Value{wrapperspb.Int32(1), wrapperspb.Int32(2)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetInt32ValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetInt64Value(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123456789")
	got, err := GetInt64Value(form, "a")
	want := wrapperspb.Int64(123456789)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetInt64Value(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetInt64ValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1", "2"}
	got, err := GetInt64ValueSlice(form, "a")
	want := []*wrapperspb.Int64Value{wrapperspb.Int64(1), wrapperspb.Int64(2)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetInt64ValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetUint32Value(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	got, err := GetUint32Value(form, "a")
	want := wrapperspb.UInt32(123)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetUint32Value(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetUint32ValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1", "2"}
	got, err := GetUint32ValueSlice(form, "a")
	want := []*wrapperspb.UInt32Value{wrapperspb.UInt32(1), wrapperspb.UInt32(2)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetUint32ValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetUint64Value(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123456789")
	got, err := GetUint64Value(form, "a")
	want := wrapperspb.UInt64(123456789)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetUint64Value(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetUint64ValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1", "2"}
	got, err := GetUint64ValueSlice(form, "a")
	want := []*wrapperspb.UInt64Value{wrapperspb.UInt64(1), wrapperspb.UInt64(2)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetUint64ValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetFloat32Value(t *testing.T) {
	form := url.Values{}
	form.Set("a", "3.14")
	got, err := GetFloat32Value(form, "a")
	want := wrapperspb.Float(3.14)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetFloat32Value(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetFloat32ValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1.1", "2.2"}
	got, err := GetFloat32ValueSlice(form, "a")
	want := []*wrapperspb.FloatValue{wrapperspb.Float(1.1), wrapperspb.Float(2.2)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetFloat32ValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetFloat64Value(t *testing.T) {
	form := url.Values{}
	form.Set("a", "3.1415")
	got, err := GetFloat64Value(form, "a")
	want := wrapperspb.Double(3.1415)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetFloat64Value(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetFloat64ValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1.1", "2.2"}
	got, err := GetFloat64ValueSlice(form, "a")
	want := []*wrapperspb.DoubleValue{wrapperspb.Double(1.1), wrapperspb.Double(2.2)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetFloat64ValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}
