package postgres

import (
	"context"
	"fmt"
	
	"github.com/Konil-Startup/go-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type TopicRepo struct {
	db *pgx.Conn
}

func NewTopicRepo(db *pgx.Conn) *TopicRepo {
	return &TopicRepo{
		db: db,
	}
}

func (t *TopicRepo) CreateTopic(ctx context.Context, topic *models.Topic) error {
	const op = "repository.postgres.CreateTopic"
	_, err := t.db.Exec(ctx, "insert into topic (name, description, avatar_url, created_at) values ($1, $2, $3, $4)", topic.Name, topic.Description, topic.AvatarURL, topic.CreatedAt)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (t *TopicRepo) DeleteTopicByID(ctx context.Context, id int) error {
	const op = "repository.postgres.DeleteTopicByID"
	_, err := t.db.Exec(ctx, "delete from topic where id = $1", id)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
