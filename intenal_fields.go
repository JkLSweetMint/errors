package errors

import (
	"sm_errors/entities"
)

// InternalFields - внутренняя реализация ошибок с полями.
type InternalFields struct {
	Internal

	fields entities.Fields
}

// Fields - получение полей.
func (instance InternalFields) Fields() (InternalFields entities.Fields) {
	return instance.fields
}

// SetError - установка изначальной ошибки.
func (instance InternalFields) SetError(err error) FieldsUniversal {
	instance.err = err
	return instance
}

// SetFields - установка значение полей.
func (instance InternalFields) SetFields(InternalFields ...entities.Field) FieldsUniversal {
	instance.fields = InternalFields
	return instance
}
