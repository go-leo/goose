package goose

import (
	"net/url"
	"strconv"

	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// FormatFloat converts a floating-point number to its string representation.
//
// Parameters:
//
//	f - floating-point number to convert
//	fmt - format specifier ('b', 'e', 'E', 'f', 'g', 'G', 'x', 'X')
//	prec - precision (number of digits after decimal point)
//	bitSize - bit size (32 or 64)
//
// Returns:
//
//	string - string representation of the floating-point number
func FormatFloat[Float constraints.Float](f Float, fmt byte, prec, bitSize int) string {
	return strconv.FormatFloat(float64(f), fmt, prec, bitSize)
}

// FormatFloatSlice converts a slice of floating-point numbers to a slice of their string representations.
//
// Parameters:
//
//	s - slice of floating-point numbers to convert
//	fmt - format specifier ('b', 'e', 'E', 'f', 'g', 'G', 'x', 'X')
//	prec - precision (number of digits after decimal point)
//	bitSize - bit size (32 or 64)
//
// Returns:
//
//	[]string - slice of string representations of floating-point numbers,
//	           returns nil if input slice is nil
func FormatFloatSlice[Float constraints.Float](s []Float, fmt byte, prec, bitSize int) []string {
	if s == nil {
		return nil
	}
	r := make([]string, 0, len(s))
	for _, f := range s {
		r = append(r, FormatFloat(float64(f), fmt, prec, bitSize))
	}
	return r
}

// ParseFloat converts a string to a floating-point number of the specified type.
// It wraps strconv.ParseFloat and converts the result to the generic type Float.
//
// Parameters:
//
//	s - the string to be parsed
//	bitSize - the size of the float (32 or 64)
//
// Returns:
//
//	Float - the parsed floating-point value
//	error - if parsing fails
func ParseFloat[Float constraints.Float](s string, bitSize int) (Float, error) {
	f, err := strconv.ParseFloat(s, bitSize)
	return Float(f), err
}

// ParseFloatSlice converts a slice of strings to a slice of floating-point numbers.
// Returns nil if the input slice is nil.
//
// Parameters:
//
//	s - the string slice to be parsed
//	bitSize - the size of the float (32 or 64)
//
// Returns:
//
//	[]Float - the parsed float slice
//	error - if any element fails to parse
func ParseFloatSlice[Float constraints.Float](s []string, bitSize int) ([]Float, error) {
	if s == nil {
		return nil, nil
	}
	r := make([]Float, 0, len(s))
	for _, str := range s {
		f, err := ParseFloat[Float](str, bitSize)
		if err != nil {
			return nil, err
		}
		r = append(r, f)
	}
	return r, nil
}

// GetFloat retrieves and parses a floating-point value from URL form values.
// If the key doesn't exist, returns zero value of the generic type Float.
// Uses ParseFloat with 64 bit size for conversion.
//
// Parameters:
//
//	form - the URL form values
//	key - the key to look up in form values
//
// Returns:
//
//	Float - the parsed floating-point value
//	error - if parsing fails
func GetFloat[Float constraints.Float](form url.Values, key string) (Float, error) {
	if _, ok := form[key]; !ok {
		var v Float
		return v, nil
	}
	return ParseFloat[Float](form.Get(key), 64)
}

// GetFloatPtr retrieves and parses a floating-point value from URL form values,
// returning a pointer to the value.
// If the key doesn't exist, returns zero value of the generic type Float.
//
// Parameters:
//
//	form - the URL form values
//	key - the key to look up in form values
//
// Returns:
//
//	*Float - pointer to the parsed floating-point value
//	error - if parsing fails
func GetFloatPtr[Float constraints.Float](form url.Values, key string) (*Float, error) {
	v, err := GetFloat[Float](form, key)
	return &v, err
}

// GetFloatSlice retrieves and parses a slice of floating-point numbers from URL form values.
// If the key doesn't exist, returns nil slice.
// Uses ParseFloatSlice with 64 bit size for conversion.
//
// Parameters:
//
//	form - the URL form values
//	key - the key to look up in form values
//
// Returns:
//
//	[]Float - the parsed float slice
//	error - if any element fails to parse
func GetFloatSlice[Float constraints.Float](form url.Values, key string) ([]Float, error) {
	if _, ok := form[key]; !ok {
		var v []Float
		return v, nil
	}
	return ParseFloatSlice[Float](form[key], 64)
}

// GetFloat32Value retrieves a float32 value wrapped in protobuf FloatValue.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*wrapperspb.FloatValue - protobuf wrapped float32
//	error - parsing error if any
func GetFloat32Value(form url.Values, key string) (*wrapperspb.FloatValue, error) {
	v, err := GetFloat[float32](form, key)
	return wrapperspb.Float(v), err
}

// GetFloat32ValueSlice retrieves a slice of float32 values wrapped in protobuf FloatValue.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	[]*wrapperspb.FloatValue - slice of protobuf wrapped float32s
//	error - parsing error if any
func GetFloat32ValueSlice(form url.Values, key string) ([]*wrapperspb.FloatValue, error) {
	v, err := GetFloatSlice[float32](form, key)
	return WrapFloat32Slice(v), err
}

// GetFloat64Value retrieves a float64 value wrapped in protobuf DoubleValue.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*wrapperspb.DoubleValue - protobuf wrapped float64
//	error - parsing error if any
func GetFloat64Value(form url.Values, key string) (*wrapperspb.DoubleValue, error) {
	v, err := GetFloat[float64](form, key)
	return wrapperspb.Double(v), err
}

// GetFloat64ValueSlice retrieves a slice of float64 values wrapped in protobuf DoubleValue.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	[]*wrapperspb.DoubleValue - slice of protobuf wrapped float64s
//	error - parsing error if any
func GetFloat64ValueSlice(form url.Values, key string) ([]*wrapperspb.DoubleValue, error) {
	v, err := GetFloatSlice[float64](form, key)
	return WrapFloat64Slice(v), err
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
