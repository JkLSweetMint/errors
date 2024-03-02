package web_addon

import (
	web_http_addon "sm_errors/addons/web/http"
	web_ws_addon "sm_errors/addons/web/ws"
	"sm_errors/entities"
)

// Http - абстракция системы web http ошибок.
type Http interface {
	ID() (id entities.ID)
	Is(err error) bool
	String() (str string)
	Error() (err string)
	Status() (status entities.Status)

	Message(options ...entities.MessageOption) (message string)
	Fields() (fields entities.Fields)

	StatusCode() (status web_http_addon.StatusCode)
}

// Ws - абстракция системы web ws ошибок.
type Ws interface {
	ID() (id entities.ID)
	Is(err error) bool
	String() (str string)
	Error() (err string)
	Status() (status entities.Status)

	Message(options ...entities.MessageOption) (message string)
	Fields() (fields entities.Fields)

	StatusCode() (status web_ws_addon.StatusCode)
}
