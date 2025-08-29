package goose

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestChain_NoMiddleware tests when no middleware is provided
func TestChain_NoMiddleware(t *testing.T) {
	// Create a mock handler
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Chain with no middlewares
	chained := Chain(mockHandler)

	// Verify the chained handler is the same as original
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	chained.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}

// TestChain_SingleMiddleware tests with a single middleware
func TestChain_SingleMiddleware(t *testing.T) {
	// Create a mock handler
	called := false
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	// Create a mock middleware that sets a header
	mockMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Test", "middleware")
			next.ServeHTTP(w, r)
		})
	}

	// Chain with one middleware
	chained := Chain(mockHandler, mockMiddleware)

	// Verify the middleware was applied
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	chained.ServeHTTP(rec, req)

	if !called {
		t.Error("Original handler was not called")
	}
	if rec.Header().Get("X-Test") != "middleware" {
		t.Error("Middleware was not applied correctly")
	}
}

// TestChain_MultipleMiddlewares tests with multiple middlewares and verifies order
func TestChain_MultipleMiddlewares(t *testing.T) {
	// Track middleware execution order
	var executionOrder []string

	// Create a mock handler
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		executionOrder = append(executionOrder, "handler")
		w.WriteHeader(http.StatusOK)
	})

	// Create mock middlewares
	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			executionOrder = append(executionOrder, "start middleware1")
			next.ServeHTTP(w, r)
			executionOrder = append(executionOrder, "end middleware1")
		})
	}

	middleware2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			executionOrder = append(executionOrder, "start middleware2")
			next.ServeHTTP(w, r)
			executionOrder = append(executionOrder, "end middleware2")
		})
	}

	// Chain with multiple middlewares (should be applied in reverse order)
	chained := Chain(mockHandler, middleware1, middleware2)

	// Verify the middlewares were applied in correct order
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	chained.ServeHTTP(rec, req)

	expectedOrder := []string{"start middleware1", "start middleware2", "handler", "end middleware2", "end middleware1"}
	if len(executionOrder) != len(expectedOrder) {
		t.Fatalf("Expected %d calls, got %d", len(expectedOrder), len(executionOrder))
	}

	for i, step := range expectedOrder {
		if executionOrder[i] != step {
			t.Errorf("At position %d, expected %s, got %s", i, step, executionOrder[i])
		}
	}
}
