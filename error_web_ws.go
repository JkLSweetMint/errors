package errors

import (
	entities_ws "sm_errors/entities/ws"
)

// WebWs - абстракция системы web ws ошибок.
type WebWs interface {
	Error

	StatusCode() (status entities_ws.StatusCode)
}

// FieldsWebWs - абстракция системы web ws ошибок с полями.
type FieldsWebWs interface {
	Fields

	StatusCode() (status entities_ws.StatusCode)
}

// ----------------------------------------------------------------

// internalWebWs - внутрення реализация web ws ошибки.
type internalWebWs struct {
	internal
}

// StatusCode - получение статус кода ошибки.
func (instance internalWebWs) StatusCode() (status entities_ws.StatusCode) {
	return instance.wsStatusCode
}

// ----------------------------------------------------------------

// internalFieldsWebWs - внутрення реализация web ws ошибок с полями.
type internalFieldsWebWs struct {
	internalFields
}

// StatusCode - получение статус кода ошибки.
func (instance internalFieldsWebWs) StatusCode() (status entities_ws.StatusCode) {
	return instance.wsStatusCode
}
