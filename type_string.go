package goose

import "google.golang.org/protobuf/types/known/wrapperspb"

// ParseBytesSlice converts a slice of strings to a slice of byte slices.
// Returns nil if the input slice is nil.
//
// Parameters:
//
//	s - the string slice to be converted
//
// Returns:
//
//	[][]byte - the resulting byte slice
func ParseBytesSlice(s []string) [][]byte {
	if s == nil {
		return nil
	}
	r := make([][]byte, 0, len(s))
	for _, str := range s {
		r = append(r, []byte(str))
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
