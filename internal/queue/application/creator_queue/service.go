package application

import (
	"ari-queue/internal/queue/domain"
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
)

type CreatorService struct {
	domain.QueueRepository
}

func (c CreatorService) Create(ctx context.Context, id, name string, agents []uuid.VoId) error {
	queue, err := domain.NewQueue(id, name, agents)
	if err != nil {
		return err
	}
	err = c.QueueRepository.Upsert(ctx, queue)
	if err != nil {
		return err
	}
	return nil
}

func NewCreatorService(repo domain.QueueRepository) CreatorService {
	return CreatorService{
		QueueRepository: repo,
	}
}
