package errors

import (
	"sm_errors/entities"
	entities_grpc "sm_errors/entities/grpc"
	entities_http "sm_errors/entities/http"
	entities_ws "sm_errors/entities/ws"
)

// FieldsConstructor - конструктор стандартных ошибок с полями.
type FieldsConstructor struct {
	ID     entities.ID
	Status entities.Status

	Err     error
	Message *entities.Message
	Fields  entities.Fields

	Grpc *GrpcConstructor
	Web  *WebConstructor
}

// FieldsGrpcConstructor - часть конструктора для создания grpc ошибок с полями.
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
	return internalFields{
		internal: internal{
			id:     constructor.ID,
			status: constructor.Status,

			message: constructor.Message,
			err:     constructor.Err,

			grpcStatusCode: constructor.Grpc.StatusCode,
			httpStatusCode: constructor.Web.HttpStatusCode,
			wsStatusCode:   constructor.Web.WsStatusCode,
		},
		fields: constructor.Fields,
	}
}

// SetError - установка внутренней ошибки.
func (constructor FieldsConstructor) SetError(err error) FieldsConstructor {
	return FieldsConstructor{
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

// SetFields - установка значение полей.
func (constructor FieldsConstructor) SetFields(internalFields ...entities.Field) FieldsConstructor {
	return FieldsConstructor{
		ID:     constructor.ID,
		Status: constructor.Status,

		Err:     constructor.Err,
		Message: constructor.Message.Clone(),
		Fields:  internalFields,

		Grpc: &GrpcConstructor{
			StatusCode: constructor.Grpc.StatusCode,
		},
		Web: &WebConstructor{
			HttpStatusCode: constructor.Web.HttpStatusCode,
			WsStatusCode:   constructor.Web.WsStatusCode,
		},
	}
}

// SetField - установка значения поля.
func (constructor FieldsConstructor) SetField(key, message string) FieldsConstructor {
	return FieldsConstructor{
		ID:     constructor.ID,
		Status: constructor.Status,

		Err:     constructor.Err,
		Message: constructor.Message.Clone(),
		Fields: append(constructor.Fields, entities.Field{
			Key:     key,
			Message: message,
		}),

		Grpc: &GrpcConstructor{
			StatusCode: constructor.Grpc.StatusCode,
		},
		Web: &WebConstructor{
			HttpStatusCode: constructor.Web.HttpStatusCode,
			WsStatusCode:   constructor.Web.WsStatusCode,
		},
	}
}
