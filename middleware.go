package gonic

import (
	"github.com/gin-gonic/gin"
)

// Chain constructs a middleware chain by combining the given handler and middlewares.
// The resulting slice will process requests in the order: middlewares -> varsMiddleware() -> handler.
//
// Parameters:
//   - handler: The main request handler that should be called at the end of the middleware chain.
//   - middlewares: A variadic list of middleware functions to be executed before the handler.
//
// Returns:
//
//	A slice of gin.HandlerFunc representing the complete middleware chain.
func Chain(handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) []gin.HandlerFunc {
	// Combine middlewares with varsMiddleware and handler in execution order
	return append(middlewares, handler)
}
