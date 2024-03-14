package web_addon

import (
	web_http_addon "sm-errors/addons/web/http"
	web_ws_addon "sm-errors/addons/web/ws"
)

// Constructor - внутренняя реализация конструктора дополнения ошибки.
type Constructor struct {
	Http *web_http_addon.Constructor
	Ws   *web_ws_addon.Constructor
}

// Clone - получение копии.
func (constructor Constructor) Clone() *Constructor {
	var newConstructor = new(Constructor)

	if constructor.Http != nil {
		newConstructor.Http = constructor.Http.Clone()
	}

	if constructor.Ws != nil {
		newConstructor.Ws = constructor.Ws.Clone()
	}

	return newConstructor
}

// Options - сбор опций ошибки.
func (constructor Constructor) Options() *WebAddonOptions {
	var options = new(WebAddonOptions)

	if constructor.Http != nil {
		options.WebHttpAddonOptions = constructor.Http.Options()
	}

	if constructor.Ws != nil {
		options.WebWsAddonOptions = constructor.Ws.Options()
	}

	return options
}
