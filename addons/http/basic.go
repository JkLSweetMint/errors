package sm_errors_http

import (
	errors "errors"
	"errors/entities"
)

const ErrorType entities.Type = "basic-http"

// Error - базовая реализация http ошибки.
type Error interface {
	ID() (id entities.ID)
	Type() (t entities.Type)
	Status() (status entities.Status)
	StatusCode() (status StatusCode)

	Message(options ...entities.MessageOption) (message string)
	Error() (err error)

	SetError(err error) Error
}

// ErrorConstructor - конструктор Error ошибки.
type ErrorConstructor struct {
	errors.ErrorConstructor
	StatusCode StatusCode
}

// basic - стандартная реализация http ошибки.
type basic struct {
	errors.Basic
	t          entities.Type
	statusCode StatusCode
}

// Build - сбор конструктора Error ошибки.
func (constructor ErrorConstructor) Build() Error {
	return &basic{
		Basic: errors.ErrorConstructor{
			ID:      constructor.ID,
			Status:  constructor.Status,
			Message: constructor.Message,
		}.Build(),
		t:          ErrorType,
		statusCode: constructor.StatusCode,
	}
}

// StatusCode - получение http статус кода ошибки.
func (instance *basic) StatusCode() (status StatusCode) {
	return instance.statusCode
}

// Type - получение типа ошибки.
func (instance *basic) Type() (t entities.Type) {
	return instance.t
}

// SetError - установка изначальной ошибки.
func (instance *basic) SetError(err error) Error {
	instance.Basic = instance.Basic.SetError(err)
	return instance
}
