package goose

import "google.golang.org/protobuf/types/known/wrapperspb"

// UnwrapBoolSlice converts a slice of BoolValue wrappers into a slice of primitive bool values.
// This is typically used for extracting primitive values from protobuf messages that use wrapper types.
//
// Parameters:
//   - s: The input slice of BoolValue wrappers. If nil, the function returns nil.
//
// Returns:
//   - []bool: A new slice containing unwrapped bool values for each input wrapper.
//     Returns nil if the input slice is nil.
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

// UnwrapInt32Slice converts a slice of Int32Value wrappers into a slice of primitive int32 values.
// This is typically used for extracting primitive values from protobuf messages that use wrapper types.
//
// Parameters:
//   - s: The input slice of Int32Value wrappers. If nil, the function returns nil.
//
// Returns:
//   - []int32: A new slice containing unwrapped int32 values for each input wrapper.
//     Returns nil if the input slice is nil.
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

// UnwrapInt64Slice converts a slice of Int64Value wrappers into a slice of primitive int64 values.
// This is typically used for extracting primitive values from protobuf messages that use wrapper types.
//
// Parameters:
//   - s: The input slice of Int64Value wrappers. If nil, the function returns nil.
//
// Returns:
//   - []int64: A new slice containing unwrapped int64 values for each input wrapper.
//     Returns nil if the input slice is nil.
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

// UnwrapFloat32Slice converts a slice of FloatValue wrappers into a slice of primitive float32 values.
// This is typically used for extracting primitive values from protobuf messages that use wrapper types.
//
// Parameters:
//   - s: The input slice of FloatValue wrappers. If nil, the function returns nil.
//
// Returns:
//   - []float32: A new slice containing unwrapped float32 values for each input wrapper.
//     Returns nil if the input slice is nil.
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

// UnwrapFloat64Slice converts a slice of DoubleValue wrappers into a slice of primitive float64 values.
// This is typically used for extracting primitive values from protobuf messages that use wrapper types.
//
// Parameters:
//   - s: The input slice of DoubleValue wrappers. If nil, the function returns nil.
//
// Returns:
//   - []float64: A new slice containing unwrapped float64 values for each input wrapper.
//     Returns nil if the input slice is nil.
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

// UnwrapUint32Slice converts a slice of UInt32Value wrappers into a slice of primitive uint32 values.
// This is typically used for extracting primitive values from protobuf messages that use wrapper types.
//
// Parameters:
//   - s: The input slice of UInt32Value wrappers. If nil, the function returns nil.
//
// Returns:
//   - []uint32: A new slice containing unwrapped uint32 values for each input wrapper.
//     Returns nil if the input slice is nil.
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

// UnwrapUint64Slice converts a slice of UInt64Value wrappers into a slice of primitive uint64 values.
// This is typically used for extracting primitive values from protobuf messages that use wrapper types.
//
// Parameters:
//   - s: The input slice of UInt64Value wrappers. If nil, the function returns nil.
//
// Returns:
//   - []uint64: A new slice containing unwrapped uint64 values for each input wrapper.
//     Returns nil if the input slice is nil.
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

// UnwrapStringSlice converts a slice of StringValue wrappers into a slice of primitive string values.
// This is typically used for extracting primitive values from protobuf messages that use wrapper types.
//
// Parameters:
//   - s: The input slice of StringValue wrappers. If nil, the function returns nil.
//
// Returns:
//   - []string: A new slice containing unwrapped string values for each input wrapper.
//     Returns nil if the input slice is nil.
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

// UnwrapBytesSlice converts a slice of BytesValue wrappers into a slice of primitive byte slice values.
// This is typically used for extracting primitive values from protobuf messages that use wrapper types.
//
// Parameters:
//   - s: The input slice of BytesValue wrappers. If nil, the function returns nil.
//
// Returns:
//   - [][]byte: A new slice containing unwrapped byte slice values for each input wrapper.
//     Returns nil if the input slice is nil.
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
