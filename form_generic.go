package goose

import (
	"net/url"

	"golang.org/x/exp/constraints"
)

// FormGetter defines a generic function type for form data retrieval
// Parameters:
//   - form: Form data
//   - key: Form field key
//
// Returns:
//   - T: Retrieved value
//   - error: Retrieval error if any
type FormGetter[T any] func(form url.Values, key string) (T, error)

// GetForm decodes form data
// Parameters:
//   - pre: Pre-existing error (if any, will be returned immediately)
//   - form: Form data
//   - key: Form field key
//   - f: Form data getter function
//
// Returns:
//   - T: Decoded value
//   - error: Decoding error if any
//
// Behavior:
//  1. If pre is not nil, returns pre error immediately
//  2. Otherwise invokes f to get form value
func GetForm[T any](pre error, form url.Values, key string, f FormGetter[T]) (T, error) {
	return BreakOnError[T](pre)(func() (T, error) { return f(form, key) })
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
