package creator_agent

import (
	"ari-queue/internal/agent/domain"
	"context"
)

type CreatorAgentService struct {
	domain.AgentRepository
}

func (c CreatorAgentService) Create(ctx context.Context, id string, code, phoneNumber int, name string) error {
	agent := domain.NewAgent(id, code, phoneNumber, name)
	err := c.AgentRepository.Save(ctx, agent)
	if err != nil {
		return err
	}
	return nil
}

func NewCreatorAgentService(repo domain.AgentRepository) CreatorAgentService {
	return CreatorAgentService{
		AgentRepository: repo,
	}
}
