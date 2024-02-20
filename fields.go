package sm_errors

import (
	"errors"
	"sm_errors/entities"
)

const FieldsType entities.Type = "fields"

// Fields - реализация ошибки с полями.
type Fields interface {
	ID() (id entities.ID)
	Type() (t entities.Type)
	Status() (status entities.Status)
	Fields() (fields entities.Fields)

	Message(options ...entities.MessageOption) (message string)
	Error() (err Error)
	Is(err error) bool

	SetError(err error) Fields
	SetFields(fields ...entities.Field) Fields
}

// fields - реализация ошибки с полями.
type fields struct {
	*basic

	fields entities.Fields
}

// Error - получение единой абстракции ошибок.
func (instance *fields) Error() (err Error) {
	return Error(instance)
}

// Is - проверка соответствия исходной ошибки.
func (instance *fields) Is(err error) bool {
	return errors.Is(instance.err, err)
}

// Fields - получение полей.
func (instance *fields) Fields() (fields entities.Fields) {
	return instance.fields
}

// SetError - установка изначальной ошибки.
func (instance *fields) SetError(err error) Fields {
	instance.err = err
	return instance
}

// SetFields - установка значение полей.
func (instance *fields) SetFields(fields ...entities.Field) Fields {
	instance.fields = fields
	return instance
}
