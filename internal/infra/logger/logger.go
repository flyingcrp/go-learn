package logger

import (
	"context"
	"log/slog"
	"os"

	"go-learn/internal/infra/middleware"

	"github.com/lmittmann/tint"
)

type contextHandler struct {
	slog.Handler
}

func (h *contextHandler) Handle(ctx context.Context, r slog.Record) error {
	if traceID := middleware.GetTraceId(ctx); traceID != "" {
		r.AddAttrs(slog.String("trace_id", traceID))
	}
	return h.Handler.Handle(ctx, r)
}

func Init(level slog.Level) {
	opts := &slog.HandlerOptions{
		Level: level,
	}
	var baseHandler slog.Handler
	if level == slog.LevelDebug {
		baseHandler = tint.NewHandler(os.Stdout, &tint.Options{Level: level})
	} else {
		baseHandler = slog.NewJSONHandler(os.Stdout, opts)
	}
	slog.SetDefault(slog.New(&contextHandler{baseHandler}))
}
