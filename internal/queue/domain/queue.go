package domain

import (
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
	"sync"
)

type QueueRepository interface {
	Save(ctx context.Context, q Queue) error
	Search(ctx context.Context, id uuid.VoId) error
	Update(ctx context.Context, q Queue) error
	Delete(ctx context.Context, id uuid.VoId) error
}

type Queue struct {
	sync.Mutex
	ID                  uuid.VoId
	CustomersWaiting    []uuid.VoId
	CustomersConnecting map[uuid.VoId]uuid.VoId
	CustomersConnected  map[uuid.VoId]uuid.VoId
	Agents              []uuid.VoId
	Musicclass          string
}

func (q Queue) AddAgent(id uuid.VoId) {
	q.Lock()
	defer q.Unlock()
	q.Agents = append(q.Agents, id)
}

func (q Queue) AddCustomer(id uuid.VoId) {
	q.Lock()
	defer q.Unlock()
	q.Customers = append(q.Customers, id)
}
