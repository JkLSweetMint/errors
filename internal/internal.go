package internal

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"sm-errors/entities"
	"sort"
)

// Internal - внутренняя реализация базовой ошибки.
type Internal struct {
	id     entities.ID
	status entities.Status
	t      string

	fields  entities.Fields
	message *entities.Message
	err     error
}

// ID - получение идентификатора ошибки.
func (instance Internal) ID() (id entities.ID) {
	return instance.id
}

// Type - получение типа ошибки.
func (instance Internal) Type() (t string) {
	return instance.t
}

// String - получение строкового представления ошибки.
func (instance Internal) String() (str string) {
	str = instance.Message()

	if str == "" {
		str = instance.err.Error()
	}

	return
}

// Status - получение статуса ошибки.
func (instance Internal) Status() (status entities.Status) {
	return instance.status
}

// Fields - получение полей.
func (instance Internal) Fields() (internalFields entities.Fields) {
	return instance.fields
}

// Message - получение сообщения ошибки.
func (instance Internal) Message(options ...entities.MessageOption) (message string) {
	if instance.err != nil {
		message = instance.err.Error()
	}

	if instance.message != nil {
		message = instance.message.String(options...)
	}

	return
}

// Error - получение единой абстракции ошибок.
func (instance Internal) Error() (err string) {
	if instance.err != nil {
		err = instance.err.Error()
	}

	return
}

// Is - проверка соответствия исходной ошибки.
func (instance Internal) Is(err error) bool {
	return errors.Is(instance.err, err)
}

// SetError - установка внутренней ошибки.
func (instance Internal) SetError(err error) Internal {
	instance.err = err
	return instance
}

// toMap - упаковка в map.
func (instance Internal) toMap() (data map[string]any) {
	data = map[string]any{
		"status":  instance.Status().String(),
		"message": instance.Message(),

		"error": map[string]any{
			"id":      instance.ID(),
			"type":    instance.Type(),
			"message": instance.Error(),
			"fields":  instance.Fields(),
		},
	}

	if v := data["error"].(map[string]any); len(v["fields"].(entities.Fields)) == 0 {
		delete(v, "fields")
	}

	return
}

// MarshalJSON - упаковка в JSON.
func (instance Internal) MarshalJSON() (data []byte, err error) {
	return json.Marshal(instance.toMap())
}

// MarshalYAML - упаковка в YAML.
func (instance Internal) MarshalYAML() (data any, err error) {
	return instance.toMap(), nil
}

// MarshalXML - упаковка в XML.
func (instance Internal) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
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
