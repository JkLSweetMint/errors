package errors

import (
	grpc_addon "sm-errors/addons/grpc"
	web_addon "sm-errors/addons/web"
	web_http_addon "sm-errors/addons/web/http"
	web_ws_addon "sm-errors/addons/web/ws"
	"sm-errors/entities"
	"sm-errors/internal"
)

// Constructor - конструктор ошибок.
type Constructor struct {
	ID     entities.ID
	Status entities.Status

	Err     error
	Message *entities.Message
	Fields  entities.Fields

	Grpc *grpc_addon.Constructor
	Web  *web_addon.Constructor
}

// Build - сбор универсальной ошибки.
func (constructor Constructor) Build() Universal {
	var t = "basic"

	if len(constructor.Fields) > 0 {
		t = "fields"
	}

	constructor = constructor.FillEmptyFields()

	return universal{
		Internal: internal.Constructor{
			ID:     constructor.ID,
			Status: constructor.Status,
			Type:   t,

			Err:     constructor.Err,
			Message: constructor.Message,
			Fields:  constructor.Fields,
		}.Build(),

		GrpcAddonOptions: constructor.Grpc.Options(),
		WebAddonOptions:  constructor.Web.Options(),
	}
}

// FillEmptyFields - заполнение пустых полей.
func (constructor Constructor) FillEmptyFields() Constructor {
	if constructor.Grpc == nil {
		constructor.Grpc = new(grpc_addon.Constructor)
	}

	if constructor.Web == nil {
		constructor.Web = new(web_addon.Constructor)
	}

	if constructor.Web.Http == nil {
		constructor.Web.Http = new(web_http_addon.Constructor)
	}

	if constructor.Web.Ws == nil {
		constructor.Web.Ws = new(web_ws_addon.Constructor)
	}

	return constructor
}

// Clone - получение копии.
func (constructor Constructor) Clone() Constructor {
	var newConstructor = Constructor{
		ID:     constructor.ID,
		Status: constructor.Status,

		Err:     constructor.Err,
		Message: nil,
		Fields:  constructor.Fields,
	}

	if constructor.Grpc != nil {
		newConstructor.Grpc = constructor.Grpc.Clone()
	}

	if constructor.Web != nil {
		newConstructor.Web = constructor.Web.Clone()
	}

	if constructor.Message != nil {
		newConstructor.Message = constructor.Message.Clone()
	}

	return newConstructor
}

// SetError - установка внутренней ошибки.
func (constructor Constructor) SetError(err error) Constructor {
	var newConstructor = constructor.Clone()
	newConstructor.Err = err

	return newConstructor
}

// SetFields - установка значение полей.
func (constructor Constructor) SetFields(internalFields ...entities.Field) Constructor {
	var newConstructor = constructor.Clone()
	newConstructor.Fields = internalFields

	return newConstructor
}

// SetField - установка значения поля.
func (constructor Constructor) SetField(key, message string) Constructor {
	var newConstructor = constructor.Clone()

	newConstructor.Fields = append(constructor.Fields, entities.Field{
		Key:     key,
		Message: message,
	})

	return newConstructor

}
