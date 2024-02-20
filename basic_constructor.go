package sm_errors

import "sm_errors/entities"

// BasicConstructor - конструктор Basic ошибки.
type BasicConstructor struct {
	ID      entities.ID
	Status  entities.Status
	Message *entities.Message
}

// Build - сбор конструктора Basic ошибки.
func (constructor BasicConstructor) Build() Basic {
	return &basic{
		id:      constructor.ID,
		t:       BasicType,
		status:  constructor.Status,
		message: constructor.Message,
		err:     nil,
	}
}
