package errors

// Universal - универсальная абстракция системы ошибок.
type Universal interface {
	Error

	ToWeb() Web
	ToGrpc() Grpc
	ToBasic() Error
}

// FieldsUniversal - универсальная абстракция системы ошибок с полями.
type FieldsUniversal interface {
	Fields

	ToWeb() FieldsWeb
	ToGrpc() FieldsGrpc
	ToBasic() Fields
}
