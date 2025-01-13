package di

import (
	"github.com/SallyKinoshita/go-rest-api/backend/internal/interfaces/controller"
	"github.com/SallyKinoshita/go-rest-api/backend/internal/middleware"
	"github.com/google/wire"
)

type DI struct {
	Controller *controller.Controller
	Middleware *middleware.Middleware
}

var diSet = wire.NewSet(
	wire.Struct(new(DI), "*"),
)
