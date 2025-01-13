package di

import (
	"github.com/SallyKinoshita/go-rest-api/backend/internal/middleware"
	"github.com/google/wire"
)

var middlewareSet = wire.NewSet(
	middleware.New,
)
