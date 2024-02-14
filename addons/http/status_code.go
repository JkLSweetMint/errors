package sm_errors_http

// StatusCode - http статус код.
type StatusCode int

// Коды состояния HTTP были украдены из net/http.
const (
	StatusNoContent       StatusCode = 204 // RFC 7231, 6.3.5
	StatusAlreadyReported StatusCode = 208 // RFC 5842, 7.1

	StatusMultipleChoices   StatusCode = 300 // RFC 7231, 6.4.1
	StatusMovedPermanently  StatusCode = 301 // RFC 7231, 6.4.2
	StatusFound             StatusCode = 302 // RFC 7231, 6.4.3
	StatusSeeOther          StatusCode = 303 // RFC 7231, 6.4.4
	StatusNotModified       StatusCode = 304 // RFC 7232, 4.1
	StatusUseProxy          StatusCode = 305 // RFC 7231, 6.4.5
	_                       StatusCode = 306 // RFC 7231, 6.4.6 (Unused)
	StatusTemporaryRedirect StatusCode = 307 // RFC 7231, 6.4.7
	StatusPermanentRedirect StatusCode = 308 // RFC 7538, 3

	StatusBadRequest                   StatusCode = 400 // RFC 7231, 6.5.1
	StatusUnauthorized                 StatusCode = 401 // RFC 7235, 3.1
	StatusPaymentRequired              StatusCode = 402 // RFC 7231, 6.5.2
	StatusForbidden                    StatusCode = 403 // RFC 7231, 6.5.3
	StatusNotFound                     StatusCode = 404 // RFC 7231, 6.5.4
	StatusMethodNotAllowed             StatusCode = 405 // RFC 7231, 6.5.5
	StatusNotAcceptable                StatusCode = 406 // RFC 7231, 6.5.6
	StatusProxyAuthRequired            StatusCode = 407 // RFC 7235, 3.2
	StatusRequestTimeout               StatusCode = 408 // RFC 7231, 6.5.7
	StatusConflict                     StatusCode = 409 // RFC 7231, 6.5.8
	StatusGone                         StatusCode = 410 // RFC 7231, 6.5.9
	StatusLengthRequired               StatusCode = 411 // RFC 7231, 6.5.10
	StatusPreconditionFailed           StatusCode = 412 // RFC 7232, 4.2
	StatusRequestEntityTooLarge        StatusCode = 413 // RFC 7231, 6.5.11
	StatusRequestURITooLong            StatusCode = 414 // RFC 7231, 6.5.12
	StatusUnsupportedMediaType         StatusCode = 415 // RFC 7231, 6.5.13
	StatusRequestedRangeNotSatisfiable StatusCode = 416 // RFC 7233, 4.4
	StatusExpectationFailed            StatusCode = 417 // RFC 7231, 6.5.14
	StatusTeapot                       StatusCode = 418 // RFC 7168, 2.3.3
	StatusMisdirectedRequest           StatusCode = 421 // RFC 7540, 9.1.2
	StatusUnprocessableEntity          StatusCode = 422 // RFC 4918, 11.2
	StatusLocked                       StatusCode = 423 // RFC 4918, 11.3
	StatusFailedDependency             StatusCode = 424 // RFC 4918, 11.4
	StatusUpgradeRequired              StatusCode = 426 // RFC 7231, 6.5.15
	StatusPreconditionRequired         StatusCode = 428 // RFC 6585, 3
	StatusTooManyRequests              StatusCode = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  StatusCode = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   StatusCode = 451 // RFC 7725, 3

	StatusInternalServerError           StatusCode = 500 // RFC 7231, 6.6.1
	StatusNotImplemented                StatusCode = 501 // RFC 7231, 6.6.2
	StatusBadGateway                    StatusCode = 502 // RFC 7231, 6.6.3
	StatusServiceUnavailable            StatusCode = 503 // RFC 7231, 6.6.4
	StatusGatewayTimeout                StatusCode = 504 // RFC 7231, 6.6.5
	StatusHTTPVersionNotSupported       StatusCode = 505 // RFC 7231, 6.6.6
	StatusVariantAlsoNegotiates         StatusCode = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           StatusCode = 507 // RFC 4918, 11.5
	StatusLoopDetected                  StatusCode = 508 // RFC 5842, 7.2
	StatusNotExtended                   StatusCode = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired StatusCode = 511 // RFC 6585, 6
)

var (
	unknownStatusCode = "Unknown Status Code"

	statusMessages = []string{
		StatusNoContent:       "No Content",
		StatusAlreadyReported: "Already Reported",

		StatusMultipleChoices:   "Multiple Choices",
		StatusMovedPermanently:  "Moved Permanently",
		StatusFound:             "Found",
		StatusSeeOther:          "See Other",
		StatusNotModified:       "Not Modified",
		StatusUseProxy:          "Use Proxy",
		StatusTemporaryRedirect: "Temporary Redirect",
		StatusPermanentRedirect: "Permanent Redirect",

		StatusBadRequest:                   "Bad Request",
		StatusUnauthorized:                 "Unauthorized",
		StatusPaymentRequired:              "Payment Required",
		StatusForbidden:                    "Forbidden",
		StatusNotFound:                     "Not Found",
		StatusMethodNotAllowed:             "Method Not Allowed",
		StatusNotAcceptable:                "Not Acceptable",
		StatusProxyAuthRequired:            "Proxy Authentication Required",
		StatusRequestTimeout:               "Request Timeout",
		StatusConflict:                     "Conflict",
		StatusGone:                         "Gone",
		StatusLengthRequired:               "Length Required",
		StatusPreconditionFailed:           "Precondition Failed",
		StatusRequestEntityTooLarge:        "Request Entity Too Large",
		StatusRequestURITooLong:            "Request URI Too Long",
		StatusUnsupportedMediaType:         "Unsupported Media Type",
		StatusRequestedRangeNotSatisfiable: "Requested Range Not Satisfiable",
		StatusExpectationFailed:            "Expectation Failed",
		StatusTeapot:                       "I'm a teapot",
		StatusMisdirectedRequest:           "Misdirected Request",
		StatusUnprocessableEntity:          "Unprocessable Entity",
		StatusLocked:                       "Locked",
		StatusFailedDependency:             "Failed Dependency",
		StatusUpgradeRequired:              "Upgrade Required",
		StatusPreconditionRequired:         "Precondition Required",
		StatusTooManyRequests:              "Too Many Requests",
		StatusRequestHeaderFieldsTooLarge:  "Request Header Fields Too Large",
		StatusUnavailableForLegalReasons:   "Unavailable For Legal Reasons",

		StatusInternalServerError:           "Internal Server Error",
		StatusNotImplemented:                "Not Implemented",
		StatusBadGateway:                    "Bad Gateway",
		StatusServiceUnavailable:            "Service Unavailable",
		StatusGatewayTimeout:                "Gateway Timeout",
		StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
		StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
		StatusInsufficientStorage:           "Insufficient Storage",
		StatusLoopDetected:                  "Loop Detected",
		StatusNotExtended:                   "Not Extended",
		StatusNetworkAuthenticationRequired: "Network Authentication Required",
	}
)

// String - получение строкового представления http статус кода.
func (code StatusCode) String() (val string) {
	val = unknownStatusCode

	if code < 100 || code > 511 {
		return
	}

	if s := statusMessages[code]; s != "" {
		val = s
	}

	return
}
