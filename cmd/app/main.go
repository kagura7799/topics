package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Konil-Startup/go-backend/internal/app"
	"github.com/Konil-Startup/go-backend/internal/config"
	"github.com/Konil-Startup/go-backend/pkg/logger"
	"github.com/Konil-Startup/go-backend/pkg/logger/sl"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to init config: %s\n", err)
	}

	l := logger.NewLogger(cfg.Env)

	application, err := app.New(cfg, l)
	if err != nil {
		l.Error("Failed to init app", sl.Err(err))
		os.Exit(1)
	}
	go func() {
		if err := application.Run(); err != nil {
			application.Logger.Error("failed stopped:", sl.Err(err))
		}
	}()

	

	application.Logger.Info("Application started",
		slog.String("env", cfg.Env),
		slog.Any("port", cfg.Port),
	)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT)
	<-sc
}
