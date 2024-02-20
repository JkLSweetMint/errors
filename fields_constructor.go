package sm_errors

import "sm_errors/entities"

// FieldsConstructor - конструктор Fields ошибки.
type FieldsConstructor struct {
	ID      entities.ID
	Status  entities.Status
	Message *entities.Message
	Fields  entities.Fields
}

// Build - сбор конструктора Fields ошибки.
func (constructor FieldsConstructor) Build() Fields {
	return &fields{
		basic: &basic{
			id:     constructor.ID,
			t:      FieldsType,
			status: constructor.Status,

			message: constructor.Message,
			err:     nil,
		},

		fields: constructor.Fields,
	}
}
