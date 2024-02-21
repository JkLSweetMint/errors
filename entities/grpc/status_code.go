package entities_grpc

// StatusCode - статус код grpc.
type StatusCode int

// Коды состояния ошибок grpc.
const (
	StatusCanceled StatusCode = iota + 1
	StatusUnknown
	StatusInvalidArgument
	StatusDeadlineExceeded
	StatusNotFound
	StatusAlreadyExists
	StatusPermissionDenied
	StatusResourceExhausted
	StatusFailedPrecondition
	StatusAborted
	StatusOutOfRange
	StatusUnimplemented
	StatusInternalServerError
	StatusUnavailable
	StatusDataLoss
	StatusUnauthenticated
)

var (
	unknownStatusCode = "Unknown Status Code"

	statusMessages = []string{
		StatusCanceled:            "Canceled",
		StatusUnknown:             "Unknown",
		StatusInvalidArgument:     "InvalidArgument",
		StatusDeadlineExceeded:    "DeadlineExceeded",
		StatusNotFound:            "NotFound",
		StatusAlreadyExists:       "AlreadyExists",
		StatusPermissionDenied:    "PermissionDenied",
		StatusResourceExhausted:   "ResourceExhausted",
		StatusFailedPrecondition:  "FailedPrecondition",
		StatusAborted:             "Aborted",
		StatusOutOfRange:          "OutOfRange",
		StatusUnimplemented:       "Unimplemented",
		StatusInternalServerError: "Internal server error",
		StatusUnavailable:         "Unavailable",
		StatusDataLoss:            "DataLoss",
		StatusUnauthenticated:     "Unauthenticated",
	}
)

// String - получение строкового представления http статус кода.
func (code StatusCode) String() (val string) {
	val = unknownStatusCode

	if code < 1 || code > 16 {
		return
	}

	if s := statusMessages[code]; s != "" {
		val = s
	}

	return
}
