package errors

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"sm_errors/entities"
	entities_grpc "sm_errors/entities/grpc"
	entities_http "sm_errors/entities/http"
	entities_ws "sm_errors/entities/ws"
	"sort"
)

// internal - внутренняя реализация ошибок.
type internal struct {
	id     entities.ID
	status entities.Status
	t      int

	message *entities.Message
	err     error

	grpcStatusCode entities_grpc.StatusCode
	httpStatusCode entities_http.StatusCode
	wsStatusCode   entities_ws.StatusCode
}

// ID - получение идентификатора ошибки.
func (instance internal) ID() (id entities.ID) {
	return instance.id
}

// String - получение строкового представления ошибки.
func (instance internal) String() (str string) {
	str = instance.Message()

	if str == "" {
		str = instance.err.Error()
	}

	return
}

// Status - получение статуса ошибки.
func (instance internal) Status() (status entities.Status) {
	return instance.status
}

// Message - получение сообщения ошибки.
func (instance internal) Message(options ...entities.MessageOption) (message string) {
	if instance.err != nil {
		message = instance.err.Error()
	}

	if instance.message != nil {
		message = instance.message.String(options...)
	}

	return
}

// Error - получение единой абстракции ошибок.
func (instance internal) Error() (err string) {
	return instance.err.Error()
}

// Is - проверка соответствия исходной ошибки.
func (instance internal) Is(err error) bool {
	return errors.Is(instance.err, err)
}

// ToGrpc - преобразования ошибки в grpc формат.
func (instance internal) ToGrpc() Grpc {
	return Grpc(internalGrpc{
		internal{
			id:     instance.id,
			status: instance.status,
			t:      typeGrpc,

			message: instance.message,
			err:     instance.err,

			grpcStatusCode: instance.grpcStatusCode,
		},
	})
}

// ToWeb - преобразования ошибки в web формат.
func (instance internal) ToWeb() Web {
	return Web(internalWeb{
		internal{
			id:     instance.id,
			status: instance.status,
			t:      typeWeb,

			message: instance.message,
			err:     instance.err,

			httpStatusCode: instance.httpStatusCode,
			wsStatusCode:   instance.wsStatusCode,
		},
	})
}

// ToBasic - преобразования ошибки в базовый формат.
func (instance internal) ToBasic() Error {
	return Error(internal{
		id:     instance.id,
		status: instance.status,
		t:      typeBasic,

		message: instance.message,
		err:     instance.err,
	})
}

// toMap - упаковка в map.
func (instance internal) toMap() (data map[string]any) {
	data = map[string]any{
		"status":  instance.status.String(),
		"message": instance.message.String(),

		"error": map[string]any{
			"id":      instance.id,
			"type":    "basic",
			"message": instance.err.Error(),
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
func (instance internal) MarshalJSON() (data []byte, err error) {
	return json.Marshal(instance.toMap())
}

// MarshalYAML - упаковка в YAML.
func (instance internal) MarshalYAML() (data any, err error) {
	return instance.toMap(), nil
}

// MarshalXML - упаковка в XML.
func (instance internal) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	var (
		tokens = make([]xml.Token, 0)
		encode func(e *xml.Encoder, input map[string]any, parent string)
	)

	// encode
	{
		encode = func(e *xml.Encoder, input map[string]any, parent string) {
			var keys []string

			// keys
			{
				for key, _ := range input {
					keys = append(keys, key)
				}

				sort.SliceStable(keys, func(i int, j int) bool {
					return keys[i] < keys[j]
				})
			}

			// Обработка
			{
				if parent != "" {
					tokens = append(tokens, xml.StartElement{Name: xml.Name{Local: parent}})
				}

				for _, key := range keys {
					var value = input[key]

					var (
						startEl = xml.StartElement{Name: xml.Name{Local: key}}
						endEl   = xml.EndElement{Name: xml.Name{Local: key}}
					)

					if reflect.TypeOf(value).String() == "map[string]interface {}" {
						encode(e, value.(map[string]any), key)
						continue
					}

					tokens = append(tokens, startEl, xml.CharData(fmt.Sprintf("%s", value)), endEl)
				}

				if parent != "" {
					tokens = append(tokens, xml.EndElement{Name: xml.Name{Local: parent}})
				}
			}
		}
	}

	encode(e, instance.toMap(), "")

	for _, t := range tokens {
		if err = e.EncodeToken(t); err != nil {
			return
		}
	}

	return e.Flush()
}
