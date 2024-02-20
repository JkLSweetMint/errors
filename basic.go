package sm_errors

import (
	"errors"
	"sm_errors/entities"
)

const BasicType entities.Type = "basic"

// Basic - базовая реализация ошибки.
type Basic interface {
	ID() (id entities.ID)
	Type() (t entities.Type)
	Status() (status entities.Status)

	Message(options ...entities.MessageOption) (message string)
	Error() (err Error)
	Is(err error) bool

	SetError(err error) Basic
}

// basic - стандартная реализация ошибки.
type basic struct {
	id     entities.ID
	t      entities.Type
	status entities.Status

	message *entities.Message
	err     error
}

// ID - получение идентификатора ошибки.
func (instance *basic) ID() (id entities.ID) {
	return instance.id
}

// Type - получение типа ошибки.
func (instance *basic) Type() (t entities.Type) {
	return instance.t
}

// Status - получение статуса ошибки.
func (instance *basic) Status() (status entities.Status) {
	return instance.status
}

// Message - получение сообщения ошибки.
func (instance *basic) Message(options ...entities.MessageOption) (message string) {
	message = instance.err.Error()

	if instance.message != nil {
		message = instance.message.String(options...)
	}

	return
}

// Error - получение единой абстракции ошибок.
func (instance *basic) Error() (err Error) {
	return Error(instance)
}

// Is - проверка соответствия исходной ошибки.
func (instance *basic) Is(err error) bool {
	return errors.Is(instance.err, err)
}

// SetError - установка изначальной ошибки.
func (instance *basic) SetError(err error) Basic {
	instance.err = err
	return instance
}
