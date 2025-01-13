package config

import "github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"

type Environment struct {
	Environment string `envconfig:"ENVIRONMENT" required:"true"`
}

const (
	Local       = "local"
	Development = "development"
	Staging     = "staging"
	Production  = "production"
)

var validEnvironments = map[string]bool{
	Local:       true,
	Development: true,
	Staging:     true,
	Production:  true,
}

func (e *Environment) validateContains() error {
	if _, ok := validEnvironments[e.Environment]; !ok {
		return cerror.New("failed to validate environment", cerror.DetailInternalServerError)
	}
	return nil
}
