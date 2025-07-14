package gonic

import (
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// GetInt32 retrieves an int32 value from URL form values.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	int32 - parsed integer value
//	error - parsing error if any
func GetInt32(form url.Values, key string) (int32, error) {
	return GetInt[int32](form, key)
}

// GetInt32Slice retrieves a slice of int32 values from the given url.Values by key.
//
// Parameters:
//   - form: url.Values containing the form data to retrieve values from
//   - key:  string key to look up in the form values
//
// Returns:
//   - []int32: slice of int32 values if successful
//   - error:   any error that occurred during parsing or retrieval
func GetInt32Slice(form url.Values, key string) ([]int32, error) {
	return GetIntSlice[int32](form, key)
}

// GetInt32Ptr retrieves an int32 value from form and returns its pointer.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*int32 - pointer to parsed value
//	error - parsing error if any
func GetInt32Ptr(form url.Values, key string) (*int32, error) {
	return GetIntPtr[int32](form, key)
}

// GetInt64 retrieves an int64 value from URL form values.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	int64 - parsed integer value
//	error - parsing error if any
func GetInt64(form url.Values, key string) (int64, error) {
	return GetInt[int64](form, key)
}

// GetInt64Slice retrieves a slice of int64 values from the given url.Values by the specified key.
//
// Parameters:
//   - form: The url.Values containing the form data to be parsed.
//   - key: The key used to lookup the values in the form data.
//
// Returns:
//   - []int64: The parsed slice of int64 values if successful.
//   - error: An error if the parsing fails or if the key is not found.
func GetInt64Slice(form url.Values, key string) ([]int64, error) {
	return GetIntSlice[int64](form, key)
}

// GetInt64Ptr retrieves an int64 value from form and returns its pointer.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*int64 - pointer to parsed value
//	error - parsing error if any
func GetInt64Ptr(form url.Values, key string) (*int64, error) {
	return GetIntPtr[int64](form, key)
}

// GetUint32 retrieves a uint32 value from URL form values.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	uint32 - parsed unsigned integer value
//	error - parsing error if any
func GetUint32(form url.Values, key string) (uint32, error) {
	return GetUint[uint32](form, key)
}

// GetUint32Slice retrieves a slice of uint32 values from the given url.Values by key.
//
// Parameters:
//   - form: url.Values containing the form data to parse
//   - key: string key to look up in the form values
//
// Returns:
//   - []uint32: slice of parsed uint32 values if successful
//   - error: any error that occurred during parsing, such as invalid format
func GetUint32Slice(form url.Values, key string) ([]uint32, error) {
	return GetUintSlice[uint32](form, key)
}

// GetUint32Ptr retrieves a uint32 value from form and returns its pointer.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*uint32 - pointer to parsed value
//	error - parsing error if any
func GetUint32Ptr(form url.Values, key string) (*uint32, error) {
	return GetUintPtr[uint32](form, key)
}

// GetUint64 retrieves a uint64 value from URL form values.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	uint64 - parsed unsigned integer value
//	error - parsing error if any
func GetUint64(form url.Values, key string) (uint64, error) {
	return GetUint[uint64](form, key)
}

// GetUint64Slice parses a URL form value as a slice of uint64 integers.
//
// Parameters:
//   - form: The URL form values containing the target key-value pairs.
//   - key: The form field key whose value should be parsed as a []uint64.
//
// Returns:
//   - []uint64: The parsed slice of uint64 integers if successful.
//   - error: An error if the parsing fails (e.g., invalid format or empty key).
//
// Note: This is a convenience wrapper around GetUintSlice[uint64] for uint64-specific parsing.
func GetUint64Slice(form url.Values, key string) ([]uint64, error) {
	return GetUintSlice[uint64](form, key)
}

// GetUint64Ptr retrieves a uint64 value from form and returns its pointer.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*uint64 - pointer to parsed value
//	error - parsing error if any
func GetUint64Ptr(form url.Values, key string) (*uint64, error) {
	return GetUintPtr[uint64](form, key)
}

// GetFloat32 retrieves a float32 value from URL form values.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	float32 - parsed floating-point value
//	error - parsing error if any
func GetFloat32(form url.Values, key string) (float32, error) {
	return GetFloat[float32](form, key)
}

// GetFloat32Ptr retrieves a float32 value from form and returns its pointer.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*float32 - pointer to parsed value
//	error - parsing error if any
func GetFloat32Ptr(form url.Values, key string) (*float32, error) {
	return GetFloatPtr[float32](form, key)
}

// GetFloat32Slice retrieves a slice of float32 values from URL form values.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	[]float32 - slice of parsed values
//	error - parsing error if any
func GetFloat32Slice(form url.Values, key string) ([]float32, error) {
	return GetFloatSlice[float32](form, key)
}

// GetFloat64 retrieves a float64 value from URL form values.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	float64 - parsed floating-point value
//	error - parsing error if any
func GetFloat64(form url.Values, key string) (float64, error) {
	return GetFloat[float64](form, key)
}

// GetFloat64Ptr retrieves a float64 value from form and returns its pointer.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*float64 - pointer to parsed value
//	error - parsing error if any
func GetFloat64Ptr(form url.Values, key string) (*float64, error) {
	return GetFloatPtr[float64](form, key)
}

// GetFloat64Slice retrieves a slice of float64 values from URL form values.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	[]float64 - slice of parsed values
//	error - parsing error if any
func GetFloat64Slice(form url.Values, key string) ([]float64, error) {
	return GetFloatSlice[float64](form, key)
}

// GetBool retrieves a boolean value from URL form values.
// Returns false if key doesn't exist without error.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	bool - parsed boolean value
//	error - parsing error if any
func GetBool(form url.Values, key string) (bool, error) {
	if _, ok := form[key]; !ok {
		return false, nil
	}
	return strconv.ParseBool(form.Get(key))
}

// GetBoolPtr retrieves a boolean value from form and returns its pointer.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	*bool - pointer to parsed value
//	error - parsing error if any
func GetBoolPtr(form url.Values, key string) (*bool, error) {
	v, err := GetBool(form, key)
	return &v, err
}

// GetBoolSlice retrieves a slice of boolean values from URL form values.
//
// Parameters:
//
//	form - URL form values containing the data
//	key - form field key to retrieve
//
// Returns:
//
//	[]bool - slice of parsed values
//	error - parsing error if any
func GetBoolSlice(form url.Values, key string) ([]bool, error) {
	if _, ok := form[key]; !ok {
		return nil, nil
	}
	return ParseBoolSlice(form[key])
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
	v, err := ParseBoolSlice(form[key])
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

// FormFromParams converts gin.Params to url.Values (form data).
//
// Parameters:
//   - m: gin.Params - A collection of URL parameters from Gin framework.
//     If nil is passed, the function returns nil.
//
// Returns:
//   - url.Values: A form data structure containing all key-value pairs from the input.
//     Returns nil if input is nil.
func FormFromParams(p gin.Params) url.Values {
	if p == nil {
		return nil
	}

	// Initialize form data and populate with all key-value pairs
	form := url.Values{}
	for _, item := range p {
		form.Add(item.Key, item.Value)
	}
	return form
}
