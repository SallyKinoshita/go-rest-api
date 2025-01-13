package di

import (
	"github.com/SallyKinoshita/go-rest-api/backend/internal/application/usecase"
	"github.com/google/wire"
)

var useCaseSet = wire.NewSet(
	usecase.NewSampleUser,

	wire.Bind(new(usecase.SampleUserUseCase), new(*usecase.SampleUser)),
)
