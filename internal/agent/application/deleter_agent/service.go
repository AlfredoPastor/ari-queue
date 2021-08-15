package application

import (
	"ari-queue/internal/agent/domain"
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
)

type DeleterAgentService struct {
	domain.AgentRepository
}

func (c DeleterAgentService) Delete(ctx context.Context, id uuid.VoId) error {
	err := c.AgentRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewDeleterAgentService(repo domain.AgentRepository) DeleterAgentService {
	return DeleterAgentService{
		AgentRepository: repo,
	}
}
