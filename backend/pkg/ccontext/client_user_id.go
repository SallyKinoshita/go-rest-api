package ccontext

import (
	"context"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cuuid"
)

type clientUserIDCtxKey struct{}

func SetClientUserID(ctx context.Context, clientUserID cuuid.CUUID) context.Context {
	return context.WithValue(ctx, clientUserIDCtxKey{}, clientUserID)
}

func GetClientUserID(ctx context.Context) (cuuid.CUUID, error) {
	val := ctx.Value(clientUserIDCtxKey{})
	if val == nil {
		return cuuid.CUUID{}, cerror.New("failed to get client user id", cerror.DetailInternalServerError)
	}

	clientUserID, ok := val.(cuuid.CUUID)
	if !ok {
		return cuuid.CUUID{}, cerror.New("failed to get client user id", cerror.DetailInternalServerError)
	}

	if err := clientUserID.Validate(); err != nil {
		return cuuid.CUUID{}, err
	}

	return clientUserID, nil
}
