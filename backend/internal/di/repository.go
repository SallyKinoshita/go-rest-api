package di

import (
	repository "github.com/SallyKinoshita/go-rest-api/backend/internal/domain/repository"
	persistencerepository "github.com/SallyKinoshita/go-rest-api/backend/internal/infrastructure/persistence/repository"
	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	persistencerepository.NewSampleUser,

	wire.Bind(new(repository.SampleUser), new(*persistencerepository.SampleUser)),
)
