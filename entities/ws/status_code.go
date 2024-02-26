package entities_ws

// StatusCode - статус код web сокетов.
type StatusCode int

// Коды состояния ошибок web сокетов.
const (
	StatusCloseNormalClosure StatusCode = iota + 1000
	StatusCloseGoingAway
	StatusCloseProtocolError
	StatusCloseUnsupportedData
	_
	StatusCloseNoStatusReceived
	StatusCloseAbnormalClosure
	StatusUnsupportedPayload
	StatusPolicyViolation
	StatusMessageTooBig
	StatusMandatoryExtension
	StatusInternalServerError
	StatusServiceRestart
	StatusTryAgainLater
	StatusBadGateway
	StatusTLSHandshakeFail
)

var (
	unknownStatusCode = "Unknown Status Code"

	statusMessages = []string{
		StatusCloseNormalClosure:    "Successful operation / regular socket shutdown. ",
		StatusCloseGoingAway:        "Client is leaving",
		StatusCloseProtocolError:    "Endpoint received a malformed frame",
		StatusCloseUnsupportedData:  "Endpoint received an unsupported frame",
		StatusCloseNoStatusReceived: "Expected close status, received none",
		StatusCloseAbnormalClosure:  "No close code frame has been receieved",
		StatusUnsupportedPayload:    "Endpoint received inconsistent message",
		StatusPolicyViolation:       "Policy violation",
		StatusMessageTooBig:         "Endpoint won't process large frame",
		StatusMandatoryExtension:    "Client wanted an extension which server did not negotiate",
		StatusInternalServerError:   "Internal server error while operating",
		StatusServiceRestart:        "Server/service is restarting",
		StatusTryAgainLater:         "Temporary server condition forced blocking client's request",
		StatusBadGateway:            "Server acting as gateway received an invalid response",
		StatusTLSHandshakeFail:      "Transport Layer Security handshake failure",
	}
)

// String - получение строкового представления http статус кода.
func (code StatusCode) String() (val string) {
	val = unknownStatusCode

	if code < 1000 || code > 1015 {
		return
	}

	if s := statusMessages[code]; s != "" {
		val = s
	}

	return
}
