package goose

import "google.golang.org/protobuf/types/known/wrapperspb"

func UnwrapBoolSlice(s []*wrapperspb.BoolValue) []bool {
	if s == nil {
		return nil
	}
	r := make([]bool, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func UnwrapInt32Slice(s []*wrapperspb.Int32Value) []int32 {
	if s == nil {
		return nil
	}
	r := make([]int32, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func UnwrapInt64Slice(s []*wrapperspb.Int64Value) []int64 {
	if s == nil {
		return nil
	}
	r := make([]int64, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func UnwrapFloat32Slice(s []*wrapperspb.FloatValue) []float32 {
	if s == nil {
		return nil
	}
	r := make([]float32, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func UnwrapFloat64Slice(s []*wrapperspb.DoubleValue) []float64 {
	if s == nil {
		return nil
	}
	r := make([]float64, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func UnwrapUint32Slice(s []*wrapperspb.UInt32Value) []uint32 {
	if s == nil {
		return nil
	}
	r := make([]uint32, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func UnwrapUint64Slice(s []*wrapperspb.UInt64Value) []uint64 {
	if s == nil {
		return nil
	}
	r := make([]uint64, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func UnwrapStringSlice(s []*wrapperspb.StringValue) []string {
	if s == nil {
		return nil
	}
	r := make([]string, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func UnwrapBytesSlice(s []*wrapperspb.BytesValue) [][]byte {
	if s == nil {
		return nil
	}
	r := make([][]byte, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}
