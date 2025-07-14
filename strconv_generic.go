package gonic

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

// ParseBool converts a string to a boolean value.
// It wraps strconv.ParseBool to provide the same functionality.
//
// Parameters:
//
//	s - the string to be parsed into a boolean
//
// Returns:
//
//	bool - the parsed boolean value
//	error - if parsing fails
func ParseBool(s string) (bool, error) {
	return strconv.ParseBool(s)
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

// ParseBoolSlice converts a slice of strings to a slice of booleans.
// Returns nil if the input slice is nil.
//
// Parameters:
//
//	s - the string slice to be parsed
//
// Returns:
//
//	[]bool - the parsed boolean slice
//	error - if any element fails to parse
func ParseBoolSlice(s []string) ([]bool, error) {
	if s == nil {
		return nil, nil
	}
	r := make([]bool, 0, len(s))
	for _, str := range s {
		b, err := strconv.ParseBool(str)
		if err != nil {
			return nil, err
		}
		r = append(r, b)
	}
	return r, nil
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
