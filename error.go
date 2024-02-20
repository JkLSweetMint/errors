package sm_errors

import "sm_errors/entities"

// Error - единая абстракция системы ошибок.
type Error interface {
	ID() (id entities.ID)
	Type() (t entities.Type)
	Is(err error) bool
	Status() (status entities.Status)

	Message(options ...entities.MessageOption) (message string)
}
