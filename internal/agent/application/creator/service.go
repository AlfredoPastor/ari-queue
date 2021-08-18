package creator

import (
	"ari-queue/internal/agent/domain"
	"context"
)

type CreatorAgentService struct {
	domain.AgentRepository
}

func (c CreatorAgentService) Create(ctx context.Context, id string, code int, name, state string, phone domain.Phone) error {
	agent, err := domain.NewAgent(id, code, name, state, phone)
	if err != nil {
		return err
	}
	err = c.AgentRepository.Save(ctx, agent)
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
