package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ConflictError struct {
	message string
}

func NewConflictError(message string) error {
	return &ConflictError{message}
}

func (e *ConflictError) Error() string {
	return e.message
}

func (e *ConflictError) gRPC() error {
	return status.Error(codes.AlreadyExists, e.message)
}
