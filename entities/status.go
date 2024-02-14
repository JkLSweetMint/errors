package entities

const (
	StatusUnknown Status = iota
	StatusFailed
	StatusError
	StatusFatal
)

// Status - статус ошибки.
type Status int

// String - получение строкового представления статуса ошибки.
func (status Status) String() (str string) {
	switch status {
	case 1:
		str = "failed"
	case 2:
		str = "error"
	case 3:
		str = "fatal"
	default:
		str = "unknown"
	}

	return
}
