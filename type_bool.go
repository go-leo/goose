package goose

import (
	"net/url"
	"strconv"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// FormatBool converts a boolean value to its string representation.
//
// Parameters:
//
//	b - boolean value to convert
//
// Returns:
//
//	string - string representation of the boolean value
func FormatBool[Bool ~bool](b Bool) string {
	return strconv.FormatBool(bool(b))
}

// FormatBoolSlice converts a slice of boolean values to a slice of their string representations.
//
// Parameters:
//
//	s - slice of boolean values to convert
//
// Returns:
//
//	[]string - slice of string representations of boolean values,
//	           returns nil if input slice is nil
func FormatBoolSlice[Bool ~bool](s []Bool) []string {
	if s == nil {
		return nil
	}
	r := make([]string, 0, len(s))
	for _, b := range s {
		r = append(r, FormatBool(b))
	}
	return r
}

// ParseBool converts a string to a boolean value of the specified type.
// It wraps strconv.ParseBool and converts the result to the generic type Bool.
//
// Parameters:
//   - s: the string to be parsed
//
// Returns:
//   - Bool: the parsed boolean value
//   - error: if parsing fails
func ParseBool[Bool ~bool](s string) (Bool, error) {
	v, err := strconv.ParseBool(s)
	return Bool(v), err
}

// ParseBoolSlice converts a slice of strings to a slice of booleans.
// Returns nil if the input slice is nil.
//
// Parameters:
//   - s: the string slice to be parsed
//
// Returns:
//   - []Bool: the parsed boolean slice
//   - error: if any element fails to parse
func ParseBoolSlice[Bool ~bool](s []string) ([]Bool, error) {
	if s == nil {
		return nil, nil
	}
	r := make([]Bool, 0, len(s))
	for _, str := range s {
		b, err := ParseBool[Bool](str)
		if err != nil {
			return nil, err
		}
		r = append(r, b)
	}
	return r, nil
}

// GetBool retrieves and parses a boolean value from URL form values.
// If the key doesn't exist, returns false.
//
// Parameters:
//   - form: the URL form values
//   - key: the key to look up in form values
//
// Returns:
//   - Bool: the parsed boolean value
//   - error: if parsing fails
func GetBool[Bool ~bool](form url.Values, key string) (Bool, error) {
	if _, ok := form[key]; !ok {
		return false, nil
	}
	return ParseBool[Bool](form.Get(key))
}

// GetBoolPtr retrieves and parses a boolean value from URL form values,
// returning a pointer to the value.
// If the key doesn't exist, returns a pointer to false.
//
// Parameters:
//   - form: the URL form values
//   - key: the key to look up in form values
//
// Returns:
//   - *Bool: pointer to the parsed boolean value
//   - error: if parsing fails
func GetBoolPtr[Bool ~bool](form url.Values, key string) (*Bool, error) {
	v, err := GetBool[Bool](form, key)
	return &v, err
}

// GetBoolSlice retrieves and parses a slice of booleans from URL form values.
// If the key doesn't exist, returns nil slice.
//
// Parameters:
//   - form: the URL form values
//   - key: the key to look up in form values
//
// Returns:
//   - []Bool: the parsed boolean slice
//   - error: if any element fails to parse
func GetBoolSlice[Bool ~bool](form url.Values, key string) ([]Bool, error) {
	if _, ok := form[key]; !ok {
		return nil, nil
	}
	return ParseBoolSlice[Bool](form[key])
}

// GetBoolValue retrieves a boolean value wrapped in protobuf BoolValue.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*wrapperspb.BoolValue - protobuf wrapped boolean
//	error - parsing error if any
func GetBoolValue(form url.Values, key string) (*wrapperspb.BoolValue, error) {
	v, err := strconv.ParseBool(form.Get(key))
	return wrapperspb.Bool(v), err
}

// GetBoolValueSlice retrieves a slice of boolean values wrapped in protobuf BoolValue.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	[]*wrapperspb.BoolValue - slice of protobuf wrapped booleans
//	error - parsing error if any
func GetBoolValueSlice(form url.Values, key string) ([]*wrapperspb.BoolValue, error) {
	v, err := ParseBoolSlice[bool](form[key])
	return WrapBoolSlice(v), err
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
