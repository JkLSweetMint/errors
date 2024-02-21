package errors

import (
	"sm_errors/entities"
	entities_grpc "sm_errors/entities/grpc"
	entities_http "sm_errors/entities/http"
	entities_ws "sm_errors/entities/ws"
)

// FieldsConstructor - конструктор стандартных ошибок с полями.
type FieldsConstructor struct {
	ID      entities.ID
	Status  entities.Status
	Message *entities.Message

	Grpc GrpcConstructor
	Http HttpConstructor
	Ws   WsConstructor
}

// FieldsGrpcConstructor - часть конструктора для создания grcp ошибок с полями.
type FieldsGrpcConstructor struct {
	StatusCode entities_grpc.StatusCode
}

// FieldsHttpConstructor - часть конструктора для создания http ошибок с полями.
type FieldsHttpConstructor struct {
	StatusCode entities_http.StatusCode
}

// FieldsWsConstructor - часть конструктора для создания ws ошибок с полями.
type FieldsWsConstructor struct {
	StatusCode entities_ws.StatusCode
}

// Build - сбор универсальной ошибки с полями.
func (constructor FieldsConstructor) Build() FieldsUniversal {
	return InternalFields{
		Internal: Internal{
			id:     constructor.ID,
			status: constructor.Status,

			grpcStatusCode: constructor.Grpc.StatusCode,
			httpStatusCode: constructor.Http.StatusCode,
			wsStatusCode:   constructor.Ws.StatusCode,

			message: constructor.Message,
			err:     nil,
		},
		fields: nil,
	}
}
