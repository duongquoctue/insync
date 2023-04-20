package logger

import (
	"context"
)

type key string

const logContextKey key = "insync_logger"

func NewLoggerCtx(ctx context.Context) *Logger {
	l, ok := ctx.Value(logContextKey).(*Logger)
	if ok {
		return l
	}

	return logger
}

func AddParamsToLoggerCtx(ctx context.Context, args ...interface{}) context.Context {
	l, ok := ctx.Value(logContextKey).(*Logger)
	if !ok {
		ctx = context.WithValue(ctx, logContextKey, logger)
		return ctx
	}

	l.sugaredLogger = l.sugaredLogger.With(args)
	return context.WithValue(ctx, logContextKey, l)
}

func NewLogger() *Logger {
	if logger != nil {
		return logger
	}

	InitLogger()
	return logger
}
