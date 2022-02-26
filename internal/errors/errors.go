package errors

import "errors"

var (
	DbError          = errors.New("error with database")
	StanConnectError = errors.New("error with nats-streaming")
	ServiceError     = errors.New("error with service")
)
