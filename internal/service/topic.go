package service

import (
	"context"
	"fmt"

	"github.com/Konil-Startup/go-backend/internal/models"
)

func (s *Service) CreateTopic(ctx context.Context, topic *models.Topic) error {
	const op = "service.CreateTopic"

	emptyTopic := models.Topic{}
	if topic == nil || *topic == emptyTopic {
		return ErrEmpty
	}

	if err := s.repo.Topic.CreateTopic(ctx, topic); err != nil {
        return fmt.Errorf("%s: %w", op, err)
    }

	return nil
}

func (s *Service) DeleteTopicByID(ctx context.Context, id int) error {
	const op = "service.DeleteTopicByID"

	if err := s.repo.Topic.DeleteTopicByID(ctx, id); err != nil {
        return fmt.Errorf("%s: %w", op, err)
    }

	return nil
}