package errors

import (
	entities_ws "sm_errors/entities/ws"
)

// Ws - абстракция системы ws ошибок.
type Ws interface {
	Error
	WsStatusCode() (status entities_ws.StatusCode)
}

// FieldsWs - абстракция системы ws ошибок с полями.
type FieldsWs interface {
	Fields
	WsStatusCode() (status entities_ws.StatusCode)
}
