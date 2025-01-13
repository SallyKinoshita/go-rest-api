package config

import (
	"context"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Environment Environment
	App         App
	DB          DB
	AWS         AWS
}

func New(ctx context.Context) (*Config, error) {
	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	if err := cfg.Environment.validateContains(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
