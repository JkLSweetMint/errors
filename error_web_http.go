package errors

import (
	entities_http "sm_errors/entities/http"
)

// WebHttp - абстракция системы web http ошибок.
type WebHttp interface {
	Error

	StatusCode() (status entities_http.StatusCode)
}

// FieldsWebHttp - абстракция системы web http ошибок с полями.
type FieldsWebHttp interface {
	Fields

	StatusCode() (status entities_http.StatusCode)
}

// ----------------------------------------------------------------

// internalWebHttp - внутрення реализация web http ошибки.
type internalWebHttp struct {
	internal
}

// StatusCode - получение статус кода ошибки.
func (instance internalWebHttp) StatusCode() (status entities_http.StatusCode) {
	return instance.httpStatusCode
}

// ----------------------------------------------------------------

// internalFieldsWebHttp - внутрення реализация web http ошибки.
type internalFieldsWebHttp struct {
	internalFields
}

// StatusCode - получение статус кода ошибок с полями.
func (instance internalFieldsWebHttp) StatusCode() (status entities_http.StatusCode) {
	return instance.httpStatusCode
}
