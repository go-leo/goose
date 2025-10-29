package goose

import (
	"net/http"
	"net/url"
	"strconv"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

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

// FormFromPath extracts specified key-value pairs from HTTP request path parameters
// and constructs them into url.Values format.
//
// Parameters:
//
//	r: HTTP request object used to retrieve path parameters
//	keys: list of path parameter key names to extract
//
// Returns:
//
//	url.Values: form data containing specified path parameter key-value pairs,
//	            returns nil if keys is nil
func FormFromPath(r *http.Request, keys ...string) url.Values {
	if keys == nil {
		return nil
	}
	form := url.Values{}
	for _, key := range keys {
		form.Add(key, r.PathValue(key))
	}
	return form
}
