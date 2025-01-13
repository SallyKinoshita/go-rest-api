package ccontext

import (
	"context"
	"time"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
)

type nowTimeCtxKey struct{}

func SetNowTime(ctx context.Context) context.Context {
	return context.WithValue(ctx, nowTimeCtxKey{}, time.Now())
}

func GetNowTime(ctx context.Context) (time.Time, error) {
	val := ctx.Value(nowTimeCtxKey{})
	if val == nil {
		return time.Time{}, cerror.New("failed to get now time", cerror.DetailInternalServerError)
	}

	now, ok := val.(time.Time)
	if !ok {
		return time.Time{}, cerror.New("failed to get now time", cerror.DetailInternalServerError)
	}

	return now, nil
}
