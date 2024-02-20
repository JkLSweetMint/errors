package sm_errors_grpc

import (
	errors "sm_errors"
	"sm_errors/entities"
)

const FieldsType entities.Type = "fields-grpc"

// Fields - базовая реализация grpc ошибки с полями.
type Fields interface {
	ID() (id entities.ID)
	Type() (t entities.Type)
	Status() (status entities.Status)
	StatusCode() (status StatusCode)
	Fields() (fields entities.Fields)

	Message(options ...entities.MessageOption) (message string)
	Error() (err Error)
	Is(err error) bool

	SetError(err error) Fields
	SetFields(fields ...entities.Field) Fields
}

// fields - стандартная реализация grpc ошибки с полями.
type fields struct {
	errors.Fields
	t          entities.Type
	statusCode StatusCode
}

// StatusCode - получение статус кода grpc ошибки.
func (instance *fields) StatusCode() (status StatusCode) {
	return instance.statusCode
}

// Type - получение типа ошибки.
func (instance *fields) Type() (t entities.Type) {
	return instance.t
}

// Error - получение единой абстракции ошибок.
func (instance *fields) Error() (err Error) {
	return Error(instance)
}

// Is - проверка соответствия исходной ошибки.
func (instance *fields) Is(err error) bool {
	return instance.Fields.Is(err)
}

// SetError - установка изначальной ошибки.
func (instance *fields) SetError(err error) Fields {
	instance.Fields = instance.Fields.SetError(err)
	return instance
}

// SetFields - установка значение полей.
func (instance *fields) SetFields(fields ...entities.Field) Fields {
	instance.Fields = instance.Fields.SetFields(fields...)
	return instance
}
