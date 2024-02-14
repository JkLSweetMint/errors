package sm_errors_ws

import (
	errors "errors"
	"errors/entities"
)

const FieldsErrorType entities.Type = "fields-ws"

// FieldsError - базовая реализация ошибки web сокетов с полями.
type FieldsError interface {
	ID() (id entities.ID)
	Type() (t entities.Type)
	Status() (status entities.Status)
	StatusCode() (status StatusCode)
	Fields() (fields entities.Fields)

	Message(options ...entities.MessageOption) (message string)
	Error() (err error)

	SetError(err error) FieldsError
	SetFields(fields ...entities.Field) FieldsError
}

// FieldsErrorConstructor - конструктор FieldsError ошибки.
type FieldsErrorConstructor struct {
	errors.FieldsErrorConstructor
	StatusCode StatusCode
}

// fields - стандартная реализация ошибки web сокетов с полями.
type fields struct {
	errors.FieldsError
	t          entities.Type
	statusCode StatusCode
}

// Build - сбор конструктора Error ошибки.
func (constructor FieldsErrorConstructor) Build() FieldsError {
	return &fields{
		FieldsError: errors.FieldsErrorConstructor{
			ID:      constructor.ID,
			Status:  constructor.Status,
			Message: constructor.Message,
			Fields:  constructor.Fields,
		}.Build(),
		t:          FieldsErrorType,
		statusCode: constructor.StatusCode,
	}
}

// StatusCode - получение статус кода ошибки web сокетов.
func (instance *fields) StatusCode() (status StatusCode) {
	return instance.statusCode
}

// Type - получение типа ошибки.
func (instance *fields) Type() (t entities.Type) {
	return instance.t
}

// SetError - установка изначальной ошибки.
func (instance *fields) SetError(err error) FieldsError {
	instance.FieldsError = instance.FieldsError.SetError(err)
	return instance
}

// SetFields - установка значение полей.
func (instance *fields) SetFields(fields ...entities.Field) FieldsError {
	instance.FieldsError = instance.FieldsError.SetFields(fields...)
	return instance
}
