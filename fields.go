package sm_errors

import "errors/entities"

const FieldsErrorType entities.Type = "fields"

// FieldsError - реализация ошибки с полями.
type FieldsError interface {
	ID() (id entities.ID)
	Type() (t entities.Type)
	Status() (status entities.Status)
	Fields() (fields entities.Fields)

	Message(options ...entities.MessageOption) (message string)
	Error() (err error)

	SetError(err error) FieldsError
	SetFields(fields ...entities.Field) FieldsError
}

// FieldsErrorConstructor - конструктор FieldsError ошибки.
type FieldsErrorConstructor struct {
	ID      entities.ID
	Status  entities.Status
	Message *entities.Message
	Fields  entities.Fields
}

// fields - реализация ошибки с полями.
type fields struct {
	*basic

	fields entities.Fields
}

// Build - сбор конструктора FieldsError ошибки.
func (constructor FieldsErrorConstructor) Build() FieldsError {
	return &fields{
		basic: &basic{
			id:     constructor.ID,
			t:      FieldsErrorType,
			status: constructor.Status,

			message: constructor.Message,
			err:     nil,
		},

		fields: constructor.Fields,
	}
}

// Fields - получение полей.
func (instance *fields) Fields() (fields entities.Fields) {
	return instance.fields
}

// SetError - установка изначальной ошибки.
func (instance *fields) SetError(err error) FieldsError {
	instance.err = err
	return instance
}

// SetFields - установка значение полей.
func (instance *fields) SetFields(fields ...entities.Field) FieldsError {
	instance.fields = fields
	return instance
}
