package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Konil-Startup/go-backend/internal/models"
	internalErrors "github.com/Konil-Startup/go-backend/pkg/errors"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (s *UserRepo) SaveUser(ctx context.Context, user *models.User) error {
	const op = "repository.postgres.SaveUser"
	_, err := s.db.Exec(ctx, "insert into users (name, email, pass_hash) values ($1, $2, $3)", user.Name, user.Email, user.Hash)
	if err != nil {
		// TODO: check for dublicate user
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *UserRepo) UserByEmail(ctx context.Context, email string) (*models.User, error) {
	const op = "repository.postgres.UserByEmail"

	row := s.db.QueryRow(ctx, "select id, name, email, pass_hash from users where email = $1")

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Hash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s: %w", op, internalErrors.ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (s *UserRepo) UserByID(ctx context.Context, id int) (*models.User, error) {
	const op = "repository.postgres.UserByID"

	row := s.db.QueryRow(ctx, "select id, name, email, pass_hash from users where id = $1")

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Hash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s: %w", op, internalErrors.ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
