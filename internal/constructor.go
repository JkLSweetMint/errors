package internal

import (
	"sm_errors/entities"
)

// Constructor - конструктор ошибок.
type Constructor struct {
	ID     entities.ID
	Status entities.Status
	Type   string

	Err     error
	Message *entities.Message
	Fields  entities.Fields
}

// Build - сбор универсальной ошибки.
func (constructor Constructor) Build() *Internal {
	return &Internal{
		id:     constructor.ID,
		status: constructor.Status,
		t:      constructor.Type,

		fields:  constructor.Fields,
		message: constructor.Message,
		err:     constructor.Err,
	}
}
