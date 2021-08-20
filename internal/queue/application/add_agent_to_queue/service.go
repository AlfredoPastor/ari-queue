package add_agent_to_queue

import (
	domainAgent "ari-queue/internal/agent/domain"
	"ari-queue/internal/queue/domain"
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
)

type AddAgentToQueueService struct {
	domain.QueueRepository
	domainAgent.AgentRepository
}

func (j AddAgentToQueueService) Add(ctx context.Context, idQueue uuid.VoId, idAgent uuid.VoId) error {
	queue, err := j.QueueRepository.Search(ctx, idQueue)
	if err != nil {
		return err
	}
	_, err = j.AgentRepository.Search(ctx, idQueue)
	if err != nil {
		return err
	}
	queue.AddAgent(idAgent)
	err = j.QueueRepository.Upsert(ctx, queue)
	if err != nil {
		return err
	}
	return nil
}

func NewAddAgentToQueueService(qrepo domain.QueueRepository, arepo domainAgent.AgentRepository) AddAgentToQueueService {
	return AddAgentToQueueService{
		QueueRepository: qrepo,
		AgentRepository: arepo,
	}
}
