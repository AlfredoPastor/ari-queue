package domain

import (
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
)

type QueueRepository interface {
	Search(ctx context.Context, id uuid.VoId) (Queue, error)
	Upsert(ctx context.Context, q Queue) error
	Delete(ctx context.Context, id uuid.VoId) error
}

type Queue struct {
	ID         uuid.VoId
	Name       string
	Customers  []uuid.VoId
	Agents     []uuid.VoId
	Musicclass string
}

func NewQueue(id, name string, agents []uuid.VoId) (Queue, error) {
	idVo, err := uuid.NewVoIdFromString(id)
	if err != nil {
		return Queue{}, err
	}
	var agentsVO []uuid.VoId
	agentsVO = append(agentsVO, agents...)
	return Queue{
		ID:        idVo,
		Name:      name,
		Agents:    agentsVO,
		Customers: []uuid.VoId{},
	}, nil
}

func (q Queue) AddAgent(id uuid.VoId) {
	q.Agents = append(q.Agents, id)
}

func (q Queue) AddCustomer(id uuid.VoId) {
	q.Customers = append(q.Customers, id)
}
