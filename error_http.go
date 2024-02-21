package errors

import (
	entities_http "sm_errors/entities/http"
)

// Http - абстракция системы http ошибок.
type Http interface {
	Error
	HttpStatusCode() (status entities_http.StatusCode)
}

// FieldsHttp - абстракция системы http ошибок с полями.
type FieldsHttp interface {
	Fields
	HttpStatusCode() (status entities_http.StatusCode)
}
