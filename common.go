package goose

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
