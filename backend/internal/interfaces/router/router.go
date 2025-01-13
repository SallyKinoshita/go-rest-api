package router

import (
	"context"
	"errors"

	"github.com/SallyKinoshita/go-rest-api/backend/internal/config"
	"github.com/SallyKinoshita/go-rest-api/backend/internal/gen/openapi"
	"github.com/SallyKinoshita/go-rest-api/backend/internal/interfaces/controller"
	"github.com/SallyKinoshita/go-rest-api/backend/internal/middleware"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

type Router struct {
	db         *bun.DB
	cfg        *config.Config
	controller *controller.Controller
	middleware *middleware.Middleware
}

func New(
	db *bun.DB,
	cfg *config.Config,
	controller *controller.Controller,
	middleware *middleware.Middleware,
) *Router {
	return &Router{
		db:         db,
		cfg:        cfg,
		controller: controller,
		middleware: middleware,
	}
}

func (r *Router) SetUp(ctx context.Context) *echo.Echo {
	e := echo.New()

	// HTTPエラーハンドラ
	e.HTTPErrorHandler = func(err error, ec echo.Context) {
		var cerr *cerror.CError
		if errors.As(err, &cerr) {
			if err := ec.JSON(cerr.StatusCode(), cerr.ConvertEchoError()); err != nil {
				return
			}
		}
	}

	// Middleware
	e.Use(echomiddleware.Recover())
	e.Use(r.middleware.CLog.Set(ctx))
	e.Use(r.middleware.ClientUser.CognitoAuth(ctx))
	e.Use(r.middleware.AccessLog.Logger(ctx))
	e.Use(r.middleware.CError.WrapCErrorAndCLogger)

	// Router
	openapi.RegisterHandlers(e, r.controller)

	// 未定義のURLにアクセスされた場合、Not Foundでログを出す
	e.Any("/*", func(c echo.Context) error {
		return cerror.New("not found url", cerror.DetailNotFound)
	})

	return e
}
