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
