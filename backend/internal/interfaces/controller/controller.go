package controller

import "github.com/SallyKinoshita/go-rest-api/backend/internal/application/usecase"

type Controller struct {
	sampleUserUseCase usecase.SampleUserUseCase
}

func New(
	sampleUserUseCase usecase.SampleUserUseCase,
) *Controller {
	return &Controller{
		sampleUserUseCase: sampleUserUseCase,
	}
}
