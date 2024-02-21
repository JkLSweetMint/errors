package errors

import "sm_errors/entities"

// Universal - универсальная абстракция системы ошибок.
type Universal interface {
	SetError(err error) Universal

	Error
	Grpc
	Http
	Ws
}

// FieldsUniversal - универсальная абстракция системы ошибок с полями.
type FieldsUniversal interface {
	SetError(err error) FieldsUniversal
	SetFields(InternalFields ...entities.Field) FieldsUniversal

	Fields
	FieldsGrpc
	FieldsHttp
	FieldsWs
}
