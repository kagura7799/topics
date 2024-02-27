package service

import (
	"errors"

	"github.com/Konil-Startup/go-backend/internal/repository"
)

var (
	ErrEmpty = errors.New("empty")
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return Service{
		repo: repo,
	}
}
