package errors

import (
	grpc_addon "sm_errors/addons/grpc"
	web_addon "sm_errors/addons/web"
	"sm_errors/entities"
)

const (
	typeBasic = iota
	typeGrpc
	typeWeb
	typeWebHttp
	typeWebWs
)

// Universal - единая универсальная абстракция системы ошибок.
type Universal interface {
	Error

	ToBasic() Error
	ToGrpc() Grpc
	ToWeb() Web
}

// Error - стандартная абстракция системы ошибок.
type Error interface {
	ID() (id entities.ID)
	Is(err error) bool
	String() (str string)
	Error() (err string)
	Status() (status entities.Status)

	Message(options ...entities.MessageOption) (message string)
	Fields() (fields entities.Fields)
}

// Grpc - абстракция системы grpc ошибок.
type Grpc interface {
	Error

	StatusCode() (status grpc_addon.StatusCode)
}

// Web - абстракция системы web ошибок.
type Web interface {
	Error

	ToHttp() web_addon.Http
	ToWs() web_addon.Ws
}
