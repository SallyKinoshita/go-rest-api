package ccontext

import (
	"context"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
)

type httpMethodCtxKey struct{}

func SetHTTPMethod(ctx context.Context, httpMethod string) context.Context {
	return context.WithValue(ctx, httpMethodCtxKey{}, httpMethod)
}

func GetHTTPMethod(ctx context.Context) (string, error) {
	val := ctx.Value(httpMethodCtxKey{})
	if val == nil {
		return "", cerror.New("failed to get http method", cerror.DetailInternalServerError)
	}

	httpMethod, ok := val.(string)
	if !ok {
		return "", cerror.New("failed to get http method", cerror.DetailInternalServerError)
	}

	return httpMethod, nil
}
