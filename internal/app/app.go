package app

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Konil-Startup/go-backend/internal/config"
	"github.com/Konil-Startup/go-backend/internal/repository"
	"github.com/Konil-Startup/go-backend/internal/repository/postgres"
	"github.com/Konil-Startup/go-backend/internal/service"
	"github.com/Konil-Startup/go-backend/internal/transport/rest"
)

type Application struct {
	Cfg    *config.Config
	Server *http.Server
	Logger *slog.Logger
	Repo   postgres.Repository
}

func New(cfg *config.Config, l *slog.Logger) (*Application, error) {
	const op = "app.New"

	dbConn, err := postgres.New(cfg.Dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %s: %w", op, "failed to init db", err)
	}

	l.Info("Successfully connected to db", slog.Any("db_stats", dbConn.Config()))

	repo := repository.NewPostgresRepo(dbConn)

	service := service.New(repo)
	restHandler := rest.New(service, l)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%v", cfg.Port),
		Handler: restHandler.Routes(),
	}

	return &Application{
		Logger: l,
		Cfg:    cfg,
		Server: s,
	}, nil
}

func (a *Application) Run() error {
	return a.Server.ListenAndServe()
}
