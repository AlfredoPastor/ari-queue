package application

import (
	"ari-queue/internal/queue/domain"
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
)

type CreatorService struct{}

func (c CreatorService) Create(ctx context.Context, id, name string, agents []uuid.VoId) error {
	queue, err := domain.NewQueue(ctx, id, name, agents)
	if err != nil {
		return err
	}

	return nil
}

func NewCreatorService() CreatorService {
	return CreatorService{}
}
