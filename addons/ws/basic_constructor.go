package sm_errors_ws

import errors "sm_errors"

// BasicConstructor - конструктор Basic ошибки.
type BasicConstructor struct {
	errors.BasicConstructor
	StatusCode StatusCode
}

// Build - сбор конструктора Basic ошибки.
func (constructor BasicConstructor) Build() Basic {
	return &basic{
		Basic: errors.BasicConstructor{
			ID:      constructor.ID,
			Status:  constructor.Status,
			Message: constructor.Message,
		}.Build(),
		t:          BasicType,
		statusCode: constructor.StatusCode,
	}
}
