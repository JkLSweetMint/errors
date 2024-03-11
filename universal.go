package errors

import (
	grpc_addon "sm-errors/addons/grpc"
	web_addon "sm-errors/addons/web"
	internal "sm-errors/internal"
)

// universal - внутренняя реализация универсальной ошибки.
type universal struct {
	*internal.Internal

	*grpc_addon.GrpcAddonOptions
	*web_addon.WebAddonOptions
}

// ToBasic - преобразования ошибки в базовый формат.
func (instance universal) ToBasic() Error {
	return Error(instance.Internal)
}

// ToGrpc - преобразования ошибки в grpc формат.
func (instance universal) ToGrpc() Grpc {
	return Grpc(grpc_addon.GrpcAddon{
		Internal:         instance.Internal,
		GrpcAddonOptions: instance.GrpcAddonOptions,
	})
}

// ToWeb - преобразования ошибки в web формат.
func (instance universal) ToWeb() Web {
	return Web(web_addon.WebAddon{
		Internal:        instance.Internal,
		WebAddonOptions: instance.WebAddonOptions,
	})
}
