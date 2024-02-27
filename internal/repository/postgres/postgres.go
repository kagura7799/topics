package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	UserRepo
	TopicRepo
}

func New(dsn string) (*pgx.Conn, error) {
	const op = "repository.postgres.New"

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to open connection to db: %w", op, err)
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("%s: failed to ping db: %w", op, err)
	}
	return conn, nil
}
