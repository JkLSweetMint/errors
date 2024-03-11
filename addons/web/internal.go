package web_addon

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
	web_http_addon "sm-errors/addons/web/http"
	web_ws_addon "sm-errors/addons/web/ws"
	"sm-errors/entities"
	"sm-errors/internal"
	"sort"
)

// WebAddon - внутренняя реализация дополнения ошибки.
type WebAddon struct {
	*internal.Internal

	*WebAddonOptions
}

// WebAddonOptions - опции внутренней реализации ошибки.
type WebAddonOptions struct {
	*web_http_addon.WebHttpAddonOptions
	*web_ws_addon.WebWsAddonOptions
}

// ToHttp - преобразования ошибки в web http формат.
func (instance WebAddon) ToHttp() Http {
	return Http(web_http_addon.WebHttpAddon{
		Internal:            instance.Internal,
		WebHttpAddonOptions: instance.WebHttpAddonOptions,
	})
}

// ToWs - преобразования ошибки в web ws формат.
func (instance WebAddon) ToWs() Ws {
	return Ws(web_ws_addon.WebWsAddon{
		Internal:          instance.Internal,
		WebWsAddonOptions: instance.WebWsAddonOptions,
	})
}

// toMap - упаковка в map.
func (instance WebAddon) toMap() (data map[string]any) {
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
func (instance WebAddon) MarshalJSON() (data []byte, err error) {
	return json.Marshal(instance.toMap())
}

// MarshalYAML - упаковка в YAML.
func (instance WebAddon) MarshalYAML() (data any, err error) {
	return instance.toMap(), nil
}

// MarshalXML - упаковка в XML.
func (instance WebAddon) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
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
