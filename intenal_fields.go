package errors

import (
	"encoding/json"
	"sm_errors/entities"
)

// internalFields - внутренняя реализация ошибок с полями.
type internalFields struct {
	internal

	fields entities.Fields
}

// Fields - получение полей.
func (instance internalFields) Fields() (internalFields entities.Fields) {
	return instance.fields
}

// ToGrpc - преобразования ошибки в grpc формат.
func (instance internalFields) ToGrpc() FieldsGrpc {
	instance.t = typeGrpc
	return FieldsGrpc(internalFieldsGrpc{
		internalFields{
			internal: internal{
				id:     instance.id,
				status: instance.status,
				t:      typeGrpc,

				message: instance.message,
				err:     instance.err,

				grpcStatusCode: instance.grpcStatusCode,
			},
			fields: instance.fields,
		},
	})
}

// ToWeb - преобразования ошибки в web формат.
func (instance internalFields) ToWeb() FieldsWeb {
	instance.t = typeWeb
	return FieldsWeb(internalFieldsWeb{
		internalFields{
			internal: internal{
				id:     instance.id,
				status: instance.status,
				t:      typeWeb,

				message: instance.message,
				err:     instance.err,

				httpStatusCode: instance.httpStatusCode,
				wsStatusCode:   instance.wsStatusCode,
			},
			fields: instance.fields,
		},
	})
}

// ToBasic - преобразования ошибки в базовый формат.
func (instance internalFields) ToBasic() Fields {
	instance.t = typeBasic
	return Fields(internalFields{
		internal: internal{
			id:     instance.id,
			status: instance.status,
			t:      typeWeb,

			message: instance.message,
			err:     instance.err,
		},
		fields: instance.fields,
	})
}

// MarshalJSON - упаковка в JSON.
func (instance internalFields) toMap() (data map[string]any) {
	data = map[string]any{
		"status":  instance.status.String(),
		"message": instance.message.String(),

		"error": map[string]any{
			"id":      instance.id,
			"type":    "fields",
			"message": instance.err.Error(),
			"fields":  instance.fields,
		},
	}

	switch instance.t {
	case typeGrpc:
		{
			data["code"] = instance.grpcStatusCode
			data["code_message"] = instance.grpcStatusCode.String()
		}
	case typeWebHttp:
		{
			data["code"] = instance.httpStatusCode
			data["code_message"] = instance.httpStatusCode.String()
		}
	case typeWebWs:
		{
			data["code"] = instance.wsStatusCode
			data["code_message"] = instance.wsStatusCode.String()
		}
	case typeWeb:
		{
		}
	}

	if data["code"] == 0 || data["code"] == nil {
		delete(data, "code")
		data["code_message"] = "Unknown"
	}

	return
}

// MarshalJSON - упаковка в JSON.
func (instance internalFields) MarshalJSON() (data []byte, err error) {
	return json.Marshal(instance.toMap())
}

// MarshalYAML - упаковка в YAML.
func (instance internalFields) MarshalYAML() (data any, err error) {
	return instance.toMap(), nil
}
