package web_ws_addon

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
	"sm-errors/entities"
	"sm-errors/internal"
	"sort"
)

// WebWsAddon - внутренняя реализация дополнения ошибки.
type WebWsAddon struct {
	*internal.Internal
	*WebWsAddonOptions
}

// WebWsAddonOptions - опции внутренней реализации ошибки.
type WebWsAddonOptions struct {
	WebWsStatusCode StatusCode
}

// StatusCode - получение статус кода ошибки.
func (instance WebWsAddon) StatusCode() (status StatusCode) {
	return instance.WebWsStatusCode
}

// toMap - упаковка в map.
func (instance WebWsAddon) toMap() (data map[string]any) {
	data = map[string]any{
		"status":  instance.Status().String(),
		"message": instance.Message(),

		"code":         instance.StatusCode(),
		"code_message": instance.StatusCode().String(),

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

	if data["code"] == 0 || data["code"] == nil {
		delete(data, "code")
		data["code_message"] = "Unknown"
	}

	return
}

// MarshalJSON - упаковка в JSON.
func (instance WebWsAddon) MarshalJSON() (data []byte, err error) {
	return json.Marshal(instance.toMap())
}

// MarshalYAML - упаковка в YAML.
func (instance WebWsAddon) MarshalYAML() (data any, err error) {
	return instance.toMap(), nil
}

// MarshalXML - упаковка в XML.
func (instance WebWsAddon) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
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
