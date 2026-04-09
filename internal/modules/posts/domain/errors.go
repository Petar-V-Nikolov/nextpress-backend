package domain

import "errors"

// Domain-level errors used across application and infrastructure.
// Infrastructure should map storage-specific errors (e.g. gorm) to these.
var (
	// ErrNotFound indicates the requested entity/resource does not exist.
	ErrNotFound = errors.New("not_found")
	// ErrConflict indicates a uniqueness or state conflict (e.g. duplicate key).
	ErrConflict = errors.New("conflict")
	// ErrInvalidArgument indicates the request is invalid (validation failure).
	ErrInvalidArgument = errors.New("invalid_argument")
)
