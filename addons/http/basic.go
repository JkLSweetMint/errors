package sm_errors_http

import (
	errors "sm_errors"
	"sm_errors/entities"
)

const BasicType entities.Type = "basic-http"

// Basic - базовая реализация http ошибки.
type Basic interface {
	ID() (id entities.ID)
	Type() (t entities.Type)
	Status() (status entities.Status)
	StatusCode() (status StatusCode)

	Message(options ...entities.MessageOption) (message string)
	Error() (err Error)
	Is(err error) bool

	SetError(err error) Basic
}

// basic - стандартная реализация http ошибки.
type basic struct {
	errors.Basic
	t          entities.Type
	statusCode StatusCode
}

// StatusCode - получение статус кода http ошибки.
func (instance *basic) StatusCode() (status StatusCode) {
	return instance.statusCode
}

// Type - получение типа ошибки.
func (instance *basic) Type() (t entities.Type) {
	return instance.t
}

// Error - получение единой абстракции ошибок.
func (instance *basic) Error() (err Error) {
	return Error(instance)
}

// Is - проверка соответствия исходной ошибки.
func (instance *basic) Is(err error) bool {
	return instance.Basic.Is(err)
}

// SetError - установка изначальной ошибки.
func (instance *basic) SetError(err error) Basic {
	instance.Basic = instance.Basic.SetError(err)
	return instance
}
