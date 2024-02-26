package errors

// Universal - универсальная абстракция системы ошибок.
type Universal interface {
	Error
	Grpc
	Http
	Ws
}

// FieldsUniversal - универсальная абстракция системы ошибок с полями.
type FieldsUniversal interface {
	Fields
	FieldsGrpc
	FieldsHttp
	FieldsWs
}
