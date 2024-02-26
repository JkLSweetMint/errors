package errors

import (
	"sm_errors/entities"
	entities_grpc "sm_errors/entities/grpc"
	entities_http "sm_errors/entities/http"
	entities_ws "sm_errors/entities/ws"
)

// Constructor - конструктор стандартных ошибок.
type Constructor struct {
	ID     entities.ID
	Status entities.Status

	Err     error
	Message *entities.Message

	Grpc *GrpcConstructor
	Http *HttpConstructor
	Ws   *WsConstructor
}

// GrpcConstructor - часть конструктора для создания grpc ошибок.
type GrpcConstructor struct {
	StatusCode entities_grpc.StatusCode
}

// HttpConstructor - часть конструктора для создания http ошибок.
type HttpConstructor struct {
	StatusCode entities_http.StatusCode
}

// WsConstructor - часть конструктора для создания ws ошибок.
type WsConstructor struct {
	StatusCode entities_ws.StatusCode
}

// Build - сбор универсальной ошибки.
func (constructor Constructor) Build() Universal {
	return Internal{
		id:     constructor.ID,
		status: constructor.Status,

		message: constructor.Message,
		err:     constructor.Err,

		grpcStatusCode: constructor.Grpc.StatusCode,
		httpStatusCode: constructor.Http.StatusCode,
		wsStatusCode:   constructor.Ws.StatusCode,
	}
}

// SetError - установка внутренней ошибки.
func (constructor Constructor) SetError(err error) Constructor {
	return Constructor{
		ID:     constructor.ID,
		Status: constructor.Status,

		Err:     err,
		Message: constructor.Message.Clone(),

		Grpc: &GrpcConstructor{
			StatusCode: constructor.Grpc.StatusCode,
		},
		Http: &HttpConstructor{
			StatusCode: constructor.Http.StatusCode,
		},
		Ws: &WsConstructor{
			StatusCode: constructor.Ws.StatusCode,
		},
	}
}
