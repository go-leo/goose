package gonic

import "google.golang.org/protobuf/types/known/wrapperspb"

// WrapInt32Slice converts a slice of int32 values into a slice of Int32Value wrappers.
// This is typically used for protobuf message construction where primitive types need to be wrapped.
//
// Parameters:
//   - s: The input slice of int32 values. If nil, the function returns nil.
//
// Returns:
//   - []*wrapperspb.Int32Value: A new slice containing Int32Value wrappers for each input value.
//     Returns nil if the input slice is nil.
func WrapInt32Slice(s []int32) []*wrapperspb.Int32Value {
	if s == nil {
		return nil
	}

	// Pre-allocate result slice with the same capacity as input for efficiency
	r := make([]*wrapperspb.Int32Value, 0, len(s))

	// Convert each int32 value to its corresponding Int32Value wrapper
	for _, v := range s {
		r = append(r, wrapperspb.Int32(v))
	}
	return r
}

// WrapInt64Slice converts a slice of int64 values into a slice of Int64Value wrappers.
// This is typically used for protobuf message construction where wrapper types are required.
//
// Parameters:
//   - s: The input slice of int64 values. If nil, the function returns nil.
//
// Returns:
//   - []*wrapperspb.Int64Value: A new slice containing Int64Value wrappers for each input value,
//     or nil if the input slice was nil.
func WrapInt64Slice(s []int64) []*wrapperspb.Int64Value {
	if s == nil {
		return nil
	}

	// Pre-allocate the result slice with the same capacity as input for efficiency
	r := make([]*wrapperspb.Int64Value, 0, len(s))

	// Convert each int64 value to its corresponding Int64Value wrapper
	for _, v := range s {
		r = append(r, wrapperspb.Int64(v))
	}
	return r
}

// WrapBoolSlice converts a slice of primitive bool values into a slice of BoolValue wrappers.
// This is useful for protobuf message fields that require wrapper types instead of primitive types.
//
// Parameters:
//
//	s []bool - the input slice of boolean values. If nil, the function returns nil.
//
// Returns:
//
//	[]*wrapperspb.BoolValue - a new slice containing wrapped boolean values.
//	The returned slice will be nil if the input is nil, otherwise it will contain
//	a BoolValue wrapper for each element in the input slice.
func WrapBoolSlice(s []bool) []*wrapperspb.BoolValue {
	if s == nil {
		return nil
	}

	// Preallocate the result slice with the same capacity as input for efficiency
	r := make([]*wrapperspb.BoolValue, 0, len(s))

	// Convert each boolean value to its wrapper equivalent
	for _, v := range s {
		r = append(r, wrapperspb.Bool(v))
	}

	return r
}

// WrapUint32Slice converts a slice of uint32 values into a slice of protocol buffer UInt32Value wrappers.
// This is useful for converting native Go types to their corresponding protobuf wrapper types for serialization.
//
// Parameters:
//   - s: The input slice of uint32 values. If nil, the function returns nil.
//
// Returns:
//   - []*wrapperspb.UInt32Value: A new slice containing protobuf UInt32Value wrappers for each uint32 value in the input.
//     Returns nil if the input slice is nil.
func WrapUint32Slice(s []uint32) []*wrapperspb.UInt32Value {
	if s == nil {
		return nil
	}

	// Pre-allocate the result slice with the same capacity as the input for efficiency
	r := make([]*wrapperspb.UInt32Value, 0, len(s))

	// Convert each uint32 value to its protobuf wrapper equivalent
	for _, v := range s {
		r = append(r, wrapperspb.UInt32(v))
	}
	return r
}

// WrapUint64Slice converts a slice of uint64 values into a slice of UInt64Value wrappers.
// This is typically used for protobuf message construction where wrapper types are required.
//
// Parameters:
//   - s: The input slice of uint64 values. If nil, the function returns nil.
//
// Returns:
//   - []*wrapperspb.UInt64Value: A new slice containing wrapped UInt64Value pointers.
//     Returns nil if the input slice is nil.
func WrapUint64Slice(s []uint64) []*wrapperspb.UInt64Value {
	if s == nil {
		return nil
	}

	// Pre-allocate the result slice with the same capacity as input for efficiency
	r := make([]*wrapperspb.UInt64Value, 0, len(s))

	// Convert each uint64 value to its wrapper type
	for _, v := range s {
		r = append(r, wrapperspb.UInt64(v))
	}
	return r
}

// WrapFloat32Slice converts a slice of float32 values into a slice of FloatValue wrappers.
// This is typically used for protobuf message construction where primitive types need to be wrapped.
//
// Parameters:
//
//	s []float32 - The input slice of float32 values. If nil, the function returns nil.
//
// Returns:
//
//	[]*wrapperspb.FloatValue - A new slice containing wrapped FloatValue pointers corresponding
//	                           to the input values. Returns nil if input is nil.
func WrapFloat32Slice(s []float32) []*wrapperspb.FloatValue {
	if s == nil {
		return nil
	}

	// Preallocate result slice with capacity matching input length
	r := make([]*wrapperspb.FloatValue, 0, len(s))

	// Convert each float32 value to its wrapped FloatValue counterpart
	for _, v := range s {
		r = append(r, wrapperspb.Float(v))
	}
	return r
}

// WrapFloat64Slice converts a slice of float64 values into a slice of DoubleValue wrappers.
// This is useful for protobuf message fields that require wrapped double values instead of plain float64.
//
// Parameters:
//   - s: The input slice of float64 values. If nil, the function returns nil.
//
// Returns:
//   - []*wrapperspb.DoubleValue: A new slice containing DoubleValue wrappers for each input value.
//     Returns nil if the input slice is nil.
func WrapFloat64Slice(s []float64) []*wrapperspb.DoubleValue {
	if s == nil {
		return nil
	}

	// Preallocate result slice with same capacity as input for efficiency
	r := make([]*wrapperspb.DoubleValue, 0, len(s))

	// Convert each float64 value to its DoubleValue wrapper
	for _, v := range s {
		r = append(r, wrapperspb.Double(v))
	}
	return r
}

// WrapStringSlice converts a slice of string values into a slice of StringValue wrappers.
//
// Parameters:
//   - s: The input string slice to be wrapped. If nil, the function returns nil.
//
// Returns:
//   - []*wrapperspb.StringValue: A new slice containing StringValue wrappers for each input string,
//     or nil if the input was nil.
func WrapStringSlice(s []string) []*wrapperspb.StringValue {
	if s == nil {
		return nil
	}

	// Pre-allocate result slice with capacity matching input length
	r := make([]*wrapperspb.StringValue, 0, len(s))

	// Convert each string to its StringValue wrapper
	for _, v := range s {
		r = append(r, wrapperspb.String(v))
	}
	return r
}
