package errors

import (
	"reflect"
	"sm-errors/entities/details"
	"sm-errors/entities/messages"
	"sm-errors/internal"
	"sm-errors/types"
)

// Базовая реализация
type (
	// Builder - функция дла построения базовой ошибки.
	Builder[T Error] func() T

	// Constructor - конструктор для построения ошибки.
	Constructor[T Error] struct {
		ID     types.ID
		Type   types.ErrorType
		Status types.Status

		Err     error
		Message types.Message
		Details types.Details
	}
)

// Build - построение ошибки.
func (c Constructor[T]) Build() (fn Builder[T]) {
	c.fillEmptyField()

	fn = func() (e T) {
		switch reflect.TypeOf(new(T)).String() {
		case "*errors.Error":
			{
				var i = internal.New(internal.Constructor{
					ID:     c.ID,
					Type:   c.Type,
					Status: c.Status,

					Err:     c.Err,
					Message: c.Message,
					Details: c.Details,
				})

				e = interface{}(i).(T)
			}
		}

		return
	}

	return
}

// fillEmptyField - заполнение пустых полей структуры.
func (c *Constructor[T]) fillEmptyField() *Constructor[T] {
	if c.Message == nil {
		c.Message = new(messages.TextMessage)
	}

	if c.Details == nil {
		c.Details = new(details.Details)
	}

	return c
}
