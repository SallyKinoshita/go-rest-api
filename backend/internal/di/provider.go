package di

import (
	"github.com/SallyKinoshita/go-rest-api/backend/internal/config"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	config.NewAWS,
	config.NewCognito,
)
