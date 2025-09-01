package goose

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

// FormatBool takes a boolean-type generic parameter `b`, converts it to a string, and returns
// the string.
func FormatBool[Bool ~bool](b Bool) string {
	return strconv.FormatBool(bool(b))
}

// FormatUint converts an unsigned integer to a string representation in a specified base.
// It does this by first converting the integer to a uint64 and then using the strconv.FormatUint
// function to format it as a string in the desired base.
func FormatUint[Unsigned constraints.Unsigned](i Unsigned, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

func FormatInt[Signed constraints.Signed](i Signed, base int) string {
	return strconv.FormatInt(int64(i), base)
}

func FormatFloat[Float constraints.Float](f Float, fmt byte, prec, bitSize int) string {
	return strconv.FormatFloat(float64(f), fmt, prec, bitSize)
}

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
