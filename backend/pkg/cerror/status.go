package cerror

import "net/http"

type status int

const (
	// 100番台
	StatusContinue           status = http.StatusContinue
	StatusSwitchingProtocols status = http.StatusSwitchingProtocols
	StatusProcessing         status = http.StatusProcessing
	StatusEarlyHints         status = http.StatusEarlyHints

	// 200番台
	StatusOK                   status = http.StatusOK
	StatusCreated              status = http.StatusCreated
	StatusAccepted             status = http.StatusAccepted
	StatusNonAuthoritativeInfo status = http.StatusNonAuthoritativeInfo
	StatusNoContent            status = http.StatusNoContent
	StatusResetContent         status = http.StatusResetContent
	StatusPartialContent       status = http.StatusPartialContent
	StatusMultiStatus          status = http.StatusMultiStatus
	StatusAlreadyReported      status = http.StatusAlreadyReported
	StatusIMUsed               status = http.StatusIMUsed

	// 300番台
	StatusMultipleChoices   status = http.StatusMultipleChoices
	StatusMovedPermanently  status = http.StatusMovedPermanently
	StatusFound             status = http.StatusFound
	StatusSeeOther          status = http.StatusSeeOther
	StatusNotModified       status = http.StatusNotModified
	StatusUseProxy          status = http.StatusUseProxy
	StatusTemporaryRedirect status = http.StatusTemporaryRedirect
	StatusPermanentRedirect status = http.StatusPermanentRedirect

	// 400番台
	StatusBadRequest                   status = http.StatusBadRequest
	StatusUnauthorized                 status = http.StatusUnauthorized
	StatusPaymentRequired              status = http.StatusPaymentRequired
	StatusForbidden                    status = http.StatusForbidden
	StatusNotFound                     status = http.StatusNotFound
	StatusMethodNotAllowed             status = http.StatusMethodNotAllowed
	StatusNotAcceptable                status = http.StatusNotAcceptable
	StatusProxyAuthRequired            status = http.StatusProxyAuthRequired
	StatusRequestTimeout               status = http.StatusRequestTimeout
	StatusConflict                     status = http.StatusConflict
	StatusGone                         status = http.StatusGone
	StatusLengthRequired               status = http.StatusLengthRequired
	StatusPreconditionFailed           status = http.StatusPreconditionFailed
	StatusRequestEntityTooLarge        status = http.StatusRequestEntityTooLarge
	StatusRequestURITooLong            status = http.StatusRequestURITooLong
	StatusUnsupportedMediaType         status = http.StatusUnsupportedMediaType
	StatusRequestedRangeNotSatisfiable status = http.StatusRequestedRangeNotSatisfiable
	StatusExpectationFailed            status = http.StatusExpectationFailed
	StatusTeapot                       status = http.StatusTeapot
	StatusMisdirectedRequest           status = http.StatusMisdirectedRequest
	StatusUnprocessableEntity          status = http.StatusUnprocessableEntity
	StatusLocked                       status = http.StatusLocked
	StatusFailedDependency             status = http.StatusFailedDependency
	StatusTooEarly                     status = http.StatusTooEarly
	StatusUpgradeRequired              status = http.StatusUpgradeRequired
	StatusPreconditionRequired         status = http.StatusPreconditionRequired
	StatusTooManyRequests              status = http.StatusTooManyRequests
	StatusRequestHeaderFieldsTooLarge  status = http.StatusRequestHeaderFieldsTooLarge
	StatusUnavailableForLegalReasons   status = http.StatusUnavailableForLegalReasons

	// 500番台
	StatusInternalServerError           status = http.StatusInternalServerError
	StatusNotImplemented                status = http.StatusNotImplemented
	StatusBadGateway                    status = http.StatusBadGateway
	StatusServiceUnavailable            status = http.StatusServiceUnavailable
	StatusGatewayTimeout                status = http.StatusGatewayTimeout
	StatusHTTPVersionNotSupported       status = http.StatusHTTPVersionNotSupported
	StatusVariantAlsoNegotiates         status = http.StatusVariantAlsoNegotiates
	StatusInsufficientStorage           status = http.StatusInsufficientStorage
	StatusLoopDetected                  status = http.StatusLoopDetected
	StatusNotExtended                   status = http.StatusNotExtended
	StatusNetworkAuthenticationRequired status = http.StatusNetworkAuthenticationRequired
)

func (s status) code() int {
	return int(s)
}

func (s status) text() string {
	return http.StatusText(int(s))
}
