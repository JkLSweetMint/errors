package errors

import (
	entities_grpc "sm_errors/entities/grpc"
)

// Grpc - абстракция системы grpc ошибок.
type Grpc interface {
	Error

	StatusCode() (status entities_grpc.StatusCode)
}

// FieldsGrpc - абстракция системы grpc ошибок с полями.
type FieldsGrpc interface {
	Fields

	StatusCode() (status entities_grpc.StatusCode)
}

// ----------------------------------------------------------------

// internalGrpc - внутрення реализация grpc ошибки.
type internalGrpc struct {
	internal
}

// StatusCode - получение статус кода ошибки.
func (instance internalGrpc) StatusCode() (status entities_grpc.StatusCode) {
	return instance.grpcStatusCode
}

// ----------------------------------------------------------------

// internalGrpc - внутрення реализация grpc ошибок с полями.
type internalFieldsGrpc struct {
	internalFields
}

// StatusCode - получение статус кода ошибки.
func (instance internalFieldsGrpc) StatusCode() (status entities_grpc.StatusCode) {
	return instance.grpcStatusCode
}
