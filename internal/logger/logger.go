package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func Get() *slog.Logger {
	if logger == nil {
		panic("logger is nil")
	}

	return logger
}

func Init(debug bool) *slog.Logger {
	if debug {
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	} else {
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return logger
}
