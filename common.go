package goose

import "errors"

// BreakOnError provides error interception functionality
// Parameters:
//   - pre: Pre-existing error
//
// Returns:
//
//	A function that will:
//	1. Return pre error immediately if pre is not nil
//	2. Otherwise execute the provided function f
func BreakOnError[T any](pre error) func(f func() (T, error)) (T, error) {
	return func(f func() (T, error)) (T, error) {
		if pre != nil {
			var v T
			return v, pre
		}
		return f()
	}
}

// ContinueOnError provides error continuation functionality
// Parameters:
//   - pre: Pre-existing error
//
// Returns:
//
//	A function that will:
//	1. Execute the provided function f
//	2. Return result and pre error if f succeeds but pre is not nil
//	3. Return result and f's error if f fails but pre is nil
//	4. Return result and joined errors if both f fails and pre is not nil
func ContinueOnError[T any](pre error) func(f func() (T, error)) (T, error) {
	return func(f func() (T, error)) (T, error) {
		v, err := f()
		if err == nil {
			return v, pre
		}
		if pre == nil {
			return v, err
		}
		joinError, ok := pre.(interface{ Unwrap() []error })
		if !ok {
			return v, errors.Join(pre, err)
		}
		return v, errors.Join(append(joinError.Unwrap(), err)...)
	}
}
