package middleware

import (
	"context"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/ccontext"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cuuid"
	"github.com/labstack/echo/v4"
)

type CLog struct{}

func newCLog() *CLog {
	return &CLog{}
}

func (c *CLog) Set(ctx context.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			traceID, err := cuuid.New()
			if err != nil {
				return err
			}

			ctx = ccontext.SetTraceID(ctx, traceID)
			ctx = ccontext.SetNowTime(ctx)
			ctx = ccontext.SetHTTPMethod(ctx, ec.Request().Method)
			ctx = ccontext.SetURL(ctx, ec.Request().URL.String())

			ec.SetRequest(ec.Request().WithContext(ctx))

			return next(ec)
		}
	}
}
