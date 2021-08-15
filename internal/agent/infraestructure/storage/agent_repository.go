package storage

import (
	"ari-queue/internal/agent/domain"
	"context"
)

type AgentRepository struct {
}

func (a AgentRepository) Save(ctx context.Context, agent domain.Agent) error {
	return nil
}
