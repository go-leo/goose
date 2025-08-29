package goose

import (
	"net/http"
)

// MiddlewareFunc is a function which receives an http.Handler and returns another http.Handler.
// Typically, the returned handler is a closure which does something with the http.ResponseWriter and http.Request passed
// to it, and then calls the handler passed as parameter to the MiddlewareFunc.
type MiddlewareFunc func(http.Handler) http.Handler

// Chain applies a series of middleware functions to an HTTP handler in reverse order.
// This allows middleware to be executed in the order they are provided (first middleware wraps the original handler,
// subsequent middlewares wrap the previous chain).
//
// Parameters:
//   - handler: The base HTTP handler to be wrapped by middlewares.
//   - middlewares: A variadic list of middleware functions to be applied. These will be executed
//     in reverse order (last middleware in the list is applied first).
//
// Returns:
//   - http.Handler: The final handler wrapped by all middlewares in the chain.
func Chain(handler http.Handler, middlewares ...MiddlewareFunc) http.Handler {
	// Iterate through middlewares in reverse order to ensure proper chaining.
	// Each middleware wraps the previous handler chain.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
