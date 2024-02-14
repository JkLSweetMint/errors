package sm_errors

import "errors/entities"

const ErrorType entities.Type = "basic"

// Basic - базовая реализация ошибки.
type Basic interface {
	ID() (id entities.ID)
	Type() (t entities.Type)
	Status() (status entities.Status)

	Message(options ...entities.MessageOption) (message string)
	Error() (err error)

	SetError(err error) Basic
}

// ErrorConstructor - конструктор Error ошибки.
type ErrorConstructor struct {
	ID      entities.ID
	Status  entities.Status
	Message *entities.Message
}

// basic - стандартная реализация ошибки.
type basic struct {
	id     entities.ID
	t      entities.Type
	status entities.Status

	message *entities.Message
	err     error
}

// Build - сбор конструктора Error ошибки.
func (constructor ErrorConstructor) Build() Basic {
	return &basic{
		id:      constructor.ID,
		t:       ErrorType,
		status:  constructor.Status,
		message: constructor.Message,
		err:     nil,
	}
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

// Error - получение изначальной ошибки.
func (instance *basic) Error() (err error) {
	return instance.err
}

// SetError - установка изначальной ошибки.
func (instance *basic) SetError(err error) Basic {
	instance.err = err
	return instance
}
