package errors

import "sm_errors/entities"

// Error - единая абстракция системы ошибок.
type Error interface {
	ID() (id entities.ID)
	Is(err error) bool
	String() (str string)
	Error() (err string)
	Status() (status entities.Status)

	Message(options ...entities.MessageOption) (message string)
}

// Fields - единая абстракция системы ошибок с полями.
type Fields interface {
	ID() (id entities.ID)
	Is(err error) bool
	String() (str string)
	Error() (err string)
	Status() (status entities.Status)

	Message(options ...entities.MessageOption) (message string)
	Fields() (fields entities.Fields)
}
