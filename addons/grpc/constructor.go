package grpc_addon

// Constructor - внутренняя реализация конструктора дополнения ошибки.
type Constructor struct {
	StatusCode StatusCode
}

// Clone - получение копии.
func (constructor Constructor) Clone() *Constructor {
	return &Constructor{
		StatusCode: constructor.StatusCode,
	}
}

// Options - сбор опций ошибки.
func (constructor Constructor) Options() *GrpcAddonOptions {
	return &GrpcAddonOptions{
		GrpcStatusCode: constructor.StatusCode,
	}
}
