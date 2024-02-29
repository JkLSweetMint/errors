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
	Web  *WebConstructor
}

// GrpcConstructor - часть конструктора для создания grpc ошибок.
type GrpcConstructor struct {
	StatusCode entities_grpc.StatusCode
}

// WebConstructor - часть конструктора для создания web ошибок.
type WebConstructor struct {
	HttpStatusCode entities_http.StatusCode
	WsStatusCode   entities_ws.StatusCode
}

// Build - сбор универсальной ошибки.
func (constructor Constructor) Build() Universal {
	return internal{
		id:     constructor.ID,
		status: constructor.Status,

		message: constructor.Message,
		err:     constructor.Err,

		grpcStatusCode: constructor.Grpc.StatusCode,
		httpStatusCode: constructor.Web.HttpStatusCode,
		wsStatusCode:   constructor.Web.WsStatusCode,
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
		Web: &WebConstructor{
			HttpStatusCode: constructor.Web.HttpStatusCode,
			WsStatusCode:   constructor.Web.WsStatusCode,
		},
	}
}
