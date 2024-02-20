package sm_errors_ws

import "sm_errors/entities"

// Error - единая абстракция системы ws ошибок.
type Error interface {
	ID() (id entities.ID)
	Type() (t entities.Type)
	Is(err error) bool
	Status() (status entities.Status)
	StatusCode() (status StatusCode)

	Message(options ...entities.MessageOption) (message string)
}
