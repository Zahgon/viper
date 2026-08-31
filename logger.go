package viper

import (
	"context"
	"log/slog"
)

func WithLogger(l *slog.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

type discardHandler struct{}

func (n *discardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *discardHandler) Handle(_ context.Context, _ slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *discardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (n *discardHandler) WithGroup(_ string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}
