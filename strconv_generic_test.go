package gonic

import (
	"reflect"
	"testing"
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
		got, err := ParseBool(tt.in)
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
		got, err := ParseBoolSlice(tt.in)
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
