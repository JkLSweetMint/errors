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
	return &Constructor{
		Http: constructor.Http.Clone(),
		Ws:   constructor.Ws.Clone(),
	}
}

// Options - сбор опций ошибки.
func (constructor Constructor) Options() *WebAddonOptions {
	return &WebAddonOptions{
		WebHttpAddonOptions: constructor.Http.Options(),
		WebWsAddonOptions:   constructor.Ws.Options(),
	}
}
