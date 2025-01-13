package middleware

import (
	"context"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/ccontext"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/clog"
	"github.com/labstack/echo/v4"
)

type AccessLog struct{}

func newAccessLog() *AccessLog {
	return &AccessLog{}
}

func (a *AccessLog) Logger(ctx context.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			ctx := ec.Request().Context()

			next(ec)

			ctx = ccontext.SetResponseTime(ctx)

			clog.Info(ctx, "access log")

			return nil
		}
	}
}
