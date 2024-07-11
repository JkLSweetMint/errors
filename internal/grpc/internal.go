package grpc

import (
	"sm-errors/internal"
)

type (
	// Internal - внутренняя реализация ошибки grpc.
	Internal struct {
		*internal.Internal
	}
)

// New - создание внутренней реализации ошибки grpc.
func New(store *internal.Store) (i *Internal) {
	i = &Internal{
		Internal: internal.New(store),
	}

	return
}
