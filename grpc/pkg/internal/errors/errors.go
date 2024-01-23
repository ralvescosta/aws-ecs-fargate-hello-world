package errors

type ApplicationError interface {
	Error() string
	ToGrpc() error
}
