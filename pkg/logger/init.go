package logger

import (
	"log/slog"
	"os"

	slogpretty "github.com/Konil-Startup/go-backend/pkg/logger/slogretty"
)

func NewLogger(env string) *slog.Logger {
	var logger *slog.Logger = slog.Default()

	switch env {
	case "local":
		logger = setupPrettySlog()
	case "dev":
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return logger
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
