package ccontext

import (
	"context"
	"time"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
)

type responseTimeCtxKey struct{}

func SetResponseTime(ctx context.Context) context.Context {
	startTime, err := GetNowTime(ctx)
	if err != nil {
		return ctx
	}

	return context.WithValue(ctx, responseTimeCtxKey{}, time.Since(startTime))
}

func GetResponseTime(ctx context.Context) (time.Duration, error) {
	val := ctx.Value(responseTimeCtxKey{})
	if val == nil {
		return 0, cerror.New("failed to get response time", cerror.DetailInternalServerError)
	}

	responseTime, ok := val.(time.Duration)
	if !ok {
		return 0, cerror.New("failed to get response time", cerror.DetailInternalServerError)
	}

	return responseTime, nil
}
