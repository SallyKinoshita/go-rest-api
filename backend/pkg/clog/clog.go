package clog

import (
	"context"
	"log/slog"
	"os"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/ccontext"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
)

type clog struct {
	logger      *slog.Logger
	logType     logType
	serviceName string
}

var cl *clog

func New(serviceName string, logType logType) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	child := logger.With(
		slog.String("service_name", serviceName),
		slog.String("log_type", string(logType)),
	)

	cl = &clog{
		logger:      child,
		serviceName: serviceName,
		logType:     logType,
	}

	return nil
}

func (c *clog) logWithLevel(ctx context.Context, msg string, err error, logLevel slog.Level) {
	slog.SetDefault(c.logger)

	logger := c.logger.With(
		slog.String("trace_id", c.getTraceID(ctx)),
		slog.String("http_method", c.getHTTPMethod(ctx)),
		slog.String("url", c.getURL(ctx)),
	)

	if clientUserID, err := ccontext.GetClientUserID(ctx); err == nil {
		logger = logger.With(slog.String("client_user_id", clientUserID.String()))
	}

	if responseTime, err := ccontext.GetResponseTime(ctx); err == nil {
		c.logger = c.logger.With(slog.String("response_time", responseTime.String()))
	}

	if err != nil {
		cerr := cerror.ConvertCError(err)
		logger = logger.With(
			slog.String("error", cerr.Error()),
			slog.String("status_code", cerr.StatusText()),
			slog.String("stack_trace", cerr.StackTrace()),
		)
	}

	switch logLevel {
	case levelInfo:
		logger.LogAttrs(ctx, slog.LevelInfo, msg)
	case levelWarn:
		logger.LogAttrs(ctx, slog.LevelWarn, msg)
	case levelError:
		logger.LogAttrs(ctx, slog.LevelError, msg)
	case levelFatal:
		logger.LogAttrs(ctx, levelFatal, msg)
	default:
		logger.LogAttrs(ctx, slog.LevelDebug, msg)
	}
}

const (
	unknown = "unknown"
)

func (c *clog) getTraceID(ctx context.Context) string {
	traceID, err := ccontext.GetTraceID(ctx)
	if err != nil {
		return unknown
	}

	return traceID.String()
}

func (c *clog) getHTTPMethod(ctx context.Context) string {
	httpMethod, err := ccontext.GetHTTPMethod(ctx)
	if err != nil {
		return unknown
	}

	return httpMethod
}

func (c *clog) getURL(ctx context.Context) string {
	url, err := ccontext.GetURL(ctx)
	if err != nil {
		return unknown
	}

	return url
}
