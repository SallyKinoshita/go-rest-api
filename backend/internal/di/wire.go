//go:build wireinject

package di

import (
	"context"

	"github.com/SallyKinoshita/go-rest-api/backend/internal/config"
	"github.com/google/wire"
	"github.com/uptrace/bun"
)

func New(ctx context.Context, cfg *config.Config, db *bun.DB) (DI, error) {
	wire.Build(
		repositorySet,
		useCaseSet,
		controllerSet,
		middlewareSet,
		diSet,
	)
	return DI{}, nil
}
