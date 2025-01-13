package ccontext

import (
	"context"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cuuid"
)

type traceIDCtxKey struct{}

func SetTraceID(ctx context.Context, traceID cuuid.CUUID) context.Context {
	return context.WithValue(ctx, traceIDCtxKey{}, traceID)
}

func GetTraceID(ctx context.Context) (cuuid.CUUID, error) {
	val := ctx.Value(traceIDCtxKey{})
	if val == nil {
		return cuuid.CUUID{}, cerror.New("failed to get trace id", cerror.DetailInternalServerError)
	}

	traceID, ok := val.(cuuid.CUUID)
	if !ok {
		return cuuid.CUUID{}, cerror.New("failed to get trace id", cerror.DetailInternalServerError)
	}

	if err := traceID.Validate(); err != nil {
		return cuuid.CUUID{}, cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	return traceID, nil
}
