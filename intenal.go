package errors

import (
	"errors"
	"sm_errors/entities"
	entities_grpc "sm_errors/entities/grpc"
	entities_http "sm_errors/entities/http"
	entities_ws "sm_errors/entities/ws"
)

// Internal - внутренняя реализация ошибок.
type Internal struct {
	id     entities.ID
	status entities.Status

	message *entities.Message
	err     error

	grpcStatusCode entities_grpc.StatusCode
	httpStatusCode entities_http.StatusCode
	wsStatusCode   entities_ws.StatusCode
}

// ID - получение идентификатора ошибки.
func (instance Internal) ID() (id entities.ID) {
	return instance.id
}

// String - получение строкового представления ошибки.
func (instance Internal) String() (str string) {
	str = instance.Message()

	if str == "" {
		str = instance.err.Error()
	}

	return
}

// Status - получение статуса ошибки.
func (instance Internal) Status() (status entities.Status) {
	return instance.status
}

// GrpcStatusCode - получение статус кода grpc ошибки.
func (instance Internal) GrpcStatusCode() (status entities_grpc.StatusCode) {
	return instance.grpcStatusCode
}

// HttpStatusCode - получение статус кода http ошибки.
func (instance Internal) HttpStatusCode() (status entities_http.StatusCode) {
	return instance.httpStatusCode
}

// WsStatusCode - получение статус кода ws ошибки.
func (instance Internal) WsStatusCode() (status entities_ws.StatusCode) {
	return instance.wsStatusCode
}

// Message - получение сообщения ошибки.
func (instance Internal) Message(options ...entities.MessageOption) (message string) {
	if instance.err != nil {
		message = instance.err.Error()
	}

	if instance.message != nil {
		message = instance.message.String(options...)
	}

	return
}

// Error - получение единой абстракции ошибок.
func (instance Internal) Error() (err string) {
	return instance.err.Error()
}

// Is - проверка соответствия исходной ошибки.
func (instance Internal) Is(err error) bool {
	return errors.Is(instance.err, err)
}
