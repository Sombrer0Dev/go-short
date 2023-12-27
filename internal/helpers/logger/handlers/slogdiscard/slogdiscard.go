package slogdiscard

import (
	"context"
	"log/slog"
)

func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

type DiscardHandler struct{}

func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	// ignore record
	return nil
}

func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	// return the same handler
	return h
}

func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	return h
}
