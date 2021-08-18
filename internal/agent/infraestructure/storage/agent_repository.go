package storage

import (
	"ari-queue/internal/agent/domain"
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
	"errors"
	"sync"
)

type AgentRepository struct {
	sync.Mutex
	agents map[uuid.VoId]domain.Agent
}

func (a *AgentRepository) Upsert(ctx context.Context, agent domain.Agent) error {
	a.Lock()
	defer a.Unlock()
	a.agents[agent.ID] = agent
	return nil
}

func (a *AgentRepository) Delete(ctx context.Context, id uuid.VoId) error {
	a.Lock()
	defer a.Unlock()
	delete(a.agents, id)
	return nil
}

func (a *AgentRepository) Search(ctx context.Context, id uuid.VoId) (domain.Agent, error) {
	a.Lock()
	defer a.Unlock()
	agent, ok := a.agents[id]
	if !ok {
		return domain.Agent{}, errors.New("agent not found")
	}
	return agent, nil
}

func NewAgentRepository() AgentRepository {
	return AgentRepository{
		agents: make(map[uuid.VoId]domain.Agent),
	}
}
