package repository

import (
	"context"

	"github.com/SallyKinoshita/go-rest-api/backend/internal/domain/model"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cuuid"
)

type SampleUser interface {
	Create(ctx context.Context, conn DBConn, sampleUser *model.SampleUser) error
	Update(ctx context.Context, conn DBConn, sampleUser *model.SampleUser) error
	Get(ctx context.Context, conn DBConn, sampleUserID cuuid.CUUID) (*model.SampleUser, error)
	List(ctx context.Context, conn DBConn, pagination model.Pagination) (model.SampleUsers, error)
	Delete(ctx context.Context, conn DBConn, sampleUserID cuuid.CUUID) error
	GetIDByCognitoUserID(ctx context.Context, conn DBConn, cognitoUserID string) (cuuid.CUUID, error)
}
