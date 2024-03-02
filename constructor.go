package errors

import (
	grpc_addon "sm_errors/addons/grpc"
	web_addon "sm_errors/addons/web"
	"sm_errors/entities"
	"sm_errors/internal"
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

// Clone - получение копии.
func (constructor Constructor) Clone() Constructor {
	var newConstructor = Constructor{
		ID:     constructor.ID,
		Status: constructor.Status,

		Err:     constructor.Err,
		Message: nil,
		Fields:  constructor.Fields,

		Grpc: constructor.Grpc.Clone(),
		Web:  constructor.Web.Clone(),
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
