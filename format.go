package goose

import (
	"strconv"

	"golang.org/x/exp/constraints"
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
