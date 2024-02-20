package sm_errors_grpc

import errors "sm_errors"

// FieldsConstructor - конструктор Fields ошибки.
type FieldsConstructor struct {
	errors.FieldsConstructor
	StatusCode StatusCode
}

// Build - сбор конструктора Fields ошибки.
func (constructor FieldsConstructor) Build() Fields {
	return &fields{
		Fields: errors.FieldsConstructor{
			ID:      constructor.ID,
			Status:  constructor.Status,
			Message: constructor.Message,
			Fields:  constructor.Fields,
		}.Build(),
		t:          FieldsType,
		statusCode: constructor.StatusCode,
	}
}
