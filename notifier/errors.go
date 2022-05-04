package notifier

import "io"

// MessageError wraps error and provides a reference to a message related to the error.
type MessageError struct {
	Message io.Reader
	Err     error
}

// Unwrap returns base error.
func (e *MessageError) Unwrap() error { return e.Err }

// Error implements error interface.
func (e *MessageError) Error() string { return e.Err.Error() }
