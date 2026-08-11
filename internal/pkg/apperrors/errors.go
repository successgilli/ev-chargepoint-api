// Package apperrors defines generic, reusable application-level error
// sentinels — such as not-found and invalid-input — shared across all
// entities and layers of the application. It contains no domain-specific
// logic; each layer wraps or checks against these sentinels using the
// standard errors.Is/errors.As mechanisms.
package apperrors

import "errors"

var (
	// ErrNotFound represents a "not found" application error
	ErrNotFound = errors.New("not found")

	// ErrInternal represents a generic internal application error
	ErrInternal = errors.New("something went wrong")
)
