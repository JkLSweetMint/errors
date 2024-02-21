package errors

import (
	entities_grpc "sm_errors/entities/grpc"
)

// Grpc - абстракция системы grpc ошибок.
type Grpc interface {
	Error
	GrpcStatusCode() (status entities_grpc.StatusCode)
}

// FieldsGrpc - абстракция системы grpc ошибок с полями.
type FieldsGrpc interface {
	Fields
	GrpcStatusCode() (status entities_grpc.StatusCode)
}
