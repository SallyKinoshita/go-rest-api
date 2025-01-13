package middleware

import (
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/clog"
	"github.com/labstack/echo/v4"
)

type CError struct{}

func newCError() *CError {
	return &CError{}
}

func (c *CError) WrapCErrorAndCLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ec echo.Context) error {
		err := next(ec)
		if err != nil {
			ctx := ec.Request().Context()
			cerr := cerror.ConvertCError(err)

			switch {
			case cerr.ContainsStatus(
				// 500番台
				cerror.StatusInternalServerError, cerror.StatusNotImplemented, cerror.StatusBadGateway, cerror.StatusServiceUnavailable, cerror.StatusGatewayTimeout, cerror.StatusHTTPVersionNotSupported, cerror.StatusVariantAlsoNegotiates, cerror.StatusInsufficientStorage, cerror.StatusLoopDetected, cerror.StatusNotExtended, cerror.StatusNetworkAuthenticationRequired,
				// 400番台
				cerror.StatusBadRequest, cerror.StatusPaymentRequired, cerror.StatusNotAcceptable, cerror.StatusProxyAuthRequired, cerror.StatusRequestTimeout, cerror.StatusLengthRequired, cerror.StatusPreconditionFailed, cerror.StatusRequestEntityTooLarge, cerror.StatusRequestURITooLong, cerror.StatusUnsupportedMediaType, cerror.StatusRequestedRangeNotSatisfiable, cerror.StatusExpectationFailed, cerror.StatusTeapot, cerror.StatusMisdirectedRequest, cerror.StatusUnprocessableEntity, cerror.StatusLocked, cerror.StatusFailedDependency, cerror.StatusTooEarly, cerror.StatusUpgradeRequired, cerror.StatusPreconditionRequired, cerror.StatusRequestHeaderFieldsTooLarge, cerror.StatusUnavailableForLegalReasons,
			):
				clog.Error(ctx, cerr)
			case cerr.ContainsStatus(
				// 400番台
				cerror.StatusUnauthorized, cerror.StatusForbidden, cerror.StatusNotFound, cerror.StatusMethodNotAllowed, cerror.StatusConflict, cerror.StatusGone, cerror.StatusTooManyRequests,
			):
				clog.WarnWithError(ctx, cerr)
			default:
				clog.InfoWithError(ctx, cerr)
			}

			return ec.JSON(cerr.StatusCode(), cerr.ConvertEchoError())
		}

		return nil
	}
}
