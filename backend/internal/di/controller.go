package di

import (
	"github.com/SallyKinoshita/go-rest-api/backend/internal/interfaces/controller"
	"github.com/google/wire"
)

var controllerSet = wire.NewSet(
	controller.New,
)
