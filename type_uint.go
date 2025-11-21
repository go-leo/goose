package goose

import (
	"net/url"
	"strconv"

	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// FormatUint converts an unsigned integer to a string representation in a specified base.
//
// Parameters:
//
//	i - unsigned integer to convert
//	base - base for string representation (2-36)
//
// Returns:
//
//	string - string representation of the unsigned integer
func FormatUint[Unsigned constraints.Unsigned](i Unsigned, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

// FormatUintSlice converts a slice of unsigned integers to a slice of their string representations
// in a specified base.
//
// Parameters:
//
//	s - slice of unsigned integers to convert
//	base - base for string representation (2-36)
//
// Returns:
//
//	[]string - slice of string representations of unsigned integers,
//	           returns nil if input slice is nil
func FormatUintSlice[Unsigned constraints.Unsigned](s []Unsigned, base int) []string {
	if s == nil {
		return nil
	}
	r := make([]string, 0, len(s))
	for _, i := range s {
		r = append(r, FormatUint(i, base))
	}
	return r
}

// ParseUint converts a string to an unsigned integer of the specified type.
// It wraps strconv.ParseUint and converts the result to the generic type Unsigned.
//
// Parameters:
//
//	s - the string to be parsed
//	base - the base for conversion (0, 2 to 36)
//	bitSize - the size of the integer (0, 8, 16, 32, 64)
//
// Returns:
//
//	Unsigned - the parsed unsigned integer value
//	error - if parsing fails
func ParseUint[Unsigned constraints.Unsigned](s string, base int, bitSize int) (Unsigned, error) {
	i, err := strconv.ParseUint(s, base, bitSize)
	return Unsigned(i), err
}

// ParseUintSlice converts a slice of strings to a slice of unsigned integers.
// Returns nil if the input slice is nil.
//
// Parameters:
//
//	s - the string slice to be parsed
//	base - the base for conversion (0, 2 to 36)
//	bitSize - the size of the integer (0, 8, 16, 32, 64)
//
// Returns:
//
//	[]Unsigned - the parsed unsigned integer slice
//	error - if any element fails to parse
func ParseUintSlice[Unsigned constraints.Unsigned](s []string, base int, bitSize int) ([]Unsigned, error) {
	if s == nil {
		return nil, nil
	}
	r := make([]Unsigned, 0, len(s))
	for _, str := range s {
		i, err := ParseUint[Unsigned](str, base, bitSize)
		if err != nil {
			return nil, err
		}
		r = append(r, i)
	}
	return r, nil
}

// GetUint retrieves and parses an unsigned integer value from URL form values.
// If the key doesn't exist, returns zero value of the generic type Unsigned.
// Uses ParseUint with base 10 and 64 bit size for conversion.
//
// Parameters:
//
//	form - the URL form values
//	key - the key to look up in form values
//
// Returns:
//
//	Unsigned - the parsed unsigned integer value
//	error - if parsing fails
func GetUint[Unsigned constraints.Unsigned](form url.Values, key string) (Unsigned, error) {
	if _, ok := form[key]; !ok {
		var v Unsigned
		return v, nil
	}
	return ParseUint[Unsigned](form.Get(key), 10, 64)
}

// GetUintPtr retrieves and parses an unsigned integer value from URL form values,
// returning a pointer to the value.
// If the key doesn't exist, returns zero value of the generic type Unsigned.
//
// Parameters:
//
//	form - the URL form values
//	key - the key to look up in form values
//
// Returns:
//
//	*Unsigned - pointer to the parsed unsigned integer value
//	error - if parsing fails
func GetUintPtr[Unsigned constraints.Unsigned](form url.Values, key string) (*Unsigned, error) {
	v, err := GetUint[Unsigned](form, key)
	return &v, err
}

// GetUintSlice retrieves and parses a slice of unsigned integers from URL form values.
// If the key doesn't exist, returns nil slice.
// Uses ParseUintSlice with base 10 and 64 bit size for conversion.
//
// Parameters:
//
//	form - the URL form values
//	key - the key to look up in form values
//
// Returns:
//
//	[]Unsigned - the parsed unsigned integer slice
//	error - if any element fails to parse
func GetUintSlice[Unsigned constraints.Unsigned](form url.Values, key string) ([]Unsigned, error) {
	if _, ok := form[key]; !ok {
		var v []Unsigned
		return v, nil
	}
	return ParseUintSlice[Unsigned](form[key], 10, 64)
}

// GetUint32Value retrieves a uint32 value wrapped in protobuf UInt32Value.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*wrapperspb.UInt32Value - protobuf wrapped uint32
//	error - parsing error if any
func GetUint32Value(form url.Values, key string) (*wrapperspb.UInt32Value, error) {
	v, err := GetUint[uint32](form, key)
	return wrapperspb.UInt32(v), err
}

// GetUint32ValueSlice retrieves a slice of uint32 values wrapped in protobuf UInt32Value.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	[]*wrapperspb.UInt32Value - slice of protobuf wrapped uint32s
//	error - parsing error if any
func GetUint32ValueSlice(form url.Values, key string) ([]*wrapperspb.UInt32Value, error) {
	v, err := GetUintSlice[uint32](form, key)
	return WrapUint32Slice(v), err
}

// GetUint64Value retrieves a uint64 value wrapped in protobuf UInt64Value.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*wrapperspb.UInt64Value - protobuf wrapped uint64
//	error - parsing error if any
func GetUint64Value(form url.Values, key string) (*wrapperspb.UInt64Value, error) {
	v, err := GetUint[uint64](form, key)
	return wrapperspb.UInt64(v), err
}

// GetUint64ValueSlice retrieves a slice of uint64 values wrapped in protobuf UInt64Value.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	[]*wrapperspb.UInt64Value - slice of protobuf wrapped uint64s
//	error - parsing error if any
func GetUint64ValueSlice(form url.Values, key string) ([]*wrapperspb.UInt64Value, error) {
	v, err := GetUintSlice[uint64](form, key)
	return WrapUint64Slice(v), err
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
