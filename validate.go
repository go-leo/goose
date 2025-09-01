package goose

import (
	"context"

	"google.golang.org/protobuf/proto"
)

// OnValidationErrCallback defines a callback function type for validation errors
// This callback will be invoked when ValidateRequest encounters a validation error
type OnValidationErrCallback func(ctx context.Context, err error)

// ValidateRequest validates the request parameters
// Parameters:
//   - ctx: Context object
//   - req: Proto.Message to validate
//   - fast: Whether to perform fast validation (skip deep validation)
//   - callback: Callback function for validation errors
//
// Returns:
//   - error: Validation error if any
//
// Behavior:
//
//	Based on fast parameter:
//	- fast=true: Attempts to call Validate() or Validate(false)
//	- fast=false: Attempts to call ValidateAll() or Validate(true) or Validate()
//	If validation fails and callback is provided, invokes the callback
func ValidateRequest(ctx context.Context, req proto.Message, fast bool, callback OnValidationErrCallback) (err error) {
	if fast {
		switch v := req.(type) {
		case interface{ Validate() error }:
			err = v.Validate()
		case interface{ Validate(all bool) error }:
			err = v.Validate(false)
		}
	} else {
		switch v := req.(type) {
		case interface{ ValidateAll() error }:
			err = v.ValidateAll()
		case interface{ Validate(all bool) error }:
			err = v.Validate(true)
		case interface{ Validate() error }:
			err = v.Validate()
		}
	}

	if err == nil {
		return nil
	}

	if callback != nil {
		callback(ctx, err)
	}
	return err
}
