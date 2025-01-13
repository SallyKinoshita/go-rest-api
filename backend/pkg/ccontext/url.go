package ccontext

import (
	"context"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
)

type urlCtxKey struct{}

func SetURL(ctx context.Context, url string) context.Context {
	return context.WithValue(ctx, urlCtxKey{}, url)
}

func GetURL(ctx context.Context) (string, error) {
	val := ctx.Value(urlCtxKey{})
	if val == nil {
		return "", cerror.New("failed to get url", cerror.DetailInternalServerError)
	}

	url, ok := val.(string)
	if !ok {
		return "", cerror.New("failed to get url", cerror.DetailInternalServerError)
	}

	return url, nil
}
