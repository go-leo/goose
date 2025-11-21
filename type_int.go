package goose

import (
	"net/url"
	"strconv"

	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// FormatInt converts a signed integer to a string representation in a specified base.
//
// Parameters:
//
//	i - signed integer to convert
//	base - base for string representation (2-36)
//
// Returns:
//
//	string - string representation of the signed integer
func FormatInt[Signed constraints.Signed](i Signed, base int) string {
	return strconv.FormatInt(int64(i), base)
}

// FormatIntSlice converts a slice of signed integers to a slice of their string representations
// in a specified base.
//
// Parameters:
//
//	s - slice of signed integers to convert
//	base - base for string representation (2-36)
//
// Returns:
//
//	[]string - slice of string representations of signed integers,
//	           returns nil if input slice is nil
func FormatIntSlice[Signed constraints.Signed](s []Signed, base int) []string {
	if s == nil {
		return nil
	}
	r := make([]string, 0, len(s))
	for _, i := range s {
		r = append(r, FormatInt(i, base))
	}
	return r
}

// ParseInt converts a string to a signed integer of the specified type.
// It wraps strconv.ParseInt and converts the result to the generic type Signed.
//
// Parameters:
//
//	s - the string to be parsed
//	base - the base for conversion (0, 2 to 36)
//	bitSize - the size of the integer (0, 8, 16, 32, 64)
//
// Returns:
//
//	Signed - the parsed integer value
//	error - if parsing fails
func ParseInt[Signed constraints.Signed](s string, base int, bitSize int) (Signed, error) {
	i, err := strconv.ParseInt(s, base, bitSize)
	return Signed(i), err
}

// ParseIntSlice converts a slice of strings to a slice of signed integers.
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
//	[]Signed - the parsed integer slice
//	error - if any element fails to parse
func ParseIntSlice[Signed constraints.Signed](s []string, base int, bitSize int) ([]Signed, error) {
	if s == nil {
		return nil, nil
	}
	r := make([]Signed, 0, len(s))
	for _, str := range s {
		i, err := ParseInt[Signed](str, base, bitSize)
		if err != nil {
			return nil, err
		}
		r = append(r, i)
	}
	return r, nil
}

// GetInt retrieves and parses a signed integer value from URL form values.
// If the key doesn't exist, returns zero value of the generic type Signed.
// Uses ParseInt with base 10 and 64 bit size for conversion.
//
// Parameters:
//
//	form - the URL form values
//	key - the key to look up in form values
//
// Returns:
//
//	Signed - the parsed integer value
//	error - if parsing fails
func GetInt[Signed constraints.Signed](form url.Values, key string) (Signed, error) {
	if _, ok := form[key]; !ok {
		var v Signed
		return v, nil
	}
	return ParseInt[Signed](form.Get(key), 10, 64)
}

// GetIntPtr retrieves and parses a signed integer value from URL form values,
// returning a pointer to the value.
// If the key doesn't exist, returns zero value of the generic type Signed.
//
// Parameters:
//
//	form - the URL form values
//	key - the key to look up in form values
//
// Returns:
//
//	*Signed - pointer to the parsed integer value
//	error - if parsing fails
func GetIntPtr[Signed constraints.Signed](form url.Values, key string) (*Signed, error) {
	v, err := GetInt[Signed](form, key)
	return &v, err
}

// GetIntSlice retrieves and parses a slice of signed integers from URL form values.
// If the key doesn't exist, returns nil slice.
// Uses ParseIntSlice with base 10 and 64 bit size for conversion.
//
// Parameters:
//
//	form - the URL form values
//	key - the key to look up in form values
//
// Returns:
//
//	[]Signed - the parsed integer slice
//	error - if any element fails to parse
func GetIntSlice[Signed constraints.Signed](form url.Values, key string) ([]Signed, error) {
	if _, ok := form[key]; !ok {
		var v []Signed
		return v, nil
	}
	return ParseIntSlice[Signed](form[key], 10, 64)
}

// GetInt32Value retrieves an int32 value wrapped in protobuf Int32Value.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*wrapperspb.Int32Value - protobuf wrapped int32
//	error - parsing error if any
func GetInt32Value(form url.Values, key string) (*wrapperspb.Int32Value, error) {
	v, err := GetInt[int32](form, key)
	return wrapperspb.Int32(v), err
}

// GetInt32ValueSlice retrieves a slice of int32 values wrapped in protobuf Int32Value.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	[]*wrapperspb.Int32Value - slice of protobuf wrapped int32s
//	error - parsing error if any
func GetInt32ValueSlice(form url.Values, key string) ([]*wrapperspb.Int32Value, error) {
	v, err := GetIntSlice[int32](form, key)
	return WrapInt32Slice(v), err
}

// GetInt64Value retrieves an int64 value wrapped in protobuf Int64Value.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*wrapperspb.Int64Value - protobuf wrapped int64
//	error - parsing error if any
func GetInt64Value(form url.Values, key string) (*wrapperspb.Int64Value, error) {
	v, err := GetInt[int64](form, key)
	return wrapperspb.Int64(v), err
}

// GetInt64ValueSlice retrieves a slice of int64 values wrapped in protobuf Int64Value.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	[]*wrapperspb.Int64Value - slice of protobuf wrapped int64s
//	error - parsing error if any
func GetInt64ValueSlice(form url.Values, key string) ([]*wrapperspb.Int64Value, error) {
	v, err := GetIntSlice[int64](form, key)
	return WrapInt64Slice(v), err
}

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
