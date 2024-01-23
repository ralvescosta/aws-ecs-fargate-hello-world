package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InternalError struct {
	message string
}

func NewInternalError(message string) ApplicationError {
	return &InternalError{message}
}

func (e *InternalError) Error() string {
	return e.message
}

func (e *InternalError) ToGrpc() error {
	return status.Error(codes.Internal, e.message)
}
